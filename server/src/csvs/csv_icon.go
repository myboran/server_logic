package csvs

import (
	"fmt"
	"server_logic/server/utils"
)

type ConfigIcon struct {
	IconId int `json:"IconId"`
	Check  int `json:"Check"`
}

var (
	ConfigIconMap         map[int]*ConfigIcon
	ConfigIconMapByRoleId map[int]*ConfigIcon
)

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)
	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)

	fmt.Println("csv_icon初始化")
	return
}

func GetIconConfig(itemId int) *ConfigIcon {
	return ConfigIconMap[itemId]
}
