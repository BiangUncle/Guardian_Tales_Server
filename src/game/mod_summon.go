package game

import (
	"fmt"
	"server/src/csvs"
)

type SummonPoolInfo struct {
	SummonPoolId int
}

type ModSummon struct {
	UpSummonInfo *SummonPoolInfo
}

func (self *ModSummon) StartOneSummon() {

	result := make(map[int]int)
	summonGroup := csvs.SummonGroupMap[1000]
	if summonGroup == nil {
		return
	}
	config := csvs.GetRandSummonNew(summonGroup)
	if config != nil {
		// todo golang的map可以返回默认值？
		result[config.Result]++
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中【%s】次数: %d", csvs.GetHeroFullName(k), v))
	}
}

func (self *ModSummon) StartTenSummon() {

	result := make(map[int]int)
	for i := 0; i < 10; i++ {
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
