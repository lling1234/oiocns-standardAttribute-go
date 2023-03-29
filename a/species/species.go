package species

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oiocns-standardAttribute-go/a/model"
	"time"
)
// 数据量不超过十条，直接通过post请求
/*
1.查询奥集能标准，获取到标准分类id,写入到表格第D3和D4,插入响应的数据写入到D5和D10。
2.读取excel数据存放到map[string,string] （map[分类代码,分类名称]） 里面
3.map数据通过post请求写入到平台
*/

// QueryStandardId 查询奥集能平台分类标准获取用户、财物、事项、商品的id
func QueryStandardId(token string) {
	// QuerySpeciesTree post请求获取响应

	sq := model.SpeciesReq{}
	sq.Module = "thing"
	sq.Action = "QuerySpeciesTree"
	sq.Params.ID = "380663455457349633"

	sqData, err := json.Marshal(sq)
	if err != nil {
		fmt.Println("sqData err", err)
		return
	}

	PostReq(sqData, token)

	// 解析json-data-nodes-id数据
	// excel文件写入数据
}

// ExcelFileParse 解析excel文件，获取map数据
func ExcelFileParse() {

}

// ExcelFileWrite excel文件写入数据
func ExcelFileWrite() {

}

func PostReq(data []byte, token string) (model.SpeciesResp, error) {
	fmt.Println("PostReq entering")
	// 发送HTTP POST请求
	var url string = "http://anyinone.com:888/orginone/kernel/rest/request"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		model.GglogFile.Info(err)
		return model.SpeciesResp{}, err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// model.GglogFile.Info(err)
		fmt.Println("err11", err)
		return model.SpeciesResp{}, err
	}
	time.Sleep(2000)
	var result1 model.SpeciesResp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err111fff111", err)
		return model.SpeciesResp{}, err
	}
	fmt.Println("body", string(body))
	err = json.Unmarshal(body, &result1)
	if err != nil {
		fmt.Println("err111fff222", err)
		return model.SpeciesResp{}, err
	}
	fmt.Println("result1", result1)
	fmt.Println("result1.Data.ID", result1.Data.ID)
	fmt.Println("len(result1.Data.Nodes)", len(result1.Data.Nodes))
	for _, v := range result1.Data.Nodes {
		fmt.Println("v.ID", v.ID)
		fmt.Println("v.Name", v.Name)

	}

	time.Sleep(2000)

	defer resp.Body.Close()

	// 输出响应结果
	// model.GglogFile.Info(resp.Status)
	return result1, err

}
func ParseResponse(response *http.Response) (*model.DictResponseInfoI, error) {
	var result *model.DictResponseInfoI
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}

// ExcelDataPostReq 解析excel的数据，通过post请求到平台
func ExcelDataPostReq() {

}

// PostRespBody 获取请求的响应数据
func PostRespBody() {

}
