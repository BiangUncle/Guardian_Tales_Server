package csvs

import "server/src/utils"

const (
	WEAPON_TYPE_DANSHOUJIAN    = 1
	WEAPON_TYPE_SHUANGSHOUJIAN = 2
	WEAPON_TYPE_BUQIANG        = 3
	WEAPON_TYPE_GONG           = 4
	WEAPON_TYPE_LANZI          = 5
	WEAPON_TYPE_QUANZHANG      = 6
	WEAPON_TYPE_SHOUTAO        = 7
	WEAPON_TYPE_DUNPAI         = 8
	WEAPON_TYPE_ZHUA           = 9
	WEAPON_TYPE_SHIPIN         = 10
	WEAPON_TYPE_QIANGHUA       = 11
)

type WeaponConfig struct {
	WeaponId   int    `json:"WeaponId"`
	WeaponName string `json:"WeaponName"`
	WeaponType int    `json:"WeaponType"`
	Star       int    `json:"Star"`
	Type       int    `json:"Type"`
}

var WeaponConfigMap map[int]*WeaponConfig

func init() {
	WeaponConfigMap = make(map[int]*WeaponConfig)
	utils.GetCsvUtilMgr().LoadCsv("weapon", &WeaponConfigMap)
	return
}

func GetWeaponConfig(weaponId int) *WeaponConfig {
	weaponConfig, ok := WeaponConfigMap[weaponId]
	if !ok {
		return nil
	}
	return weaponConfig
}
