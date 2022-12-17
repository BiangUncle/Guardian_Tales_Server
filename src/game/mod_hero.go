package game

import (
	"fmt"
	"server/src/csvs"
)

const (
	ITEM_HERO_CRYSTAL         = 5000001
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

// AddHero 增加英雄
// @param heroId: 英雄id
func (m *ModHero) AddHero(heroId int, player *Player) {
	// 判断是否存在英雄
	if m.HasHero(heroId) {
		// 转成英雄水晶
		m.Transform2HeroCrystal(m.HeroInfos[heroId], player)
	} else {
		// 获取英雄配置
		heroInfo := csvs.GetHeroConfig(heroId)
		if heroInfo == nil {
			fmt.Println(MODHERO_SYSTEM_COLOR_INFO, "非法英雄Id")
		}
		m.HeroInfos[heroId] = &HeroInfo{
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

// HasHero 判断该角色是否有这个英雄
func (m *ModHero) HasHero(heroId int) bool {
	_, ok := m.HeroInfos[heroId]
	if !ok {
		return false
	}
	return true
}

// Transform2HeroCrystal 转换到英雄水晶
func (m *ModHero) Transform2HeroCrystal(heroInfo *HeroInfo, player *Player) {
	switch heroInfo.Star {
	case 1:
		player.ModBag.AddNormalItem(ITEM_HERO_CRYSTAL, csvs.OneStarHeroHeroCrystal)
	case 2:
		player.ModBag.AddNormalItem(ITEM_HERO_CRYSTAL, csvs.TwoStarHeroHeroCrystal)
	case 3:
		player.ModBag.AddNormalItem(ITEM_HERO_CRYSTAL, csvs.ThreeStarHeroHeroCrystal)
	}
}
