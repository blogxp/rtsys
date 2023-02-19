package cache

import (
	. "rtsys/common/daos/system"
	. "rtsys/common/models/system"
	"rtsys/utils/core"

	"github.com/go-redis/redis"
)

type ConfigCache struct {
	key   string        //redis存储的键
	redis *redis.Client //redis连接对象
}

func NewConfigCache() *ConfigCache {
	return &ConfigCache{
		key:   "b5_config_list",
		redis: core.G_Redis.Conn(),
	}
}

func (c *ConfigCache) GetValue(key string) (result string) {
	result = ""
	if key == "" {
		return
	}

	if res, err := c.redis.HGet(c.key, key).Result(); err != nil {
		doRes, _ := core.G_Single.Do(c.key, func() (interface{}, error) {
			if val := NewConfigDao().GetInfoByType(key); val != nil {
				c.redis.HSet(c.key, key, result)
				return val, nil
			}
			return "", nil
		})
		result = doRes.(string)
	} else {
		result = res
	}
	return
}

func (c *ConfigCache) FlushAll() {
	c.redis.Del(c.key)
}

func (c *ConfigCache) BuildAll() {
	c.FlushAll()
	model := NewConfigModel()
	list := model.NewSlice()
	err := core.NewDao(model).SetField("type,value").Lists(list, "")
	if err != nil {
		return
	}
	var fields = map[string]any{}
	for _, item := range *list {
		fields[item.Type] = item.Value
	}
	c.redis.HMSet(c.key, fields)
}
