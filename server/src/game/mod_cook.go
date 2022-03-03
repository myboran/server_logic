package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type Cook struct {
	CookId int
}

type ModCook struct {
	CookInfo map[int]*Cook
}

func (self *ModCook) addItem(cookId int) {
	_, ok := self.CookInfo[cookId]
	if ok {
		fmt.Println("已学习过: ", csvs.GetItemConfig(cookId).ItemName)
		return
	}

	config := csvs.GetCookConfig(cookId)
	if config == nil {
		fmt.Println("非法食谱id: ", cookId)
		return
	}

	self.CookInfo[cookId] = &Cook{
		CookId: cookId,
	}
	fmt.Println("学会了新食谱: ", csvs.GetItemConfig(cookId).ItemName)
}
