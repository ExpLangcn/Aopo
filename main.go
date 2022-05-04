package main

import (
	"Aopo/common"
	"Aopo/gologger"
	host "Aopo/lib"
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"time"
)

func main() {
	banner := `
                            
       /\                     
      /  \   ___  _ __   ___  
     / /\ \ / _ \| '_ \ / _ \ 
    / ____ \ (_) | |_) | (_) |
   /_/    \_\___/| .__/ \___/ 
                 | |    By：` + aurora.Red("ExpLang").String() + `      
                 |_|    Version: ` + aurora.Green(common.Version).String() + `
 Github: ` + aurora.Blue("github.com/ExpLangcn/Aopo").String() + `

`
	print(banner)
	time.Sleep(time.Duration(1) * time.Second)
	gologger.Printf(`法律免责声明
  本工具仅面向合法授权的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。
  在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。请勿对非授权目标进行扫描。
  如果发现上述禁止行为，我们将保留追究您法律责任的权利。

  如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任.
  您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。`)
	gologger.Printf("")
	time.Sleep(time.Duration(1) * time.Second)
	start := time.Now()
	var Info common.Flaginfo
	common.Flag(&Info)
	if Info.Passdic != "" {
		zhfile(Info.Passdic)
	}
	if Info.All != false {
		if Info.Ip != "" || Info.Ipfile != "" {
			gologger.Infof("错误：参数过多")
			os.Exit(0)
		}
		host.Host("192.168.0.0/16,10.0.0.0/8,172.0.0.0/8", "", "")
	} else if Info.Ip != "" {
		if Info.Ipfile != "" {
			gologger.Infof("错误：参数过多")
			os.Exit(0)
		}
		host.Host(Info.Ip, "", "")
	} else if Info.Ipfile != "" {
		if Info.Ip != "" {
			gologger.Infof("错误：参数过多")
			os.Exit(0)
		}
		host.Host("", Info.Ipfile, "")
	}
	t := time.Now().Sub(start)
	gologger.Printf("================================== 结束END ==================================")
	gologger.Infof("扫描结束,耗时: " + fmt.Sprint(t) + "\n")
}

func zhfile(file string) {
	fp, err := os.Open(file)
	if err != nil {
		fmt.Println(err) //打开文件错误
		return
	}
	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		if IsContain(common.Passwords, line) {
			continue
		} else {
			common.Passwords = append(common.Passwords, line)
			gologger.Infof("成功合并: " + line + " 到内置字典")
		}
	}
}
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
