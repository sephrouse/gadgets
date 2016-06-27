//copyright 2016. all rights reserved.
//programmed by sephrouse.
//copy some functions from https://github.com/paulstuart/ping/blob/master/ping.go
//transfer this program to my gadgets reponsitory.

package main 

import (
    "fmt"
    "time"
    "strconv"
    "os"
    "bytes"
    "net"
)

const (
	icmpv4EchoRequest = 8
	icmpv4EchoReply   = 0
	icmpv6EchoRequest = 128
	icmpv6EchoReply   = 129
)

type icmpMessage struct {
	Type     int             // type
	Code     int             // code
	Checksum int             // checksum
	Body     icmpMessageBody // body
}

type icmpMessageBody interface {
	Len() int
	Marshal() ([]byte, error)
}

var fin chan string

// Marshal returns the binary enconding of the ICMP echo request or
// reply message m.
func (m *icmpMessage) Marshal() ([]byte, error) {
	b := []byte{byte(m.Type), byte(m.Code), 0, 0}
	if m.Body != nil && m.Body.Len() != 0 {
		mb, err := m.Body.Marshal()
		if err != nil {
			return nil, err
		}
		b = append(b, mb...)
	}
	switch m.Type {
	case icmpv6EchoRequest, icmpv6EchoReply:
		return b, nil
	}
	csumcv := len(b) - 1 // checksum coverage
	s := uint32(0)
	for i := 0; i < csumcv; i += 2 {
		s += uint32(b[i+1])<<8 | uint32(b[i])
	}
	if csumcv&1 == 0 {
		s += uint32(b[csumcv])
	}
	s = s>>16 + s&0xffff
	s = s + s>>16
	// Place checksum back in header; using ^= avoids the
	// assumption the checksum bytes are zero.
	b[2] ^= byte(^s & 0xff)
	b[3] ^= byte(^s >> 8)
	return b, nil
}

// imcpEcho represenets an ICMP echo request or reply message body.
type icmpEcho struct {
	ID   int    // identifier
	Seq  int    // sequence number
	Data []byte // data
}

func (p *icmpEcho) Len() int {
	if p == nil {
		return 0
	}
	return 4 + len(p.Data)
}

// Marshal returns the binary enconding of the ICMP echo request or
// reply message body p.
func (p *icmpEcho) Marshal() ([]byte, error) {
	b := make([]byte, 4+len(p.Data))
	b[0], b[1] = byte(p.ID>>8), byte(p.ID&0xff)
	b[2], b[3] = byte(p.Seq>>8), byte(p.Seq&0xff)
	copy(b[4:], p.Data)
	return b, nil
}

func Ping(address string, timeout int) bool {
	err := Pinger(address, timeout)
	return err == nil
}

func Pinger(address string, timeout int) error {
	c, err := net.Dial("ip4:icmp", address)
	if err != nil {
		return err
	}
	c.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	defer c.Close()

	typ := icmpv4EchoRequest
	xid, xseq := os.Getpid()&0xffff, 1
	wb, err := (&icmpMessage{
		Type: typ, Code: 0,
		Body: &icmpEcho{
			ID: xid, Seq: xseq,
			Data: bytes.Repeat([]byte("Go Go Gadget Ping!!!"), 1),
		},
	}).Marshal()
	if err != nil {
		return err
	}
	if _, err = c.Write(wb); err != nil {
        fmt.Println("Pinger Error")
		return err
	}
	return nil
}

func doPing(host string, timeout int){
    Ping(host, timeout)
}

func worker(beginIp string) {
    var cntEverySecond int = 0
    var ipaddr string
    var startIpaddr string=beginIp+".0.0.1"
    var ts=time.Now().Second()
    for i1:=0;i1<=255;i1++ {
        for i2:=0;i2<=255;i2++ {
            for i3:=1;i3<255;i3++ {
                ipaddr=beginIp+"."+strconv.Itoa(i1)+"."+strconv.Itoa(i2)+"."+strconv.Itoa(i3)
                cntEverySecond++

                doPing(ipaddr,5)
                
                tn:=time.Now().Second()
                if(ts!=tn) {
                    ts=tn
                    fmt.Println(time.Now(), " ", cntEverySecond, "ips have been pinged. from ",startIpaddr, "to ", ipaddr)
                    startIpaddr=ipaddr
                    cntEverySecond=0
                }
            }
        }
    }
    
    fin<-beginIp
}


func main()  {
    var workerNum int = 2
    if(len(os.Args)<2) {
        fmt.Println("starting. default worker number is 2.")
    }else{
        wNum, err:=strconv.Atoi(os.Args[1])
        fmt.Println("starting. the worker number is ", wNum)
        if(err != nil) {
            fmt.Println(err)
        }
        workerNum=wNum
        if(workerNum>100) {
            fmt.Println("too many workers. stop working.")
            return
        }
    }
    
    for i:=1;i<=workerNum;i++ {
        go worker(strconv.Itoa(i))
    }

    for j:=1;j<=workerNum;j++ {
        <-fin
    }
}