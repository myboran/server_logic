package csvs

import "server_logic/server/utils"

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterId"`
}

var (
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

func init() {
	utils.GetCsvUtilMgr("PlayerLevel.csv", &ConfigPlayerLevelSlice)
	return
}

func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	if level < 0 || level >= len(ConfigPlayerLevelSlice) {
		return nil
	}
	return ConfigPlayerLevelSlice[level]
}
