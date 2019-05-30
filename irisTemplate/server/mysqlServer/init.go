package mysqlServer

import (
	"../../config"
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Mysql *sql.DB
)

func init() {
	user := config.Config.Get("mysql.user").(string)
	password := config.Config.Get("mysql.password").(string)
	database := config.Config.Get("mysql.database").(string)
	host := config.Config.Get("mysql.host").(string)
	port := config.Config.Get("mysql.port").(string)
	//db, err := sqlHandle.Open("mysql", "数据库账号:数据库密码@tcp(数据库Ip:端口)/数据库名称?parseTime=true")
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	Mysql = db
}
