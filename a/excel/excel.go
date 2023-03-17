package excel

import (
	"bytes"
	"encoding/json"
	"net/http"
	"oiocns-standardAttribute-go/a/model"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func OpenExcelFile(p model.PayloadData) {
	// 打开Excel文件
	f, err := excelize.OpenFile(p.FileName)
	if err != nil {
		model.GglogFile.Info(err)
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
	model.GglogFile.Info("dateStruct", dateStruct)
	jsonData, err := json.Marshal(dateStruct)
	if err != nil {
		model.GglogFile.Info(err)
		return
	}
	model.GglogFile.Info("jsonData", jsonData)
	var jsonInfo []model.StandardAttributeReq
	err2 := json.Unmarshal(jsonData, &jsonInfo)
	if err2 != nil {
		model.GglogFile.Info("json解析错误", err2)
		return
	}
	model.GglogFile.Info("jsonInfo[0].Name", jsonInfo[0].Name)
	model.GglogFile.Info("len(jsonInfo)", len(jsonInfo))

	HttpPostReqLoop(jsonInfo)

}
func HttpPostReqLoop(jsonInfo []model.StandardAttributeReq) {
	model.GglogFile.Info("-------------HttpPostRequest start---------------")
	for i := 0; i < len(jsonInfo); i++ {

		jsonData1 := model.StandardAttribute{Module: "thing", Action: "CreateAttribute", Params: jsonInfo[i]}
		jsonDataBody, err := json.Marshal(jsonData1)
		if err != nil {
			model.GglogFile.Info("json.Marshal err", err)
			return
		}
		HttpPostRequest(jsonDataBody)
	}
	model.GglogFile.Info("-------------HttpPostRequest end---------------")
}

func HttpPostRequest(jsonData []byte) {
	model.GglogFile.Info("HttpPostRequest start")
	var jsonInfo model.StandardAttribute
	err2 := json.Unmarshal(jsonData, &jsonInfo)
	if err2 != nil {
		model.GglogFile.Info("json解析错误", err2)
		return
	}
	model.GglogFile.Info("jsonInfo", jsonInfo)
	// 发送HTTP POST请求
	req, err := http.NewRequest("POST", "http://anyinone.com:800/orginone/kernel/rest/request", bytes.NewBuffer(jsonData))
	if err != nil {
		model.GglogFile.Info(err)
		return
	}
	req.Header.Set("Authorization", model.PD.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		model.GglogFile.Info(err)
		return
	}
	defer resp.Body.Close()

	// 输出响应结果
	model.GglogFile.Info(resp.Status)
	// time.Sleep(1000);
}
