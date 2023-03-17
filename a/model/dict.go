package model

type DictResponseInfo struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Data    []Data `json:"data"`
	Msg     string `json:"msg"`
}
type Children struct {
	ID          string `json:"id"`
	ParentID    string `json:"parentId"`
	HasChildren bool   `json:"hasChildren"`
	Title       string `json:"title"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	OrgID       string `json:"orgId"`
}
type Data struct {
	ID          string     `json:"id"`
	ParentID    string     `json:"parentId"`
	Children    []Children `json:"children"`
	HasChildren bool       `json:"hasChildren"`
	Title       string     `json:"title"`
	Key         string     `json:"key"`
	Value       string     `json:"value"`
	OrgID       string     `json:"orgId"`
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
