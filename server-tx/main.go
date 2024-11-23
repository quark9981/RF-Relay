package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var pipeFile = "/tmp/pipetx.ipc"

func main() {
	var ip string
	var port string
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&ip, "i", "xxx", "ip，默认为xxx")
	flag.StringVar(&port, "p", "xxx", "ip，默认为xxx")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	// 输出结果
	// fmt.Printf("licence：%v\n", licence)
	if ip == "xxx" || port == "xxx" {
		fmt.Println("ip或port输入错误")
		os.Exit(3)
	}
	ipport := ip + ":" + port
	os.Remove(pipeFile)
	err1 := syscall.Mkfifo(pipeFile, 0777)
	if err1 != nil {
		log.Fatal("create named pipe error:", err1)
	}
	go hackrftx()
	file, err2 := os.OpenFile(pipeFile, os.O_RDWR, 0777)
	if err2 != nil {
		log.Fatalf("error opening file: %v", err2)
	}

	listener, err4 := net.Listen("tcp", ipport)
	if err4 != nil {
		fmt.Println("net.Listen err4:", err4)
		return
	}
	defer listener.Close()

	conn, err5 := listener.Accept()
	if err5 != nil {
		fmt.Println("accept err5", err5)
		return
	}
	defer conn.Close()

	for {

		buf := make([]byte, 4096)
		n, err3 := conn.Read(buf)
		if err3 != nil {
			fmt.Println("conn Read err3", err3)
			return
		}
		write(file, buf[:n])
		time.Sleep(time.Microsecond * 10)
	}

}
func write(f *os.File, bytedata []byte) {
	f.Write(bytedata)
}
func hackrftx() {
	time.Sleep(time.Microsecond * 2000)
	cmd := exec.Command("sudo", "hackrf_transfer", "-t", "/tmp/pipetx.ipc", "-f", "433000000", "-x", "32", "-a", "1", "-p", "1", "-s", "8000000", "-b", "4000000")
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Println("TX OK")
	}

}
