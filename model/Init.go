package model

import (
	"GoPanClient/InitPack"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	DB *gorm.DB //全局DB变量保证只初始化一次DB
	err error //附赠品 用于检查数据库连接是否错误
	DB_query *sql.DB //全局变量普通DB_query保证只初始化一次DB (待改善)
	cfg *InitPack.Config
)


func GetConfig() *InitPack.Config{
	//path,err:=os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//}
	cfg,err := InitPack.ParseConfig("./config/Config.json") //配置文件加载
	if err != nil {
		fmt.Println(err)
	}
	return cfg
}

type BetterTime struct{ //因为gorm自带转化time不是想要的标准格式 故用自定义time
	time.Time
}

func (t BetterTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t BetterTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *BetterTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = BetterTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func RegisterDB(userName string,password string, host string, port string, dbName string) (*gorm.DB,*sql.DB){
	//fmt.Println(userName+":"+password+"@tcp("+host+port+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", userName+":"+password+"@tcp("+host+port+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local") //连接MYSQL
	if err != nil {
		log.Fatal(err.Error())
	}

	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(100)

	connstr := userName+":"+password+"@tcp("+host+port+")/"+dbName//原生MYSQL
	DB_query, err = sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return DB,DB_query
}

func QuerySql(sqlString string) (string, error) {
	stmt, err := DB_query.Prepare(sqlString)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
