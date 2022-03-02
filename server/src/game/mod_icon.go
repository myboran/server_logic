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
		return
	}

	self.IconInfo[iconId] = &Icon{
		Id: iconId,
	}
	item := csvs.GetItemConfig(iconId)
	fmt.Println("获得头像: ", item.ItemName)
}

func (self *ModIcon) CheckGetIcon(roleId int) {
	config := csvs.GetIconConfigByRoleId(roleId)
	if config == nil {
		fmt.Println("没找到头像")
		return
	}
	self.AddItem(config.IconId)
}
