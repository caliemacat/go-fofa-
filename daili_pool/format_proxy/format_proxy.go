package format_proxy

import (
	"bufio"
	"fmt"
	"os"
)

var socks5_proxy []string

func Format_proxy(address string) {
	file1, err := os.OpenFile(address, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	scanner := bufio.NewScanner(file1)
	file2, err2 := os.OpenFile("format-socks5.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err2 != nil {
		fmt.Println("创建文件失败")
		return
	}
	for scanner.Scan() {
		tmp1 := scanner.Text()
		tmp2 := "socks5://" + tmp1
		_, err3 := file2.Write([]byte(tmp2))
		_, err4 := file2.WriteString("\n")
		if err3 != nil {
			fmt.Println("写入错误")
			return
		}
		if err4 != nil {
			fmt.Println("写入错误")
			return
		}
	}
}
