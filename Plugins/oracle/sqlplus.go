package oracle

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"os/exec"
)

func SqlplusAuth(host string, port string, user string, pass string, db string) bool {
	cmd := exec.Command("sqlplus", "-s", user+"/"+pass+"@"+host+":"+port+"/"+db)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
		return false
	}
	if string(output) == "" {
		return true
	}
	return false
}

func SqlPlusScan2(ScanType string, Target string) {
Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			fmt.Println("Check... " + Target + " " + u + " " + p)
			res := SqlplusAuth(Target, "1521", u, p, "orcl")
			if res {
				logger.PrintIsok2(ScanType, Target, "1521", u, p)
				break Loop
			}
		}
	}
}
