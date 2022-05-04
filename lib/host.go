package host

import (
	"Aopo/Plugins/ssh"
	"Aopo/common"
	"Aopo/gologger"
	"fmt"
	"time"
)

var Hostlist []string

func Host(host string, filename string, nohosts ...string) {
	time.Sleep(time.Duration(1) * time.Second)
	if host != "" {
		gologger.Infof("当前加载的目标: " + host)
	}
	if filename != "" {
		gologger.Infof("当前加载的目标文件: " + filename)
	}
	gologger.Infof("当前密码字典数量: " + fmt.Sprintln(len(common.Passwords)))
	gologger.Infof("当前已启动模块: ICMP存活探测 ｜ 主机服务探测 ｜ 主机WEB服务资产探测 ｜ WEB服务资产漏洞验证 ｜ 主机服务弱口令爆破")
	time.Sleep(time.Duration(1) * time.Second)
	//fmt.Println(aurora.Yellow("==================================" + aurora.Blue("正在进行存活资产探测").String() + aurora.Yellow("==================================").String()))
	gologger.Printf("================================== 正在进行存活资产探测 ==================================")
	time.Sleep(time.Duration(1) * time.Second)
	Hostlist, _ = common.ParseIP(host, filename, nohosts...)
	gologger.Printf("================================== 正在进行存活资产统一服务探测 ==================================")
	//fmt.Println(aurora.Yellow("==================================" + aurora.Blue("正在进行存活资产统一服务探测").String() + aurora.Yellow("==================================").String()))
	time.Sleep(time.Duration(1) * time.Second)
	for i := range Hostlist {
		common.Portwg.Add(1)
		go common.ScanPort(Hostlist[i])
	}
	common.Portwg.Wait()
	//gologger.Printf("================================== 正在进行存活资产WEB服务探测 ==================================")
	////fmt.Println(aurora.Yellow("==================================" + aurora.Blue("正在进行存活资产WEB服务探测").String() + aurora.Yellow("==================================").String()))
	//time.Sleep(time.Duration(1) * time.Second)
	//for i := range Hostlist {
	//	common.Portwebwg.Add(1)
	//	go common.ScanwebPort(Hostlist[i])
	//}
	//common.Portwebwg.Wait()
	//common.Rwmap.RLock()
	//for i := range common.Rwmap.Portweblist {
	//	for p := range common.Rwmap.Portweblist[i] {
	//		common.Httpwg.Add(1)
	//		go common.Httpscan(i, common.Rwmap.Portweblist[i][p])
	//	}
	//	common.Httpwg.Wait()
	//}
	//common.Rwmap.RUnlock()
	//gologger.Printf("================================== 正在进行存活资产WEB服务漏洞探测 ==================================")
	////fmt.Println(aurora.Yellow("==================================" + aurora.Blue("正在进行存活资产WEB服务漏洞探测").String() + aurora.Yellow("==================================").String()))
	//time.Sleep(time.Duration(1))
	//for i := range common.Xraylist {
	//	common.Xraywg.Add(1)
	//	//gologger.Infof(aurora.Green("当前目标：" + common.Xraylist[i]).String())
	//	go xrayscan.Xrayscan(common.Xraylist[i])
	//}
	//common.Xraywg.Wait()
	gologger.Printf("================================== 正在进行存活资产服务弱口令探测 ==================================")
	//fmt.Println(aurora.Yellow("==================================" + aurora.Blue("正在进行存活资产服务弱口令探测").String() + aurora.Yellow("==================================").String()))
	time.Sleep(time.Duration(1) * time.Second)
	ipnumber := 1
	common.Rwmap.RLock()
	for i := range common.Rwmap.Portlist {
		common.Statuses = append(common.Statuses, false)
		go Auth(i, ipnumber)
		ipnumber += 1
	}
	common.Weakpass.Wait()
	common.Rwmap.RUnlock()
}
func Auth(hlist string, sum int) {
	//fmt.Println(hlist, port)
	//defer common.Weakpass.Done()
	//gologger.Infof("当前探测目标信息:" + fmt.Sprint(hlist) + "｜" + fmt.Sprint(port))
	for p := range common.Rwmap.Portlist[hlist] {
		if common.Rwmap.Portlist[hlist][p] == common.PORTList["ssh"] {
			sshchek(hlist, sum)
		}
	}
	//if port == common.PORTList["ssh"] {
	//gologger.Infof("正在爆破 " + hlist + " 的SSH服务")
	//sshchek(hlist)
	//} else if port == common.PORTList["mysql"] {
	//	//gologger.Infof("正在爆破 " + hlist + " 的Mysql服务")
	//	for u := range common.Userdict["mysql"] {
	//		for p := range common.Passwords {
	//			common.Mysqlwg.Add(1)
	//			go mysql.ScanMysql(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["mysql"]), common.Userdict["mysql"][u], common.Passwords[p])
	//		}
	//	}
	//	common.Mysqlwg.Wait()
	//} else if port == common.PORTList["mssql"] {
	//	//gologger.Infof("正在爆破 " + hlist + " 的Mssql服务")
	//	for u := range common.Userdict["mssql"] {
	//		for p := range common.Passwords {
	//			common.Mssqlwg.Add(1)
	//			go mssql.MssqlAuth(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["mssql"]), common.Userdict["mssql"][u], common.Passwords[p])
	//		}
	//	}
	//	common.Mssqlwg.Wait()
	//} else if port == common.PORTList["redis"] {
	//	//gologger.Infof("正在爆破 " + hlist + " 的Redis服务")
	//	common.Rediswg.Add(1)
	//	go redis.RedisNullAuth(fmt.Sprint(hlist), common.PORTList["redis"])
	//	common.Rediswg.Wait()
	//} else if port == common.PORTList["mgo"] {
	//	//gologger.Infof("正在爆破 " + hlist + " 的Mongodb服务")
	//	for u := range common.Userdict["mongodb"] {
	//		for p := range common.Passwords {
	//			common.Mongowg.Add(1)
	//			go mgo.MongoAuth(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["mgo"]), common.Userdict["mongodb"][u], common.Passwords[p])
	//		}
	//	}
	//	common.Mongowg.Wait()
	//} else if port == common.PORTList["oracle"] {
	//	//gologger.Infof("正在爆破 " + hlist + " 的Oracle服务")
	//	for u := range common.Userdict["oracle"] {
	//		for p := range common.Passwords {
	//			common.Oraclewg.Add(1)
	//			go oracle.OracleAuth(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["oracle"]), common.Userdict["oracle"][u], common.Passwords[p])
	//		}
	//	}
	//	common.Oraclewg.Wait()
	//} else if port == common.PORTList["smb"] {
	//	//gologger.Infof("正在爆破 " + hlist + " 的SMB服务")
	//	for u := range common.Userdict["smb"] {
	//		for p := range common.Passwords {
	//			common.Smbwg.Add(1)
	//			go smb.SmbAuth(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["smb"]), common.Userdict["smb"][u], common.Passwords[p])
	//		}
	//	}
	//	common.Smbwg.Wait()
	//}
}
func sshchek(hlist string, sum int) int {
	for u := range common.Userdict["ssh"] {
		for p := range common.Passwords {
			if common.Statuses[sum] == false {
				common.Sshwg.Add(1)
				fmt.Println(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["ssh"]), common.Userdict["ssh"][u], common.Passwords[p], sum)
				go ssh.ScanSsh(fmt.Sprint(hlist), fmt.Sprint(common.PORTList["ssh"]), common.Userdict["ssh"][u], common.Passwords[p], sum)
			}
		}
	}
	for {
		time.Sleep(10000 * time.Millisecond)
		if common.Statuses[sum] == true {
			fmt.Println(common.Statuses[sum])
			common.Sshwg.Wait()
			common.Weakpass.Done()
			return 1
		} else {
			common.Sshwg.Wait()
			common.Weakpass.Done()
			return 0
		}
	}
	//return 0
}
