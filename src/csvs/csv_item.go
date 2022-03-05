package csvs

import (
	"server/src/utils"
)

const (
	ITEMCONFIG_SYSTEM_INFO = "【ItemConfig工具】"

	ITEMTYPE_GUARDIAN  = 1
	ITEMTYPE_ROLE      = 2
	ITEMTYPE_ICON      = 3
	ITEMTYPE_WEAPON    = 4
	ITEMTYPE_NORMAL    = 5
	ITEMTYPE_COSTUME   = 6
	ITEMTYPE_ACCESSORY = 7
	ITEMTYPE_MAGICCARD = 8
	ITEMTYPE_INVENTORY = 9
)

type ItemConfig struct {
	ItemId   int    `json:"ItemId"`
	SortType int    `json:"SortType"`
	ItemName string `json:"ItemName"`
}

var ItemConfigMap map[int]*ItemConfig

func init() {
	ItemConfigMap = make(map[int]*ItemConfig)
	utils.GetCsvUtilMgr().LoadCsv("item", &ItemConfigMap)
	return
}

func GetItemConfig(itemId int) *ItemConfig {
	itemConfig, ok := ItemConfigMap[itemId]
	if !ok {
		return nil
	}
	return itemConfig
}
