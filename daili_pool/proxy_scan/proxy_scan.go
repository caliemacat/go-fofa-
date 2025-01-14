package proxy_scan

import (
	"fmt"
	"github.com/imroc/req/v3"
	"os"
)

type Res struct {
	Error           bool       `json:"error"`
	ConsumedFpoint  int        `json:"consumed_fpoint"`
	RequiredFpoints int        `json:"required_fpoints"`
	Size            int        `json:"size"`
	Tip             string     `json:"tip"`
	Page            int        `json:"page"`
	Mode            string     `json:"mode"`
	Query           string     `json:"query"`
	Results         [][]string `json:"results"`
}

var res Res

func Proxy_scan(key string, size int) {
	fmt.Println("开始爬取fofa数据")
	client := req.C()
	//默认5000
	fofaurl := fmt.Sprintf("https://fofa.info/api/v1/search/all?key=%s&qbase64=cHJvdG9jb2w9PSJzb2NrczUiICYmICJWZXJzaW9uOjUgTWV0aG9kOk5vIEF1dGhlbnRpY2F0aW9uKDB4MDApIiAmJiBjb3VudHJ5PSJDTiI=&size=%d", key, size)
	//fofaurl := "https://fofa.info/api/v1/search/all?key=自己的key&qbase64=cHJvdG9jb2w9PSJzb2NrczUiICYmICJWZXJzaW9uOjUgTWV0aG9kOk5vIEF1dGhlbnRpY2F0aW9uKDB4MDApIiAmJiBjb3VudHJ5PSJDTiI=&size=10"
	fofaresp, err := client.
		R().
		SetSuccessResult(&res).
		Get(fofaurl)
	if err != nil {
		fmt.Println("fofa点数不够")
		return
	}
	file1, err1 := os.OpenFile("port_open.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err1 != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer file1.Close()

	if fofaresp.IsSuccessState() {
		for _, data := range res.Results {
			tmp := data[0]
			res1 := []byte(tmp)
			_, err2 := file1.Write(res1)
			_, err3 := file1.WriteString("\n")
			if err2 != nil {
				fmt.Println("写入文件失败")
				return
			}
			if err3 != nil {
				fmt.Println("写入文件失败")
				return
			}
		}
	}
	fmt.Println("爬取成功")

}
