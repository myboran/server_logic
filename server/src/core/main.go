package main

import (
	"encoding/json"
	"fmt"
	"server_logic/server/src/csvs"
	"server_logic/server/src/game"
	"time"
)

func main() {

	// 加载配置 TODO

	csvs.CheckLoadCsv()

	fmt.Println("数据测试-----start")

	go game.GetManageBanWord().Run()

	play := game.NewTestPlayer()

	play.RecvSetIcon(0)
	play.RecvSetCard(1)
	play.RecvSetName("张三")
	play.RecvSetSign("张三真厉害")
	play.RecvSetName("张原")

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%3 == 0 {
				play.RecvSetName("专业代练")
			} else if time.Now().Unix()%5 == 0 {
				play.RecvSetName("正常名字")
			}
		}
	}

	//GetInfo(play)
}

func GetInfo(play *game.Player) {
	bstr, err := json.Marshal(play)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bstr))
}
