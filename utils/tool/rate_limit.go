// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

////////   简单实现的 固定窗口和浮动窗口限流 方法  令牌桶和漏桶需要定时不写了    ///////////

package tool

import (
	"log"
	"rtsys/utils/core"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

type RateLimitKind string

const (
	FIX_RATE   RateLimitKind = "fix"   //固定窗口
	SLIDE_RATE RateLimitKind = "slide" //浮动窗口
)

type IRateLimit interface {
	CheckLimit(key string) bool
}

type RateLimit struct { //基础机构
	RedisClient *redis.Client
	sync.Mutex
}

func NewRateLimit(kind RateLimitKind, args map[string]int) IRateLimit {
	if kind == FIX_RATE {
		return NewFixRateLimit(args)
	} else if kind == SLIDE_RATE {
		return NewSlideRateLimit(args)
	}
	return &InValidRateLimit{}
}

// FixRateLimit 固定窗口限流  用于Expire秒内，最多运行Limit次请求
type FixRateLimit struct {
	RateLimit
	Limit  int  //限制请求次数 小于1 代表不限制  优先级高于Expire
	Expire int  //秒显示时间范围  小于1 代表永远拒绝
	useLua bool //是否使用lua脚本  ，测试了下 lua还不如直接操作
}

func NewFixRateLimit(args map[string]int) *FixRateLimit {
	l := &FixRateLimit{useLua: false}
	if args != nil {
		if limit, ok := args["limit"]; ok {
			l.Limit = limit
		}
		if expire, ok := args["exp"]; ok {
			l.Expire = expire
		}
	}
	l.RedisClient = core.G_Redis.Conn()
	return l
}

func (l *FixRateLimit) CheckLimit(key string) bool {
	if key == "" || l.Limit < 1 {
		return true
	}
	if l.Expire < 1 {
		return false
	}
	l.Lock()
	defer l.Unlock()

	if l.useLua {
		//使用lua脚本的原子性
		script := `
local redis_val = redis.call('get',KEYS[1])
if redis_val == false then
	redis.call('setex',KEYS[1],tonumber(ARGV[1]),1)
else
	if tonumber(redis_val) >= tonumber(ARGV[2]) then
		return '0'
	end
	redis.call('incr',KEYS[1])
end
return '1'
`
		if result, err := l.RedisClient.Eval(script, []string{key}, l.Expire, l.Limit).Result(); err == nil {
			if v, ok := result.(string); ok {
				if v == "0" {
					return false
				}
			}
		} else {
			log.Println("redis执行lua错误:", err)
		}
	} else {
		if res, err := l.RedisClient.Get(key).Result(); err == nil {
			num, _ := strconv.Atoi(res)
			if num >= l.Limit {
				return false
			}
			l.RedisClient.Incr(key) //不会清除过期时间
			return true
		}
		l.RedisClient.Set(key, 1, time.Duration(l.Expire)*time.Second)
	}
	return true
}

// SlideRateLimit 浮动窗口限流
// 使用zset有序集合实现，创建有效期为Expire*Length的集合
// 插入当前时间戳和Expire
// NewRateLimit(SLIDE_RATE, map[string]int{"limit":2,"exp":3,"len":5,"total":5})
//
//	5个窗口 每个窗口间隔3秒 最多访问2次， 5个窗口内最多访问5次
type SlideRateLimit struct {
	RateLimit
	Limit  int  // 每个窗口限制请求次数 小于1 代表不限制  优先级高于Expire
	Expire int  // 每个窗口的时间间隔  小于1 代表永远拒绝
	Length int  // 窗口数量 小于1则默认为1个窗口，相当于固定窗口了
	Total  int  // 总窗口限制请求次数 默认为 Limit * Length
	useLua bool //是否使用lua脚本  ，测试了下 lua还不如直接操作
}

func NewSlideRateLimit(args map[string]int) *SlideRateLimit {
	l := &SlideRateLimit{useLua: false}
	if args != nil {
		if limit, ok := args["limit"]; ok {
			l.Limit = limit
		}
		if expire, ok := args["exp"]; ok {
			l.Expire = expire
		}
		if length, ok := args["len"]; ok {
			l.Length = length
		}
		if total, ok := args["total"]; ok {
			l.Total = total
		}
	}
	l.RedisClient = core.G_Redis.Conn()
	return l
}

// 哈希表实现
func (s *SlideRateLimit) CheckLimit(key string) bool {
	if s.Limit < 1 {
		return true
	}
	if s.Expire < 1 {
		return false
	}
	if s.Length <= 1 {
		f := &FixRateLimit{Limit: s.Limit, Expire: s.Expire}
		f.RedisClient = s.RedisClient
		return s.CheckLimit(key)
	}
	s.Lock()
	defer s.Unlock()
	total := s.Total //总限制次数
	if total < 1 {
		total = s.Limit * s.Length
	}
	cycle := s.Length * s.Expire       //总周期
	nowTime := time.Now().Unix()       //当前时间戳
	lastTime := nowTime - int64(cycle) //保存在字段的最小的值，小于该值的删除

	//当前的窗口信息
	var suplus = nowTime % int64(s.Expire) //余数
	var currentIndex = nowTime - suplus    //当前窗口的score
	var currentIndexStr = strconv.FormatInt(currentIndex, 10)

	if s.useLua {
		script := `
local num = 0
local count = 0

local current_val = redis.call('hget',KEYS[1],KEYS[2])
if current_val ~= false then
	if tonumber(current_val) >= tonumber(ARGV[1]) then
		return '0'
	else
		num = tonumber(current_val)
	end
end

local list = redis.call('hgetall',KEYS[1])
if list ~= false then
	for k, v in pairs(list) do
		if tonumber(k) <= tonumber(ARGV[2]) then
			redis.call('hdel',KEYS[1],KEYS[2])
		else
			count = count+ tonumber(v)
		end
	end
end

if count >= tonumber(ARGV[3]) then
	return '0'
end

redis.call('hset',KEYS[1],KEYS[2],num+1)
redis.call('expire',KEYS[1],tonumber(ARGV[4]))
return '1'
`
		if result, err := s.RedisClient.Eval(script, []string{key, currentIndexStr}, s.Limit, lastTime, total, cycle*2).Result(); err == nil {
			if v, ok := result.(string); ok {
				if v == "0" {
					return false
				}
			}
		} else {
			log.Println("redis执行lua错误:", err)
		}
	} else {
		var num = 0   //当前窗口次数
		var count = 0 //总次数

		//查询当前窗口信息
		if currentVal, err := s.RedisClient.HGet(key, currentIndexStr).Result(); err == nil {
			val, _ := strconv.Atoi(currentVal)
			if val >= s.Limit {
				return false
			}
			num = val
		}

		if result, err := s.RedisClient.HGetAll(key).Result(); err == nil {
			for field, val := range result {
				fieldInt, _ := strconv.Atoi(field)
				if int64(fieldInt) <= lastTime { //过期删除
					s.RedisClient.HDel(key, field)
					continue
				}
				vInt, _ := strconv.Atoi(val)
				count += vInt
			}
		}
		if count >= total {
			return false
		}

		s.RedisClient.HSet(key, currentIndexStr, num+1)
		s.RedisClient.Expire(key, time.Duration(cycle*2)*time.Second)
	}

	return true
}

// 有序集合实现 性能不理想
func (s *SlideRateLimit) CheckLimit11(key string) bool {
	if s.Limit < 1 {
		return true
	}
	if s.Expire < 1 {
		return false
	}
	if s.Length <= 1 {
		f := &FixRateLimit{Limit: s.Limit, Expire: s.Expire}
		f.RedisClient = s.RedisClient
		return s.CheckLimit(key)
	}
	s.Lock()
	defer s.Unlock()
	total := s.Total //总限制次数
	if total < 1 {
		total = s.Limit * s.Length
	}

	var precision float64 = 100000 //精度 存储的 score使用时间戳+num/precison 5位小数表示
	precisionDec := decimal.NewFromFloat(precision)

	cycle := s.Length * s.Expire //总周期

	nowTime := time.Now().Unix() //当前时间戳

	//保存在集合中的最小的score，小于该值的删除
	lastTime := nowTime - int64(cycle) + 1

	//当前的窗口信息
	var suplus = nowTime % int64(s.Expire) //余数
	var currentIndex = nowTime - suplus    //当前窗口的score
	var currentIndexStr = strconv.FormatInt(currentIndex, 10)

	//查找key再redis是否存在
	if exists, err := s.RedisClient.Exists(key).Result(); err == nil {
		num := decimal.NewFromInt(0)   //当前窗口次数
		count := decimal.NewFromInt(0) //总次数

		if exists == 1 {
			//删除lastTime之前的数据
			s.RedisClient.ZRemRangeByScore(key, "0", strconv.FormatInt(lastTime, 10))

			//获取当前所有有效的集合
			if result, err1 := s.RedisClient.ZRevRangeWithScores(key, 0, -1).Result(); err1 == nil {
				for _, v := range result {
					score := decimal.NewFromFloat(v.Score)

					numItem := score.Sub(score.Floor()).Mul(precisionDec).Floor() //当前窗口的已访问次数

					if v.Member.(string) == currentIndexStr {
						if numItem.Sub(decimal.NewFromInt(int64(s.Limit))).Sign() >= 0 { //当前窗口超出
							return false
						}
						num = numItem
					}
					count = count.Add(numItem)
				}
			} else {
				log.Println("redis执行错误1:", err)
			}
			if count.Sub(decimal.NewFromInt(int64(total))).Sign() >= 0 {
				return false
			}
		}

		//更新当前的窗口数据
		newNum := num.Add(decimal.NewFromInt(1)).Div(precisionDec)
		upNum, _ := decimal.NewFromInt(currentIndex).Add(newNum).Float64()
		s.RedisClient.ZAdd(key, redis.Z{Member: currentIndexStr, Score: upNum})

		//总有效期，为所有窗口的和的2倍，防止间隔时间将近和时，前面的窗口数据丢失
		s.RedisClient.Expire(key, time.Duration(cycle*2)*time.Second)
	} else {
		log.Println("redis执行错误:", err)
	}
	return true
}

// InValidRateLimit 定义的无效的 永远返回true
type InValidRateLimit struct {
}

func (l *InValidRateLimit) CheckLimit(key string) bool {
	return true
}
