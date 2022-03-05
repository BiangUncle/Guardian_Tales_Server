package utils

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	err := os.Chdir("C:\\Users\\Biang\\Desktop\\github.com\\BiangUncle\\Guardian_Tales_Server")
	if err != nil {
		panic(err)
	}
}

func TestCsvUtilMgr_readCsv(t *testing.T) {

	result := GetCsvUtilMgr().readCsv("csv/GuardianLevel.csv")
	for _, r := range result {
		fmt.Println(r)
	}
}

type Test struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

var tests []*Test

func TestCsvUtilMgr_getValueType(t *testing.T) {

	typ := GetCsvUtilMgr().getValueType(&tests)
	fmt.Println(typ)
}

func TestCsvUtilMgr_LoadCsv(t *testing.T) {
	GetCsvUtilMgr().LoadCsv("test", &tests)
}
