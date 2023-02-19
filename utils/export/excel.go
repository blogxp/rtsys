// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package export

import (
	"errors"
	"rtsys/utils/core"
	"rtsys/utils/tool"
	"rtsys/utils/types"
	"time"

	"github.com/xuri/excelize/v2"
)

type ExcelExport struct {
	List         []map[string]any
	Labels       []types.KeyVal
	File         *excelize.File
	StreamWriter *excelize.StreamWriter
	Path         string
	FileName     string
	Url          string
}

func (e *ExcelExport) New(isStream bool) error {
	if len(e.List) < 1 {
		return errors.New("无数据导出")
	}
	if e.Labels == nil || len(e.Labels) < 1 {
		return errors.New("未设置标题")
	}

	if err := e.setSavePath(); err != nil {
		return err
	}

	e.File = excelize.NewFile()
	if isStream {
		streamWriter, err := e.File.NewStreamWriter("Sheet1")
		if err != nil {
			return errors.New("创建流式写入器失败" + err.Error())
		}
		e.StreamWriter = streamWriter
	}
	return nil
}

func (e *ExcelExport) Export() error {
	err := e.steamSetHeader()
	if err != nil {
		return errors.New("标题行" + err.Error())
	}
	err1 := e.steamSetData()
	if err1 != nil {
		return errors.New("数据遍历失败" + err1.Error())
	}
	return e.steamSaveFile()
}

func (e *ExcelExport) setSavePath() error {
	times := time.Now()
	e.Path = "/static/export/excel/" + times.Format("2006") + "/" + times.Format("01")
	e.FileName = core.G_GENID.New().String() + ".xlsx"
	e.Url = e.Path + "/" + e.FileName

	if err := tool.DirCreate(e.Path); err != nil {
		return errors.New("保存目录创建失败" + err.Error())
	}
	return nil
}

// steamHeader 添加首行标题
func (e *ExcelExport) steamSetHeader() error {
	cols := len(e.Labels)
	i := 0
	headers := make([]any, cols)
	for _, item := range e.Labels {
		headers[i] = item.Value
		i++
	}
	return e.steamSaveRow(1, headers)
}

// steamSetData 保存数据
func (e *ExcelExport) steamSetData() error {
	rowID := 2
	for _, item := range e.List {
		i := 0
		row := make([]interface{}, len(e.Labels))
		for _, label := range e.Labels {
			row[i] = item[label.Key]
			i++
		}
		err := e.steamSaveRow(rowID, row)
		if err != nil {
			return err
		}
		rowID++
	}
	return nil
}

func (e *ExcelExport) steamSaveRow(rowID int, values []any) error {
	cell, err := excelize.CoordinatesToCellName(1, rowID)
	if err != nil {
		return errors.New("设置行错误：" + err.Error())
	}
	if err := e.StreamWriter.SetRow(cell, values); err != nil {
		return errors.New("保存行数据错误：" + err.Error())
	}
	return nil
}

func (e *ExcelExport) steamSaveFile() error {
	if err := e.StreamWriter.Flush(); err != nil {
		return errors.New("结束excel流失败" + err.Error())
	}
	if err := e.File.SaveAs("." + e.Url); err != nil {
		return errors.New("保存excel文件失败" + err.Error())
	}
	return nil
}
