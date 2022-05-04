package ssh

import (
	"Aopo/common"
	"Aopo/gologger"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

//type Node struct {
//	User     string
//	Password string
//}
//
//func (this *Node) Conn(addr string, port string) (*ssh.Client, error) {
//
//	var authMethods []ssh.AuthMethod
//
//	keyboardInteractiveChallenge := func(
//		user,
//		instruction string,
//		questions []string,
//		echos []bool,
//	) (answers []string, err error) {
//		if len(questions) == 0 {
//			return []string{}, nil
//		}
//		return []string{this.Password}, nil
//	}
//
//	authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
//	authMethods = append(authMethods, ssh.Password(this.Password))
//
//	sshConfig := &ssh.ClientConfig{
//		User:    this.User,
//		Auth:    authMethods,
//		Timeout: 3 * time.Second,
//		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
//			return nil
//		},
//	}
//
//	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:"+port, addr), sshConfig)
//
//	if err != nil {
//		//return nil, fmt.Errorf(fmt.Sprintf("Connect fail \n(%s)", addr, err.Error()))
//	}
//
//	return client, nil
//}
//func SshAuth(host string, port string, user string, pass string) (result bool, err error) {
//	defer common.Sshwg.Done()
//	result = false
//	var authMethods []ssh.AuthMethod
//
//	keyboardInteractiveChallenge := func(
//		user,
//		instruction string,
//		questions []string,
//		echos []bool,
//	) (answers []string, err error) {
//		if len(questions) == 0 {
//			return []string{}, nil
//		}
//		return []string{pass}, nil
//	}
//
//	authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
//	authMethods = append(authMethods, ssh.Password(pass))
//
//	sshConfig := &ssh.ClientConfig{
//		User: user,
//		Auth: authMethods,
//		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
//			return nil
//		},
//	}
//
//	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", host, port), sshConfig)
//	if err == nil {
//		defer func(client *ssh.Client) {
//			_ = client.Close()
//		}(client)
//		session, err := client.NewSession()
//		errRet := session.Run("echo ISOK")
//		if err == nil && errRet == nil {
//			defer func(session *ssh.Session) {
//				_ = session.Close()
//			}(session)
//			result = true
//		}
//	}
//	if result != false {
//		gologger.Infof("SSH 爆破成功 " + host + ":" + port + " " + user + " " + pass)
//	}
//	//common.Sshwg.Done()
//	return result, err
//}

func ScanSsh(host string, port string, user string, pass string, sum int) {
	//defer common.Sshwg.Done()
	//fmt.Println(host, port, user, pass)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		Timeout: 2 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	fmt.Println(common.Statuses[sum])
	if common.Statuses[sum] == false {
		client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", host, port), config)
		fmt.Println(fmt.Sprintf("%v:%v", host, port), config)
		if err == nil {
			session, err := client.NewSession()
			client.Close()
			session.Close()
			if err == nil {
				gologger.Infof("SSH 爆破成功 " + host + ":" + port + " " + user + " " + pass)
				common.Statuses[sum] = true
				//common.Rwmap.Sshlist[host] = append(common.Rwmap.Sshlist[host], true)
			}
		}
		//session.Close()
	}
	common.Sshwg.Done()
}
