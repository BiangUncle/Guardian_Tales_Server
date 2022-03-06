package csvs

import (
	"testing"
)

func Test_ConfigGuardianLevel_Init(t *testing.T) {
	heroInfo := LoadAllHeroInfo()
	CreateNormalSummonGroup(heroInfo)
	print("DONE")
}
