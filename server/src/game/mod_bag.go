package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

type ModBag struct {
	BagInfo map[int]*ItemInfo
}

func (self *ModBag) AddItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		// 不存在的物品
		fmt.Println("物品不存在 id: ", itemId)
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		fmt.Println("普通物品: ", itemConfig.ItemName)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色: ", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		fmt.Println("头像: ", itemConfig.ItemName)

		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		fmt.Println("名片: ", itemConfig.ItemName)
		player.ModCard.AddItem(itemId, 9)
	}
}
