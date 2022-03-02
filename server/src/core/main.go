package main

import (
	"fmt"
	"server_logic/server/src/csvs"
	"server_logic/server/src/game"
)

func main() {

	// 加载配置 TODO
	csvs.CheckLoadCsv()

	go game.GetManageBanWord().Run()

	playerGM := game.NewTestPlayer()
	//fmt.Println("---修改头像---")
	//playerGM.ModPlayer.SetIcon(3000001, playerGM)
	//playerGM.ModBag.AddItem(3000001,2,  playerGM)
	//playerGM.ModBag.AddItem(3000002,2,  playerGM)
	//playerGM.ModPlayer.SetIcon(3000001, playerGM)
	//playerGM.ModPlayer.SetIcon(3000002, playerGM)
	////playerGM.ModBag.AddItem(1000005, playerGM)
	//fmt.Println("---修改名片---")
	//playerGM.ModPlayer.SetCard(4000001, playerGM)
	//playerGM.ModBag.AddItem(4000001, 2, playerGM)
	//playerGM.ModPlayer.SetCard(4000001, playerGM)
	//fmt.Println("---背包---")
	//playerGM.ModBag.RemoveItemToBagGM(1000014, 15)
	//playerGM.ModBag.AddItemToBag(1000014, 15)
	//playerGM.ModBag.RemoveItemToBagGM(1000014, 13)
	//playerGM.ModBag.AddItemToBag(1000015, 15)
	//playerGM.ModBag.AddItemToBag(1000015, 18)
	//fmt.Println("-")
	//playerGM.ModBag.RemoveItemToBag(1000016, 10)
	//playerGM.ModBag.AddItemToBag(1000016, 10)
	//playerGM.ModBag.RemoveItemToBag(1000016, 11)
	//playerGM.ModBag.RemoveItemToBag(1000016, 7)
	//fmt.Println("---角色---")
	//playerGM.ModRole.AddItem(2000017, 1)
	//playerGM.ModRole.AddItem(2000017, 2)
	//playerGM.ModRole.AddItem(2000017, 3)
	//playerGM.ModRole.AddItem(2000017, 4)
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
	playerGM.Run()
	fmt.Println("hello world")
	return
}
