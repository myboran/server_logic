package utils

import (
	"fmt"
	"github.com/zhnxin/csvreader"
	"reflect"
)

func GetCsvUtilMgrSlice(fileName string, bean interface{}) {
	//准备读取文件
	if err := csvreader.New().UnMarshalFile("./server/utils/"+fileName+".csv", bean); err != nil {
		panic(err)
	}
}

func GetCsvUtilMgr(fileName string, bean interface{}) {
	//准备读取文件
	beanT := reflect.ValueOf(bean)

	fmt.Println(beanT.Type())
	if err := csvreader.New().UnMarshalFile("./server/utils/"+fileName+".csv", bean); err != nil {
		panic(err)
	}
}
