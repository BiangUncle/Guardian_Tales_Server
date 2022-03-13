package test

import (
	"fmt"
	"server/src/csvs"
	"server/src/game"
	"server/src/utils"
)

func RUN_FLOAT_CITY_TEST(player *game.Player) {
	for {
		fmt.Println(utils.Cyan("# 输入要进行的操作"))
		fmt.Println(utils.Cyan("# 1. 查询建筑 2. 建造 3. 退出"))
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			player.ModModFloatCity.ShowAllBuilding()
		case 2:
			RUN_BUILD_PRODUCT_BUILD(player)
		case 3:
			return
		default:
			fmt.Println(utils.Cyan("# 无效操作"))
		}
	}
}

func RUN_BUILD_PRODUCT_BUILD(player *game.Player) {
	fmt.Println(utils.Cyan("# 输入要建筑的建筑"))
	var buildId int
	fmt.Scan(&buildId)
	if !csvs.HaveBuild(buildId) {
		fmt.Println(utils.Cyan("# 无效建筑"))
		return
	}

	fmt.Println(utils.Cyan("# 输入要建筑的的位置"))
	var loc int
	fmt.Scan(&loc)

	player.ModModFloatCity.Build(game.BUILD_TYPE_PRODUCT, buildId, loc)
}
