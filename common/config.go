package common

import (
	"sync"
	"time"
)

var Icmpwg sync.WaitGroup
var Icmp8wg sync.WaitGroup
var Portwg sync.WaitGroup
var Portwebwg sync.WaitGroup
var Httpwg sync.WaitGroup
var Xraywg sync.WaitGroup
var Weakpass sync.WaitGroup
var Sshwg sync.WaitGroup
var Mysqlwg sync.WaitGroup
var Mssqlwg sync.WaitGroup
var Rediswg sync.WaitGroup
var Mongowg sync.WaitGroup
var Oraclewg sync.WaitGroup
var Smbwg sync.WaitGroup
var Statuses []bool
var Chanip = make(chan string)
var Timeout = time.After(3 * time.Second)

var Rwmap = struct {
	sync.RWMutex
	Portweblist map[string][]int
	Portlist    map[string][]int
	Sshlist     map[string][]bool
}{Portweblist: make(map[string][]int), Portlist: make(map[string][]int), Sshlist: make(map[string][]bool)}

var Xraylist []string
var Version = "1.0.2"

type Flaginfo struct {
	Ipfile  string
	Ip      string
	All     bool
	Passdic string
	Threads int
}

// =====================================================================================================================

var Userdict = map[string][]string{
	//"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data", "test", "user"},
	"mysql":      {"root", "mysql", "user"},
	"mssql":      {"sa", "sql", "user"},
	"smb":        {"administrator", "admin", "guest", "user"},
	"rdp":        {"administrator", "admin", "guest", "user"},
	"postgresql": {"postgres", "admin", "user"},
	"ssh":        {"root", "test", "admin", "user"},
	"mongodb":    {"root", "admin", "user"},
	"oracle":     {"sys", "system", "admin", "test", "web", "orcl", "user"},
}

// =====================================================================================================================

var Passwords = []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "{user}@123#4", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e", "Charge123", "Aa123456789"}

// =====================================================================================================================

var PORTList = map[string]int{
	//"ftp":    21,
	"ssh":    22,
	"smb":    445,
	"mssql":  1433,
	"oracle": 1521,
	"mysql":  3306,
	"redis":  6379,
	"mgo":    27017,
}

// =====================================================================================================================

var Webport = []int{80, 81, 82, 88, 91, 443, 2018, 2019, 7000, 7001, 7002, 8000, 8001, 8008, 8014, 8042, 8069, 8080, 8081, 8082, 8087, 8088, 8090, 8091, 8100, 8118, 8123, 8188, 8172, 8222, 8243, 8251, 8280, 8281, 8333, 8443, 8500, 8834, 8879, 8880, 8888, 8983, 9000, 9043, 9060, 9080, 9090, 9091, 9200, 9443, 9800, 9981, 10000, 10001, 10003, 10002, 12443, 15672, 16080, 18091, 18080, 18082, 18092, 20720, 28017, 83, 84, 85, 86, 89, 90, 99, 800, 801, 808, 880, 888, 889, 1000, 1010, 1980, 3000, 3505, 5555, 6080, 6677, 7003, 7005, 7007, 7070, 7080, 7200, 7777, 7890, 8002, 8003, 8004, 8006, 8009, 8010, 8011, 8012, 8016, 8020, 8028, 8053, 8060, 8070, 8083, 8084, 8085, 8086, 8089, 8099, 8161, 8180, 8181, 8200, 8300, 8360, 8484, 8800, 8881, 8899, 8989, 9001, 9002, 9081, 9085, 9999, 38501, 3888880, 87, 92, 98, 1080, 1081, 1082, 1099, 1118, 1888, 2008, 2020, 2100, 2375, 2379, 3008, 3128, 6648, 6868, 7004, 7008, 7071, 7074, 7078, 7088, 7680, 7687, 7688, 8018, 8030, 8038, 8044, 8046, 8048, 8092, 8093, 8094, 8095, 8096, 8097, 8098, 8101, 8108, 8244, 8258, 8288, 8448, 8838, 8848, 8858, 8868, 9008, 9010, 9082, 9083, 9084, 9086, 9087, 9088, 9089, 9092, 9093, 9094, 9095, 9096, 9097, 9098, 9099, 9100, 9448, 9986, 9988, 9998, 10004, 10008, 10010, 10250, 12018, 14000, 18000, 18001, 18002, 18004, 18008, 18088, 18090, 18098, 19001, 20000, 21000, 21501, 21502, 28018, 20880} // Top130
