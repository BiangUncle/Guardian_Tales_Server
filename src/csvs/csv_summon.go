package csvs

import (
	"fmt"
	"math/rand"
	"server/src/utils"
	"time"
)

const (
	THREESTAR_HERO_RATE = 2750
	TWOSTAR_HERO_RATE   = 19000
	ONESTAR_HERO_RATE   = 78250

	TWOSTAR_HERO_SURE_RATE   = 97250
	THREESTAR_HERO_SURE_RATE = 2750

	ONESTAR_HERO_HERO_CRTSTAL   = 1
	TWOSTAR_HERO_HERO_CRTSTAL   = 8
	THREESTAR_HERO_HERO_CRTSTAL = 50
)

type SummonConfig struct {
	SummonId int `json:"SummonId"`
	Weight   int `json:"Weight"`
	Result   int `json:"Result"`
	IsEnd    int `json:"IsEnd"`
}

var SummonConfigSlice []*SummonConfig

func init() {
	utils.GetCsvUtilMgr().LoadCsv("hero_summon", &SummonConfigSlice)
	return
}

type SummonGroup struct {
	SummonId      int
	WeightAll     int
	SummonConfigs []*SummonConfig
}

var SummonGroupMap map[int]*SummonGroup

func CheckLoadCsv() {
	rand.Seed(time.Now().Unix())
	SummonGroupMap = make(map[int]*SummonGroup)
	fmt.Println(CONFIG_SYSTEM_INFO, "配置检测完成")

	for _, v := range SummonConfigSlice {
		summonGroup, ok := SummonGroupMap[v.SummonId]
		// todo 这里可以使用同名指针赋值，检查一下原理如何
		if !ok {
			summonGroup = new(SummonGroup)
			summonGroup.SummonId = v.SummonId
			SummonGroupMap[v.SummonId] = summonGroup
		}
		summonGroup.WeightAll += v.Weight
		summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, v)
	}
	//RandDropTest()
	return
}

func RandDropTest() {

	summonGroup := SummonGroupMap[1000]
	if summonGroup == nil {
		return
	}
	num := 0
	for {
		config := GetRandSummonNew(summonGroup)
		if config.IsEnd == 1 {
			fmt.Println(GetHeroFullName(config.Result))
			num++
			summonGroup = SummonGroupMap[1000]
			if num >= 100 {
				break
			}
			continue
		}
		summonGroup = SummonGroupMap[config.Result]
		if summonGroup == nil {
			break
		}
	}
}

func GetRandSummonNew(summonGroup *SummonGroup) *SummonConfig {
	randNum := rand.Intn(summonGroup.WeightAll)
	randNow := 0

	for _, v := range summonGroup.SummonConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == 1 {
				return v
			}

			summonGroup = SummonGroupMap[v.Result]
			if summonGroup == nil {
				return nil
			}
			return GetRandSummonNew(summonGroup)
		}
	}
	return nil
}

func GetRandSummon(summonGroup *SummonGroup) *SummonConfig {
	randNum := rand.Intn(summonGroup.WeightAll)
	randNow := 0

	for _, v := range summonGroup.SummonConfigs {
		randNow += v.Weight
		if randNum < randNow {
			return v
		}
	}
	return nil
}

func LoadAllHeroInfo() map[int][]*HeroConfig {
	allHeroINfo := make(map[int][]*HeroConfig)

	for _, v := range HeroConfigMap {
		switch v.Star {
		case 3:
			allHeroINfo[3] = append(allHeroINfo[3], v)
		case 2:
			allHeroINfo[2] = append(allHeroINfo[2], v)
		default:
			allHeroINfo[1] = append(allHeroINfo[1], v)
		}
	}
	return allHeroINfo
}

func CreateNormalSummonGroup(allHeroINfo map[int][]*HeroConfig) {
	start := 1000
	SummonGroupMap = make(map[int]*SummonGroup)

	// 一星池
	summonGroup := new(SummonGroup)
	summonGroup.SummonId = start
	summonGroup.WeightAll = ONESTAR_HERO_RATE + TWOSTAR_HERO_RATE + THREESTAR_HERO_RATE
	SummonGroupMap[start] = summonGroup

	// 一星池
	summonConfig1 := new(SummonConfig)
	summonConfig1.SummonId = start
	summonConfig1.Result = start*10 + 1
	summonConfig1.Weight = ONESTAR_HERO_RATE
	summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, summonConfig1)

	// 二星池
	summonConfig2 := new(SummonConfig)
	summonConfig2.SummonId = start
	summonConfig2.Result = start*10 + 2
	summonConfig2.Weight = TWOSTAR_HERO_RATE
	summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, summonConfig2)

	// 三星池
	summonConfig3 := new(SummonConfig)
	summonConfig3.SummonId = start
	summonConfig3.Result = start*10 + 3
	summonConfig3.Weight = THREESTAR_HERO_RATE
	summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, summonConfig3)

	for i := 1; i < 4; i++ {
		cur_start := start*10 + i

		subSummonGroup := new(SummonGroup)
		subSummonGroup.SummonId = cur_start
		SummonGroupMap[cur_start] = subSummonGroup

		for _, config := range allHeroINfo[i] {
			summonConfig := new(SummonConfig)
			summonConfig.SummonId = cur_start
			summonConfig.IsEnd = 1
			summonConfig.Weight = 1
			summonConfig.Result = config.HeroId

			subSummonGroup.WeightAll++
			subSummonGroup.SummonConfigs = append(subSummonGroup.SummonConfigs, summonConfig)
		}
	}

}

func CreateUpSummonGroup(heroId int) {
}
