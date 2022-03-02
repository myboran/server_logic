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

func (self *ModRole) AddItem(roleId int, num int) {
	for i := 0; i < num; i++ {
		_, ok := self.RoleInfo[roleId]
		if !ok {
			self.RoleInfo[roleId] = &RoleInfo{
				RoleId:   roleId,
				GetTimes: 1,
			}
		} else {
			// 判定实际获得的东西
			fmt.Println("获得实际物品-------")
			self.RoleInfo[roleId].GetTimes++
		}
	}
	itemConfig := csvs.GetItemConfig(roleId)
	if itemConfig != nil {
		fmt.Println("获得", itemConfig.ItemName, "-------", self.RoleInfo[roleId].GetTimes, "次")
	}
}
