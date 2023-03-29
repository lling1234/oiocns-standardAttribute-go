package excel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"oiocns-standardAttribute-go/a/model"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func OpenExcelFile(p model.PayloadData) {
	// 打开Excel文件
	f, err := excelize.OpenFile(p.FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取数据并生成JSON
	rows := f.GetRows(p.ExcelSheet)
	var dateStruct []model.StandardAttributeReq
	for i, row := range rows {
		// 跳过表头
		if i == 0 || i == 1 {
			continue
		}

		dateStruct = append(dateStruct, model.StandardAttributeReq{
			// SequenceNumber: row[1],
			SpeciesID:   p.SpeciesID,
			SpeciesCode: p.SpeciesCode,
			Public:      true,
			ValueType:   row[5],
			Name:        row[3],
			Code:        row[2],
			BelongID:    p.BelongID,
			AuthID:      p.AuthID,
			Remark:      row[8],
		})

	}
	fmt.Println("dateStruct", dateStruct)
	jsonData, err := json.Marshal(dateStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("jsonData", jsonData)
	var jsonInfo []model.StandardAttributeReq
	err2 := json.Unmarshal(jsonData, &jsonInfo)
	if err2 != nil {
		fmt.Println("json解析错误", err2)
		return
	}
	fmt.Println("jsonInfo[0].Name", jsonInfo[0].Name)
	fmt.Println("len(jsonInfo)", len(jsonInfo))

	HttpPostReqLoop(jsonInfo,p.Token)

}
func HttpPostReqLoop(jsonInfo []model.StandardAttributeReq,token string) {
	fmt.Println("-------------HttpPostRequest start---------------")
	for i := 0; i < len(jsonInfo); i++ {

		jsonData1 := model.StandardAttribute{Module: "thing", Action: "CreateAttribute", Params: jsonInfo[i]}
		jsonDataBody, err := json.Marshal(jsonData1)
		if err != nil {
			fmt.Println("json.Marshal err", err)
			return
		}
		HttpPostRequest(jsonDataBody,token)
		time.Sleep(1000)
	}
	fmt.Println("-------------HttpPostRequest end---------------")
}

func HttpPostRequest(jsonData []byte,token string) {
	fmt.Println("HttpPostRequest start")
	var jsonInfo model.StandardAttribute
	err2 := json.Unmarshal(jsonData, &jsonInfo)
	if err2 != nil {
		fmt.Println("json解析错误", err2)
		return
	}
	fmt.Println("jsonInfo", jsonInfo)
	// 发送HTTP POST请求
	req, err := http.NewRequest("POST", "http://anyinone.com:888/orginone/kernel/rest/request", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	time.Sleep(1000)

	// 输出响应结果
	fmt.Println(resp.Status)
	// time.Sleep(1000);
}
