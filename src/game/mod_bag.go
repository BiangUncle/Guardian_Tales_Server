package game

import (
	"fmt"
	"server/src/csvs"
)

const (
	MODBAG_SYSTEM_INFO = "【背包模组】"

	MP       = 1000001
	GOLD     = 1000002
	FREEGEMS = 1000003
	PAIDGEMS = 1000004
)

type Item struct {
	ItemId   int
	ItemNum  int64
	ItemName string
}

type ModBag struct {
	BagInfos map[int]*Item
	Gold     int64
	FreeGems int64
	PaidGems int64
	MP       int64 // todo 以后要知道体力的英文
}

func (self *ModBag) AddItem(itemId int, num int64, player *Player) {

	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(MODBAG_SYSTEM_INFO, "物品不存在")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_GUARDIAN:
		self.AddGuardianItem(itemId, num)
	case csvs.ITEMTYPE_ROLE:
		player.ModHero.AddHero(itemId, player)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddIcon(itemId)
	case csvs.ITEMTYPE_WEAPON:
		player.ModWeapon.AddWeapon(itemId)
	case csvs.ITEMTYPE_NORMAL:
		self.AddNormalItem(itemId, num)
	case csvs.ITEMTYPE_COSTUME:
		player.ModCostume.AddCostume(itemId)
	case csvs.ITEMTYPE_ACCESSORY:
		player.ModAccessory.AddAccessory(itemId)
	case csvs.ITEMTYPE_MAGICCARD:
		player.ModMagicCard.AddMagicCard(itemId)
	case csvs.ITEMTYPE_INVENTORY:
		player.ModInventory.AddInventory(itemId)
	default:
		fmt.Println(MODBAG_SYSTEM_INFO, "未知物品")
	}

	return
}

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {

	item, ok := self.BagInfos[itemId]
	if !ok {
		fmt.Println(MODBAG_SYSTEM_INFO, "背包不存在该物品")
		return false
	}

	if item.ItemNum < num {
		return false
	}

	return true
}

func (self *ModBag) AddNormalItem(itemId int, num int64) {

	item, ok := self.BagInfos[itemId]
	if !ok {
		itemInfo := csvs.GetItemConfig(itemId)
		self.BagInfos[itemId] = &Item{ItemId: itemInfo.ItemId, ItemNum: num, ItemName: itemInfo.ItemName}
		fmt.Println(MODBAG_SYSTEM_INFO, "添加物品成功，物品名：", itemInfo.ItemName, " 物品数量：", num)
		return
	}
	item.ItemNum += num
	fmt.Println(MODBAG_SYSTEM_INFO, "添加物品成功，物品名：", item.ItemName, " 物品数量：", num)
	return
}

func (self *ModBag) RemoveNormalItem(itemId int, num int64) {

	if !self.HasEnoughItem(itemId, num) {
		fmt.Println(MODBAG_SYSTEM_INFO, "物品数量不足，无法扣除")
	}

	item, _ := self.BagInfos[itemId]
	item.ItemNum -= num
	fmt.Println(MODBAG_SYSTEM_INFO, "添加物品成功，物品名：", item.ItemName, " 物品数量：", num)
	if item.ItemNum == 0 {
		delete(self.BagInfos, itemId) // 删除该物品信息
	}
	return
}

func (self *ModBag) AddGuardianItem(itemId int, num int64) {

	switch itemId {
	case MP:
		self.MP += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加体力成功，当前体力：", self.MP)
	case GOLD:
		self.Gold += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加金币成功，当前金币：", self.Gold)
	case FREEGEMS:
		self.FreeGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加免费钻石成功，当前免费钻石：", self.FreeGems)
	case PAIDGEMS:
		self.PaidGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加付费钻石成功，当前付费钻石：", self.PaidGems)
	default:
		fmt.Println(MODBAG_SYSTEM_INFO, "未知背包物品")
	}
}

func (self *ModBag) RemoveGuardianItem(itemId int, num int64) {

	switch itemId {
	case MP:
		if self.MP >= num {
			self.MP -= num
		}
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除体力成功，当前体力：", self.MP)
	case GOLD:
		if self.Gold >= num {
			self.Gold -= num
		}
		self.Gold += num
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除金币成功，当前金币：", self.Gold)
	case FREEGEMS:
		if self.FreeGems >= num {
			self.FreeGems -= num
		}
		self.FreeGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除免费钻石成功，当前免费钻石：", self.FreeGems)
	case PAIDGEMS:
		if self.PaidGems >= num {
			self.PaidGems -= num
		}
		self.PaidGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除付费钻石成功，当前付费钻石：", self.PaidGems)
	default:
		fmt.Println(MODBAG_SYSTEM_INFO, "未知背包物品")
	}
}
