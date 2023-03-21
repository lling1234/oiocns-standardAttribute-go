package model

import "github.com/gogf/gf/os/glog"

var PD PayloadData 
var GglogFile *glog.Logger

type PayloadData struct {
	Token       string `json:"token"`
	FileName    string `json:"fileName"`
	ExcelSheet  string `json:"excelSheet"`
	SpeciesCode string `json:"speciesCode"`
	SpeciesID   string `json:"speciesId"`
	BelongID    string `json:"belongId"`
	AuthID      string `json:"authId"`
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
