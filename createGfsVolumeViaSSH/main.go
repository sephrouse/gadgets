package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

const user = "root"
const pwd = "k8sgogogo"
const dstIP = "10.8.65.156:22"

func main() {
	action := os.Args[1]
	volumeName := os.Args[2]

	switch action {
	case "c":
		createGfsVolume(volumeName)
	case "d":
		deleteGfsVolume(volumeName)
	}

}

func deleteGfsVolume(vName string) error {
	pw := []ssh.AuthMethod{ssh.Password(pwd)}
	conf := ssh.ClientConfig{
		User: user,
		Auth: pw,
	}

	client, _ := ssh.Dial("tcp", dstIP, &conf)
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("deleteGfsVolume error. ", err)
		return err
	}

	defer session.Close()

	fmt.Println("prepare to execute cmd via ssh.")

	// var commonResult bytes.Buffer
	// var errorResult bytes.Buffer
	//
	// session.Stdout = &commonResult
	// session.Stderr = &errorResult

	cmd := fmt.Sprintf("/usr/sbin/gluster volume delete %s_vol", vName)
	fmt.Println("cmd is: ", cmd)

	stdinBuf, _ := session.StdinPipe()
	stdoutBuf, _ := session.StdoutPipe()

	session.Shell()
	stdinBuf.Write([]byte(cmd + "\n"))

	stdinBuf.Write([]byte("y\n"))

	done1 := make(chan bool, 1)
	var result []byte
	go func() {
		out1 := bufio.NewReader(stdoutBuf)
		result, _, _ = out1.ReadLine()
		done1 <- true
		fmt.Println("end of go func()")
	}()
	a := <-done1
	fmt.Println("done go func()", a, string(result))

	return nil
}

func createGfsVolume(vName string) error {
	pw := []ssh.AuthMethod{ssh.Password(pwd)}
	conf := ssh.ClientConfig{
		User: user,
		Auth: pw,
	}

	client, _ := ssh.Dial("tcp", dstIP, &conf)
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("createGfsVolume error. ", err)
		return err
	}

	defer session.Close()

	fmt.Println("prepare to execute cmd via ssh.")

	var commonResult bytes.Buffer
	var errorResult bytes.Buffer

	session.Stdout = &commonResult
	session.Stderr = &errorResult

	cmd := fmt.Sprintf("/usr/sbin/gluster volume create %s_vol replica 2 10.8.65.{156,157}:/home/zxc/gfsdata/%s_vol", vName, vName)
	fmt.Println("cmd is: ", cmd)

	err = session.Run(cmd)
	fmt.Println(err)

	fmt.Println("create common: ", commonResult.String())
	fmt.Println("create error: ", errorResult.String())

	return nil
}
