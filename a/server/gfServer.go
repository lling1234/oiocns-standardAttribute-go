package server

import (
	"oiocns-standardAttribute-go/a/excel"
	"oiocns-standardAttribute-go/a/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func GFServer() {

	s := g.Server()

	model.GglogFile = g.Log("oiocns-standardAttribute-go")

	model.GglogFile.SetLevel(glog.LEVEL_INFO)
	model.GglogFile.SetPath("log")
	model.GglogFile.Info("1111")
	s.SetPort(8081) // 设置端口号为8080
	s.BindHandler("/upload", func(r *ghttp.Request) {
		if r.Method == "POST" {
			file := r.GetUploadFile("file")
			if file != nil {
				// defer file.Close()
				glog.Info("file.Filename", file.Filename)
				ext := gfile.ExtName(file.Filename)
				glog.Info("ext", ext)
				if ext == "xlsx" {
					// 将上传的 Excel 文件保存到本地
					filename, err := file.Save("./files", false)
					if err != nil {
						r.Response.Write("保存文件失败")
					} else {
						r.Response.Write("文件保存成功,文件名：" + filename + "正在解析excel导入数据中")

						model.PD.FileName = "./files/" + filename
						glog.Info("model.PD", model.PD)

						go excel.OpenExcelFile(model.PD)
					}
				} else {
					r.Response.Write("只支持上传 .xlsx 文件")
				}
			} else {
				r.Response.Write("未上传文件")
			}
		} else {
			r.Response.Write("请使用 POST 方法上传文件")
		}
	})
	s.Run()
}
