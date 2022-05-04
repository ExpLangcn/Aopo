package common

import "flag"

func Flag(Info *Flaginfo) {
	flag.StringVar(&Info.Ip, "ip", "", "指定IP，支持格式：192.168.1.1｜192.168.1.1/24｜192.168.1.1/16｜192.168.1.1/8｜192.168.1.1,192.168.1.2")
	flag.StringVar(&Info.Ipfile, "ipf", "", "从文件读取目标，支持格式如上 注意：一行一个")
	flag.BoolVar(&Info.All, "all", false, "自动化探测内网资产，默认扫描：10 A段｜172 A段｜192 B段")
	flag.StringVar(&Info.Passdic, "addpass", "", "自定义字典合并到内置程序字典库中 (自动去重)")
	//flag.IntVar(&Info.Threads, "t", 500, "自定义线程 默认500线程")
	flag.Parse()
}
