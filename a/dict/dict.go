package dict

import (
	"encoding/json"
	"fmt"
	"oiocns-standardAttribute-go/a/model"

	"github.com/gogf/gf/os/gfile"
)

func DictParse() model.DictResponseInfo {
	var d, di string
	content := gfile.GetContents("./c.json")
	fmt.Printf("content: %v\n", content)
	var dri model.DictResponseInfo
	json.Unmarshal([]byte(content), &dri)
	dData := dri.Data
	fmt.Println("len(dData)", len(dData))
	for i, _ := range dData {
		d = dData[i].Title    //字典名称
		d = dData[i].ParentID //字典代码，a.json的id 根据c.json的id获取字典代码
		fmt.Println("s", d)

		fmt.Println("len(dData[i].Children)", len(dData[i].Children))
		for j, _ := range dData[i].Children {
			// 序号
			di = dData[i].Children[j].Title // 字典项名称
			di = dData[i].Children[j].ID // // 字典项值,a.json的id 根据c.json的id获取字典代码

		}
		fmt.Println("di", di)

	}
	return dri
}
