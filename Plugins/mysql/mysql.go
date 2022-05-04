package mysql

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlAuth(ip string, port string, user string, pass string) (result bool, err error) {
	//defer common.Mysqlwg.Done()
	result = false
	db, _ := sql.Open("mysql", user+":"+pass+"@tcp("+ip+":"+port+")/mysql?charset=utf8")
	db.Close()
	if db.Ping() == nil {
		result = true
		gologger.Infof("Mysql 爆破成功: " + ip + ":" + port + " " + user + " " + pass)
	}
	common.Mysqlwg.Done()
	return result, err
}
func ScanMysql(ip string, port string, user string, pass string) (err error, result bool) {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", user,
		pass, ip, port, "mysql")
	db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		err = db.Ping()
		if db.Ping() == nil {
			db.Close()
			result = true
			gologger.Infof("Mysql 爆破成功: " + ip + ":" + port + " " + user + " " + pass)
		}
	}
	return err, result
}
