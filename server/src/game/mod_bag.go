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
	case csvs.ItemTypeNormal:
		fmt.Println("普通物品: ", itemConfig.ItemName)

		self.AddItemToBag(itemId, 1)
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
	case csvs.ItemTypeRelics:
		fmt.Println("圣遗物: ", itemConfig.ItemName)

		player.ModRelics.addItem(itemId, num)
	case csvs.ItemTypeCook:
		fmt.Println("学习食谱: ", itemConfig.ItemName)

		player.ModCook.addItem(itemId)

	case csvs.ItemTypeCookbook:
		fmt.Println("获得食谱: ", itemConfig.ItemName)

		self.AddItemToBag(itemId, 1)

	case csvs.ItemTypeHomeItem:
		fmt.Println("获得家具: ", itemConfig.ItemName)

		player.ModHome.addItem(itemId, num, player)
	default:
		fmt.Println("无法识别此物品")
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
		fmt.Println("rm普通物品:", itemConfig.ItemName)

		self.RemoveItemToBag(itemId, num, player)
	case csvs.ItemTypeCookbook:
		fmt.Println("rm食谱:", itemConfig.ItemName)
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

	config := csvs.GetItemConfig(itemId)
	if config == nil {
		fmt.Println("没有此物品")
		return
	}
	_, ok := self.BagInfo[itemId]

	if ok {
		if !self.HasEnoughItem(itemId, num) {
			fmt.Println(config.ItemName, "数量不足, 当前数量: ", self.BagInfo[itemId].ItemNum)
			return
		}
		self.BagInfo[itemId].ItemNum -= num
		fmt.Println("减少物品:", config.ItemName, "数量:", num, "当前数量:", self.BagInfo[itemId].ItemNum)
		return
	} else {
		fmt.Println("没有该商品: ", itemId)

		return
	}
}

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {
	return self.BagInfo[itemId].ItemNum >= num
}

func (self *ModBag) useItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		// 不存在的物品
		fmt.Println("物品不存在 id: ", itemId)
		return
	}

	_, ok := self.BagInfo[itemId]
	if ok {
		if !self.HasEnoughItem(itemId, num) {
			fmt.Println(itemConfig.ItemName, "数量不足, 当前数量: ", self.BagInfo[itemId].ItemNum)

			return
		}
	} else {
		fmt.Println("没有该商品: ", itemId)

		return
	}

	switch itemConfig.SortType {
	case csvs.ItemTypeCookbook:
		fmt.Println("使用食谱: ", csvs.GetItemConfig(itemId).ItemName)
		self.useCookBook(itemId, num, player)
	case csvs.ItemTypeFood:
		fmt.Println("使用食物: ", csvs.GetItemConfig(itemId).ItemName)

	default:
		fmt.Println("此物品无法使用")
	}
}

func (self *ModBag) useCookBook(itemId int, num int64, player *Player) {
	cookBookConfig := csvs.GetCookBookInfo(itemId)
	if cookBookConfig == nil {
		fmt.Println("没有该物品")
		return
	}

	self.RemoveItem(itemId, num, player)
	player.ModCook.addItem(cookBookConfig.Reward)
}
