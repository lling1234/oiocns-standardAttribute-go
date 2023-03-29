package model

type SpeciesReq struct {
	Module string `json:"module"`
	Action string `json:"action"`
	Params Params `json:"params"`
}
type Page struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Filter string `json:"filter"`
}
type Params struct {
	ID   string `json:"id"`
	Page Page   `json:"page"`
}

type SpeciesResp struct {
	Code    int         `json:"code"`
	Data    SpeciesData `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

type SpeciesData struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Code       string         `json:"code"`
	Remark     string         `json:"remark"`
	Public     bool           `json:"public"`
	Status     int            `json:"status"`
	CreateUser string         `json:"createUser"`
	UpdateUser string         `json:"updateUser"`
	Version    string         `json:"version"`
	CreateTime string         `json:"createTime"`
	UpdateTime string         `json:"updateTime"`
	Nodes      []SpeciesNodes `json:"nodes"`
}
type SpeciesNodes struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	Remark     string  `json:"remark"`
	Public     bool    `json:"public"`
	ParentID   string  `json:"parentId"`
	Status     int     `json:"status"`
	CreateUser string  `json:"createUser"`
	UpdateUser string  `json:"updateUser"`
	Version    string  `json:"version"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
	Nodes      []Nodes `json:"nodes,omitempty"`
}
type Belong struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	TypeName   string `json:"typeName"`
	BelongID   string `json:"belongId"`
	ThingID    string `json:"thingId"`
	Status     int    `json:"status"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
	Version    string `json:"version"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
type Nodes struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Remark     string `json:"remark"`
	Public     bool   `json:"public"`
	ParentID   string `json:"parentId"`
	Status     int    `json:"status"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
	Version    string `json:"version"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	BelongID   string `json:"belongId,omitempty"`
	AuthID     string `json:"authId,omitempty"`
	Belong     Belong `json:"belong,omitempty"`
}
