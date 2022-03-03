package csvs

import "server_logic/server/utils"

type ConfigHomeItem struct {
	HomeItemId int `json:"HomeItemId"`
	Type       int `json:"Type"`
}

var (
	ConfigHomeItemMap map[int]*ConfigHomeItem
)

func init() {
	ConfigHomeItemMap = make(map[int]*ConfigHomeItem)
	utils.GetCsvUtilMgr().LoadCsv("Home", &ConfigHomeItemMap)
	return
}

func GetHomeItemConfig(homeItemId int) *ConfigHomeItem {
	return ConfigHomeItemMap[homeItemId]
}
