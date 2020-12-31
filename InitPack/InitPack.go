package InitPack

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	DBHost string `json:"db_host"`
	DBPort string `json:"db_port"`
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBName string `json:"db_name"`
}

var _cfg *Config=nil

func ParseConfig(path string) (*Config,error) {
	file,err :=os.Open(path)
	//fmt.Println(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	decoder:=json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg);err !=nil{
		return nil,err
	}
	//fmt.Println(_cfg)
	return _cfg,nil
}