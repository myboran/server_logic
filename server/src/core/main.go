package main

import (
	"encoding/json"
	"fmt"
	"server_logic/server/src/game"
)

func main() {
	fmt.Println("数据测试-----start")
	play := game.NewTestPlayer()

	play.RecvSetIcon(0)
	play.RecvSetCard(1)
	play.RecvSetName("张三")
	play.RecvSetSign("张三真厉害")
	play.RecvSetName("张原")

	bstr, err := json.Marshal(play)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bstr))
}
