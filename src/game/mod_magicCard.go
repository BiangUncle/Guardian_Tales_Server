package game

import (
	"fmt"
	"sync"
)

const MODMAGICCARD_SYSTEM_INFO = "【时装模块】"

type MagicCardInfo struct {
	MagicCardId int
}

type ModMagicCard struct {
	MagicCards map[int]*MagicCardInfo
	Locker     *sync.RWMutex
}

func (self *ModMagicCard) HavaMagicCard(magicCardId int) bool {
	_, ok := self.MagicCards[magicCardId]
	return ok
}

func (self *ModMagicCard) AddMagicCard(magicCardId int) {
	_, ok := self.MagicCards[magicCardId]
	if ok {
		fmt.Println(MODMAGICCARD_SYSTEM_INFO, "已存在该魔卡")
	}

	self.MagicCards[magicCardId] = &MagicCardInfo{MagicCardId: magicCardId}
	fmt.Println(MODMAGICCARD_SYSTEM_INFO, "添加魔卡成功")
	return
}
