package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type Relics struct {
	RelicsId int
	KeyId    int
}

type ModRelics struct {
	RelicsInfo map[int]*Relics
	MaxKey     int
}

func (self *ModRelics) addItem(relicsId int, num int64) {
	config := csvs.GetRelicsConfig(relicsId)
	if config == nil {
		fmt.Println("圣遗物不存在")
		return
	}

	if len(self.RelicsInfo)+int(num) > 1500 {
		fmt.Println("圣遗物个数超出上限")
		return
	}

	for i := int64(0); i < num; i++ {
		relics := new(Relics)
		relics.RelicsId = relicsId
		self.MaxKey++
		relics.KeyId = self.MaxKey

		self.RelicsInfo[relics.KeyId] = relics
		fmt.Println("获得圣遗物 ", csvs.GetItemConfig(relicsId).ItemName, "keyId", relics.KeyId)
	}
}
