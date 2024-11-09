package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var 下载url string //= "http://mb1.hcbyj.xyz:8080/project"

func main() {
	初始化设置()
	下载url = 全局_地址 + ":" + 全局_端口 + "/" + 全局_项目
	var 服务器所有文件 map[string]string
	var r string
	for i := 0; i < 100; i++ {
		r = httpPost(下载url+"/getFileList", url.Values{"password": {全局_密码}})
		服务器所有文件 = json转map(r)
		if len(服务器所有文件) > 1 {
			break
		}
		fmt.Println(`※※  更新失败  ※※10秒后再次尝试※※` + strconv.Itoa(i))
		time.Sleep(time.Second * 10)
	}

	fmt.Println(len(r))
	fmt.Println(永无BUG)
	本地所有文件 := 获取本地文件(".")
	fmt.Println("本地:", len(本地所有文件))
	fmt.Println("云端:", len(服务器所有文件))
	fmt.Println(永无BUG)
	url := 下载url + "/getFile" // 下载文件的 URL
	需要下载的文件 := []string{}
	for k, 服务器md5 := range 服务器所有文件 {
		if 服务器md5 != 本地所有文件[k] {
			// fmt.Println("需要下载", k)
			// fmt.Println(url)
			// fileName := k
			需要下载的文件 = append(需要下载的文件, k) // 下载的文件名
			// err := DownloadFile(url, fileName)
			// if err != nil {
			// 	fmt.Println("下载文件失败:", err)
			// 	return
			// }
			// fmt.Print("下载完成!")

		} else {

		}
	}
	fmt.Println("需要下载文件数量:", len(需要下载的文件))
	for num, fileName := range 需要下载的文件 {
		fmt.Println("下载:", num+1, "/", len(需要下载的文件), "文件", fileName)
		err := DownloadFile(url, fileName)
		if err != nil {
			fmt.Println("下载文件失败:", err)
			return
		}

	}
	fmt.Println(`下载结束`, "文件数量:", len(服务器所有文件))
	if len(服务器所有文件) == 0 {
		for i := 1; i < 5; i++ {

			fmt.Println(`※※※※※※※※※  更新失败  ※※※※※※※※※`)
		}

	}

}

// DownloadFile 从给定的 URL 下载文件到本地，并显示下载进度条。
func DownloadFile(url2 string, fileName string) error {
	resp, err := http.PostForm(url2, url.Values{"fileName": {fileName}, "password": {全局_密码}})
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()
	{

		dirPath := filepath.Dir(fileName)
		// 判断目录是否存在
		_, err := os.Stat(dirPath)
		if os.IsNotExist(err) {
			// 目录不存在，创建目录
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				return err
			}
		}
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("创建文件失败")
		fmt.Println(fileName)
		return err
	}
	defer file.Close()
	// 获取文件大小
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	// 创建进度条
	barLength := 50
	progressBar := NewProgressBar(size, barLength)
	// 创建多个写入器，用于写入文件和进度条
	writers := []io.Writer{file, progressBar}
	writer := io.MultiWriter(writers...)
	// 复制响应体到文件和进度条
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// DownloadFile 从给定的 URL 下载文件到本地，并显示下载进度条。
func DownloadFile2(url string, fileName string) error {
	a := map[string]string{}
	a["fileName"] = fileName
	b, _ := json.Marshal(a)
	// 创建 HTTP 客户端
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// 发送 HTTP 请求，并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 创建文件
	// fileName = "/" + strings.Replace(fileName, "\\", "/", -1)
	// os.MkdirAll()
	{

		dirPath := filepath.Dir(fileName)
		// 判断目录是否存在
		_, err := os.Stat(dirPath)
		if os.IsNotExist(err) {
			// 目录不存在，创建目录
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				return err
			}
		}
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("创建文件失败")
		fmt.Println(fileName)
		return err
	}
	defer file.Close()
	// 获取文件大小
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	// 创建进度条
	barLength := 50
	progressBar := NewProgressBar(size, barLength)
	// 创建多个写入器，用于写入文件和进度条
	writers := []io.Writer{file, progressBar}
	writer := io.MultiWriter(writers...)
	// 复制响应体到文件和进度条
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// ProgressBar 用于显示下载进度条
type ProgressBar struct {
	total    int
	current  int
	barLen   int
	barStart string
	barEnd   string
}

// NewProgressBar 创建一个新的 ProgressBar 对象
func NewProgressBar(total int, barLen int) *ProgressBar {
	return &ProgressBar{
		total:    total,
		barLen:   barLen,
		barStart: "[",
		barEnd:   "]",
	}
}

// Write 方法实现 io.Writer 接口，用于向终端窗口写入进度条
func (p *ProgressBar) Write(b []byte) (int, error) {
	n := len(b)
	p.current += n
	p.draw()
	return n, nil
}

// draw 用于计算并显示进度条
func (p *ProgressBar) draw() {
	// 计算百分比和进度条长度
	percent := float64(p.current) / float64(p.total) * 100
	bars := int(percent / 100 * float64(p.barLen))
	// 绘制进度条
	fmt.Printf("\r%.2f%%%s%s%s]", percent, p.barStart,
		strings.Repeat(">", bars), strings.Repeat(" ", p.barLen-bars))
	// 如果全部完成，输出换行符
	if p.current == p.total {
		fmt.Println()
	}
}
