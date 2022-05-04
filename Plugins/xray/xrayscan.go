package xrayscan

import (
	"Aopo/Plugins/xray/lib"
	"Aopo/common"
	"embed"
	"net/http"
	"strings"
)

//go:embed pocs
var Pocs embed.FS
var Pocpath string
var PocName string

func Xrayscan(Target string) {
	defer common.Xraywg.Done()
	lib.Inithttp() //初始化httpclient
	req, err := http.NewRequest("GET", Target, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")

	if Pocpath == "" {
		lib.CheckBuiltinPoc(req, Pocs, 100, PocName)
	} else {
		if strings.HasSuffix(Pocpath, ".yml") || strings.HasSuffix(Pocpath, ".yaml") {
			lib.CheckSinglePoc(req, Pocpath)
		} else {
			lib.CheckExternalPoc(req, Pocpath, 100, PocName)
		}
	}
}
