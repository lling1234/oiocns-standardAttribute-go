package main

import "oiocns-standardAttribute-go/a/model"
import "oiocns-standardAttribute-go/a/server"

func main() {

	model.PD.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImNvbXBhbnlJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImV4cCI6MTY4MDAwMTk5NywiaWF0IjoxNjc5OTk0Nzk3LCJ1c2VySWQiOiIzNTg2MjY1Nzg2MTc0NzA5NzYifQ.SB1ONWutSvTRZO_IPg-Ldw2GqEAZ7VaAAb5hYIsNIRE"
	// pd.FileName = ""
	model.PD.ExcelSheet = "登记管理-资产登记-房产登记"
	model.PD.SpeciesID = "428607427370422272"
	model.PD.SpeciesCode = "propertyBuildingPropertyRegistration"
	model.PD.BelongID = "380663455457349633"
	model.PD.AuthID = "361356410044420096"

	 server.GFServer()

}
