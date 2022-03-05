package game

import (
	"fmt"
	"sync"
)

const MODACCESSORY_SYSTEM_INFO = "【周边模块】"

type AccessoryInfo struct {
	AccessoryId int
}

type ModAccessory struct {
	Accessories map[int]*AccessoryInfo
	Locker      *sync.RWMutex
}

func (self *ModAccessory) HavaAccessory(assessoryId int) bool {
	_, ok := self.Accessories[assessoryId]
	return ok
}

// todo 魔卡是可以拿很多个的，跟装备逻辑一样
func (self *ModAccessory) AddAccessory(assessoryId int) {
	_, ok := self.Accessories[assessoryId]
	if ok {
		fmt.Println(MODACCESSORY_SYSTEM_INFO, "已存在该周边")
	}

	self.Accessories[assessoryId] = &AccessoryInfo{AccessoryId: assessoryId}
	fmt.Println(MODACCESSORY_SYSTEM_INFO, "添加周边成功")
	return
}
