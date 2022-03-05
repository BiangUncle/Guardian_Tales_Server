package game

import (
	"sync"
)

const SYSTEM_INFO = "【系统通知】"

type Player struct {
	ModGuardian    *ModGuardian
	ModIcon        *ModIcon
	ModAchievement *ModAchievement
	ModBag         *ModBag
	ModHero        *ModHero
	ModWeapon      *ModWeapon
	ModCostume     *ModCostume
	ModAccessory   *ModAccessory
	ModMagicCard   *ModMagicCard
	ModInventory   *ModInventory
}

/*
当前：角色模块
1. 角色等级
2. 头像
3. 昵称
4. 唯一Id
5. 守护者积分（升级）
6. 守护者等级加成
7. 守护者等级特殊技能
*/

/*
当前：背包模块
1. 物品识别（周边 Accessory，时装 Costume，魔卡 Magic Card，库存 Inventory）
2. 物品添加
3. 物品消耗
4. 物品消耗
*/

func NewTestPlayer() *Player {
	player := new(Player)
	//****************************
	player.ModGuardian = new(ModGuardian)
	player.ModGuardian.Level = 1

	//****************************
	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfos = make(map[int]*IconInfo)

	//****************************
	player.ModAchievement = new(ModAchievement)
	player.ModAchievement.Achievements = make(map[int]*AchievementInfo)
	player.ModAchievement.Locker = new(sync.RWMutex)

	//****************************
	player.ModBag = new(ModBag)
	player.ModBag.BagInfos = make(map[int]*Item)

	//****************************
	player.ModHero = new(ModHero)
	player.ModHero.HeroInfos = make(map[int]*HeroInfo)

	//****************************
	player.ModWeapon = new(ModWeapon)
	player.ModWeapon.WeaponInfos = make(map[int]*WeaponInfo)

	//****************************
	player.ModCostume = new(ModCostume)
	player.ModCostume.Costumes = make(map[int]*CostumeInfo)
	player.ModCostume.Locker = new(sync.RWMutex)

	//****************************
	player.ModAccessory = new(ModAccessory)
	player.ModAccessory.Accessories = make(map[int]*AccessoryInfo)
	player.ModAccessory.Locker = new(sync.RWMutex)

	//****************************
	player.ModMagicCard = new(ModMagicCard)
	player.ModMagicCard.MagicCards = make(map[int]*MagicCardInfo)
	player.ModMagicCard.Locker = new(sync.RWMutex)

	//****************************
	player.ModInventory = new(ModInventory)
	player.ModInventory.Inventories = make(map[int]*InventoryInfo)
	player.ModInventory.Locker = new(sync.RWMutex)

	return player
}

func (self *Player) RecvSetIcon(iconId int) {
	self.ModIcon.SetIcon(iconId, self)
}

func (self *Player) RecvSetNickName(nickName string) {
	self.ModGuardian.SetNickName(nickName, self)
}

func (self *Player) RecvUpgrade() {
	self.ModGuardian.SetUpgrade(self)
}
