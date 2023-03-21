package dict

import (
	"encoding/json"
	"fmt"
	"oiocns-standardAttribute-go/a/model"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gogf/gf/os/gfile"
)

func ExcelParse() []model.Dict {
	// 打开Excel文件
	// f, err := excelize.OpenFile("基础配置-业务字典.xlsx")
	f, err := excelize.OpenFile("基础配置-业务字典last2.xlsx")
	if err != nil {
		model.GglogFile.Info(err)
		return nil
	}

	// 读取数据并生成JSON
	rows := f.GetRows("Sheet1")
	var excelData []model.Dict
	for i, row := range rows {
		// 跳过表头
		if i == 0 {
			continue
		}

		excelData = append(excelData, model.Dict{
			DictName: row[1],
			DictCode: row[2],

			DictItemId:   row[6],
			DictItemName: row[7],
			DictItemCode: row[8],
		})

	}
	fmt.Println("excelData", excelData)
	fmt.Println("len(excelData)", len(excelData))

	return excelData
}

func DictParse() []model.Dict {
	var dict model.Dict
	var dictArr []model.Dict

	content := gfile.GetContents("./aCopy.json")
	// content := gfile.GetContents("./a.json")
	fmt.Printf("content: %v\n", content)
	var dri model.DictResponseInfo
	json.Unmarshal([]byte(content), &dri)
	dData := dri.Data
	fmt.Println("len(dData)", len(dData))

	for i, _ := range dData {
		dict.DictName = dData[i].DictValue //字典名称
		dict.DictCode = dData[i].Code

		fmt.Println("len(dData[i].Children)", len(dData[i].Children))
		if len(dData[i].Children) > 0 {
			for j, _ := range dData[i].Children {
				jj := strconv.Itoa(j + 1)
				dict.DictItemId = jj                               // 字典项序号
				dict.DictItemName = dData[i].Children[j].DictValue // 字典项名称
				dict.DictItemCode = dData[i].Code + jj             // // 字典项值,a.json的id 根据c.json的id获取字典代码

				dictArr = append(dictArr, dict)
			}
		}

	}
	return dictArr
}
func ExcelWrite(data []model.Dict) {
	// 创建一个新的excel文件
	f := excelize.NewFile()

	// 创建一个工作表
	sheetName := "Sheet1"
	index := f.NewSheet(sheetName)
	fmt.Println("index:", index)

	// 设置工作表的列名
	headers := []string{"序号", "字典名称", "字典代码", "字典制定组织", "字典向下级组织公开", "字典备注",
		"字典项序号", "字典项名称", "字典项值", "字典项指定组织", "字典项备注"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 填充测试数据
	for i, row := range data {
		cellDictName := fmt.Sprintf("%s%d", string(rune('B')), i+2)
		f.SetCellValue(sheetName, cellDictName, row.DictName)
		cellDictCode := fmt.Sprintf("%s%d", string(rune('C')), i+2)
		f.SetCellValue(sheetName, cellDictCode, row.DictCode)

		cellDictItemId := fmt.Sprintf("%s%d", string(rune('G')), i+2)
		f.SetCellValue(sheetName, cellDictItemId, row.DictItemId)
		cellDictItemName := fmt.Sprintf("%s%d", string(rune('H')), i+2)
		f.SetCellValue(sheetName, cellDictItemName, row.DictItemName)
		cellDictItemCode := fmt.Sprintf("%s%d", string(rune('I')), i+2)
		f.SetCellValue(sheetName, cellDictItemCode, row.DictItemCode)

	}

	// 保存文件
	err := f.SaveAs("test.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
