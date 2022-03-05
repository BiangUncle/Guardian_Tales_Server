package game

import (
	"fmt"
	"server/src/csvs"
)

const (
	MODWEAPON_SYSTEM_INFO = "【武器模块】"

	MAX_WEAPON_NUM = 800
)

type WeaponInfo struct {
	WeaponId   int    `json:"WeaponId"`
	WeaponName string `json:"WeaponName"`
	WeaponType int    `json:"WeaponType"`
	Star       int    `json:"Star"`
	Type       int    `json:"Type"`
}

type ModWeapon struct {
	WeaponInfos map[int]*WeaponInfo
	key         int
}

func (self *ModWeapon) AddWeapon(weaponId int) {

	if len(self.WeaponInfos) > MAX_WEAPON_NUM {
		fmt.Println(MODWEAPON_SYSTEM_INFO, "背包已满")
	}

	weaponConfig := csvs.GetWeaponConfig(weaponId)
	if weaponConfig == nil {
		fmt.Println(MODWEAPON_SYSTEM_INFO, "非法武器Id")
	}

	self.key++
	weaponInfo := &WeaponInfo{
		WeaponId:weaponConfig.WeaponId,
		WeaponName:weaponConfig.WeaponName,
		WeaponType:weaponConfig.WeaponType,
		Star:weaponConfig.Star,
		Type:weaponConfig.Type,
	}
	self.WeaponInfos[self.key] = weaponInfo

}
