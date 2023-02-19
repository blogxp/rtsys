// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

// ** 后台公共操作方法增删改查实现 ** //

import (
	"fmt"
	"net/http"
	. "rtsys/utils/core"
	. "rtsys/utils/export"
	. "rtsys/utils/trans"
	"rtsys/utils/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HookAction struct {
	BeforeAction func(ctx *gin.Context, action string) bool //该controller所有路由进行请求前的逻辑判断处理

	IndexRender  func(ctx *gin.Context) (map[string]any, error)                  //首页渲染前进行数据分配及逻辑处理
	IndexAfter   func(ctx *gin.Context, list []any) []any                        //首页列表返回前对返回的数据进行处理
	AddRender    func(ctx *gin.Context)                                          //添加页渲染前进行数据分配及逻辑处理
	EditRender   func(ctx *gin.Context, model IModel)                            //编辑页渲染前进行数据分配及逻辑处理
	SaveBefore   func(ctx *gin.Context, model IModel, operate string) error      // 添加编辑保存前
	SaveAfter    func(ctx *gin.Context, model IModel, operate string, id string) // 添加编辑保存后
	DropBefore   func(ctx *gin.Context, model IModel) error                      //删除前的逻辑判断处理
	DropAfter    func(ctx *gin.Context, model IModel)                            //删除后的操作
	ExportBefore func(ctx *gin.Context, list []any) *ExcelExport                 //导出前处理
}

func (ca *Controller) Index(ctx *gin.Context) {
	var data map[string]any
	//渲染前的钩子方法
	if ca.IndexRender != nil {
		if re, err := ca.IndexRender(ctx); err != nil {
			ca.Jump(ctx, err)
			return
		} else {
			data = re
		}
	}
	fmt.Println(data)
	ca.Render(ctx, "index", data)
}

// GetIndex 处理列表查询操作的子方法，必须在具体控制中调用
// @params  list any 是一个指针型的Model数组
func (ca *Controller) GetIndex(ctx *gin.Context, list any, where string, replace []any, order []types.KeyVal) {
	if replace == nil {
		replace = make([]any, 0)
	}
	if list == nil {
		ca.Error(ctx, "列表model实例错误")
		return
	}
	//条件拼接
	data := &SearchFormData{}
	_ = ctx.ShouldBindBodyWith(data, binding.JSON)

	formSearch := NewFormSearch(data)
	formSearch.Bind(WithBeforeWhere(where), WithBeforeArgs(replace))
	formSearch.Bind(WithBeforeOrder(order))
	formSearch.Build()

	//获取总数
	totalInfo := &MTotal{}
	if formSearch.Data.IsExport != "1" && formSearch.Data.IsTree != "1" {
		err0 := NewDao(ca.IModel).SetField("count(*) as total").First(totalInfo, formSearch.Where, formSearch.Args...)
		if err0 != nil {
			ca.Error(ctx, "查询错误:"+err0.Error())
			return
		}
	}

	err := NewDao(ca.IModel).SetOrderBy(formSearch.Order).SetLimit(formSearch.Limit).Lists(list, formSearch.Where, formSearch.Args...)
	//将any转为[]any
	lists := AnyToSlice(list)
	if err != nil {
		fmt.Println(err)
	} else {
		if ca.IndexAfter != nil { //渲染前对列表进行处理
			lists = ca.IndexAfter(ctx, lists)
		}
	}
	if formSearch.Data.IsExport == "1" {
		//结果查询后的处理
		if ca.ExportBefore == nil {
			ca.Error(ctx, "未设置ExportBefore")
			return
		}
		excel := ca.ExportBefore(ctx, lists)
		err10 := excel.New(true)
		if err10 != nil {
			ca.Error(ctx, err10.Error())
			return
		}
		err11 := excel.Export()
		if err11 != nil {
			ca.Error(ctx, err11.Error())
			return
		}
		ca.Success(ctx, excel.Url)
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{"code": RSuccess, "msg": "操作成功", "data": lists, "total": totalInfo.Total})
	//ca.Success("", list)
}

// Add 通用添加页渲染
// 还可以使用AddRender 定义前置操作
func (ca *Controller) Add(ctx *gin.Context) {
	var data map[string]any

	//渲染前的钩子方法
	if ca.AddRender != nil {
		ca.AddRender(ctx)
		return
	}
	ca.Render(ctx, "add", data)
}

// AddSave 通用保存操作
func (ca *Controller) AddSave(ctx *gin.Context) {
	dao := NewDao(ca.IModel)
	model := dao.INew()

	err := ctx.ShouldBindBodyWith(model, binding.JSON)
	if err != nil {
		ca.Error(ctx, "表单绑定失败："+err.Error())
		return
	}
	err1 := G_Validate.Struct(model)
	if err1 != nil {
		ca.Error(ctx, G_Validate.GetError(err1))
		return
	}
	if ca.SaveBefore != nil {
		if err4 := ca.SaveBefore(ctx, model, "add"); err4 != nil {
			ca.Error(ctx, err4.Error())
			return
		}
	}
	id, err2 := dao.InsertNamed(GetDAOMapFromStruct(model, []string{"id"}, true)) //named形式
	//_, err2 := dao.Insert(GetStructDBKV(model, []string{"id"}, true))
	if err2 != nil {
		ca.Error(ctx, "插入失败："+err2.Error())
		return
	}
	if ca.SaveAfter != nil {
		ca.SaveAfter(ctx, model, "add", id)
	}
	ca.Success(ctx, "信息保存成功")
}

// Edit 通用编辑页渲染
func (ca *Controller) Edit(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ca.Error(ctx, "参数错误")
		return
	}
	dao := NewDao(ca.IModel)
	model := dao.INew()
	err := dao.First(model, id)
	if err != nil {
		ca.Error(ctx, "信息查询不存在")
		return
	}

	if ca.EditRender != nil {
		ca.EditRender(ctx, model)
		return
	}
	data := map[string]any{"info": model}
	ca.Render(ctx, "edit", data)
}

// EditSave 通用编辑保存操作
func (ca *Controller) EditSave(ctx *gin.Context) {
	idForm := &types.SimpleId{}
	_ = ctx.ShouldBindBodyWith(idForm, binding.JSON)
	if idForm.Id == "" {
		ca.Error(ctx, "Id参数丢失")
		return
	}
	dao := NewDao(ca.IModel)
	model := dao.INew()
	err2 := dao.First(model, idForm.Id)
	if err2 != nil {
		ca.Error(ctx, "查询错误"+err2.Error())
		return
	}
	err := ctx.ShouldBindBodyWith(model, binding.JSON)
	if err != nil {
		ca.Error(ctx, "表单绑定失败："+err.Error())
		return
	}
	err1 := G_Validate.Struct(model)
	if err1 != nil {
		ca.Error(ctx, G_Validate.GetError(err1))
		return
	}
	if ca.SaveBefore != nil {
		if err4 := ca.SaveBefore(ctx, model, "edit"); err4 != nil {
			ca.Error(ctx, err4.Error())
			return
		}
	}
	upMap := GetDAOMapFromStruct(model, []string{"id", "create_time"}, true)
	effected, err3 := dao.UpdateNamed(upMap, map[string]any{"id": model.GetId()})
	//fields, values := GetStructDBKV(model, []string{"id", "create_time"}, true)
	//effected, err3 := dao.Update(fields,values,id)
	if err3 != nil {
		ca.Error(ctx, "更新失败："+err3.Error())
		return
	}
	if effected > 0 {
		if ca.SaveAfter != nil {
			ca.SaveAfter(ctx, model, "edit", model.GetId())
		}
		ca.Success(ctx, "信息更新成功")
	} else {
		ca.Success(ctx, "操作成功，数据未变化")
	}
}

// Drop 通用单条删除操作
func (ca *Controller) Drop(ctx *gin.Context) {
	formModel := &types.SimpleId{}
	err := ctx.ShouldBindJSON(formModel)
	if err != nil {
		ca.Error(ctx, "参数绑定失败："+err.Error())
		return
	}
	if formModel.Id == "" {
		ca.Error(ctx, "Id参数丢失")
		return
	}
	dao := NewDao(ca.IModel)
	model := dao.INew()

	err1 := dao.First(model, formModel.Id) //查询信息
	if err1 != nil {
		ca.Error(ctx, "信息不存在或已删除")
		return
	}
	if ca.DropBefore != nil { //删除前的操作
		if err4 := ca.DropBefore(ctx, model); err4 != nil {
			ca.Error(ctx, err4.Error())
			return
		}
	}
	affected, err2 := dao.Delete(model.GetId()) //删除
	if err2 != nil {
		ca.Error(ctx, "删除失败："+err2.Error())
		return
	}
	if affected > 0 && ca.DropAfter != nil { //删除后操作
		ca.DropAfter(ctx, model)
	}
	ca.Success(ctx, fmt.Sprintf("成功删除%d条记录", affected))
}

// DropAll 通用批量删除操作
func (ca *Controller) DropAll(ctx *gin.Context) {
	formModel := &types.SimpleId{}
	err := ctx.ShouldBindJSON(formModel)
	if err != nil {
		ca.Error(ctx, "参数绑定失败："+err.Error())
		return
	}
	if formModel.Ids == "" {
		ca.Error(ctx, "Id参数丢失")
		return
	}
	idList := strings.Split(formModel.Ids, ",")
	if len(idList) == 0 {
		ca.Error(ctx, "Id参数丢失")
		return
	}
	dao := NewDao(ca.IModel)
	i := 0
	for _, id := range idList {
		if id == "" {
			continue
		}
		model := dao.INew()
		err1 := dao.First(model, id) //查询信息
		if err1 != nil {
			continue
		}
		if ca.DropBefore != nil { //删除前的操作
			if err4 := ca.DropBefore(ctx, model); err4 != nil {
				continue
			}
		}
		affected, err2 := dao.Delete(model.GetId()) //删除
		if err2 != nil {
			continue
		}
		if affected < 1 {
			continue
		}
		i++
		if ca.DropAfter != nil { //删除后操作
			ca.DropAfter(ctx, model)
		}
	}
	ca.Success(ctx, fmt.Sprintf("成功删除%d条记录", i))
}

// SetStatus 通用状态修改操作
func (ca *Controller) SetStatus(ctx *gin.Context) {
	setStatusData := &SetStatusData{}
	err := ctx.ShouldBindJSON(setStatusData)
	if err != nil {
		ca.Error(ctx, "参数绑定失败："+err.Error())
		return
	}
	if setStatusData.Id == "" {
		ca.Error(ctx, "缺少编辑主键条件")
		return
	}
	if setStatusData.Status != "0" && setStatusData.Status != "1" {
		ca.Error(ctx, "状态错误")
		return
	}
	dao := NewDao(ca.IModel)
	model := dao.INew()
	err2 := dao.First(model, setStatusData.Id)
	if err2 != nil {
		ca.Error(ctx, "信息查询错误"+err2.Error())
		return
	}
	field := setStatusData.Field
	if field == "" {
		field = "status"
	}
	title := setStatusData.Name
	if title == "" {
		if setStatusData.Status == "1" {
			title = "启用"
		} else {
			title = "停用"
		}
	}

	infoMap := StructToMap(model, true)
	infoStatus, ok := infoMap["status"]
	if !ok {
		ca.Error(ctx, "信息不含有"+field+"字段")
		return
	}
	if infoStatus == setStatusData.Status {
		ca.Success(ctx, title+"成功")
		return
	}
	if ca.SaveBefore != nil {
		if err4 := ca.SaveBefore(ctx, model, "status"); err4 != nil {
			ca.Error(ctx, err4.Error())
			return
		}
	}
	upFields := []string{field}
	upValues := []any{setStatusData.Status}
	_, ok1 := infoMap["update_time"]
	if ok1 {
		upFields = append(upFields, "update_time")
		upValues = append(upValues, TimeFormat(time.Now(), ""))
	}
	effected, err3 := dao.Update(upFields, upValues, "`id` = ?", model.GetId())
	if err3 != nil {
		ca.Error(ctx, title+"失败："+err3.Error())
		return
	}
	if effected > 0 {
		if ca.SaveAfter != nil {
			ca.SaveAfter(ctx, model, "status", model.GetId())
		}
		ca.Success(ctx, title+"成功")
	} else {
		ca.Success(ctx, "操作成功，数据未变化")
	}
}
