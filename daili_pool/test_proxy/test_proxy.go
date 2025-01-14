package test_proxy

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch = make(chan string, 1)
var open_proxies []string

func Task_test() {
	defer wg.Done()
	scan := <-ch
	conn, err := net.DialTimeout("tcp", scan, time.Second*2)
	if err != nil {
		fmt.Printf("%s连接不通\n", scan)
		return
	}
	defer conn.Close()
	open_proxies = append(open_proxies, scan)
}

func Test_proxy() {
	file1, err := os.OpenFile("port_open.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("proxy文件不存在")
		return
	}
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		scan := scanner.Text()
		ch <- scan
		wg.Add(1)
		go Task_test()
	}
	wg.Wait()
	file2, err2 := os.OpenFile("port_success.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err2 != nil {
		fmt.Println("创建文件失败")
		return
	}
	for _, proxy := range open_proxies {
		tmp := []byte(proxy)
		_, err3 := file2.Write(tmp)
		_, err4 := file2.WriteString("\n")
		if err3 != nil {
			fmt.Println("写入文件失败")
			return
		}
		if err4 != nil {
			fmt.Println("写入文件失败")
			return
		}
	}
}
