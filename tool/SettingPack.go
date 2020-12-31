package tool

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func SetSavePlace() bool {
	fmt.Println("设置saveplace")
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path += "\\download" //下载位置
	_, err = os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.Mkdir(path, 0666)
		if err != nil {
			fmt.Println(err)
		}
	}
	key, err := registry.OpenKey(registry.CLASSES_ROOT, "gopan", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
	}
	defer key.Close()

	setting, err := registry.OpenKey(key, "Setting", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
	}
	defer setting.Close()
	err = setting.SetStringValue("savePlace", path)
	fmt.Println("设置完成saveplace")
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("设置目录成功，在根目录下的download文件夹")
		return true
	}
}

func SetEXEUrl() bool {
	fmt.Println("设置执行文件目录")
	filepath, err := os.Getwd()
	str := filepath
	str += "\\GoPan.exe"
	if err != nil {
		fmt.Println(err)
	}
	key, err := registry.OpenKey(registry.CLASSES_ROOT, "gopan", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
	}
	defer key.Close()

	setting, err := registry.OpenKey(key, "Setting", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
	}
	defer setting.Close()
	err = setting.SetStringValue("EXEUrl", str)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("设置执行目录成功")
		return true
	}
}
