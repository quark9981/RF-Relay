package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var pipeFile = "/tmp/piperx.ipc"

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
	go hackrfrx()
	file, _ := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
	conn, err := net.Dial("tcp", ipport)
	if err != nil {
		fmt.Println("net.Dail err", err)
		return
	}
	defer conn.Close()

	for {
		conn.Write(read(file))
		time.Sleep(time.Microsecond * 10)
	}

}
func read(file *os.File) []byte {

	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print("load string: false")

	}
	return line
}

func hackrfrx() {
	time.Sleep(time.Microsecond * 2000)
	cmd := exec.Command("sudo", "hackrf_transfer", "-r", "/tmp/piperx.ipc", "-f", "315000000", "-g", "30", "-l", "24", "-a", "1", "-p", "1", "-s", "8000000", "-b", "4000000")
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println("test", err1.Error())
	} else {
		fmt.Println("RX OK")
	}

}
