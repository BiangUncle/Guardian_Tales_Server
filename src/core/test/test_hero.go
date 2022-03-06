package test

import (
	"fmt"
	"server/src/game"
	"server/src/utils"
)

func RUN_HERO_TEST(player *game.Player) {
	for {
		fmt.Println(utils.Cyan("# 输入要进行的操作"))
		fmt.Println(utils.Cyan("# 1. 查询 2. 退出"))
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("#####################")
			for _, v := range player.ModHero.HeroInfos {
				fmt.Println("# ", v.NickName, v.HeroName)
			}
			fmt.Println("#####################")
		case 2:
			return
		}
	}
}
