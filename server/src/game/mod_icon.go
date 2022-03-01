package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type Icon struct {
	Id int
}

type ModIcon struct {
	IconInfo map[int]*Icon
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := self.IconInfo[iconId]
	if ok {
		return true
	}
	return false
}

func (self *ModIcon) AddItem(iconId int) {
	_, ok := self.IconInfo[iconId]
	if ok {
		fmt.Println("已经存在此头像")
		return
	}

	// 验证配置的有效性
	config := csvs.GetIconConfig(iconId)
	if config == nil {
		fmt.Println("非法头像: ", iconId)
	}

	self.IconInfo[iconId] = &Icon{
		Id: iconId,
	}
	fmt.Println("获得头像: ", iconId)
}
