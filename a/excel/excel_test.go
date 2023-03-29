package excel

import (
	"oiocns-standardAttribute-go/a/model"
	"testing"
)

func TestOpenExcelFile(t *testing.T) {
	type args struct {
		p model.PayloadData
	}
	tests := []struct {
		name string
		args args
	}{
		{"openexcel", args{model.PayloadData{
			FileName:    "../files/1.xlsx",
			Token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImNvbXBhbnlJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImV4cCI6MTY4MDA4ODk3MCwiaWF0IjoxNjgwMDgxNzcwLCJ1c2VySWQiOiIzNTg2MjY1Nzg2MTc0NzA5NzYifQ.1Km4zyWkEFox4Q-gmrq32cKYxDlxxg2h2KGiuyUNIkQ",
			ExcelSheet:  "登记管理-资产登记-房产登记",
			SpeciesID:   "428971093794099200",
			SpeciesCode: "propertyBuildingPropertyRegistration",
			BelongID:    "380663455457349633",
			AuthID:      "361356410044420096",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenExcelFile(tt.args.p)
		})
	}
}
