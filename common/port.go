package common

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/gologger"
	"net"
	"strconv"
	"sync"
	"time"
)

var scan sync.WaitGroup

func ScanwebPort(host string) {
	defer Portwebwg.Done()
	for _, p := range Webport {
		scan.Add(1)
		//CheckPort(net.ParseIP(host),p)
		go PortwebCheck(host, p)
	}
	scan.Wait()
	//Portwebwg.Done()
}
func ScanPort(host string) {
	defer Portwg.Done()
	for p := range PORTList {
		scan.Add(1)
		//CheckPort(net.ParseIP(host),p)
		go PortCheck(host, PORTList[p])
	}
	scan.Wait()
	//Portwg.Done()
}
func PortwebCheck(host string, port int) bool {
	defer scan.Done()
	p := strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", host+":"+p, time.Second*3)
	if err != nil {
		//fmt.Println(host, p, "Close")
		return false
	} else {
		gologger.Infof("WEB服务: " + host + " " + p + " Open")
		//_, ok := Portlist[host]
		Rwmap.Lock()
		Rwmap.Portweblist[host] = append(Rwmap.Portweblist[host], port)
		Rwmap.Unlock()
		conn.Close()
		//scan.Done()
		return true
	}
}
func PortCheck(host string, port int) bool {
	defer scan.Done()
	p := strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", host+":"+p, time.Second*3)
	if err != nil {
		//fmt.Println(host, p, "Close")
		return false
	} else {
		gologger.Infof("主机服务: " + host + " " + p + " Open")
		//_, ok := Portlist[host]
		Rwmap.Lock()
		Rwmap.Portlist[host] = append(Rwmap.Portlist[host], port)
		Rwmap.Unlock()
		conn.Close()
		return true
	}
}
