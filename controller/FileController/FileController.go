package FileController

import (
	"GoPanClient/model/FileModel"
	"GoPanClient/tool"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FileList(page int) {
	tool.ClearTerm()
	file, dataCount, maxPage := GetFileList(page)
	for {
		//fmt.Println("id\t|\tname\t|\t\t\t\t\t\t\t\t\turl\t\t\t\t\t\t\t\t|\ttype\t|")
		fmt.Println("id | name |\t\t\t\t\turl\t\t\t\t\t| type |")
		for i := 0; i < len(file); i++ {
			fmt.Println(strconv.Itoa(file[i].ID) + " | " + file[i].Name + " | " + file[i].Url + " | " + file[i].Type + " |\n")
		}
		fmt.Printf("查询完毕，共计页数%v，一共有%v条记录", maxPage, dataCount)
		fmt.Println("请选择文件id进行下载\t输入up或者down翻页\t输入0退出此界面\t输入-1退出程序")
		var option string
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		option = input.Text()
		mightInt, err := strconv.Atoi(option)
		if err != nil {
			tool.ClearTerm()
			switch option {
			case "down":
				if page == maxPage {
					//fmt.Printf("输入页数%v\n",page)
					//fmt.Printf("最大页数%v\n",maxPage)
					fmt.Println("已经是最后一页了")
					break
				} else {
					page += 1
					break
				}
			case "up":
				if page == 0 {
					fmt.Println("已经是第一页了")
					//fmt.Printf("输入页数%v\n",page)
					//fmt.Printf("最大页数%v\n",maxPage)
					break
				} else {
					page -= 1
					break
				}
			}
			file, dataCount, maxPage = GetFileList(page)
		} else {
			if mightInt == 0 {
				fmt.Println("退出文件列表")
				break
				//break
			} else if mightInt == -1 {
				fmt.Println("退出程序")
				os.Exit(0)
			} else if FileModel.CheckFileIdExist(mightInt) {
				fileInfo := GetFileById(mightInt)
				//fmt.Println(dataCount)
				//fmt.Println(mightInt)
				//tool.StartDownload(url)
				_, err := tool.StartDownload(fileInfo.Url, fileInfo.Type)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("输入错误id信息")
			}
			file, dataCount, maxPage = GetFileList(page)
		}
	}
}

func GetFileList(page int) ([]FileModel.File, int, int) {
	page = page - 1
	//fmt.Printf("页数%v\n", page)
	fmt.Println("正在查询文件列表")
	fmt.Println("————————————————————")
	file, dataCount, maxPage := FileModel.GetFileList(page)
	//fmt.Println("执行完getfilelist")
	//fmt.Println(file)
	return file, dataCount, maxPage
}

func GetFileById(id int) FileModel.File {
	file := FileModel.GetFileUrlById(id)
	if file.ID == 0 {
		return FileModel.File{Url: "http://10.24.2.84:8080/static/img/image.png", Type: ".png"}
	} else {
		return file
	}
}

func SyncWithHaoshen() {
	fmt.Println("正在同步")
	FileModel.SyncWithHaoshen()
	fmt.Println("同步完成")
}

//func DirectDownload() {
//	tool.ClearTerm()
//	fmt.Println("请输入链接\t例如http://10.24.2.84:8080/static/file/test.zip")
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	url := input.Text()
//	if _,err:=tool.StartDownload(url);err != nil {
//		fmt.Println("下载失败")
//	}
//}
