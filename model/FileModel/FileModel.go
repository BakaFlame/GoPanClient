package FileModel

import (
	"GoPanClient/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

//type Haoshen struct {
//	ID int `gorm:"column:id" json:"id"`
//	Name string `gorm:"column:name" json:"name"`
//	Url string `gorm:"column:url" json:"url"`
//	Type string `gorm:"column:type" json:"type"`
//	//"id": 1,
//	//"video_url": "1608726823584.mp4",
//	//"video_title": "【GUMI】谣言【ポリスピカデリー】",
//	//"createDate": "2020-12-25 20:14:01",
//	//"createBy": "cbf5625a443911ebb9f500ff85fbad2a",
//	//"synopsis": "sm35669777 https://www.youtube.com/watch?v=DWwFK2gjwa8\nRumor - Police Piccadilly feat. GUMI\n\nポリスピカデリー （mylist/53849103）\nTwitter : https://twitter.com/_policep\n\nOff Vox : のちほど",
//	//"media_cover": "1608726818112.jpg",
//	//"username": null
//}

func GetFileList(page int) ([]File,int,int){
	var dataCount int
	file:=[]File{}
	sql:="select id,name,url,create_time,type from files order by create_time desc limit 10 offset " + strconv.Itoa(page*10)
	//fmt.Println(sql)
	model.DB.Raw(sql).Scan(&file)
	model.DB.Model(&File{}).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 10)
	return file,dataCount, int(maxPage)
}

func GetFileUrlById(id int) File {
	file:=File{}
	sql:="select id,url,type from files where id = ?"
	model.DB.Raw(sql,id).Scan(&file)
	return file
}

func CheckFileIdExist(id int) bool {
	file:=File{}
	sql:="select id,url,type from files where id = ?"
	model.DB.Raw(sql,id).Scan(&file)
	if file.ID == 0 {
		return false
	} else {
		return true
	}
}

func SyncWithHaoshen() (error){
	request, err:=http.Get("http://10.24.1.240:8081/video/queryAll")
	//request, err:=http.Get("http://127.0.0.1:8080/getdata")
	if err != nil {
		log.Fatal(err)
	}
	defer request.Body.Close()
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	var result = map[string]interface{}{}
	err = json.Unmarshal(data, &result)	//将拿到的json从字节转成正常的interface类型
	jsonData:=result["data"].([]interface{})	//再将里面的interface断言成切片类型的interface(因为有多个所以是切片)
	for _, v := range jsonData {	//开始便利切片 切片下每一个value都是一个map[string]interface{}
		tempv:=v.(map[string]interface{})	//类型断言拿到string键下的interface{}值
		name:=tempv["video_title"].(string)	//开始将里面的值转化成string
		floatId:=tempv["id"].(float64)	//id因为是数字所以自动识别为float64 暂时转成float
		id:=int(floatId)	//通过float转回int 去掉小数点
		url:="http://10.24.1.240:8081/video/playVideo?id="+strconv.Itoa(id)	//拼接字符串
		fmt.Println(name)
		fmt.Println(url)
		//开始插入
		sql:="insert files(name,url,type) select "+"'"+name+"','"+url+"','.mp4' from dual where not EXISTS (select * from files where url = '"+url+"')"
		err=model.DB.Exec(sql).Error
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}
