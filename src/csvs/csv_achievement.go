package csvs

import "server/src/utils"

type ConfigAchievement struct {
	AchievementId int `json:"AchievementId"`
	OpenLevel     int `json:"OpenLevel"`
	Condition     int `json:"Condition"`
}

var ConfigAchievementMap map[int]*ConfigAchievement

func init() {
	ConfigAchievementMap = make(map[int]*ConfigAchievement)
	utils.GetCsvUtilMgr().LoadCsv("Achievement", &ConfigAchievementMap)
	return
}
