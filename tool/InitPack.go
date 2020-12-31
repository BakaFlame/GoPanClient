package tool

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"runtime"
)

func GetVars() (string, string, int) {
	//获取下载目录
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
	savePlace, _, err := setting.GetStringValue("savePlace")
	if err != nil {
		fmt.Println(err)
	}
	//获取cpu线程数
	cpuThreads := runtime.NumCPU()

	//获取启动项参数
	var GoPanLink string
	if len(os.Args) == 2 {
		GoPanLink = os.Args[1]
	} else if len(os.Args) == 3 {
		GoPanLink = os.Args[1]
	} else {
		GoPanLink = ""
	}

	return savePlace, GoPanLink, cpuThreads
}

func BootFrom() string {
	argsLen := len(os.Args)
	//fmt.Println("参数长度" + strconv.Itoa(argsLen))
	if argsLen >= 2 {
		//fmt.Println("参数1:" + os.Args[0])
		return "web"
	} else {
		return "pc"
	}
}

func CreateRegister() {
	fmt.Println("检测到是第一次启动,正在初始化内容")
	//创建主键值
	key, exists, err := registry.CreateKey(registry.CLASSES_ROOT, "gopan", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}

	//创建settings
	setting, exists, err := registry.CreateKey(key, "Setting", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer setting.Close()

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}

	//创建shell
	shell, exists, err := registry.CreateKey(key, "Shell", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer shell.Close()

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}

	//创建open
	open, exists, err := registry.CreateKey(shell, "Open", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer open.Close()

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}

	dst, exists, err := registry.CreateKey(open, "command", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}

	defer dst.Close()

	//"D:\迅雷\Program\Thunder.exe" "%1" -StartType:thunder
	str := `"`
	filepath, _ := os.Getwd()
	str += filepath
	str += `\GoPan.exe" "%1"`

	err = dst.SetExpandStringValue("", str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)

	setting.SetStringValue("bootTime", "1")

	// fmt.Println("请设置你的保存盘符,默认为本目录下的download文件夹")
	// var savePlace string
	// _, err = fmt.Scanf("%s", &savePlace)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("开始设置下载目录")
	// if SettingPack.SetSavePlace() {
	// 	fmt.Println("设置下载目录成功")
	// } else {
	// 	fmt.Println("设置失败,将使用本目录下的download文件夹作为目录")
	// }
}
