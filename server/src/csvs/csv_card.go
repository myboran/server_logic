package csvs

import (
	"fmt"
	"server_logic/server/utils"
)

type ConfigCard struct {
	CardId       int `json:"CardId"`
	Friendliness int `json:"Friendliness"`
	Check        int `json:"Check"`
}

var (
	ConfigCardMap         map[int]*ConfigCard
	ConfigCardMapByRoleId map[int]*ConfigCard
)

func init() {

	ConfigCardMap = make(map[int]*ConfigCard)
	utils.GetCsvUtilMgr().LoadCsv("Card", &ConfigCardMap)

	//for _, v := range ConfigCardMap {
	//	ConfigCardMapByRoleId[v.Check] = v
	//}
	fmt.Println("csv_card初始化")
	return
}

func GetCardConfig(cardId int) *ConfigCard {
	return ConfigCardMap[cardId]
}
