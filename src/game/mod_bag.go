package game

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"server/src/csvs"
)

const (
	MODBAG_SYSTEM_INFO       = "【背包模组】"
	MODBAG_SYSTEM_COLOR_INFO = "\x1b[0;31m【背包模组】\x1b[0m"

	MP       = 1000001
	GOLD     = 1000002
	FREEGEMS = 1000003
	PAIDGEMS = 1000004
)

// Item 物品结构体
type Item struct {
	ItemId   int    `json:"item_id"`   // 物品ID
	ItemNum  int64  `json:"item_num"`  // 物品数量
	ItemName string `json:"item_name"` // 物品名称
}

// ModBag 背包结构体
type ModBag struct {
	BagInfos map[int]*Item `json:"bag_infos"` // 背包信息
	Gold     int64         `json:"gold"`      // 金币
	FreeGems int64         `json:"free_gems"` // 免费钻石
	PaidGems int64         `json:"paid_gems"` // 付费钻石
	MP       int64         `json:"mp"`        // 体力 // todo 以后要知道体力的英文
	MPLimit  int64         `json:"mp_limit"`  // 体力上限
}

// NewBag 创建创建背包模组
func NewBag() *ModBag {
	bag := new(ModBag)
	bag.BagInfos = make(map[int]*Item)
	return bag
}

func (m *ModBag) AddItem(itemId int, num int64, player *Player) {

	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(MODBAG_SYSTEM_COLOR_INFO, "物品不存在")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_GUARDIAN:
		m.AddGuardianItem(itemId, num)
	case csvs.ITEMTYPE_ROLE:
		player.ModHero.AddHero(itemId, player)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddIcon(itemId)
	case csvs.ITEMTYPE_WEAPON:
		player.ModWeapon.AddWeapon(itemId)
	case csvs.ITEMTYPE_NORMAL:
		m.AddNormalItem(itemId, num)
	case csvs.ITEMTYPE_COSTUME:
		player.ModCostume.AddCostume(itemId)
	case csvs.ITEMTYPE_ACCESSORY:
		player.ModAccessory.AddAccessory(itemId)
	case csvs.ITEMTYPE_MAGICCARD:
		player.ModMagicCard.AddMagicCard(itemId)
	case csvs.ITEMTYPE_INVENTORY:
		player.ModInventory.AddInventory(itemId)
	default:
		fmt.Println(MODBAG_SYSTEM_COLOR_INFO, "未知物品")
	}

	return
}

func (m *ModBag) HasEnoughItem(itemId int, num int64) bool {

	item, ok := m.BagInfos[itemId]
	if !ok {
		fmt.Println(MODBAG_SYSTEM_COLOR_INFO, "背包不存在该物品")
		return false
	}

	if item.ItemNum < num {
		return false
	}

	return true
}

func (m *ModBag) AddNormalItem(itemId int, num int64) {

	item, ok := m.BagInfos[itemId]
	if !ok {
		itemInfo := csvs.GetItemConfig(itemId)
		m.BagInfos[itemId] = &Item{ItemId: itemInfo.ItemId, ItemNum: num, ItemName: itemInfo.ItemName}
		fmt.Println(MODBAG_SYSTEM_INFO, "添加物品成功，物品名：", itemInfo.ItemName, " 物品数量：", num)
		return
	}
	item.ItemNum += num
	fmt.Println(MODBAG_SYSTEM_INFO, "添加物品成功，物品名：", item.ItemName, " 物品数量：", num)
	return
}

func (m *ModBag) RemoveNormalItem(itemId int, num int64) {

	if !m.HasEnoughItem(itemId, num) {
		fmt.Println(MODBAG_SYSTEM_INFO, "物品数量不足，无法扣除")
	}

	item, _ := m.BagInfos[itemId]
	item.ItemNum -= num
	fmt.Println(MODBAG_SYSTEM_INFO, "添加物品成功，物品名：", item.ItemName, " 物品数量：", num)
	if item.ItemNum == 0 {
		delete(m.BagInfos, itemId) // 删除该物品信息
	}
	return
}

func (m *ModBag) AddGuardianItem(itemId int, num int64) {

	switch itemId {
	case MP:
		m.MP += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加体力成功，当前体力：", m.MP)
	case GOLD:
		m.Gold += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加金币成功，当前金币：", m.Gold)
	case FREEGEMS:
		m.FreeGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加免费钻石成功，当前免费钻石：", m.FreeGems)
	case PAIDGEMS:
		m.PaidGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "增加付费钻石成功，当前付费钻石：", m.PaidGems)
	default:
		fmt.Println(MODBAG_SYSTEM_INFO, "未知背包物品")
	}
}

func (m *ModBag) RemoveGuardianItem(itemId int, num int64) {

	switch itemId {
	case MP:
		if m.MP >= num {
			m.MP -= num
		}
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除体力成功，当前体力：", m.MP)
	case GOLD:
		if m.Gold >= num {
			m.Gold -= num
		}
		m.Gold += num
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除金币成功，当前金币：", m.Gold)
	case FREEGEMS:
		if m.FreeGems >= num {
			m.FreeGems -= num
		}
		m.FreeGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除免费钻石成功，当前免费钻石：", m.FreeGems)
	case PAIDGEMS:
		if m.PaidGems >= num {
			m.PaidGems -= num
		}
		m.PaidGems += num
		fmt.Println(MODBAG_SYSTEM_INFO, "扣除付费钻石成功，当前付费钻石：", m.PaidGems)
	default:
		fmt.Println(MODBAG_SYSTEM_INFO, "未知背包物品")
	}
}

func (m *ModBag) GetCurGuardianItem() int64 {
	return 0
}

// Save 保存背包信息
func (m *ModBag) Save() error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(dir+"/"+"local"+"bag.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	defer func() {
		err = f.Close()
		log.Errorln(err)
	}()
	if err != nil {
		log.Errorln(err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Errorln(err)
	}
	log.Infoln("写入文件成功")
	return nil
}
