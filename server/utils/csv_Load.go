package utils

import "github.com/zhnxin/csvreader"

func GetCsvUtilMgr(fileName string, bean interface{}) {
	//准备读取文件
	if err := csvreader.New().UnMarshalFile("./server/utils/"+fileName, bean); err != nil {
		panic(err)
	}

}
