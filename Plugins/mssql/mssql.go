package mssql

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func MssqlAuth(ip, port, user, pass string) (result bool, err error) {
	//defer common.Mssqlwg.Done()
	result = false
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;encrypt=disable", ip, user, pass, port)
	db, err := sql.Open("mssql", connString)
	if err == nil {
		db.Close()
		err = db.Ping()
		if err == nil {
			result = true
		}
	}
	if result != false {
		gologger.Infof("Mssql 爆破成功 " + ip + ":" + port + " " + user + " " + pass)
		//go gologger.Infof(aurora.Red("Mssql 爆破成功 %v:%v %v %v").String(), ip, port, user, pass)
	}
	common.Mssqlwg.Done()
	return result, err
}
