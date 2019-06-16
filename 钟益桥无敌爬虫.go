package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func SaveImg(i,idx int, url string, page chan int)  {
	path :="D:/buf/100/" + strconv.Itoa(i)+"/"+strconv.Itoa(idx+1) + ".jpg"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(" http.Get err:", err)
		return
	}
	defer f.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(" http.Get err:", err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		f.Write(buf[:n])
	}
	page <- idx
}

func spider2(i int,url string)  {


	// 爬取 整个页面，将整个页面全部信息，保存在result
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	// 解析编译正则
	ret := regexp.MustCompile(`src="(.*?)" class="">`)
	// 提取每一张图片的 url
	alls := ret.FindAllStringSubmatch(result, -1)

	page := make(chan int)
	n := len(alls)

	for idx, imgURL := range alls {
		//fmt.Println("imgURL:", imgURL[1])
		go SaveImg(i,idx, imgURL[1], page)
	}

	for i:=0; i<n; i++ {
		fmt.Printf("下载第 %d 张图片完成\n", <- page)
	}
	fmt.Printf("第%d完成",i)
}
func workinggogogo(start,end int)  {
	for i:=start;i<=end ;i++  {
		err:=os.Mkdir("D:/buf/100/"+strconv.Itoa(i),0777)
		if err!=nil{fmt.Println(err)}
		url := "https://movie.douban.com/top250?start="+strconv.Itoa((i-1)*25)+"&filter="
		spider2(i,url)
	}
}
func main()  {
	var start,end  int
	fmt.Print("输入起始页")
	fmt.Scan(&start)
	fmt.Print("输入尾页")
	fmt.Scan(&end)

	workinggogogo(start,end)

}

// 获取一个网页所有的内容， result 返回
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}