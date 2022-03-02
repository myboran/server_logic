package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type Weapon struct {
	WeaponId int
	KeyId    int
}

type ModWeapon struct {
	WeaponInfo map[int]*Weapon
	MaxKey     int
}

func (self *ModWeapon) addItem(itemId int, num int64) {
	// 容量 2000
	if len(self.WeaponInfo)+int(num) > csvs.WeaponMaxCount {
		fmt.Println("武器已经到达上限了")
		return
	}

	for i := int64(0); i < num; i++ {
		weapon := new(Weapon)
		weapon.WeaponId = itemId
		self.MaxKey++
		weapon.KeyId = self.MaxKey

		self.WeaponInfo[weapon.KeyId] = weapon
		fmt.Println("获得武器 ", csvs.GetItemConfig(itemId).ItemName, " keyId", weapon.KeyId)
	}
}
