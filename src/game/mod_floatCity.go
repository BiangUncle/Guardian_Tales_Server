package game

import (
	"fmt"
	"server/src/csvs"
	"server/src/utils"
	"time"
)

const (
	FLOATCITY_SYSTEM_INFO = "【浮游城模组】"

	PRODUCT_BUILD_NUM  = 36
	LANDMARK_BUILD_NUM = 6
	PLANT_NUM          = 17

	PLANT_STATUS_GROWING = 1
	PLANT_STATUS_MATURE  = 2
	PLANT_GROWING_TIME   = time.Hour * 24

	INN_LOCATION = 0
)

type PlantInfo struct {
	location        int
	NextCollectTime int64
	Status          int
}

func (self *PlantInfo) GetLocation() int {
	return self.location
}

func (self *PlantInfo) GetLastCollectTime() int64 {
	return self.NextCollectTime
}

func (self *PlantInfo) GetStatus() int {
	return self.Status
}

type ModLandMarkBuild struct {
	Inn           *Inn
	SkyGarden     *SkyGarden
	PowerTower    *PowerTower
	DefenseTower  *DefenseTower
	HeroStatue    *HeroStatue
	Shop          *Shop
	BuildIdxSlice [6]int
}

type ModFloatCity struct {
	FoodOutput  int
	DrinkOutput int
	EnterOutput int
	SplitPoint  int

	WorkerNum        int
	WorkingWorkerNum int

	HeroNum int

	SplitPointOutput int

	SplitPointLimitCity int

	GoldOutput int

	FriendNum     int
	VisitRange    int
	LastVisitTime int64

	BuildKey           int
	ProductBuildInfos  [PRODUCT_BUILD_NUM]*ProductBuild
	ProductBuildNumMap map[int]int
	//LandMarkBuildInfos [LANDMARK_BUILD_NUM]*Build
	ModLandMarkBuild *ModLandMarkBuild
	PlantInfos       [PLANT_NUM]*PlantInfo
}

func (m *ModFloatCity) FloatCityInit() {

	// map初始化
	m.ProductBuildNumMap = make(map[int]int)
	m.ModLandMarkBuild = new(ModLandMarkBuild)
	for i := 0; i < 6; i++ {
		m.ModLandMarkBuild.BuildIdxSlice[i] = -1
	}

	// 初始化植物
	for i := 0; i < PLANT_NUM; i++ {
		plant := new(PlantInfo)
		plant.Status = PLANT_STATUS_MATURE
		plant.NextCollectTime = time.Now().Unix()
		plant.location = i
		m.PlantInfos[i] = plant
	}

	innConfig := csvs.GetLandmarkBuildConfig(1)
	if innConfig == nil {
		panic("FloatCityInit error")
	}

	// 初始化旅馆
	inn := new(Inn)
	inn.Location = INN_LOCATION
	inn.Name = innConfig.BuildName
	inn.Status = BUILD_STATUS_FINISH
	inn.BuildType = BUILD_TYPE_LANDMARK
	inn.SplitPointLimitHero = 200
	inn.HeroNumLimit = 20
	inn.Level = 1
	inn.BuildLevelLimit = 1

	m.BuildKey++
	inn.KeyId = m.BuildKey
	m.ModLandMarkBuild.Inn = inn

}

func (m *ModFloatCity) GetBuildLevelLimit() int {
	return m.ModLandMarkBuild.Inn.BuildLevelLimit
}

func (m *ModFloatCity) GetHeroNumLimit() int {
	return m.ModLandMarkBuild.Inn.HeroNumLimit
}

func (m *ModFloatCity) GetSplitPointLimitHero() int {
	return m.ModLandMarkBuild.Inn.SplitPointLimitHero
}

func (m *ModFloatCity) Build(buildType int, buildId int, loc int) {
	if buildType == BUILD_TYPE_PRODUCT {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "建造生产建筑")
		m.BuildProductBuild(buildId, loc)
		return
	}
	if buildType == BUILD_TYPE_LANDMARK {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "建造地标建筑")
		return
	}
	fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "未知建筑类型")
}

func (m *ModFloatCity) BuildLandMarkBuild(buildId int, loc int) {
	if loc == INN_LOCATION || buildId == INN_BUILD_ID {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "不准动旅馆")
		return
	}

	if m.ModLandMarkBuild.BuildIdxSlice[loc] != -1 {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "该地已有建筑物")
		return
	}

	landmarkConfig := csvs.GetLandmarkBuildConfig(buildId)
	if landmarkConfig == nil {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "未知地标建筑物")
		return
	}
	build := new(Build)
	build.Level = 1
	build.Status = BUILD_STATUS_BUILDING
	build.KeyId = buildId
	build.Location = loc
	build.BuildType = BUILD_TYPE_LANDMARK
	build.Name = landmarkConfig.BuildName
	build.BuildFinishTime = time.Now().Add(time.Duration(landmarkConfig.BuildCostTime) * time.Second).Unix()

	switch buildId {
	case SKYGARDEN_KEY:
		if m.ModLandMarkBuild.SkyGarden != nil {
			fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "已有该建筑物")
			return
		}
		skyGarden := new(SkyGarden)
		skyGarden.Build = *build
		skyGarden.SplitPointLimitCity = 100000
		m.ModLandMarkBuild.SkyGarden = skyGarden
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), skyGarden.GetFunction())
	case DEFENSETOWER_KEY:
		if m.ModLandMarkBuild.DefenseTower != nil {
			fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "已有该建筑物")
			return
		}
		defenseTower := new(DefenseTower)
		defenseTower.Build = *build
		defenseTower.DefAdd = 1
		m.ModLandMarkBuild.DefenseTower = defenseTower
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), defenseTower.GetFunction())
	case POWERTOWER_KEY:
		if m.ModLandMarkBuild.PowerTower != nil {
			fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "已有该建筑物")
			return
		}
		powerTower := new(PowerTower)
		powerTower.Build = *build
		powerTower.AtkAdd = 1
		m.ModLandMarkBuild.PowerTower = powerTower
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), powerTower.GetFunction())
	case HEROSTATUE_KEY:
		if m.ModLandMarkBuild.HeroStatue != nil {
			fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "已有该建筑物")
			return
		}
		heroStatue := new(HeroStatue)
		heroStatue.Build = *build
		heroStatue.VisitRangeMinus = 20
		m.ModLandMarkBuild.HeroStatue = heroStatue
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), heroStatue.GetFunction())
	case SHOP_KEY:
		if m.ModLandMarkBuild.Shop != nil {
			fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "已有该建筑物")
			return
		}
		shop := new(Shop)
		shop.Build = *build
		shop.UpdateTime = time.Now().Add(24 * time.Hour).Unix()
		shop.GoodsLevel = 1
		m.ModLandMarkBuild.Shop = shop
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), shop.GetFunction())
	default:
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "未知地标")
		return
	}
}

func (m *ModFloatCity) BuildProductBuild(buildId int, loc int) {

	if loc >= PRODUCT_BUILD_NUM {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "错误建筑索引")
		return
	}

	if m.ProductBuildInfos[loc] != nil {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "该地点已有建筑物")
		return
	}

	buildConfig := csvs.GetProductBuildConfig(buildId)
	if buildConfig == nil {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "错误建筑物")
		return
	}

	if v, ok := m.ProductBuildNumMap[buildId]; ok {
		if v >= buildConfig.Limit {
			fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "建筑物达到上限")
			return
		}
	} else {
		m.ProductBuildNumMap[buildId] = 0
	}

	m.BuildKey++
	productBuild := &ProductBuild{
		Build: Build{
			Location:        loc,
			Status:          BUILD_STATUS_BUILDING,
			KeyId:           m.BuildKey,
			Level:           1,
			Name:            buildConfig.BuildName,
			BuildType:       BUILD_STATUS_BUILDING,
			BuildFinishTime: time.Now().Add(time.Second * 5).Unix(),
		},
		Type:       buildConfig.Type,
		TypeOutput: 20,
	}
	m.ProductBuildInfos[loc] = productBuild
	fmt.Println(fmt.Sprintf("%s[%d][%s] Lv:【%d】 产量:【%d】 状态:【%s】",
		utils.Red(FLOATCITY_SYSTEM_INFO),
		productBuild.GetLocation(),
		productBuild.GetName(),
		productBuild.GetLevel(),
		productBuild.GetTypeOutput(),
		productBuild.GetRemainTime(),
	))

	switch buildConfig.Type {
	case csvs.PRODUCT_BUILD_TYPE_FOOD:
		m.FoodOutput += 20
	case csvs.PRODUCT_BUILD_TYPE_DRINK:
		m.DrinkOutput += 20
	case csvs.PRODUCT_BUILD_TYPE_ENTERTAINMENT:
		m.EnterOutput += 20
	}
	m.ProductBuildNumMap[buildId]++
}

func (m *ModFloatCity) ShowAllBuilding() {

	fmt.Println("= 生活类建筑 =======================================================")

	for _, build := range m.ProductBuildInfos {
		if build != nil {
			fmt.Println(fmt.Sprintf("[%d][%s] Lv:【%d】 产量:【%d】 状态:【%s】",
				build.GetLocation(),
				build.GetName(),
				build.GetLevel(),
				build.GetTypeOutput(),
				build.GetRemainTime(),
			))
		}
	}

	//fmt.Println("= 地标类建筑 =======================================================")
	//
	//for _, build := range m.LandMarkBuildInfos {
	//	if build != nil {
	//		if lanmarkBuild, ok := (*build).(LandMarkBuild); ok {
	//
	//		}
	//		fmt.Println(fmt.Sprintf("[%d][%s] Lv:【%d】 功能:【%s】 状态:【%d】",
	//			build.GetLocation(),
	//			build.GetName(),
	//			build.GetLevel(),
	//			build.GetFunction(),
	//			build.GetStatus(),
	//		))
	//	}
	//}

	fmt.Println("= 植物类建筑 =======================================================")

	for _, plant := range m.PlantInfos {
		if plant != nil {
			fmt.Println(fmt.Sprintf("[%d] 状态:【%d】",
				plant.GetLocation(),
				plant.GetStatus(),
			))
		}
	}

	fmt.Println("================================================================")
}

func (m *ModFloatCity) UpdatePlant(loc int) {
	if time.Now().Unix() >= m.PlantInfos[loc].NextCollectTime {
		m.PlantInfos[loc].Status = PLANT_STATUS_MATURE
	}
}

func (m *ModFloatCity) CollectPlant(loc int) {
	if m.PlantInfos[loc].Status != PLANT_STATUS_MATURE {
		fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "植物未成熟")
	}
	plant := m.PlantInfos[loc]
	plant.Status = PLANT_STATUS_GROWING
	plant.NextCollectTime = time.Now().Add(10 * time.Second).Unix()
	fmt.Println(utils.Red(FLOATCITY_SYSTEM_INFO), "植物采集成功")
	// todo 获得某些物品
}
