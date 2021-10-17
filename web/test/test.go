package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.POST("/upload", UploadImage)
	router.Run("localhost:8005")
}

//单张图片上传
func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	name := c.PostForm("user_id")

	//filename := file.Filename
	filename := name + ".png"
	if err := c.SaveUploadedFile(file, "/Users/zh/ImageServer/"+filename); err != nil {
		//自己完成信息提示
		return
	}
	c.String(200, "Success")
}
