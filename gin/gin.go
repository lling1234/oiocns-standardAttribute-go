package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // 接收上传的 .xlsx 文件
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.String(http.StatusBadRequest, fmt.Sprintf("上传文件错误：%s", err.Error()))
            return
        }

        // 检查文件类型
        if file.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
            c.String(http.StatusBadRequest, "上传文件必须是 .xlsx 类型")
            return
        }

        // 将文件保存到本地
        err = c.SaveUploadedFile(file, "./files/"+file.Filename)
        if err != nil {
            c.String(http.StatusBadRequest, fmt.Sprintf("保存文件错误：%s", err.Error()))
            return
        }

        c.String(http.StatusOK, fmt.Sprintf("上传成功，文件名：%s", file.Filename))
    })

    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
