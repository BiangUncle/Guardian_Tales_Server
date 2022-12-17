package game

import (
	log "github.com/sirupsen/logrus"
)

const ModIconSystemInfo = "【头像模块】"

// IconInfo 图标信息
type IconInfo struct {
	IconId int
}

// ModIcon 图标模组
type ModIcon struct {
	IconInfos map[int]*IconInfo
}

// IsHasIcon 是否有该图标
func (m *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := m.IconInfos[iconId]
	if ok {
		return true
	}
	return false
}

// SetIcon 设置图标
func (m *ModIcon) SetIcon(iconId int, player *Player) {

	if !m.IsHasIcon(iconId) {
		log.Infoln(SYSTEM_INFO, "该用户没有该头像")
		return
	}

	log.Infoln(SYSTEM_INFO, "设置用户头像成功")
	player.ModGuardian.Icon = iconId

	return
}

// AddIcon 增加图标
func (m *ModIcon) AddIcon(iconId int) {

	_, ok := m.IconInfos[iconId]
	if ok {
		log.Infoln(ModIconSystemInfo, "已存在该头像")
		return
	}

	m.IconInfos[iconId] = &IconInfo{IconId: iconId}
	log.Infoln(ModIconSystemInfo, "添加头像成功")
}
