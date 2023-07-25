package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
 * @Time    : 2023/7/25 11:25
 * @File    : crawl1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 爬虫Demo-并发
 */

// HandleError 异常处理函数
func HandleError(err error, s string) {
	if err != nil {
		fmt.Println(s, err)
	}
}

// downloadFile 根据传入的URL以及图片名称下载图片
func downloadFile(url string, filename string) (ok bool) {
	// 获取图片数据
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		HandleError(err, "Body.Close")
	}(resp.Body)
	// 读取图片数据
	bytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// 将图片数据写入文件
	filename = "/Users/saitama/go/chapter6/crawl/img/" + filename
	err = os.WriteFile(filename, bytes, 0666)
	if err != nil {
		fmt.Println("os.WriteFile err:", err)
		return false
	} else {
		return true
	}
}

var (
	// 存放图片链接的通道
	imageUrlChan chan string
	chanTask     chan string
	waitGroup    sync.WaitGroup
	baseUrl      string = "https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/"
	reImg        string = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	imageUrlChan = make(chan string, 100000)
	chanTask = make(chan string, 26)

	// 开启26个协程，往图片链接通道里写数据
	for i := 1; i <= 27; i++ {
		waitGroup.Add(1)
		// 生成每个页面的URL
		curUlr := baseUrl + strconv.Itoa(i) + ".html"
		go writeUrl(curUlr)
	}
	// 任务统计协程
	waitGroup.Add(1)
	go checkTask()
	// 开始下载
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go downloadImage()
	}
	waitGroup.Wait()
}

func downloadImage() {
	for imageUrl := range imageUrlChan {
		fileName := getFileNameFromUrl(imageUrl)
		ok := downloadFile(imageUrl, fileName)
		if ok {
			fmt.Printf("%s下载成功", fileName)
		} else {
			fmt.Printf("%s下载失败", fileName)
		}
	}
	waitGroup.Done()
}

func getFileNameFromUrl(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	fileName := url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	fileName += "-" + timePrefix
	return fileName
}

func checkTask() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s已完成任务", url)
		count++
		if count == 26 {
			close(imageUrlChan)
			break
		}
	}
	waitGroup.Done()
}

// writeUrl 根据传入的URL获取当前页面的所有图片链接，并写入图片链接通道
func writeUrl(url string) {
	var urls []string = getImgUrls(url)
	for _, curUrl := range urls {
		imageUrlChan <- curUrl
	}
	// 标志当前协程的任务完成
	chanTask <- url
	waitGroup.Done()
}

// getImgUrls 获取当前页面的所有图片链接(字符串)
func getImgUrls(url string) (rets []string) {
	var pageStr string = GetPageStr(url)
	reg := regexp.MustCompile(reImg)
	results := reg.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		fmt.Println(result[0])
		rets = append(rets, result[0])
	}
	return rets
}

// GetPageStr 根据传入的URL获取页面内容
func GetPageStr(url string) string {
	// 封装请求
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	HandleError(err, "http.NewRequest")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	request.Header.Add("Referer", "https://www.bizhizu.cn/")
	request.Header.Add("Content-Type", "image/jpeg")
	// 发起请求
	resp, err := client.Do(request)
	HandleError(err, "http.Get url")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		HandleError(err, "Body.Close")
	}(resp.Body)
	// 读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	pageStr := string(pageBytes)
	return pageStr
}
