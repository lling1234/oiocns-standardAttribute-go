package model

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

	DictItemId       string    `json:"dict`
	DictItemName     string `json:"dict`
	DictItemCode     string `json:"dict`
	DictItemBelongID string `json:"belongId"`
	DictItemRemark   string `json:"remark"`
}
