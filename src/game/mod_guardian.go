package game

import (
	"fmt"
	"server/src/csvs"
	"time"
)

const GUARDIAN_SYSTEM_INFO = "【Guardian模组】"

type ModGuardian struct {
	UUID     int     // 玩家唯一ID
	Icon     int     // 玩家头像
	Level    int     // 玩家守护者等级
	Exp      int     // 玩家当前经验值
	NickName string  // 玩家昵称
	AttackUp float64 // 攻击提升
	DefendUp float64 // 防御提升
	HPUp     float64 // 生命提升

	Prohibit int64 // 是否封号
	IsGM     int   // 是否是管理号
}

func (self *ModGuardian) SetNickName(nickName string, player *Player) {

	if GetManageBanWord().IsBanWord(nickName) {
		fmt.Println(GUARDIAN_SYSTEM_INFO, "存在违禁词")
		return
	}

	player.ModGuardian.NickName = nickName
	fmt.Println(GUARDIAN_SYSTEM_INFO, "设置昵称成功")
}

func (self *ModGuardian) SetUpgrade(player *Player) {

	config := csvs.GetCurConfigGuardianLevel(self.Level)
	if config == nil {
		fmt.Println(GUARDIAN_SYSTEM_INFO, "没有对应的等级配置")
		return
	}

	if self.Exp < config.GuardianExp || config.GuardianExp == 0 {
		fmt.Println(GUARDIAN_SYSTEM_INFO, "角色经验不够或者已经满级了")
		return
	}

	// todo 需要完成突破任务
	if config.AchievementId > 0 && !player.ModAchievement.IsAchievementDone(config.AchievementId) {
		fmt.Println(GUARDIAN_SYSTEM_INFO, "角色未完成某些成就")
		return
	}

	self.Exp -= config.GuardianExp
	self.Level++
	fmt.Println(GUARDIAN_SYSTEM_INFO, "当前等级：", self.Level, " 当前经验：", self.Exp)
	return
}

func (self *ModGuardian) AddExp(exp int) {

	if exp > 0 {
		self.Exp += exp
		fmt.Println(GUARDIAN_SYSTEM_INFO, "添加经验成功，当前经验：", self.Exp)
	}

}

func (self *ModGuardian) SetProhibit(prohibit int64) {
	self.Prohibit = prohibit
}

func (self *ModGuardian) SetGM(isGM int) {
	self.IsGM = isGM
}

func (self *ModGuardian) IsProhibit() bool {
	return self.Prohibit < time.Now().Unix()
}
