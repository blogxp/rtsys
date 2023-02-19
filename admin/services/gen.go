// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	"errors"
	"io/fs"
	"os"
	"rtsys/utils/tool"
	"strings"
)

type GenService struct {
	Db        string              `form:"db"`      //数据库标识
	Table     string              `form:"table"`   //表名
	Class     string              `form:"class"`   //类名
	Modules   string              `form:"modules"` //模块名
	Type      string              `form:"type"`    //类型
	Title     string              `form:"title"`   //标题
	Schema    *tool.SchemaOperate `form:"-"`
	ClassName string              `form:"-"` //处理后的类名
}

func (s *GenService) Save() error {
	if err := s.checkParams(); err != nil {
		return err
	}
	errList := []string{}
	if err := s.createModel(); err != nil {
		errList = append(errList, err.Error())
	}
	if err := s.createController(); err != nil {
		errList = append(errList, err.Error())
	}
	if err := s.createIndexHtml(); err != nil {
		errList = append(errList, err.Error())
	}
	if err := s.createAddHtml(); err != nil {
		errList = append(errList, err.Error())
	}
	if err := s.createEditHtml(); err != nil {
		errList = append(errList, err.Error())
	}
	if len(errList) > 0 {
		return s.error(strings.Join(errList, "<br/>"))
	}
	return nil
}

// createModel 创建Model文件
func (s *GenService) createModel() error {
	if s.Type != "0" && s.Type != "4" {
		return nil
	}
	//生成路径和文件
	genPath := "common/models/" + s.Modules + "/"
	genFile := genPath + s.Class + ".go"
	if tool.FileExist(genFile) {
		return s.error("model文件:" + s.Class + ".go 已存在")
	}

	//字段处理
	fieldHtml := ""
	fields := s.Schema.TableFields(s.Table)

	timeReplace := ""
	primaryReplace := "\"\""
	for _, field := range fields {
		humpName := tool.StrToHump(field.Name)
		columnType := "string"
		formName := field.Name
		if field.Type == "datetime" || field.Type == "date" {
			columnType = "*time.Time"
			timeReplace = "\"time\""
		}
		if field.Name == "create_time" || field.Name == "update_time" {
			formName = "-"
		}
		if field.Name == "id" || field.Key == "PRI" {
			primaryReplace = "m." + humpName
		}
		fieldHtml += "	" + humpName + "    " + columnType + "     `db:\"" + field.Name + "\" json:\"" + field.Name + "\" form:\"" + formName + "\"` // " + *field.Comment + " \n"
	}

	//模板路径及内容
	tempPath := "template/admin/demo/gen/create/model.tpl"
	tempFile, err := os.ReadFile(tempPath)
	if err != nil {
		return s.error("model模板文件读取失败:" + err.Error())
	}
	tempContent := string(tempFile)

	//替换占位信息
	tempContent = strings.ReplaceAll(tempContent, "__TIME__", timeReplace)
	tempContent = strings.ReplaceAll(tempContent, "__GROUP__", s.Modules)
	tempContent = strings.ReplaceAll(tempContent, "__MODEL__", s.ClassName+"Model")
	tempContent = strings.ReplaceAll(tempContent, "__TABLE__", s.Table)
	tempContent = strings.ReplaceAll(tempContent, "__FIELD__", fieldHtml)
	tempContent = strings.ReplaceAll(tempContent, "__PRIMARY__", primaryReplace)

	//生成路径
	if err := tool.DirCreate(genPath); err != nil {
		return s.error("创建model路径失败:" + err.Error())
	}

	//创建保存文件
	if err := os.WriteFile(genFile, []byte(tempContent), fs.ModePerm); err != nil {
		return errors.New("创建model文件失败:" + err.Error())
	}

	return nil
}

// createController 创建控制器文件
func (s *GenService) createController() error {
	if s.Type != "0" && s.Type != "1" && s.Type != "3" {
		return nil
	}

	//生成路径和文件
	genPath := "admin/controller/" + s.Modules + "/"
	genFile := genPath + s.Class + ".go"
	if tool.FileExist(genFile) {
		return s.error("controller文件:" + s.Class + ".go 已存在")
	}

	//模板路径及内容
	tempPath := "template/admin/demo/gen/create/controller.tpl"
	tempFile, err := os.ReadFile(tempPath)
	if err != nil {
		return s.error("controller模板文件读取失败:" + err.Error())
	}
	tempContent := string(tempFile)

	//判断模型文件是否存在
	modelPackage := ""
	modelNew := ""
	modelSlice := "nil"
	modelPath := "common/models/" + s.Modules + "/" + s.Class + ".go"
	if tool.FileExist(modelPath) {
		modelPackage = ". \"rtsys/common/models/" + s.Modules + "\""
		modelNew = "c.IModel = New" + s.ClassName + "Model()"
		modelSlice = "New" + s.ClassName + "Model().NewSlice()"
	}

	//替换占位信息
	tempContent = strings.ReplaceAll(tempContent, "__GROUP__", s.Modules)
	tempContent = strings.ReplaceAll(tempContent, "__MODEL_PACKAGE__", modelPackage)
	tempContent = strings.ReplaceAll(tempContent, "__CONTROLLER__", s.ClassName+"Controller")
	tempContent = strings.ReplaceAll(tempContent, "__ID__", s.Class)
	tempContent = strings.ReplaceAll(tempContent, "__MODEL_NEW__", modelNew)
	tempContent = strings.ReplaceAll(tempContent, "__MODEL_SLICE__", modelSlice)

	//生成路径
	if err := tool.DirCreate(genPath); err != nil {
		return s.error("创建controller路径失败:" + err.Error())
	}

	//创建保存文件
	if err := os.WriteFile(genFile, []byte(tempContent), fs.ModePerm); err != nil {
		return errors.New("创建controller文件失败:" + err.Error())
	}
	return nil
}

// createIndexHtml 创建index.html
func (s *GenService) createIndexHtml() error {
	if s.Type != "0" && s.Type != "2" && s.Type != "3" {
		return nil
	}
	//生成路径和文件
	genPath := "template/admin/" + s.Modules + "/" + s.Class + "/"
	genFile := genPath + "index.html"
	if tool.FileExist(genFile) {
		return s.error("index.html文件: 已存在")
	}

	//模板路径及内容
	tempPath := "template/admin/demo/gen/create/index.tpl"
	tempFile, err := os.ReadFile(tempPath)
	if err != nil {
		return s.error("index.html模板文件读取失败:" + err.Error())
	}
	tempContent := string(tempFile)

	//字段处理
	viewer := ""
	fieldHtml := ""
	fields := s.Schema.TableFields(s.Table)
	for _, field := range fields {
		title := *field.Comment
		if title == "" {
			title = field.Name
		}
		if field.Name == "id" {
			continue
		}

		//状态栏字段
		if field.Name == "status" || field.Name == "state" {
			fieldHtml += "				{\n" +
				"                	field: '" + field.Name + "',\n" +
				"                	title: '" + title + "',\n" +
				"					sortable: true,\n" +
				"					formatter: function (value, row, index) {\n" +
				"						return $.view.statusShow(row,false);\n" +
				"					}\n" +
				"				},\n"
			continue
		}

		//图片字段
		if strings.Index(field.Name, "img") > -1 {
			viewer = "{{block \"viewer\" .}}{{end}}"
			fieldHtml += "				{\n" +
				"                	field: '" + field.Name + "',\n" +
				"                	title: '" + title + "',\n" +
				"					formatter: function (value, row, index) {\n" +
				"						return $.table.imageView(row,'" + field.Name + "');\n" +
				"					}\n" +
				"				},\n"
			continue
		}
		if field.Name == "url" || field.Name == "link" {
			fieldHtml += "				{\n" +
				"                	field: '" + field.Name + "',\n" +
				"                	title: '" + title + "',\n" +
				"					formatter: function (value, row, index) {\n" +
				"						return $.table.tooltip(value,25,'link');\n" +
				"					}\n" +
				"				},\n"
			continue
		}

		//过长的字符串
		if field.Type == "varchar" && (field.Max != nil && *field.Max > 50) {
			fieldHtml += "				{\n" +
				"                	field: '" + field.Name + "',\n" +
				"                	title: '" + title + "',\n" +
				"					formatter: function (value, row, index) {\n" +
				"						return $.table.tooltip(value,25);\n" +
				"					}\n" +
				"				},\n"
			continue
		}

		//普通字段
		fieldHtml += "				{field: '" + field.Name + "',title: '" + title + "',align: 'center'},\n"
	}
	//替换占位信息
	tempContent = strings.ReplaceAll(tempContent, "__PATH__", s.Modules+"/"+s.Class+"/index")
	tempContent = strings.ReplaceAll(tempContent, "__TITLE__", s.Title)
	tempContent = strings.ReplaceAll(tempContent, "__VIEWER__", viewer)
	tempContent = strings.ReplaceAll(tempContent, "___REPLACE___", fieldHtml)

	//生成路径
	if err := tool.DirCreate(genPath); err != nil {
		return s.error("创建index.html路径失败:" + err.Error())
	}

	//创建保存文件
	if err := os.WriteFile(genFile, []byte(tempContent), fs.ModePerm); err != nil {
		return errors.New("创建index.html文件失败:" + err.Error())
	}
	return nil
}
func (s *GenService) createAddHtml() error {
	if s.Type != "0" && s.Type != "2" && s.Type != "3" {
		return nil
	}
	//生成路径和文件
	genPath := "template/admin/" + s.Modules + "/" + s.Class + "/"
	genFile := genPath + "add.html"
	if tool.FileExist(genFile) {
		return s.error("add.html文件: 已存在")
	}

	//模板路径及内容
	tempPath := "template/admin/demo/gen/create/add.tpl"
	tempFile, err := os.ReadFile(tempPath)
	if err != nil {
		return s.error("add.html模板文件读取失败:" + err.Error())
	}
	tempContent := string(tempFile)

	//字段处理
	fieldHtml := ""
	fields := s.Schema.TableFields(s.Table)
	for _, field := range fields {
		if field.Name == "id" || field.Name == "create_time" || field.Name == "update_time" {
			continue
		}
		title := *field.Comment
		if title == "" {
			title = field.Name
		}
		//状态栏字段
		if field.Name == "status" || field.Name == "state" {
			fieldHtml += "	<div class=\"form-group\">\n" +
				"		<label class=\"col-sm-3 control-label is-required\">" + title + "：</label>\n" +
				"		<div class=\"col-sm-8\">\n" +
				"			<label class=\"radio-box\">\n" +
				"				<input type=\"radio\" name=\"status\" value=\"0\"/> 隐藏\n" +
				"			</label>\n" +
				"			<label class=\"radio-box\">\n" +
				"				<input type=\"radio\" name=\"status\" value=\"1\" checked/> 显示\n" +
				"			</label>\n" +
				"		</div>\n" +
				"	</div>\n"
			continue
		}

		//文本框
		if field.Name == "remark" || field.Name == "note" || field.Name == "desc" {
			fieldHtml += "	<div class=\"form-group\">\n" +
				"		<label class=\"col-sm-3 control-label is-required\">" + title + "：</label>\n" +
				"		<div class=\"col-sm-8\">\n" +
				"			<textarea type=\"text\" name=\"" + field.Name + "\" class=\"form-control\"></textarea>\n" +
				"		</div>\n" +
				"	</div>\n"
			continue
		}

		//普通
		fieldHtml += "	<div class=\"form-group\">\n" +
			"		<label class=\"col-sm-3 control-label is-required\">" + title + "：</label>\n" +
			"		<div class=\"col-sm-8\">\n" +
			"			<input type=\"text\" name=\"" + field.Name + "\" class=\"form-control\" required autocomplete=\"off\"/>\n" +
			"		</div>\n" +
			"	</div>\n"
	}

	//替换占位信息
	tempContent = strings.ReplaceAll(tempContent, "__PATH__", s.Modules+"/"+s.Class+"/add")
	tempContent = strings.ReplaceAll(tempContent, "___REPLACE___", fieldHtml)

	//生成路径
	if err := tool.DirCreate(genPath); err != nil {
		return s.error("创建add.html路径失败:" + err.Error())
	}

	//创建保存文件
	if err := os.WriteFile(genFile, []byte(tempContent), fs.ModePerm); err != nil {
		return errors.New("创建add.html文件失败:" + err.Error())
	}
	return nil
}
func (s *GenService) createEditHtml() error {
	if s.Type != "0" && s.Type != "2" && s.Type != "3" {
		return nil
	}
	//生成路径和文件
	genPath := "template/admin/" + s.Modules + "/" + s.Class + "/"
	genFile := genPath + "edit.html"
	if tool.FileExist(genFile) {
		return s.error("edit.html文件: 已存在")
	}

	//模板路径及内容
	tempPath := "template/admin/demo/gen/create/edit.tpl"
	tempFile, err := os.ReadFile(tempPath)
	if err != nil {
		return s.error("edit.html模板文件读取失败:" + err.Error())
	}
	tempContent := string(tempFile)

	//字段处理
	fieldHtml := ""
	fields := s.Schema.TableFields(s.Table)
	for _, field := range fields {
		if field.Name == "id" || field.Name == "create_time" || field.Name == "update_time" {
			continue
		}
		humpName := tool.StrToHump(field.Name)
		title := *field.Comment
		if title == "" {
			title = field.Name
		}
		//状态栏字段
		if field.Name == "status" || field.Name == "state" {
			fieldHtml += "	<div class=\"form-group\">\n" +
				"		<label class=\"col-sm-3 control-label is-required\">" + title + "：</label>\n" +
				"		<div class=\"col-sm-8\">\n" +
				"			<label class=\"radio-box\">\n" +
				"				<input type=\"radio\" name=\"status\" value=\"0\" {{if eq $.info." + humpName + " \"0\"}} checked {{end}}/> 隐藏\n" +
				"			</label>\n" +
				"			<label class=\"radio-box\">\n" +
				"				<input type=\"radio\" name=\"status\" value=\"1\" {{if eq $.info." + humpName + " \"1\"}} checked {{end}}/> 显示\n" +
				"			</label>\n" +
				"		</div>\n" +
				"	</div>\n"
			continue
		}

		//文本框
		if field.Name == "remark" || field.Name == "note" || field.Name == "desc" {
			fieldHtml += "	<div class=\"form-group\">\n" +
				"		<label class=\"col-sm-3 control-label is-required\">" + title + "：</label>\n" +
				"		<div class=\"col-sm-8\">\n" +
				"			<textarea type=\"text\" name=\"" + field.Name + "\" class=\"form-control\">{{.info." + humpName + "}}</textarea>\n" +
				"		</div>\n" +
				"	</div>\n"
			continue
		}

		//普通
		fieldHtml += "	<div class=\"form-group\">\n" +
			"		<label class=\"col-sm-3 control-label is-required\">" + title + "：</label>\n" +
			"		<div class=\"col-sm-8\">\n" +
			"			<input type=\"text\" name=\"" + field.Name + "\" value=\"{{.info." + humpName + "}}\" class=\"form-control\" required autocomplete=\"off\"/>\n" +
			"		</div>\n" +
			"	</div>\n"
	}
	//替换占位信息
	tempContent = strings.ReplaceAll(tempContent, "__PATH__", s.Modules+"/"+s.Class+"/edit")
	tempContent = strings.ReplaceAll(tempContent, "___REPLACE___", fieldHtml)

	//生成路径
	if err := tool.DirCreate(genPath); err != nil {
		return s.error("创建edit.html路径失败:" + err.Error())
	}

	//创建保存文件
	if err := os.WriteFile(genFile, []byte(tempContent), fs.ModePerm); err != nil {
		return errors.New("创建edit.html文件失败:" + err.Error())
	}
	return nil
}

func (s *GenService) checkParams() error {
	if s.Db == "" {
		return s.error("请选择数据库")
	}
	if s.Table == "" {
		return s.error("请选择表")
	}
	if s.Class == "" || s.hasPath(s.Class) {
		return s.error("类名不能为空或包含路径符号")
	}
	if s.Modules == "" || s.hasPath(s.Modules) {
		return s.error("模块名不能为空或包含路径符号")
	}
	if s.Type == "" {
		return s.error("请选择创建模式")
	}
	schema := tool.NewSchemaOperate(s.Db)
	if !schema.TableExist(s.Table) {
		return s.error("表不存在")
	}
	s.Schema = schema
	s.Modules = strings.ToLower(s.Modules)
	s.Class = strings.ToLower(strings.ReplaceAll(s.Class, "-", "_"))
	s.ClassName = tool.StrToHump(s.Class) //类名转大驼峰
	return nil
}

func (s *GenService) error(err string) error {
	return errors.New(err)
}

func (s *GenService) hasPath(str string) bool {
	if strings.Index(str, "/") > -1 || strings.Index(str, "\\") > -1 {
		return true
	}
	return false
}
func (s *GenService) hasSep(str string) bool {
	if strings.Index(str, "_") > -1 || strings.Index(str, "-") > -1 {
		return true
	}
	return false
}
