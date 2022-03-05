package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

// todo 这个是在 window 电脑下的配置，肯定需要改
const CSV_DIR = "C:\\Users\\Biang\\Desktop\\github.com\\BiangUncle\\Guardian_Tales_Server\\csv\\"

type CsvManager struct {
}

var csvManager *CsvManager

func GetCSVManager() *CsvManager {
	if csvManager == nil {
		csvManager = &CsvManager{}
	}
	return csvManager
}

func (self *CsvManager) LoadCsv(filename string, config interface{}) {

	f, err := os.Open(CSV_DIR + filename + ".csv")
	if err != nil {
		// todo 需要报错返回
		return
	}

	reader := csv.NewReader(f)
	line, err := reader.Read()
	if err != nil {
		// todo 需要报错返回
		return
	}

	for _, v := range line {
		fmt.Println(v)
	}

	wtype := reflect.TypeOf(config)
	fmt.Println(wtype)
}
