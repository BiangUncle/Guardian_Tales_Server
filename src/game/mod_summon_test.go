package game

import (
	"server/src/csvs"
	"testing"
)

func init() {
	csvs.CheckLoadCsv()
}

func TestSummonTenTimes(t *testing.T) {
	got, err := SummonTenTimes()
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
}
