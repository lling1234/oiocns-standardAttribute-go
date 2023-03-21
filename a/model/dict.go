package model

var DPD DictPayloadData = DictPayloadData{
	Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImNvbXBhbnlJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImV4cCI6MTY3OTMwNjI1MiwiaWF0IjoxNjc5Mjk5MDUyLCJ1c2VySWQiOiIzNTg2MjY1Nzg2MTc0NzA5NzYifQ.S06p40d5G_CTBiOHAm2oeJQKDGTW6Rx37MsmxA8Djng",
	BelongID:  "380663455457349633",
	SpeciesID: "425658752432214016",
}

type DictPayloadData struct {
	Token     string `json:"token"`
	BelongID  string `json:"belongId"`
	SpeciesID string `json:"speciesId"`
}
type DictResponseInfoI struct {
	Code    int          `json:"code"`
	Success bool         `json:"success"`
	Data    DictRespData `json:"data"`
	Msg     string       `json:"msg"`
}

type DictRespData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Remark     string `json:"remark"`
	Public     bool   `json:"public"`
	SpeciesID  string `json:"speciesId"`
	BelongID   string `json:"belongId"`
	Status     int    `json:"status"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
	Version    string `json:"version"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type DictResponseInfo struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Data    []Data `json:"data"`
	Msg     string `json:"msg"`
}

type Data struct {
	ID          string     `json:"id"`
	TenantID    string     `json:"tenantId"`
	BizAppID    string     `json:"bizAppId"`
	ParentID    string     `json:"parentId"`
	Code        string     `json:"code"`
	DictKey     string     `json:"dictKey"`
	DictValue   string     `json:"dictValue"`
	Sort        int        `json:"sort"`
	Remark      string     `json:"remark"`
	IsSealed    int        `json:"isSealed"`
	IsRequired  int        `json:"isRequired"`
	IsDeleted   int        `json:"isDeleted"`
	Children    []Children `json:"children"`
	ParentName  string     `json:"parentName"`
	BizAppCode  string     `json:"bizAppCode"`
	HasChildren bool       `json:"hasChildren"`
}
type Children struct {
	ID          string `json:"id"`
	TenantID    string `json:"tenantId"`
	BizAppID    string `json:"bizAppId"`
	ParentID    string `json:"parentId"`
	Code        string `json:"code"`
	DictKey     string `json:"dictKey"`
	DictValue   string `json:"dictValue"`
	Sort        int    `json:"sort"`
	Remark      string `json:"remark"`
	IsSealed    int    `json:"isSealed"`
	IsRequired  int    `json:"isRequired"`
	IsDeleted   int    `json:"isDeleted"`
	ParentName  string `json:"parentName"`
	BizAppCode  string `json:"bizAppCode"`
	HasChildren bool   `json:"hasChildren"`
}

type Dict struct {
	DictName  string `json:"dict`
	DictCode  string `json:"dict`
	BelongID  string `json:"belongId"`
	Remark    string `json:"remark"`
	SpeciesID string `json:"speciesId"`

	DictItemId       string `json:"dict`
	DictItemName     string `json:"dict`
	DictItemCode     string `json:"dict`
	DictItemBelongID string `json:"belongId"`
	DictItemRemark   string `json:"remark"`
}

type PostDict struct {
	Module string         `json:"module"`
	Action string         `json:"action"`
	Params PostDictParams `json:"params"`
}
type PostDictParams struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	BelongID  string `json:"belongId"`
	Remark    string `json:"remark"`
	Public    bool   `json:"public"`
	SpeciesID string `json:"speciesId"`
}

type PostDictItem struct {
	Module string             `json:"module"`
	Action string             `json:"action"`
	Params PostDictItemParams `json:"params"`
}
type PostDictItemParams struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	BelongID string `json:"belongId"`
	Remark   string `json:"remark"`
	DictID   string `json:"dictId"`
}
