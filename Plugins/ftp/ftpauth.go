package ftp

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func FtpAuth(ip string, port string, user string, pass string) {
	//fmt.Println(ip+":"+port, time.Duration(5)*time.Second)
	conn, _ := ftp.DialTimeout(ip+":"+port, time.Duration(5)*time.Second)
	err := conn.Login(user, pass)
	fmt.Println(err)
	if err == nil {
		defer conn.Logout()
		defer common.Bp.Done()
		gologger.Infof("FTP 爆破成功 " + ip + ":" + port + " " + user + " " + pass)
	}

}
