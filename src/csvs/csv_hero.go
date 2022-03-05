package csvs

import "server/src/utils"

const (
	TYPE_BASIC = 1
	TYPE_LIGHT = 2
	TYPE_DARK  = 3
	TYPE_WATER = 4
	TYPE_FIRE  = 5
	TYPE_EARTH = 6

	HERO_TYPE_WARRIOR = 1
	HERO_TYPE_RANGED  = 2
	HERO_TYPE_TANKER  = 3
	HERO_TYPE_SUPPORT = 4
)

type HeroConfig struct {
	HeroId   int    `json:"HeroId"`
	NickName string `json:"NickName"`
	HeroName string `json:"HeroName"`
	HeroType int    `json:"HeroType"`
	Star     int    `json:"Star"`
	Type     int    `json:"Type"`
}

var HeroConfigMap map[int]*HeroConfig

func init() {
	HeroConfigMap = make(map[int]*HeroConfig)
	utils.GetCsvUtilMgr().LoadCsv("hero", &HeroConfigMap)
	return
}

func GetHeroConfig(heroId int) *HeroConfig {
	heroInfo, ok := HeroConfigMap[heroId]
	if !ok {
		return nil
	}
	return heroInfo
}
