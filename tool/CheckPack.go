package tool

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func CheckBootTime() bool {
	fmt.Println("正在检查是否第一次启动")
	key, err := registry.OpenKey(registry.CLASSES_ROOT, "gopan", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		fmt.Println("找不到gopan")
		return true
	}
	defer key.Close()

	setting, err := registry.OpenKey(key, "Setting", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		fmt.Println("找不到setting")
		return true
	}
	defer setting.Close()

	value, _, err := setting.GetStringValue("bootTime")
	if value == "" {
		fmt.Println("bootTime为空")
		return true
	} else {
		fmt.Println("有bootTime\t" + value)
		return false
	}
}

func CheckSavePlace() bool {
	key, err := registry.OpenKey(registry.CLASSES_ROOT, "gopan", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		fmt.Println("gopan找不到")
	}
	defer key.Close()
	setting, err := registry.OpenKey(key, "Setting", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		fmt.Println("setting找不到")
	}
	defer setting.Close()
	value, _, err := setting.GetStringValue("savePlace")
	if value == "" {
		fmt.Println("无savePlace")
		return true
	} else {
		fmt.Println("有savePlace\t" + value)
		if err != nil {
			fmt.Println(err)
		}
		_, err = os.Stat(value)
		if err != nil {
			fmt.Println(err)
		}
		if os.IsNotExist(err) {
			// 创建文件夹
			fmt.Println("在该目录下没找到download文件夹")
			return true
		}
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path += "\\download" //下载位置
		_, err = os.Stat(path)
		if value != path {
			return true
		} else {
			return false
		}
	}
}
