// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package trans

import (
	"reflect"
	. "rtsys/utils/tool"
	"strings"
	"time"
)

// AnyToSlice 将any(interface{})转为切片
func AnyToSlice(arr any) []any {
	v := reflect.ValueOf(arr)
	ptr := false
	l := 0
	if v.Kind() == reflect.Ptr {
		if v.Elem().Kind() != reflect.Slice {
			return nil
		}
		ptr = true
		l = v.Elem().Len()
	} else {
		if v.Kind() != reflect.Slice {
			return nil
		}
		l = v.Len()
	}
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		if ptr {
			ret[i] = v.Elem().Index(i).Interface()
		} else {
			ret[i] = v.Index(i).Interface()
		}
	}
	return ret
}

func AnyToMapList(arr any) []map[string]any {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return nil
	}
	l := v.Len()
	ret := make([]map[string]any, l)
	for i := 0; i < l; i++ {
		item := v.Index(i).Interface()
		var data = make(map[string]any)
		vv := reflect.ValueOf(item)
		if vv.Kind() == reflect.Map {
			keys := vv.MapKeys()
			for _, elem := range keys {
				data[elem.String()] = vv.MapIndex(elem).Interface()
			}
		} else if vv.Kind() == reflect.Struct {
			tt := reflect.TypeOf(item)
			ll := tt.NumField()
			for ii := 0; ii < ll; ii++ {
				data[tt.Field(ii).Name] = vv.Field(ii).Interface()
			}
		} else {
			data = nil
		}
		ret[i] = data
	}
	return ret
}

// StructListToMapList 将结构体切片转为[]map[sting]any
func StructListToMapList(dest []any, isJson bool) []map[string]any {
	list := make([]map[string]any, len(dest))
	for index, item := range dest {
		list[index] = StructToMap(item, isJson)
	}
	return list
}

// StructToMap 结构体转map
func StructToMap(dest any, isJson bool) map[string]any {
	t := reflect.TypeOf(dest)
	ptr := false
	l := 0

	if t.Kind() == reflect.Ptr {
		ptr = true
		if t.Elem().Kind() == reflect.Map {
			return dest.(map[string]any)
		}
		if t.Elem().Kind() != reflect.Struct {
			return nil
		}
		l = t.Elem().NumField()
	} else {
		if t.Kind() == reflect.Map {
			return dest.(map[string]any)
		}
		if t.Kind() != reflect.Struct {
			return nil
		}
		l = t.NumField()
	}

	var data = make(map[string]any)
	v := reflect.ValueOf(dest)
	for i := 0; i < l; i++ {
		var key string
		var val any
		if ptr {
			key = t.Elem().Field(i).Name
			val = v.Elem().Field(i).Interface()
			if isJson {
				jsonKey := t.Elem().Field(i).Tag.Get("json")
				if jsonKey != "" {
					key = jsonKey
				}
			}
		} else {
			key = t.Field(i).Name
			val = v.Field(i).Interface()
			if isJson {
				jsonKey := t.Field(i).Tag.Get("json")
				if jsonKey != "" {
					key = jsonKey
				}
			}
		}
		data[key] = val
	}
	return data
}

// GetStructDBKV 将结构体进行拆分成键数组和值数组
// except 不需要的键（小写）
// 会对create_time/update_time 进行默认赋值
func GetStructDBKV(model any, except []string, autoTime bool) ([]string, []any) {
	v := reflect.ValueOf(model)
	t := v.Type()
	var fields []string
	var values []any
	for i := 0; i < v.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		var key string
		if field.Tag.Get("db") != "" {
			key = field.Tag.Get("db")
		} else {
			key = strings.ToLower(field.Name)
		}
		if ok := InArray(key, except); ok {
			continue
		}
		fields = append(fields, key)
		val := v.Elem().Field(i).Interface()
		if autoTime && (key == "create_time" || key == "update_time") {
			val = TimeFormat(time.Now(), "")
		}
		values = append(values, val)
	}
	return fields, values
}

// GetDAOMapFromStruct 将结构体转为InsertNamed需要的map
// 键为数据库字段名
func GetDAOMapFromStruct(model any, except []string, autoTime bool) map[string]any {
	key, val := GetStructDBKV(model, except, autoTime)
	m := make(map[string]any)
	for i := 0; i < len(key); i++ {
		m[key[i]] = val[i]
	}
	return m
}
