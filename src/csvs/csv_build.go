package csvs

import (
	"fmt"
	"server/src/utils"
)

const (
	PRODUCT_BUILD_TYPE_FOOD          = 1
	PRODUCT_BUILD_TYPE_DRINK         = 2
	PRODUCT_BUILD_TYPE_ENTERTAINMENT = 3
)

type ProductBuildConfig struct {
	BuildId   int    `json:"BuildId"`
	BuildName string `json:"BuildName"`
	Type      int    `json:"Type"`
	Star      int    `json:"Star"`
	InnLevel  int    `json:"InnLevel"`
	Cost      int    `json:"Cost"`
	Limit     int    `json:"Limit"`
}

var ProductBuildConfigMap map[int]*ProductBuildConfig

type LandmarkBuildConfig struct {
	BuildId       int    `json:"BuildId"`
	BuildName     string `json:"BuildName"`
	BuildCostTime int    `json:"BuildCostTime"`
}

var LandmarkBuildConfigMap map[int]*LandmarkBuildConfig

func init() {
	ProductBuildConfigMap = make(map[int]*ProductBuildConfig)
	LandmarkBuildConfigMap = make(map[int]*LandmarkBuildConfig)
	utils.GetCsvUtilMgr().LoadCsv("productBuild", &ProductBuildConfigMap)
	utils.GetCsvUtilMgr().LoadCsv("landmarkBuild", &LandmarkBuildConfigMap)
	fmt.Println(LandmarkBuildConfigMap)
	return
}

func GetProductBuildConfig(buildKey int) *ProductBuildConfig {
	buildConfig, ok := ProductBuildConfigMap[buildKey]
	if !ok {
		return nil
	}
	return buildConfig
}

func GetProductBuildLimit(buildKey int) int {

	buildConfig, ok := ProductBuildConfigMap[buildKey]
	if !ok {
		return 0
	}
	return buildConfig.Limit
}

func HaveBuild(buildKey int) bool {
	_, ok := ProductBuildConfigMap[buildKey]
	if !ok {
		return false
	}
	return true
}

func GetLandmarkBuildConfig(buildId int) *LandmarkBuildConfig {
	landmarkBuildConfig, ok := LandmarkBuildConfigMap[buildId]
	if !ok {
		return nil
	}
	return landmarkBuildConfig
}
