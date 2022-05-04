package common

import (
	"Aopo/gologger"
	"bytes"
	"fmt"
	"github.com/k8gege/LadonGo/str"
	"io"
	"net/http"
	"regexp"
	"time"

	//"os"
	"strings"
)

func Httpscan(url string, port int) {
	defer Httpwg.Done()
	//fmt.Println(url, port)
	HttpBanner(url, port)
}
func IsUrl(url string, port string) (result string) {
	if !strings.Contains(url, "http") {
		url := "http://" + url + ":" + port
		return url
	} else if !strings.Contains(url, "https") {
		url := "https://" + url + ":" + port
		return url
	}
	return url
}
func HttpBanner(url string, port int) (result bool, err error) {

	url2 := IsUrl(url, fmt.Sprint(port))
	response, err := http.Head(url2)
	title := ScanTitle(url2)
	if err != nil {
		//fmt.Println(err.Error())
		//os.Exit(2)
		return false, err
	}

	//fmt.Println(response)
	//fmt.Println(response.Status)
	Xraylist = append(Xraylist, url2)
	gologger.Infof("WEB资产: " + url2 + " ｜ " + fmt.Sprint(response.StatusCode) + " ｜ " + fmt.Sprint(response.ContentLength) + " ｜ " + title)
	//gologger.Infof("%v | %v | %v | [%v]", url2, aurora.Yellow(response.StatusCode), aurora.Blue(response.ContentLength), aurora.Green(title))
	return result, err

}
func GetHtml(url string) string {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		//panic(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			//panic(err)
			return ""
		}
	}
	return result.String()
}

func GetTitle(html string) string {
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", "", -1)
	title := strings.Trim(str.GetBetween(html, "<title>", "</title>"), " ")
	return title
}

func ScanTitle(host string) (title string) {
	if strings.Contains(host, ":") {
		title = GetTitle(GetHtml(host))
	} else {
		url := "http://" + host
		title = GetTitle(GetHtml(url))

		url = "https://" + host
		title = GetTitle(GetHtml(url))
	}
	return title
}
