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

var tmpDictID string = ""

func PostBlp(data []model.Dict) {
	fmt.Println("PostBlp data", data)
	// var tmpDictCOde string = ""
	var postDict model.PostDict

	for i, _ := range data {
		fmt.Println("for i",i)

		// 第一行，先插入分类，在插入分类项
		if i == 0 {
			// 先插入分类
			// inster dict start
			postDict.Module = "thing"
			postDict.Action = "CreateDict"

			postDict.Params.Name = data[i].DictName
			postDict.Params.Code = data[i].DictCode
			postDict.Params.BelongID = "380663455457349633"
			postDict.Params.Remark = data[i].DictName + " " + data[i].DictCode
			postDict.Params.SpeciesID = "425706835434147840"
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
			// inster dict end

			time.Sleep(2000)
			// 在插入分类项
			// insert dictItem start
			// 添加字典分类项
			var postDictItem model.PostDictItem
			postDictItem.Module = "thing"
			postDictItem.Action = "CreateDictItem"

			postDictItem.Params.Name = data[i].DictItemName
			postDictItem.Params.Value = data[i].DictItemCode
			postDictItem.Params.BelongID = model.DPD.BelongID
			postDictItem.Params.Remark = data[i].DictItemName + " " + data[i].DictItemCode
			postDictItem.Params.DictID = tmpDictID

			pdi, err := json.Marshal(postDictItem)
			if err != nil {
				model.GglogFile.Info(err)
				return
			}
			fmt.Println("pdi", pdi)

			PostReqDict(pdi)
			// insert dictItem end
			time.Sleep(2000)

			continue

		}
		// 第二行，判断data[0].DictCode是否等于data[1].DictCode，相等就插入分类项，不相等先插入分类，在插入分类项
		if data[i+1].DictCode == data[i].DictCode {
			// 相等就插入分类项
			// insert dict start
			// 添加字典分类项
			var postDictItem model.PostDictItem
			postDictItem.Module = "thing"
			postDictItem.Action = "CreateDictItem"

			postDictItem.Params.Name = data[i].DictItemName
			postDictItem.Params.Value = data[i].DictItemCode
			postDictItem.Params.BelongID = model.DPD.BelongID
			postDictItem.Params.Remark = data[i].DictItemName + " " + data[i].DictItemCode
			postDictItem.Params.DictID = tmpDictID

			pdi, err := json.Marshal(postDictItem)
			if err != nil {
				model.GglogFile.Info(err)
				return
			}
			fmt.Println("pdi", pdi)

			PostReqDict(pdi)
			// insert dict end
			continue

		} else {
			// 不相等先插入分类，在插入分类项
			// 先插入分类
			// inster dict start
			postDict.Module = "thing"
			postDict.Action = "CreateDict"

			postDict.Params.Name = data[i].DictName
			postDict.Params.Code = data[i].DictCode
			postDict.Params.BelongID = "380663455457349633"
			postDict.Params.Remark = data[i].DictName + " " + data[i].DictCode
			postDict.Params.SpeciesID = "425706835434147840"
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
			// inster dict end

			// 再插入分类项
			// insert dict start
			// 添加字典分类项
			var postDictItem model.PostDictItem
			postDictItem.Module = "thing"
			postDictItem.Action = "CreateDictItem"

			postDictItem.Params.Name = data[i].DictItemName
			postDictItem.Params.Value = data[i].DictItemCode
			postDictItem.Params.BelongID = model.DPD.BelongID
			postDictItem.Params.Remark = data[i].DictItemName + " " + data[i].DictItemCode
			postDictItem.Params.DictID = tmpDictID

			pdi, err := json.Marshal(postDictItem)
			if err != nil {
				model.GglogFile.Info(err)
				return
			}
			fmt.Println("pdi", pdi)

			PostReqDict(pdi)
			// insert dict end
			continue

		}
		// ----------------------------------------------

	}

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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM1ODYyNjU3ODYxNzQ3MDk3NiIsImNvbXBhbnlJZCI6IjAiLCJleHAiOjE2NzkzMTM4NzQsImlhdCI6MTY3OTMwNjY3NCwidXNlcklkIjoiMzU4NjI2NTc4NjE3NDcwOTc2In0.3y_y9rUGLTBueD5aSO26-pmmZ8GsZ5VZiUkejFdu0rg"
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// model.GglogFile.Info(err)
		fmt.Println("err11", err)

	}
	var result1 model.DictResponseInfoI
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result1)
	}
	fmt.Println("result1", result1)
	fmt.Println("result1.Data.ID", result1.Data.ID)
	tmpDictID = result1.Data.ID

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
