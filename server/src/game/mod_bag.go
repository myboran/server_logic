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

func (self *ModBag) AddItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		// 不存在的物品
		fmt.Println("物品不存在 id: ", itemId)
		return
	}

	switch itemConfig.SortType {
	//case csvs.ITEMTYPE_NORMAL:
	//	fmt.Println("普通物品: ", itemConfig.ItemName)
	//
	//	self.AddItemToBag(itemId, 1)
	case csvs.ItemTypeRole:
		fmt.Println("角色: ", itemConfig.ItemName)

		player.ModRole.AddItem(itemId, num, player)
	case csvs.ItemTypeIcon:
		fmt.Println("头像: ", itemConfig.ItemName)

		player.ModIcon.AddItem(itemId)
	case csvs.ItemTypeCard:
		fmt.Println("名片: ", itemConfig.ItemName)

		player.ModCard.AddItem(itemId, 9)
	case csvs.ItemTypeWeapon:
		fmt.Println("武器: ", itemConfig.ItemName)

		player.ModWeapon.addItem(itemId, num)

	default: // 同普通物品
		fmt.Println("普通物品: ", itemConfig.ItemName)

		self.AddItemToBag(itemId, num)
	}
}

func (self *ModBag) AddItemToBag(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum += num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得物品:", config.ItemName, "数量:", num, "总数:", self.BagInfo[itemId].ItemNum)
	}
}

func (self *ModBag) RemoveItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		// 不存在的物品
		fmt.Println("物品不存在 id:", itemId)
		return
	}

	switch itemConfig.SortType {
	case csvs.ItemTypeNormal:
		fmt.Println("普通物品:", itemConfig.ItemName)

		self.RemoveItemToBag(itemId, num, player)
	default:
		//self.RemoveItemToBag(itemId, 1)
	}
}

func (self *ModBag) RemoveItemToBagGM(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum -= num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: 0 - num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("减少物品:", config.ItemName, "数量:", num, "总数:", self.BagInfo[itemId].ItemNum)
	}
}

func (self *ModBag) RemoveItemToBag(itemId int, num int64, player *Player) {

	_, ok := self.BagInfo[itemId]

	if ok {
		if !self.HasEnoughItem(itemId, num) {
			config := csvs.GetItemConfig(itemId)
			if config != nil {
				fmt.Println(config.ItemName, "数量不足, 当前数量: ", self.BagInfo[itemId].ItemNum)

				return
			}
		}

		self.BagInfo[itemId].ItemNum -= num
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			fmt.Println("减少物品:", config.ItemName, "数量:", num, "当前数量:", self.BagInfo[itemId].ItemNum)

			return
		}
	} else {
		fmt.Println("没有该商品: ", itemId)

		return
	}
}

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {
	return self.BagInfo[itemId].ItemNum > num
}
