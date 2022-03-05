package csvs

import "fmt"

const CONFIG_SYSTEM_INFO = "【配置通知】"

type ConfigBanWord struct {
	Id  int    // 违禁词Id
	Txt string // 违禁词
}

var ConfigBanWordSlice []*ConfigBanWord

func init() {

	ConfigBanWordSlice = append(ConfigBanWordSlice,
		&ConfigBanWord{Id: 0, Txt: "作弊"},
		&ConfigBanWord{Id: 1, Txt: "外挂"},
	)

	fmt.Println(CONFIG_SYSTEM_INFO, "初始化完成")
}

func GetBandWordBase() []string {
	relString := make([]string, 0)
	for _, v := range ConfigBanWordSlice {
		relString = append(relString, v.Txt)
	}
	return relString
}
