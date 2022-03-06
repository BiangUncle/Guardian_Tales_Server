package game

import (
	"fmt"
	"server/src/csvs"
)

const (
	ITEM_HERO_CRYSTAL = 5000001

	MODHERO_SYSTEM_INFO       = "【英雄模块】"
	MODHERO_SYSTEM_COLOR_INFO = "\x1b[0;31m【英雄模块】\x1b[0m"
)

type HeroInfo struct {
	HeroId   int
	NickName string
	HeroName string
	HeroType int
	Star     int
	Type     int
}

type ModHero struct {
	HeroInfos map[int]*HeroInfo
}

func (self *ModHero) AddHero(heroId int, player *Player) {
	if self.HasHero(heroId) {
		self.Transform2HeroCrystal(self.HeroInfos[heroId], player)
	} else {
		heroInfo := csvs.GetHeroConfig(heroId)
		if heroInfo == nil {
			fmt.Println(MODHERO_SYSTEM_COLOR_INFO, "非法英雄Id")
		}
		self.HeroInfos[heroId] = &HeroInfo{
			HeroId:   heroId,
			NickName: heroInfo.NickName,
			HeroName: heroInfo.HeroName,
			HeroType: heroInfo.HeroType,
			Star:     heroInfo.Star,
			Type:     heroInfo.Type,
		}
		fmt.Println(MODHERO_SYSTEM_COLOR_INFO, "获取英雄成功，英雄为：", heroInfo.NickName, heroInfo.HeroName)
	}

	// 获得英雄，同时获得英雄头像
	iconConfig := csvs.GetIconConfig(heroId)
	if iconConfig == nil {
		fmt.Println(MODHERO_SYSTEM_COLOR_INFO, "配置错误，该英雄没有对应的英雄头像")
	}
	player.ModIcon.AddIcon(iconConfig.IconId)
}

func (self *ModHero) HasHero(heroId int) bool {
	_, ok := self.HeroInfos[heroId]
	if !ok {
		return false
	}
	return true
}

func (self *ModHero) Transform2HeroCrystal(heroInfo *HeroInfo, player *Player) {
	switch heroInfo.Star {
	case 1:
		player.ModBag.AddNormalItem(ITEM_HERO_CRYSTAL, csvs.ONESTAR_HERO_HERO_CRTSTAL)
	case 2:
		player.ModBag.AddNormalItem(ITEM_HERO_CRYSTAL, csvs.TWOSTAR_HERO_HERO_CRTSTAL)
	case 3:
		player.ModBag.AddNormalItem(ITEM_HERO_CRYSTAL, csvs.THREESTAR_HERO_HERO_CRTSTAL)
	}
}
