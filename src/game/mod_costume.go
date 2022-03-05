package game

import (
	"fmt"
	"sync"
)

const MODCOSTUME_SYSTEM_INFO = "【时装模块】"

type CostumeInfo struct {
	CostumeId int
}

type ModCostume struct {
	Costumes map[int]*CostumeInfo
	Locker   *sync.RWMutex
}

func (self *ModCostume) HavaCostume(costumeId int) bool {
	_, ok := self.Costumes[costumeId]
	return ok
}

func (self *ModCostume) AddCostume(costumeId int) {
	_, ok := self.Costumes[costumeId]
	if ok {
		fmt.Println(MODCOSTUME_SYSTEM_INFO, "已存在该时装")
	}

	self.Costumes[costumeId] = &CostumeInfo{CostumeId: costumeId}
	fmt.Println(MODCOSTUME_SYSTEM_INFO, "添加时装成功")
	return
}
