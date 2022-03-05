package game

import (
	"fmt"
)

const MODICON_SYSTEM_INFO = "【头像模块】"

type IconInfo struct {
	IconId int
}

type ModIcon struct {
	IconInfos map[int]*IconInfo
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := self.IconInfos[iconId]
	if ok {
		return true
	}
	return false
}

func (self *ModIcon) SetIcon(iconId int, player *Player) {

	if !self.IsHasIcon(iconId) {
		fmt.Println(SYSTEM_INFO, "该用户没有该头像")
		return
	}

	fmt.Println(SYSTEM_INFO, "设置用户头像成功")
	player.ModGuardian.Icon = iconId

	return
}

func (self *ModIcon) AddIcon(iconId int) {

	_, ok := self.IconInfos[iconId]
	if ok {
		fmt.Println(MODICON_SYSTEM_INFO, "已存在该头像")
		return
	}

	self.IconInfos[iconId] = &IconInfo{IconId: iconId}
	fmt.Println(MODICON_SYSTEM_INFO, "添加头像成功")
}
