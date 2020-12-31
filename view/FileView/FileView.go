package FileView

import (
"GoPanClient/controller/FileController"
	"GoPanClient/tool"
	"bufio"
	"fmt"
"os"
)

//var (
//	option chan string
//)

func ShowUI() {
	//option = make(chan string, 1)
	for {
		fmt.Println("初始化界面")
		fmt.Println("————————————————————")
		fmt.Println("0.退出程序\t1.显示文件列表\t2.更改下载目录(还没做呢)\t3.开始同步数据库")
		var inputStr string
		input:= bufio.NewScanner(os.Stdin)
		input.Scan()
		inputStr = input.Text()
		switch inputStr {
		case "0":
			os.Exit(0)
		case "1":
			FileController.FileList(1)
			break
		case "2":
			tool.ClearTerm()
			fmt.Println("还没做呢")
			break
		case "3":
			tool.ClearTerm()
			FileController.SyncWithHaoshen()
			break
		}
	}
}

