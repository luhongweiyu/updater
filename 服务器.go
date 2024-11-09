package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func 服务器文件夹(c *gin.Context) string {
	s := c.Param("服务器文件夹")
	s = regexp.MustCompile(`\w+`).FindString(s)
	if len(s) < 2 {
		return "空文件夹"
	}
	return s
}
func uploadFile(c *gin.Context) {
	路径 := 服务器文件夹(c)
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	// 上传文件到服务器代码
	filePath := c.PostForm("filePath")
	fmt.Println(filePath)

	dst := "./" + 路径 + "/" + filePath //file.Filename
	fmt.Println("保存路径:", dst)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))
}

func getFileList(c *gin.Context) {
	// 用户端获取文件列表
	路径 := 服务器文件夹(c)
	context, _ := ioutil.ReadFile("./" + 路径 + "/fileList.json")
	c.String(http.StatusOK, string(context))
}

func getFileListAdmin(c *gin.Context) {
	路径 := 服务器文件夹(c)
	a := 获取本地文件(路径)
	b := map[string]string{}
	for k, v := range a {
		k = strings.Replace(k, 路径+"\\", "", 1)
		k = strings.Replace(k, 路径+"/", "", 1)
		b[k] = v
	}
	c.JSON(http.StatusOK, b)
}
func getFile(c *gin.Context) {
	路径 := 服务器文件夹(c)
	// a := map[string]string{}
	// c.ShouldBindJSON(&a)

	filePath := "./" + 路径 + "/" + c.PostForm("fileName") //a["fileName"]
	fmt.Println(filePath)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filePath))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filePath)

}
func 密码验证(c *gin.Context) {
	if c.PostForm("password") != 全局_密码 {
		fmt.Println("密码错", 全局_密码, ",", c.PostForm("password"))
		c.String(http.StatusOK, "密码错误")
		c.Abort()
		return
	}
}
func main() {
	初始化设置()
	r := gin.Default()
	r.POST("/:服务器文件夹/updata112233", 密码验证, uploadFile)
	r.POST("/:服务器文件夹/getFileList", 密码验证, getFileList)
	r.POST("/:服务器文件夹/getFile", 密码验证, getFile)
	r.POST("/:服务器文件夹/getFileListAdmin", 密码验证, getFileListAdmin)
	r.Run(":" + 全局_端口)
}

// go env -w GOOS=linux;go build .;go env -w GOOS=windows
// go env -w GOOS=linux;go build main.go 服务器.go;go env -w GOOS=windows
