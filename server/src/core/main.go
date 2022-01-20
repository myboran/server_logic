package main

import (
	"fmt"
	"server_logic/server/src/game"
)

func main() {
	fmt.Println("数据测试-----start")
	play := game.NewTestPlayer()

	play.RecvSetIcon(0)
	play.RecvSetCard(1)

	fmt.Println(play.ModPlayer.Icon)
	fmt.Println(play.ModPlayer.Card)
}
