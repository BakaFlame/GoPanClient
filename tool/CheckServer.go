package tool

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func CheckServer(url string) bool{
	newUrl:=strings.Split(url,"http://")	//分隔服务器地址
	fmt.Println(newUrl)
	newUrl=strings.Split(newUrl[1],"/")
	timeout := 5 * time.Second
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", newUrl[0], timeout)
	fmt.Println("waist time :", time.Now().Sub(t1))
	if err != nil {
		fmt.Println("链接网络错误，无法下载，error: ", err)
		return false
	} else {
		fmt.Println("服务器连接成功，开始下载")
		return true
	}
}
