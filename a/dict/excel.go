package dict

import (
	"fmt"
	"oiocns-standardAttribute-go/a/model"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ExcelWrite2(data [][]model.Dict) {
	// 创建一个新的excel文件
	f := excelize.NewFile()

	// 创建一个工作表
	sheetName := "Sheet1"
	index := f.NewSheet(sheetName)
	fmt.Println("index:", index)

	// 设置工作表的列名
	headers := []string{"序号", "字典名称", "字典代码", "字典制定组织", "字典向下级组织公开", "字典备注", "字典项序号", "字典项名称", "字典项值", "字典项指定组织", "字典项备注"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 填充测试数据
	for i, row := range data {
		for j, value := range row {
			cell := fmt.Sprintf("%s%d", string(rune('A'+j)), i+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 保存文件
	err := f.SaveAs("test.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
