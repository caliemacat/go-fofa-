package fofaview

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

var res [][]string

// fofaviewer 导出来的是xlsx文件 解析它 然后提取ip正常还是到文件里面
func Fofaview_read(address string) {
	if xlFile, err := xlsx.OpenFile(address); err == nil {
		for index, sheet := range xlFile.Sheets {
			//第一个sheet
			if index == 0 {
				temp := make([][]string, len(sheet.Rows))
				for k, row := range sheet.Rows {
					var data []string
					for _, cell := range row.Cells {
						data = append(data, cell.Value)
					}
					temp[k] = data
				}
				res = append(res, temp...)
			}

		}
	} else {
		fmt.Println("错误")
		return
	}
	//fmt.Println(res)
	file1, err := os.OpenFile("port_open.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer file1.Close()
	for _, data := range res[2:] {
		tmp := []byte(data[0])
		_, err1 := file1.Write(tmp)
		_, err2 := file1.WriteString("\n")
		if err1 != nil {
			fmt.Println("写入文件失败")
			return
		}
		if err2 != nil {
			fmt.Println("写入换行失败")
			return
		}
	}

}
