package csvs

import (
	"os"
	"server/src/utils"
)

type ConfigAchievement struct {
	AchievementId int `json:"AchievementId"`
	OpenLevel     int `json:"OpenLevel"`
	Condition     int `json:"Condition"`
}

var ConfigAchievementMap map[int]*ConfigAchievement

func init() {

	// todo 需要改动，不能默认这个地方来初始化文件目录
	err := os.Chdir("C:\\Users\\Biang\\Desktop\\github.com\\BiangUncle\\Guardian_Tales_Server")
	if err != nil {
		panic(err)
	}
	ConfigAchievementMap = make(map[int]*ConfigAchievement)
	utils.GetCsvUtilMgr().LoadCsv("Achievement", &ConfigAchievementMap)
	return
}
