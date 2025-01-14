package main

import (
	"daili_pool/fofaview"
	"daili_pool/format_proxy"
	"daili_pool/proxy_scan"
	"daili_pool/test_alive"
	"daili_pool/test_proxy"
	"flag"
	"fmt"
)

var (
	Fofa   string
	Test   bool
	Alive  bool
	help   bool
	View   string
	Size   int
	Format string
)

func main() {
	//如果有fofa F点 可以再这查询
	//proxy_scan.Proxy_scan()
	//fofaview.Fofaview_read()
	//test_proxy.Test_proxy()
	//test_alive.Test_alive()
	flag.StringVar(&Fofa, "key", "", "")
	flag.StringVar(&View, "extra", "", "")
	flag.BoolVar(&Test, "test", false, "")
	flag.BoolVar(&Alive, "alive", false, "")
	flag.BoolVar(&help, "h", false, "")
	flag.IntVar(&Size, "s", 5000, "")
	flag.StringVar(&Format, "format", "", "")
	flag.Usage = func() {
		fmt.Println(`
		Usage:
			-key 传入fofa密钥参数  -s 指定查找的数量 默认5000条
			-extra 无F点 使用fofa-view 导出xlsx文件 然后使用这个提取 传入xlsx文件路径
			-test 测试端口是否开放 (先测试端口是否开放 再-alive 测试代理是否能使用)
			-alive 测试代理是否能使用访问
			-format 用于将测试成功的代理 转化为格式 可以使用这位师傅工具 将代理 转化为对应的格式 传入相关文件的地址
			(https://github.com/honmashironeko/ProxyCat/blob/main/ProxyCat-Manual/Operation%20Manual.md)
			-h 帮助
		Options:
			[-key]  [-test] ([-alive]) -s
			[-extra]  [-test] ([-alive])
			[-format]
`)
	}
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if Fofa != "" {
		proxy_scan.Proxy_scan(Fofa, Size)
	} else if View != "" {
		fofaview.Fofaview_read(View)
	} else if Test {
		test_proxy.Test_proxy()
	} else if Alive {
		test_alive.Test_alive()
	} else if Format != "" {
		format_proxy.Format_proxy(Format)
	}

}
