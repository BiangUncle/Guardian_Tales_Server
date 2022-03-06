package test

import (
	"fmt"
	"server/src/game"
	"server/src/utils"
)

func RUN_BAG_TEST(player *game.Player) {

	for {
		fmt.Println(utils.Cyan("# 输入要进行的操作"))
		fmt.Println(utils.Cyan("# 1. 查询 2. 添加 3. 退出"))
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			RUN_CHECK_ITEM(player)
		case 2:
			RUN_ADD_ITEM(player)
		case 3:
			return
		default:
			fmt.Println(utils.Cyan("# 未知操作"))
		}
	}
}

func RUN_CHECK_ITEM(player *game.Player) {
	for {
		fmt.Println(utils.Cyan("# 输入要查询的东西"))
		fmt.Println(utils.Cyan("# 1. 体力,钻石,金币 2. 退出"))
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println(fmt.Sprintf("体力 %d/%d 钻石 %d 金币 %d",
				player.ModBag.MP,
				player.ModBag.MPLimit,
				player.ModBag.FreeGems+player.ModBag.PaidGems,
				player.ModBag.Gold),
			)
		case 2:
			return
		}
	}
}

func RUN_ADD_ITEM(player *game.Player) {
	for {
		fmt.Println(utils.Cyan("# 输入要添加的物品ID"))
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
			return
		}
	}
}
