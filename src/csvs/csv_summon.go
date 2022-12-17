package csvs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"server/src/utils"
	"time"
)

// todo 改成配置文件
const (
	ThreeStarHeroRate = 2750  // 三星英雄抽中几率
	TwoStarHeroRate   = 19000 // 二星英雄抽中几率
	OneStarHeroRate   = 78250 // 一星英雄抽中几率

	OneStarHeroHeroCrystal   = 1  // 一星英雄英雄水晶
	TwoStarHeroHeroCrystal   = 8  // 二星英雄英雄水晶
	ThreeStarHeroHeroCrystal = 50 // 三星英雄英雄水晶
)

// SummonConfig 抽奖池配置
type SummonConfig struct {
	SummonId int `json:"SummonId"` // 奖池ID
	Weight   int `json:"Weight"`   // 权重
	Result   int `json:"Result"`   // 结果
	IsEnd    int `json:"IsEnd"`    // 是否是节点
}

// SummonConfigSlice 奖池全局变量
var SummonConfigSlice []*SummonConfig

// init 初始化函数
// todo 改成主动初始化
func init() {
	utils.GetCsvUtilMgr().LoadCsv("hero_summon", &SummonConfigSlice)
	return
}

// SummonGroup 奖池组
type SummonGroup struct {
	SummonId      int             // 奖池ID
	WeightAll     int             // 奖池总权重
	SummonConfigs []*SummonConfig // 奖池配置
}

// SummonGroupMap  奖池组全局配置
var SummonGroupMap map[int]*SummonGroup

func CheckLoadCsv() {
	rand.Seed(time.Now().Unix())
	SummonGroupMap = make(map[int]*SummonGroup)

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
	log.Infoln(CONFIG_SYSTEM_INFO, "配置检测完成")

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
	summonGroup.WeightAll = OneStarHeroRate + TwoStarHeroRate + ThreeStarHeroRate
	SummonGroupMap[start] = summonGroup

	// 一星池
	summonConfig1 := new(SummonConfig)
	summonConfig1.SummonId = start
	summonConfig1.Result = start*10 + 1
	summonConfig1.Weight = OneStarHeroRate
	summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, summonConfig1)

	// 二星池
	summonConfig2 := new(SummonConfig)
	summonConfig2.SummonId = start
	summonConfig2.Result = start*10 + 2
	summonConfig2.Weight = TwoStarHeroRate
	summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, summonConfig2)

	// 三星池
	summonConfig3 := new(SummonConfig)
	summonConfig3.SummonId = start
	summonConfig3.Result = start*10 + 3
	summonConfig3.Weight = ThreeStarHeroRate
	summonGroup.SummonConfigs = append(summonGroup.SummonConfigs, summonConfig3)

	for i := 1; i < 4; i++ {
		curStart := start*10 + i

		subSummonGroup := new(SummonGroup)
		subSummonGroup.SummonId = curStart
		SummonGroupMap[curStart] = subSummonGroup

		for _, config := range allHeroINfo[i] {
			summonConfig := new(SummonConfig)
			summonConfig.SummonId = curStart
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
