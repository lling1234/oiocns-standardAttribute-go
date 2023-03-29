package dict

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oiocns-standardAttribute-go/a/model"
	"time"
)

var token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImNvbXBhbnlJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImV4cCI6MTY3OTkxNTc2MSwiaWF0IjoxNjc5OTA4NTYxLCJ1c2VySWQiOiIzNTg2MjY1Nzg2MTc0NzA5NzYifQ.BsmT7czhCoGWHiamxkj84a-14wocojy4XWuqFdEAm5g"
var tmpDictID string = ""

func PostBlp(data []model.Dict) {
	fmt.Println("PostBlp data", data)
	var currentDictCode string = ""
	// var postDict model.PostDict

	for i, _ := range data {
		fmt.Println("for i----", i)

		if currentDictCode != data[i].DictCode {
			fmt.Println("1111 start")
			InsertDIct(data[i].DictName, data[i].DictCode)
			currentDictCode = data[i].DictCode
			time.Sleep(2000)
			fmt.Println("tmpDictID", tmpDictID)
			fmt.Println("1111 end")
		} else {
			fmt.Println("2222 start")
			fmt.Println("tmpDictID", tmpDictID)

			InsertDIctItem(data[i].DictItemName, data[i].DictItemCode)
			time.Sleep(2000)
			fmt.Println("2222 end")
			continue
		}
		fmt.Println("3333 start")
		fmt.Println("tmpDictID", tmpDictID)

		InsertDIctItem(data[i].DictItemName, data[i].DictItemCode)
		time.Sleep(2000)
		fmt.Println("3333 end")

		continue
	}

}
func InsertDIct(dictName, dictCode string) {
	var postDict model.PostDict

	postDict.Module = "thing"
	postDict.Action = "CreateDict"

	postDict.Params.Name = dictName
	postDict.Params.Code = dictCode
	postDict.Params.BelongID = "380663455457349633"
	postDict.Params.Remark = dictName + " " + dictCode
	postDict.Params.SpeciesID = "428241051162120192"
	postDict.Params.Public = true

	pd, err := json.Marshal(postDict)
	fmt.Println("err111", err)
	if err != nil {
		model.GglogFile.Info(err)
		return
	}
	fmt.Println("PostReqDict start---------------")
	PostReqDict(pd)
	fmt.Println("PostReqDict end---------------")
}
func InsertDIctItem(dictItemName, dictItemCode string) {
	// insert dictItem start
	// 添加字典分类项
	var postDictItem model.PostDictItem
	postDictItem.Module = "thing"
	postDictItem.Action = "CreateDictItem"

	postDictItem.Params.Name = dictItemName
	postDictItem.Params.Value = dictItemCode
	postDictItem.Params.BelongID = "380663455457349633"
	postDictItem.Params.Remark = dictItemName + " " + dictItemCode
	postDictItem.Params.DictID = tmpDictID

	pdi, err := json.Marshal(postDictItem)
	if err != nil {
		model.GglogFile.Info(err)
		return
	}
	// fmt.Println("pdi", pdi)

	PostReqDictItem(pdi)
	// insert dictItem end
}
func PostReqDict(data []byte) {
	fmt.Println("PostReqDict entering")
	// 发送HTTP POST请求
	var url string = "http://anyinone.com:800/orginone/kernel/rest/request"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		model.GglogFile.Info(err)
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// model.GglogFile.Info(err)
		fmt.Println("err11", err)

	}
	time.Sleep(2000)
	var result1 model.DictResponseInfoI
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println("err111fff111", err)
	}
	err = json.Unmarshal(body, &result1)
	if err == nil {
		fmt.Println("err111fff222", err)
	}
	fmt.Println("result1", result1)
	fmt.Println("result1.Data.ID", result1.Data.ID)
	tmpDictID = result1.Data.ID
	time.Sleep(2000)

	defer resp.Body.Close()

	// 输出响应结果
	// model.GglogFile.Info(resp.Status)
}

func PostReqDictItem(data []byte) {
	fmt.Println("PostReqDict entering")
	// 发送HTTP POST请求
	var url string = "http://anyinone.com:800/orginone/kernel/rest/request"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		model.GglogFile.Info(err)
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// model.GglogFile.Info(err)
		fmt.Println("err11", err)

	}
	time.Sleep(2000)
	var result1 model.DictResponseInfoI
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println("err111fff111", err)
	}
	err = json.Unmarshal(body, &result1)
	if err == nil {
		fmt.Println("err111fff222", err)
	}
	fmt.Println("result1", result1)
	fmt.Println("result1.Data.ID", result1.Data.ID)
	time.Sleep(2000)

	defer resp.Body.Close()

	// 输出响应结果
	// model.GglogFile.Info(resp.Status)
}
func ParseResponse(response *http.Response) (*model.DictResponseInfoI, error) {
	var result *model.DictResponseInfoI
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}
