package main

import (
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
	playerGM.ModBag.AddItem(1000005)
	playerGM.ModBag.AddItem(2000005)
	playerGM.ModBag.AddItem(3000005)
	playerGM.ModBag.AddItem(4000005)
	playerGM.ModBag.AddItem(2000085)
	//playerGM.ModPlayer.AddExp(10000000, playerGM)
	//playerGM.SetBirth(2228)
	//playerGM.SetBirth(228)
	//playerGM.SetBirth(1228)
	//
	//playerGM.SetShowCard([]int{100, 200, 200, 300})
	//playerGM.ModPlayer.SetShowTeam([]int{1001, 1002}, playerGM)
	//
	//ticker := time.NewTicker(time.Second * 1)
	//for {
	//	select {
	//	case <-ticker.C:
	//		if time.Now().Unix()%3 == 0 {
	//			playerGM.ReduceWorldLevel()
	//		} else if time.Now().Unix()%5 == 0 {
	//			playerGM.ReturnWorldLevel()
	//		}
	//	}
	//}
	//ticker := time.NewTicker(time.Second * 10)
	//for {
	//	select {
	//	case <-ticker.C:
	//		player := game.NewTestPlayer()
	//		go player.Run()
	//	}
	//
	//}

	// 背包系统:   增加 角色模块
	// 1 物品识别
	// 2 物品增加     权利下放概念
	// 3 物品消耗
	// 4 物品使用
	// 5 角色模块--->头像模块

	fmt.Println("hello world")
	return
}
