package smb

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"encoding/hex"
)

var (
	negotiateProtocolRequest, _  = hex.DecodeString("00000085ff534d4272000000001853c00000000000000000000000000000fffe00004000006200025043204e4554574f524b2050524f4752414d20312e3000024c414e4d414e312e30000257696e646f777320666f7220576f726b67726f75707320332e316100024c4d312e325830303200024c414e4d414e322e3100024e54204c4d20302e313200")
	sessionSetupRequest, _       = hex.DecodeString("00000088ff534d4273000000001807c00000000000000000000000000000fffe000040000dff00880004110a000000000000000100000000000000d40000004b000000000000570069006e0064006f007700730020003200300030003000200032003100390035000000570069006e0064006f007700730020003200300030003000200035002e0030000000")
	treeConnectRequest, _        = hex.DecodeString("00000060ff534d4275000000001807c00000000000000000000000000000fffe0008400004ff006000080001003500005c005c003100390032002e003100360038002e003100370035002e003100320038005c00490050004300240000003f3f3f3f3f00")
	transNamedPipeRequest, _     = hex.DecodeString("0000004aff534d42250000000018012800000000000000000000000000088ea3010852981000000000ffffffff0000000000000000000000004a0000004a0002002300000007005c504950455c00")
	trans2SessionSetupRequest, _ = hex.DecodeString("0000004eff534d4232000000001807c00000000000000000000000000008fffe000841000f0c0000000100000000000000a6d9a40000000c00420000004e0001000e000d0000000000000000000000000000")
)

// func main() {
// host := flag.String("i", "", "single ip address")
// timeout := flag.Duration("t", 1, "timeout on connection, in seconds")
// netCIDR := flag.String("n", "", "CIDR notation of a network")
// flag.Parse()

// if *host != "" {
// checkHost(*host, *timeout)
// return
// }

// if *netCIDR != "" && *host == "" {
// ip, ipNet, err := net.ParseCIDR(*netCIDR)
// if err != nil {
// fmt.Println("invalid CIDR")
// return
// }
// var wg sync.WaitGroup

// for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
// wg.Add(1)
// go func(ip string) {
// defer wg.Done()
// checkHost(ip, *timeout)
// }(ip.String())
// }

// wg.Wait()
// }
// }
