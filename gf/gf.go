package main

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

func main() {
	s := g.Server()
	s.SetPort(8080) // 设置端口号为8080
	s.BindHandler("/upload", func(r *ghttp.Request) {
		if r.Method == "POST" {
			file := r.GetUploadFile("file")
			if file != nil {
				// defer file.Close()
				fmt.Println("file.Filename", file.Filename)
				ext := gfile.ExtName(file.Filename)
				fmt.Println("ext",ext)
				if ext == "xlsx" {
					// 将上传的 Excel 文件保存到本地
					filename, err := file.Save("./files", false)
					if err != nil {
						r.Response.Write("保存文件失败")
					} else {
						r.Response.Write("文件保存成功,文件名：" + filename)
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
