package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//程序入口
func main() {
	var baseUrl = "http://jandan.net/ooxx/page-"

	doProcess(baseUrl, baseUrl, 2932)

}

func doProcess(baseUrl string, url string, pagenum int) {
	pageNum := pagenum
	cc := make(chan string)
	next_url := parseContent(cc, url, pageNum)
	if pageNum < 2940 {
		doProcess(baseUrl, baseUrl, next_url)
	}
}

func parseContent(cc chan string, url string, pagenum int) int {
	var innerImg = []string{}
	doc, _ := goquery.NewDocument(url + strconv.Itoa(pagenum))
	div_body := doc.Find("div.row")
	//fmt.Println("div row:", div_body)
	div_body.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		img_path, _ := s.Attr("href")
		innerImg = append(innerImg, "http:"+img_path)
		//fmt.Println(img_path)
	})

	downPic(innerImg, pagenum)
	return pagenum + 1
}

//下载功能
func downPic(paths []string, pagenum int) {
	fmt.Println("page : " + strconv.Itoa(pagenum))
	var base_dir = "C:/pic/"

	for _, path := range paths {
		fmt.Println("single path:", path)
		res, _ := http.Get(path)
		file_name, dir := splitPath(path)
		os.Mkdir(base_dir+dir, 0777)

		//fmt.Println("save dir:", base_dir+dir)
		save_dir := base_dir + dir + "/"
		fmt.Println("file:", save_dir+file_name)

		file, _ := os.Create(save_dir + file_name)
		io.Copy(file, res.Body)
	}
}

//路径处理
func splitPath(path string) (string, string) {
	ss := strings.Split(path, "/")
	lenth := len(ss)
	file_name := ss[lenth-1]
	dir_name := ss[lenth-3] + ss[lenth-2]
	return file_name, dir_name
}
