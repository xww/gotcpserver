package main

import (
	"fmt"
	"net"
	"time"
	"encoding/binary"
	//"bytes"
)

const (
	ip = ""
	port = 3333
)

func main() {

	listen,err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接...")
	Server(listen)
}

func Server(listen *net.TCPListener) {
	for {
		//listen.SetDeadline(time.Now().Add(2))
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				fmt.Println("receive data len is : ",i)
				fmt.Print (time.Unix(time.Now().Unix(), 0).String())
				//fmt.Println("  客户端发来数据:", string(data[0:i]))
				fmt.Println("客户端发来数据:", data[0:i])
				var a uint16
				//binary.Read(bytes.NewBuffer(data),binary.BigEndian,&a)
				a=binary.BigEndian.Uint16(data[0:2])
				fmt.Println(a)
				var b uint8
				//binary.Read(bytes.NewBuffer(data),binary.BigEndian,&b)
				b=data[2]
				fmt.Println(b)
				var c uint32
				//binary.Read(bytes.NewBuffer(data),binary.BigEndian,&c)
				c=binary.BigEndian.Uint32(data[3:7])
				fmt.Println(c)
				if err != nil {
					fmt.Println("读取客户端数据错误:", err.Error())
					break
				}
				//conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
			}

		}()
	}
}
