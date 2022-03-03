package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type HomeItem struct {
	HomeItemId  int
	HomeItemNum int64
	KeyId       int
}

type ModHome struct {
	HomeItemIdInfo map[int]*HomeItem
}

func (self *ModHome) addItem(itemId int, num int64, player *Player) {
	_, ok := self.HomeItemIdInfo[itemId]
	if ok {
		self.HomeItemIdInfo[itemId].HomeItemNum += num
	} else {
		self.HomeItemIdInfo[itemId] = &HomeItem{
			HomeItemId:  itemId,
			HomeItemNum: num,
		}
	}

	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得家园物品: ", config.ItemName, "数量:", num, "当前总数:", self.HomeItemIdInfo[itemId].HomeItemNum)
	}
}
