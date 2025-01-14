package test_alive

import (
	"bufio"
	"fmt"
	"github.com/imroc/req/v3"
	"os"
	"sync"
	"time"
)

var ch = make(chan string, 1)
var ok_proxies []string
var wg sync.WaitGroup

func Task_alive() {
	defer wg.Done()
	scan := <-ch
	client := req.C().SetProxyURL("socks5://" + scan).SetTimeout(2 * time.Second)
	resp, err := client.R().Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println(scan, "连接不通")
		return
	}
	if resp.IsSuccessState() {
		fmt.Println(scan, "成功")
		ok_proxies = append(ok_proxies, scan)
	}
}

func Test_alive() {
	file1, err := os.OpenFile("port_success.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		scan := scanner.Text()
		ch <- scan
		wg.Add(1)
		go Task_alive()
	}
	wg.Wait()
	file2, err2 := os.OpenFile("ok_open.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err2 != nil {
		fmt.Println("创建文件失败")
		return
	}
	for _, proxy := range ok_proxies {
		tmp := []byte(proxy)
		_, err3 := file2.Write(tmp)
		_, err4 := file2.WriteString("\n")
		if err3 != nil {
			fmt.Println("写入文件失败")
		}
		if err4 != nil {
			fmt.Println("写入文件失败")
		}
	}
}
