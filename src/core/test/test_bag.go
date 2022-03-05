package test

import (
	"fmt"
	"server/src/game"
	"server/src/utils"
)

func RUN_BAG_TEST(player *game.Player) {

	for {
		//fmt.Println("###############################")
		fmt.Println(utils.Cyan("# 输入要添加的物品ID             #"))
		var itemId int
		fmt.Scan(&itemId)
		fmt.Println(utils.Cyan("# 输入要添加的物品数量             #"))
		var num int64
		fmt.Scan(&num)
		player.ModBag.AddItem(itemId, num, player)
		fmt.Println(utils.Cyan("# 是否继续添加物品 1. 是  2. 退出   #"))
		var exit int
		fmt.Scan(&exit)
		if exit != 1 {
			break
		}
	}
}
