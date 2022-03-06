package game

import (
	"server/src/csvs"
	"testing"
)

func TestModSummon_StartSummon(t *testing.T) {
	csvs.CheckLoadCsv()
	modSummon := &ModSummon{UpSummonInfo: &SummonPoolInfo{SummonPoolId: 1000}}
	modSummon.StartSummon()
}
