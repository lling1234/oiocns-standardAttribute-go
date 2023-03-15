package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImNvbXBhbnlJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImV4cCI6MTY3ODc5MDY5OCwiaWF0IjoxNjc4NzgzNDk4LCJ1c2VySWQiOiIzNTg2MjY1Nzg2MTc0NzA5NzYifQ.xTJMi2s_k-lb0NBaxK6XDmCsS_JqC-1QphET-ZXebRk"
const sheet = "登记管理-资产登记-房产登记"
const SpeciesCodeData = "buildingPropertyRegistration"

type PayloadData struct {
	Token           string `json:"token"`
	FileName        string `json:"fileName"`
	ExcelSheet      string `json:"excelSheet"`
	SpeciesCodeData string `json:"speciesCodeData"`
}

type StandardAttribute struct {
	Module string               `json:"module"`
	Action string               `json:"action"`
	Params StandardAttributeReq `json:"params"`
}

type StandardAttributeReq struct {
	// SequenceNumber string `json:"sequenceNumber"`
	SpeciesID   string `json:"speciesId"`
	SpeciesCode string `json:"speciesCode"`
	Public      bool   `json:"public"`
	ValueType   string `json:"valueType"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	BelongID    string `json:"belongId"`
	AuthID      string `json:"authId"`
	Remark      string `json:"remark"`
}

type StandardAttributeResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	ValueType  string `json:"valueType"`
	Public     bool   `json:"public"`
	Remark     string `json:"remark"`
	SpeciesID  string `json:"speciesId"`
	BelongID   string `json:"belongId"`
	AuthID     string `json:"authId"`
	Status     int    `json:"status"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
	Version    string `json:"version"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func main() {

}

func OpenExcelFile(p PayloadData) {
	// 打开Excel文件
	f, err := excelize.OpenFile(p.FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取数据并生成JSON
	rows := f.GetRows(p.ExcelSheet)
	var data []map[string]string
	var dateStruct []StandardAttributeReq
	for i, row := range rows {
		// 跳过表头
		if i == 0 || i == 1 {
			continue
		}
		fmt.Println("i", i)
		fmt.Println("row", row)
		fmt.Println("------------")
		fmt.Println("row[1]", row[1])
		type1 := reflect.TypeOf(row[1])
		fmt.Println("type", type1)
		id, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Println("转换失败：", err)
			return
		}
		fmt.Println("id2", id)
		fmt.Println("序号", id)
		fmt.Println("特性编号", row[2])
		fmt.Println("特性名称", row[3])
		fmt.Println("特性分类", row[4])
		fmt.Println("特性类型", row[5])
		fmt.Println("特性类型", row[8])

		dateStruct = append(dateStruct, StandardAttributeReq{
			// SequenceNumber: row[1],
			SpeciesID:   "423521602211287040",
			SpeciesCode: SpeciesCodeData,
			Public:      true,
			ValueType:   row[5],
			Name:        row[3],
			Code:        row[2],
			BelongID:    "380663455457349633",
			AuthID:      "361356410044420096",
			Remark:      row[8],
		})

		// 解析每行数据
		item := make(map[string]string)
		item["sequenceNumber"] = row[2] //序号
		item["name"] = row[3]           //特性编号
		item["code"] = row[4]           //特性名称
		item["classify"] = row[5]       //特性分类
		item["valueType"] = row[6]      //特性类型
		item["belongId"] = row[7]       //选择字典
		item["belongId"] = row[8]       //共享组织
		item["remark"] = row[9]         //特性定义
		data = append(data, item)
	}
	fmt.Println("data", data)
	fmt.Println("dateStruct", dateStruct)
	jsonData, err := json.Marshal(dateStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("jsonData", jsonData)
	var jsonInfo []StandardAttributeReq
	err2 := json.Unmarshal(jsonData, &jsonInfo)
	if err2 != nil {
		fmt.Println("json解析错误", err2)
		return
	}
	fmt.Println("jsonInfo[0].Name", jsonInfo[0].Name)
	fmt.Println("len(jsonInfo)", len(jsonInfo))
	fmt.Println("-------------HttpPostRequest start---------------")
	for i := 0; i < len(jsonInfo); i++ {
		// jsonData, err := json.Marshal(jsonInfo[i])
		// if err != nil {
		// 	fmt.Println("json.Marshal err", err)
		// 	return
		// }

		jsonData1 := StandardAttribute{Module: "thing", Action: "CreateAttribute", Params: jsonInfo[i]}
		jsonDataBody, err := json.Marshal(jsonData1)
		if err != nil {
			fmt.Println("json.Marshal err", err)
			return
		}
		HttpPostRequest(jsonDataBody)
	}
	fmt.Println("-------------HttpPostRequest end---------------")
}

func HttpPostRequest(jsonData []byte) {
	fmt.Println("HttpPostRequest start")
	var jsonInfo StandardAttribute
	err2 := json.Unmarshal(jsonData, &jsonInfo)
	if err2 != nil {
		fmt.Println("json解析错误", err2)
		return
	}
	fmt.Println("jsonInfo", jsonInfo)
	// 发送HTTP POST请求
	req, err := http.NewRequest("POST", "http://anyinone.com:800/orginone/kernel/rest/request", bytes.NewBuffer(jsonData))
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

	// 输出响应结果
	fmt.Println(resp.Status)
	// time.Sleep(1000);
}
