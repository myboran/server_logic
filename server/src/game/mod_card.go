package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type Card struct {
	CardId int
}

type ModCard struct {
	CardInfo map[int]*Card
}

func (self *ModCard) IsHasCard(cardId int) bool {
	_, ok := self.CardInfo[cardId]
	if ok {
		return true
	}
	return false
}

func (self *ModCard) AddItem(cardId int, friendliness int) {
	_, ok := self.CardInfo[cardId]
	if ok {
		fmt.Println("已经存在此名片")
		return
	}

	// 验证配置的有效性
	config := csvs.GetCardConfig(cardId)
	if config == nil {
		fmt.Println("非法名片: ", cardId)
		return
	}

	if friendliness < config.Friendliness {
		fmt.Println("好感度不足: ", cardId)
		return
	}

	self.CardInfo[cardId] = &Card{
		CardId: cardId,
	}
	item := csvs.GetItemConfig(cardId)
	fmt.Println("获得名片: ", item.ItemName)
}

func (self *ModCard) CheckGetCard(roleId int, friendliness int) {
	config := csvs.GetCardConfigByRoleId(roleId)
	if config == nil {
		fmt.Println("没找到名片")
		return
	}
	self.AddItem(config.CardId, friendliness)
}
