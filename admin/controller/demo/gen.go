package demo

import (
	"rtsys/admin/lib"
	"rtsys/admin/services"
	"rtsys/utils/core"
	"rtsys/utils/tool"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *GenController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("index", false, c.Index))
	group.POST(c.Dispatch("table_list", false, c.TableList))
	group.POST(c.Dispatch("create", false, c.Create))
}

type GenController struct {
	lib.Controller
}

func NewGenController() *GenController {
	c := &GenController{}
	c.Id = "gen"
	c.Group = "demo"
	return c
}

func (c *GenController) Index(ctx *gin.Context) {
	dbList := core.G_CONFIG.DataBase
	c.Render(ctx, "index", gin.H{"dbList": dbList})
}

// TableList 获取数据库表列表
func (c *GenController) TableList(ctx *gin.Context) {
	db := ctx.DefaultQuery("db", "default")
	service := tool.NewSchemaOperate(db)

	lists := make([]string, 0)
	tables := service.TableList()

	for _, table := range tables {
		if strings.Index(table, "b5net_") != 0 && strings.Index(table, "demo_") != 0 {
			lists = append(lists, table)
		}
	}
	c.Success(ctx, "", lists)
}

// Create 创建操作
func (c *GenController) Create(ctx *gin.Context) {
	service := &services.GenService{}
	if err := ctx.ShouldBindJSON(service); err != nil {
		c.Error(ctx, "数据绑定失败:"+err.Error())
		return
	}
	if err := service.Save(); err != nil {
		c.Error(ctx, err.Error())
		return
	}
	c.Success(ctx, "创建成功")
}
