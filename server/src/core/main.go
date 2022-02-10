package main

import (
	"encoding/json"
	"fmt"
	"server_logic/server/src/csvs"
	"server_logic/server/src/game"
)

func main() {

	// 加载配置 TODO

	csvs.CheckLoadCsv()

	fmt.Println("数据测试-----start")

	go game.GetManageBanWord().Run()

	playerGM := game.NewTestPlayer()

	playerGM.ModPlayer.AddExp(10000000, playerGM)

	//ticker := time.NewTicker(time.Second * 3)
	//for {
	//	select {
	//	case <-ticker.C:
	//		fmt.Println("加了 2000 经验")
	//		playerGM.ModPlayer.AddExp(2000)
	//	}
	//}
	for {
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
