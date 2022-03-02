package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type RoleInfo struct {
	RoleId   int
	GetTimes int
	// 等级 经验 圣遗物
}

type ModRole struct {
	RoleInfo map[int]*RoleInfo
}

func (self *ModRole) IsHasRole(roleId int) bool {
	return true
}

func (self *ModRole) GetRoleLevel(roleId int) int {
	return 80
}

func (self *ModRole) AddItem(roleId int, num int64, player *Player) {
	itemConfig := csvs.GetRoleConfig(roleId)

	for i := 0; i < int(num); i++ {
		_, ok := self.RoleInfo[roleId]
		if !ok {
			fmt.Println("获得实际物品---", itemConfig.ItemName)
			self.RoleInfo[roleId] = &RoleInfo{
				RoleId:   roleId,
				GetTimes: 1,
			}
			player.ModIcon.CheckGetIcon(roleId)
			player.ModCard.CheckGetCard(roleId, 10)
		} else {
			// 判定实际获得的东西
			self.RoleInfo[roleId].GetTimes++
			if self.RoleInfo[roleId].GetTimes >= csvs.ADD_ROLE_TIME_NORMAL_MIN &&
				self.RoleInfo[roleId].GetTimes <= csvs.ADD_ROLE_TIME_NORMAL_MAX {
				player.ModBag.AddItemToBag(itemConfig.Stuff, itemConfig.StuffNum)
				player.ModBag.AddItemToBag(itemConfig.StuffItem, itemConfig.StuffItemNum)
			} else {
				player.ModBag.AddItemToBag(itemConfig.MaxStuffItem, itemConfig.MaxStuffItemNum)
			}
		}
	}
}
