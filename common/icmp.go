package common

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/gologger"
	"net"
	"time"
)

var Icmplist []string

func Icmpscan(host string) {
	defer Icmpwg.Done()
	var size int
	var timeout int64
	var seq int16 = 1
	const ECHO_REQUEST_HEAD_LEN = 8

	size = 1024
	timeout = 1000

	starttime := time.Now()
	conn, err := net.DialTimeout("ip4:icmp", host, time.Duration(timeout*1000*1000))
	if err != nil {
		return
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	id0, id1 := genidentifier(host)

	var msg = make([]byte, size+ECHO_REQUEST_HEAD_LEN)
	msg[0] = 8                        // echo
	msg[1] = 0                        // code 0
	msg[2] = 0                        // checksum
	msg[3] = 0                        // checksum
	msg[4], msg[5] = id0, id1         //identifier[0] identifier[1]
	msg[6], msg[7] = gensequence(seq) //sequence[0], sequence[1]

	length := size + ECHO_REQUEST_HEAD_LEN

	check := checkSum(msg[0:length])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	conn.SetDeadline(starttime.Add(time.Duration(timeout * 1000 * 1000)))
	_, err = conn.Write(msg[0:length])

	const ECHO_REPLY_HEAD_LEN = 20

	var receive = make([]byte, ECHO_REPLY_HEAD_LEN+length)
	n, err := conn.Read(receive)
	_ = n
	var endduration = int(int64(time.Since(starttime)) / (1000 * 1000))

	if err != nil || receive[ECHO_REPLY_HEAD_LEN+4] != msg[4] || receive[ECHO_REPLY_HEAD_LEN+5] != msg[5] || receive[ECHO_REPLY_HEAD_LEN+6] != msg[6] || receive[ECHO_REPLY_HEAD_LEN+7] != msg[7] || endduration >= int(timeout) || receive[ECHO_REPLY_HEAD_LEN] == 11 {
		//
	} else {
		Icmplist = append(Icmplist, host)
		//fmt.Println(host)
		gologger.Infof("ICMP存活主机: " + host)
	}
}
func Scan(host string) {
	defer Icmp8wg.Done()
	var size int
	var timeout int64
	var seq int16 = 1
	const ECHO_REQUEST_HEAD_LEN = 8

	size = 3
	timeout = 500

	starttime := time.Now()
	conn, err := net.DialTimeout("ip4:icmp", host, time.Duration(timeout*1000*1000))
	if err != nil {
		return
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	id0, id1 := genidentifier(host)

	var msg = make([]byte, size+ECHO_REQUEST_HEAD_LEN)
	msg[0] = 8                        // echo
	msg[1] = 0                        // code 0
	msg[2] = 0                        // checksum
	msg[3] = 0                        // checksum
	msg[4], msg[5] = id0, id1         //identifier[0] identifier[1]
	msg[6], msg[7] = gensequence(seq) //sequence[0], sequence[1]

	length := size + ECHO_REQUEST_HEAD_LEN

	check := checkSum(msg[0:length])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	conn.SetDeadline(starttime.Add(time.Duration(timeout * 1000 * 1000)))
	_, err = conn.Write(msg[0:length])

	const ECHO_REPLY_HEAD_LEN = 20

	var receive = make([]byte, ECHO_REPLY_HEAD_LEN+length)
	n, err := conn.Read(receive)
	_ = n
	var endduration = int(int64(time.Since(starttime)) / (1000 * 1000))

	if err != nil || receive[ECHO_REPLY_HEAD_LEN+4] != msg[4] || receive[ECHO_REPLY_HEAD_LEN+5] != msg[5] || receive[ECHO_REPLY_HEAD_LEN+6] != msg[6] || receive[ECHO_REPLY_HEAD_LEN+7] != msg[7] || endduration >= int(timeout) || receive[ECHO_REPLY_HEAD_LEN] == 11 {
		//
	} else {
		verip = append(verip, host+"/24")
		gologger.Infof("ICMP存活网段: " + host + "/24")
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	length := len(msg)
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256 // notice here, why *256?
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	var answer = uint16(^sum)
	return answer
}

func gensequence(v int16) (byte, byte) {
	ret1 := byte(v >> 8)
	ret2 := byte(v & 255)
	return ret1, ret2
}

func genidentifier(host string) (byte, byte) {
	return host[0], host[1]
}
