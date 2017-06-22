package main

import (
	"fmt"
	"net"
	"time"
	//"strconv"
	//"os"

	"math/rand"
	//"encoding/binary"
	//"bytes"
	"encoding/binary"
	"bytes"
)

const (
	addr = "127.0.0.1:3333"
)

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max - min) + min
}

func test () []byte {
	//buffer := make([]byte,0)
	bb := bytes.NewBuffer(nil)
	binary.Write(bb,binary.BigEndian,uint16(1))
	binary.Write(bb,binary.BigEndian,uint8(2))
	binary.Write(bb,binary.BigEndian,uint32(101234))
	/*fmt.Println(bb.Bytes())
	var a uint8
	binary.Read(bb,binary.BigEndian,&a)
	fmt.Println(a)
	var b uint8
	binary.Read(bb,binary.BigEndian,&b)
	fmt.Println(b)
	var c uint16
	binary.Read(bb,binary.BigEndian,&c)
	fmt.Println(c)*/
	return bb.Bytes()
}
func main() {
	content := test()
	//fmt.Println(content)
	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100))
	//fmt.Println(os.Args[1])
	//b, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < 1; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("连接服务端失败:", err.Error())
			return
		}
		//fmt.Println("已连接服务器")
		defer conn.Close()
		go Client(conn, i,content)
	}
	select {}
}

func Client(conn net.Conn, i int,content []byte) {
	//sms := make([]byte, 128)

	for {
		//time.Sleep(time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)) * time.Second)
		time.Sleep(3*time.Second)
		fmt.Println(i)
		//conn.Write([]byte("i am " + strconv.Itoa(i) + " client."))
		conn.Write(content)
		fmt.Println("send data is: ",content)
		buff := bytes.NewBuffer(content)
		var a uint16
		binary.Read(buff,binary.BigEndian,&a)
		fmt.Println(a)
		var b uint8
		binary.Read(buff,binary.BigEndian,&b)
		fmt.Println(b)
		var c uint32
		binary.Read(buff,binary.BigEndian,&c)
		fmt.Println(c)

		//buf := make([]byte, 128)
		//_, err := conn.Read(buf)
		//if err != nil {
		//	fmt.Println("读取服务器数据异常:", err.Error())
		//}
		//fmt.Println(string(buf[0:c]))
	}

}
