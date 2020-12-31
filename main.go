package main

import (
	"GoPanClient/model"
	"GoPanClient/tool"
	"GoPanClient/view/FileView"
	"fmt"
	"strings"
	"sync"
)

var (
	wg         sync.WaitGroup
	savePlace  string
	goPanLink  string
	cpuThreads int
)

func main() {
	fmt.Println("正在检查是否第一次启动和下载目录正常与否")
	if  tool.CheckBootTime() {
		tool.CreateRegister()                             //第一次打开写入注册表启动
	}
	if tool.CheckSavePlace() {
		tool.SetSavePlace()                               //设置下载存放目录
	}
	savePlace, goPanLink, cpuThreads = tool.GetVars()
	if tool.BootFrom() == "web" {
		fmt.Println("web")
		fileTypeList:=strings.Split(goPanLink,".")
		fileType:="."+fileTypeList[1]
		tool.StartDownload(goPanLink,fileType)
		} else if tool.BootFrom() == "pc" {
			fmt.Println("pc")
		}
	cfg:=model.GetConfig()
	gormDB,sqlDB:=model.RegisterDB(cfg.DBUsername,cfg.DBPassword,cfg.DBHost,cfg.DBPort,cfg.DBName)
	defer gormDB.Close()
	defer sqlDB.Close()
	FileView.ShowUI()
}
