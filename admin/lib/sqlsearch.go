// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

import (
	"rtsys/utils/core"
	"rtsys/utils/types"
	"strconv"
	"strings"
	"time"
)

// ** 简单封装的对后台列表查询的条件拼装 ** //

type FormSearch struct {
	Data *SearchFormData

	Where    string         //最终生成的 where 条件字符串
	whereArr []string       //最终生成的 where 条件字数组
	Args     []any          //最终生成的where 条件占位对应值数组
	Order    []types.KeyVal // 最终生成order 集合
	Limit    string

	BeforeWhere string
	BeforeArgs  []any
}

type SearchFormData struct {
	Like      map[string]string            `form:"like"`
	Where     map[string]string            `form:"where"`
	In        map[string]string            `form:"in"`
	FindInSet map[string]string            `form:"findinset"`
	Between   map[string]map[string]string `form:"between"` // 范围使用between  只有一个使用 > <
	Date      map[string]map[string]string `form:"date"`    //日期格式范围
	Time      map[string]map[string]string `form:"time"`    //时间格式范围

	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序规则
	PageNum       string `form:"pageNum"`       //页码
	PageSize      string `form:"pageSize"`      //每页长度
	IsTree        string `form:"isTree"`        //是否开启分页
	IsExport      string `form:"isExport"`      //是否导出

	Extend map[string]string `form:"extend"` //扩展自定义参数
}

func NewFormSearch(data *SearchFormData) *FormSearch {
	return &FormSearch{Where: "", Args: nil, Data: data}
}

func (fs *FormSearch) Bind(args ...FormSearchWithFunc) {
	for _, t := range args {
		t(fs)
	}
}
func (fs *FormSearch) Build() {
	fs.parseWhere()
	fs.parseLike()
	fs.parseFindInSet()
	fs.parseIn()
	fs.parseBetween()
	fs.parseDate()
	fs.parseTime()
	fs.parseOrder()
	fs.parsePage()
	fs.joinBefore()
}

func (fs *FormSearch) joinBefore() {
	whereStr := ""
	if len(fs.whereArr) > 0 {
		whereStr = "(" + strings.Join(fs.whereArr, " and ") + ")"
	}
	if fs.BeforeWhere != "" {
		if whereStr != "" {
			whereStr = whereStr + " and (" + fs.BeforeWhere + ")"
		} else {
			whereStr = fs.BeforeWhere
		}
	}
	fs.Where = whereStr

	if fs.BeforeArgs != nil && len(fs.BeforeArgs) > 0 {
		for _, v := range fs.BeforeArgs {
			fs.Args = append(fs.Args, v)
		}
	}
}
func (fs *FormSearch) parseWhere() {
	if fs.Data.Where == nil {
		return
	}
	for key, val := range fs.Data.Where {
		val = strings.Trim(val, " ")
		if val != "" && key != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` = ?")
			fs.Args = append(fs.Args, val)
		}
	}
}
func (fs *FormSearch) parseLike() {
	if fs.Data.Like == nil {
		return
	}
	for key, val := range fs.Data.Like {
		val = strings.Trim(val, " ")
		if val != "" && key != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` like ?")
			fs.Args = append(fs.Args, "%"+val+"%")
		}
	}
}
func (fs *FormSearch) parseFindInSet() {
	if fs.Data.FindInSet == nil {
		return
	}
	for key, val := range fs.Data.FindInSet {
		val = strings.Trim(val, " ")
		if val != "" && key != "" {
			fs.whereArr = append(fs.whereArr, "find_in_set ( ? ,`"+key+"`)")
			fs.Args = append(fs.Args, val)
		}
	}
}
func (fs *FormSearch) parseIn() {
	if fs.Data.In == nil {
		return
	}
	for key, val := range fs.Data.In {
		val = strings.Trim(val, " ")
		if val != "" && key != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` in (?)")
			fs.Args = append(fs.Args, strings.Split(val, ","))
		}
	}
}
func (fs *FormSearch) parseBetween() {
	if fs.Data.Between == nil {
		return
	}
	for key, val := range fs.Data.Between {
		start := ""
		end := ""
		if s, sok := val["start"]; sok && s != "" {
			start = strings.Trim(s, " ")
		}
		if e, eok := val["end"]; eok && e != "" {
			end = strings.Trim(e, " ")
		}
		if start != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` >= ?")
			fs.Args = append(fs.Args, start)
		}
		if end != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` <= ?")
			fs.Args = append(fs.Args, end)
		}
	}
}
func (fs *FormSearch) parseDate() {
	for key, val := range fs.Data.Date {
		start := ""
		end := ""
		if s, sok := val["start"]; sok && s != "" {
			start = strings.Trim(s, " ")
		}
		if e, eok := val["end"]; eok && e != "" {
			end = strings.Trim(e, " ")
		}
		if start == "" && end == "" {
			continue
		}
		if start != "" && end != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` between ? and ?")
			fs.Args = append(fs.Args, start)
			fs.Args = append(fs.Args, end)
		} else if start != "" {
			parse, err := time.Parse(core.G_DATE, start)
			if err == nil {
				start = parse.Format(core.G_TIME)

				fs.whereArr = append(fs.whereArr, "UNIX_TIMESTAMP(`"+key+"`) >= UNIX_TIMESTAMP(?)")
				fs.Args = append(fs.Args, start)
			}
		} else if end != "" {
			parse, err := time.Parse(core.G_DATE, end)
			if err == nil {
				end = parse.Format(core.G_TIME)
				fs.whereArr = append(fs.whereArr, "UNIX_TIMESTAMP(`"+key+"`) <= UNIX_TIMESTAMP(?)")
				fs.Args = append(fs.Args, end)
			}
		}
	}
}

func (fs *FormSearch) parseTime() {
	for key, val := range fs.Data.Time {
		start := ""
		end := ""
		if s, sok := val["start"]; sok && s != "" {
			start = strings.Trim(s, " ")
		}
		if e, eok := val["end"]; eok && e != "" {
			end = strings.Trim(e, " ")
		}
		if start == "" && end == "" {
			continue
		}
		if start != "" && end != "" {
			fs.whereArr = append(fs.whereArr, "`"+key+"` between ? and ?")
			fs.Args = append(fs.Args, start)
			fs.Args = append(fs.Args, end)
		} else if start != "" {
			parse, err := time.Parse(core.G_TIME, start)
			if err == nil {
				start = parse.Format(core.G_TIME)

				fs.whereArr = append(fs.whereArr, "UNIX_TIMESTAMP(`"+key+"`) >= UNIX_TIMESTAMP(?)")
				fs.Args = append(fs.Args, start)
			}
		} else if end != "" {
			parse, err := time.Parse(core.G_TIME, end)
			if err == nil {
				end = parse.Format(core.G_TIME)
				fs.whereArr = append(fs.whereArr, "UNIX_TIMESTAMP(`"+key+"`) <= UNIX_TIMESTAMP(?)")
				fs.Args = append(fs.Args, end)
			}
		}
	}
}
func (fs *FormSearch) parseOrder() {
	if fs.Data.OrderByColumn == "" {
		return
	}
	orderByAsc := "asc"
	if fs.Data.IsAsc != "" {
		orderByAsc = fs.Data.IsAsc
	}
	order := fs.Order
	if order == nil {
		order = make([]types.KeyVal, 0)
	}
	has := false
	for _, v := range order {
		if v.Key == fs.Data.OrderByColumn {
			has = true
			break
		}
	}
	if !has {
		order = append(order, types.KeyVal{Key: fs.Data.OrderByColumn, Value: orderByAsc})
	}
	fs.Order = order
}
func (fs *FormSearch) parsePage() {
	if fs.Data.IsTree == "1" || fs.Data.IsExport == "1" {
		fs.Limit = ""
		return
	}
	pageSize := "10"
	pageNum := "1"
	if fs.Data.PageSize != "" {
		pageSize = fs.Data.PageSize
	}
	if fs.Data.PageNum != "" {
		pageNum = fs.Data.PageNum
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}
	pageNumInt, err := strconv.Atoi(pageNum)
	if err != nil {
		pageNumInt = 1
	}
	offset := (pageNumInt - 1) * pageSizeInt
	fs.Limit = strconv.Itoa(offset) + "," + strconv.Itoa(pageSizeInt)
}

type FormSearchWithFunc func(*FormSearch)

func WithBeforeWhere(args string) FormSearchWithFunc {
	return func(rs *FormSearch) {
		rs.BeforeWhere = args
	}
}
func WithBeforeArgs(args []any) FormSearchWithFunc {
	return func(rs *FormSearch) {
		rs.BeforeArgs = args
	}
}
func WithBeforeOrder(args []types.KeyVal) FormSearchWithFunc {
	return func(rs *FormSearch) {
		rs.Order = args
	}
}
