package csvs

import "testing"

func Test_ConfigGuardianLevel_Init(t *testing.T) {
	ConfigGuardianLevel_Init()
	config := ConfigGuardianLevelSlice
	println(config)
}

func Test_ConfigAchievement_Init(t *testing.T) {
	ConfigAchievement_Init()
}

func Test_ConfigItem_Init(t *testing.T) {
	ItemConfig_Init()
}
