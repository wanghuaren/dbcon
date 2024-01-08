package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func main() {
	initConf()
	initDB()
	initHttp()
}

// var db_host = "rm-6nn6p824cmnlxjc81ho.mysql.rds.aliyuncs.com"
// var db_port = "3306"
// var db_user = "proxyweb"
// var db_pwd = "dMLvZyPWpB@dci8m"
// var db_name = "hhweb"
// var cgi_name = "dbctl"
// var host = "192.168.1.8"
// var port = "8090"

var db_host string
var db_port string
var db_user string
var db_pwd string
var db_name string
var cgi_name string
var host string
var port string

var key = "db"

var sqlName = "aaa"
var sqlMoney = "777"

var mIni *ini.File
var mDB *sql.DB

func initDB() {
	mDB, _ = sql.Open("mysql", db_user+":"+db_pwd+"@tcp("+db_host+":"+db_port+")/"+db_name)
	if err := mDB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

}

func initHttp() {
	r := gin.Default()
	r.POST(cgi_name, dbCtl)
	httpListenerStr := host + ":" + port
	r.Run(httpListenerStr)
}

func dbCtl(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	sqlName = json["account"].(string)
	sqlMoney = json["money"].(string)
	var sqlStr = "INSERT INTO player_rechange (player_id,proxy_id,money,pay_time,player_account,group_name,order_number,title) SELECT id,(SELECT proxy_id FROM player_user WHERE account='" + sqlName + "' LIMIT 1)," + sqlMoney + ",NOW(),'" + sqlName + "','ZZJH',(DATE_FORMAT(NOW(),'%Y%m%" + "" + "d%H%i%" + "" + "s')),(CONCAT('玩家:','" + sqlName + "','充值:'," + sqlMoney + ")) FROM account WHERE name=CONCAT('" + sqlName + "','@game.sohu.com')"
	fmt.Println(sqlStr)
	result, err := mDB.Exec(sqlStr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result.RowsAffected())
	}
}

func initConf() {
	mIni, _ = ini.Load("config.ini")

	db_host = getString("db_host")
	db_port = getString("db_port")
	db_user = getString("db_user")
	db_pwd = getString("db_pwd")
	db_name = getString("db_name")
	cgi_name = getString("cgi_name")
	host = getString("host")
	port = getString("port")
}

func getString(field string) string {
	_cv := mIni.Section(key).Key(field).String()
	return _cv
}
