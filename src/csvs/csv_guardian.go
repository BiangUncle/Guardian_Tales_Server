package csvs

import "server/src/utils"

type ConfigGuardianLevel struct {
	GuardianLevel int `json:"GuardianLevel"`
	GuardianExp   int `json:"GuardianExp"`
	AchievementId int `json:"AchievementId"`
}

var ConfigGuardianLevelSlice []*ConfigGuardianLevel

func init() {
	utils.GetCsvUtilMgr().LoadCsv("GuardianLevel", &ConfigGuardianLevelSlice)
	return
}

func GetCurConfigGuardianLevel(level int) *ConfigGuardianLevel {
	if level <= 0 || level > len(ConfigGuardianLevelSlice) {
		return nil
	}
	return ConfigGuardianLevelSlice[level-1]
}
