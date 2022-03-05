package game

import (
	"fmt"
	"sync"
)

const MODINVENTORY_SYSTEM_INFO = "【库存模块】"

type InventoryInfo struct {
	InventoryId int
}

type ModInventory struct {
	Inventories map[int]*InventoryInfo
	Locker      *sync.RWMutex
}

func (self *ModInventory) HavaInventory(inventoryId int) bool {
	_, ok := self.Inventories[inventoryId]
	return ok
}

// todo 库存是可以拿很多个的，但是东西可以重叠
func (self *ModInventory) AddInventory(inventoryId int) {
	_, ok := self.Inventories[inventoryId]
	if ok {
		fmt.Println(MODINVENTORY_SYSTEM_INFO, "已存在该周边")
	}

	self.Inventories[inventoryId] = &InventoryInfo{InventoryId: inventoryId}
	fmt.Println(MODINVENTORY_SYSTEM_INFO, "添加周边成功")
	return
}
