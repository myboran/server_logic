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

	playerGM := game.NewTestPlayer()

	playerGM.ModPlayer.AddExp(10000000, playerGM)

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%3 == 0 {
				playerGM.ReduceWorldLevel()
			} else if time.Now().Unix()%5 == 0 {
				playerGM.ReturnWorldLevel()
			}
		}
	}

	//GetInfo(playerGM)
}

func GetInfo(play *game.Player) {
	bstr, err := json.Marshal(play)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bstr))
}
