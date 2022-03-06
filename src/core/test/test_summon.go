package test

import (
	"fmt"
	"server/src/game"
	"server/src/utils"
)

func RUN_SUMMON_TEST(player *game.Player) {
	for {
		fmt.Println(utils.Cyan("# 输入要抽奖的次数: 1. 一次 2. 十次 10. 返回"))
		var times int
		fmt.Scan(&times)
		switch times {
		case 1:
			player.ModSummon.StartOneSummon(player)
		case 2:
			player.ModSummon.StartTenSummon(player)
		case 10:
			return
		default:
			fmt.Println(utils.Cyan("# 未知操作"))
		}
	}
}
