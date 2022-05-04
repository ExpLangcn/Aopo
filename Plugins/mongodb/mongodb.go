package mgo

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"gopkg.in/mgo.v2"
	"time"
)

func MongoAuth(ip string, port string, username string, password string) (result bool, err error) {
	//defer common.Mongowg.Done()
	session, err := mgo.DialWithTimeout("mongodb://"+username+":"+password+"@"+ip+":"+port+"/"+"admin", time.Second*3)
	if err == nil && session.Ping() == nil {
		session.Close()
		if err == nil && session.Run("serverStatus", nil) == nil {
			result = true
		}
	}
	if result != false {
		gologger.Infof("MongoDB 爆破成功 " + ip + ":" + port + " " + username + " " + password)

		//gologger.Infof(aurora.Red("MongoDB 爆破成功 %v:%v %v %v").String(), ip, port, username, password)
	}
	common.Mongowg.Done()
	return result, err
}
