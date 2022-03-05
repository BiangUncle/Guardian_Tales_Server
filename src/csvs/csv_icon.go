package csvs

import "server/src/utils"

type IconConfig struct {
	HeroId int `json:"HeroId"`
	IconId int `json:"IconId"`
}

var IconConfigMap map[int]*IconConfig

func init() {
	IconConfigMap = make(map[int]*IconConfig)
	utils.GetCsvUtilMgr().LoadCsv("icon", &IconConfigMap)
	return
}

func GetIconConfig(heroId int) *IconConfig {
	return IconConfigMap[heroId]
}
