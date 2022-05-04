package oracle

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"database/sql"
	_ "github.com/godror/godror"
)

func OracleAuth(host string, port string, user string, pass string) (result bool) {
	//defer common.Oraclewg.Done()
	db, err := sql.Open("godror", user+"/"+pass+"@"+host+":"+port+"/orcl")
	//if err != nil {
	//panic(err)
	//return false
	//}
	if err == nil {
		if db.Ping() == nil {
			db.Close()
			result = true
			return result
		}
	}
	result = false
	if result != false {
		gologger.Infof("Oracle 爆破成功 " + host + ":" + port + " " + user + " " + pass)

		//go gologger.Infof(aurora.Red("Oracle 爆破成功 %v:%v %v %v").String(), host, port, user, pass)
	}
	common.Oraclewg.Done()
	return result
}
