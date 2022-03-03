package csvs

import "server_logic/server/utils"

type ConfigCookBook struct {
	CookBookId int `json:"CookBookId"`
	Reward     int `json:"Reward"`
}

var (
	ConfigCookBookMap map[int]*ConfigCookBook
)

func init() {
	ConfigCookBookMap = make(map[int]*ConfigCookBook)
	utils.GetCsvUtilMgr().LoadCsv("CookBook", &ConfigCookBookMap)
	return
}

func GetCookBookInfo(cookBookId int) *ConfigCookBook {
	return ConfigCookBookMap[cookBookId]
}
