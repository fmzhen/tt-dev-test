package env

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// connect db
	dbHostName := beego.AppConfig.String("db_host_name")
	dbPort := beego.AppConfig.String("db_post")
	dbName := beego.AppConfig.String("db_name")
	dbUser := beego.AppConfig.String("db_user")
	dbPassword := beego.AppConfig.String("db_password")
	dbConnStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUser, dbPassword, dbHostName, dbPort, dbName)
	orm.RegisterDataBase("default", "postgres", dbConnStr)

	//log
	logFile := beego.AppConfig.String("log_file")
	logConf := fmt.Sprintf(`{"filename":"%+v"}`, logFile)
	beego.SetLogger("file", logConf)
}
