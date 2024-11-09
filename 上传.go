package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// var 更新url string = "http://127.0.0.1:8080"
var 更新url string //= "http://mb1.hcbyj.xyz:8080"

func main() {
	// http.HandleFunc("/upload", uploadFile)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// http.Post("http://127.0.0.1:8080/getFileList")
	初始化设置()
	更新url = 全局_地址 + ":" + 全局_端口 + "/" + 全局_项目
	fmt.Println(更新url)

	r := httpPost(更新url+"/getFileListAdmin", url.Values{"password": {全局_密码}})
	if r == "nil" {

	}
	服务器文件 := json转map(r)
	本地文件 := 获取本地文件("本地文件夹")
	delete(本地文件, "fileList.json")
	delete(本地文件, "本地文件夹/fileList.json")
	// fmt.Println("服务器", 服务器文件, "本地", 本地文件)
	fmt.Println(("开始上传文件:-------------------"))
	新的文件表 := map[string]string{}
	// 本地文件 := map[string]string{}
	for 本地文件名, 本地md5 := range 本地文件 {
		服务器文件名 := strings.Replace(本地文件名, "本地文件夹/", "", 1)
		// 本地文件[] = v
		新的文件表[服务器文件名] = 本地md5

		if 服务器文件[服务器文件名] != 本地md5 {
			fmt.Println("上传文件", 本地md5, ",", 服务器文件[服务器文件名], 服务器文件名)
			POST上传文件(更新url+"/updata112233", 本地文件名, 服务器文件名)
		} else {
			// fmt.Println("无需上传")
		}

	}
	s, _ := json.Marshal(新的文件表)
	err := os.WriteFile("本地文件夹/fileList.json", s, 0777)
	if err != nil {
		fmt.Println("本地文件夹/fileList.json.写入失败")
		fmt.Println(err)
	}
	// ioutil.WriteFile("本地文件夹/fileList.json", []byte("999"), 777)

	POST上传文件(更新url+"/updata112233", "本地文件夹/fileList.json", "fileList.json")

	fmt.Println(("结束上传文件:-------------------"))
	// fmt.Println(新的文件表)

}
func POST上传文件(url string, 文件名 string, 保存路径 string) {
	// url := "http://www.example.com/api/upload"
	file, err := os.Open(文件名)
	if err != nil {
		// 处理错误
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", 文件名)

	if err != nil {
		// 处理错误
	}
	_, err = io.Copy(part, file)
	if err != nil {
		// 处理错误
	}
	writer.WriteField("filePath", 保存路径)
	writer.WriteField("password", 全局_密码)
	err = writer.Close()
	if err != nil {
		// 处理错误
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// 处理错误
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
	}
	defer resp.Body.Close()
	fmt.Println("响应状态:", resp.Status)
}
