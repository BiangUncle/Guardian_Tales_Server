package game

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"server/src/csvs"
	"server/src/utils"
)

const ModSummonSystemInfo = "【抽奖模块】"

type SummonPoolInfo struct {
	SummonPoolId int
}

type ModSummon struct {
	UpSummonInfo *SummonPoolInfo
}

// StartOneSummon 单抽
func (m *ModSummon) StartOneSummon(player *Player) {

	summonGroup := csvs.SummonGroupMap[1000]
	if summonGroup == nil {
		return
	}
	config := csvs.GetRandSummonNew(summonGroup)
	if config != nil {
		fmt.Println(fmt.Sprintf("%s抽中【%s】", utils.Red(ModSummonSystemInfo), csvs.GetHeroFullName(config.Result)))
		player.ModBag.AddItem(config.Result, 1, player)
	}

}

func (m *ModSummon) StartTenSummon(player *Player) {

	for i := 0; i < 10; i++ {
		summonGroup := csvs.SummonGroupMap[1000]
		if summonGroup == nil {
			return
		}
		config := csvs.GetRandSummonNew(summonGroup)
		if config != nil {
			fmt.Println(fmt.Sprintf("%s第%d抽中【%s】", utils.Red(ModSummonSystemInfo), i+1, csvs.GetHeroFullName(config.Result)))
			player.ModBag.AddItem(config.Result, 1, player)
		}
	}
}

// StartSummonWithTimes 测试用
func (m *ModSummon) StartSummonWithTimes(times int) {

	result := make(map[int]int)
	for i := 0; i < times; i++ {
		summonGroup := csvs.SummonGroupMap[1000]
		if summonGroup == nil {
			return
		}
		config := csvs.GetRandSummonNew(summonGroup)
		if config != nil {
			// todo golang的map可以返回默认值？
			result[config.Result]++
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中【%s】次数: %d", csvs.GetHeroFullName(k), v))
	}
}

// SummonTenTimes HTTP请求抽奖10次英雄
func SummonTenTimes() ([]string, error) {

	heroList := make([]string, 10)

	for i := 0; i < 10; i++ {
		summonGroup := csvs.SummonGroupMap[1000]
		if summonGroup == nil {
			return nil, errors.New("获取到空的角色")
		}
		config := csvs.GetRandSummonNew(summonGroup)
		if config != nil {
			log.Infoln(fmt.Sprintf("%s第%d抽中【%s】", utils.Red(ModSummonSystemInfo), i+1, csvs.GetHeroFullName(config.Result)))
		}
		heroList[i] = csvs.GetHeroFullName(config.Result)
	}

	return heroList, nil
}
