package main

import (
	"fmt"
	"os"
	"server/src/core/test"
	"server/src/csvs"
	"server/src/game"
	"server/src/utils"
)

func init() {
	err := os.Chdir("C:\\Users\\Biang\\Desktop\\github.com\\BiangUncle\\Guardian_Tales_Server")
	if err != nil {
		panic(err)
	}
}

func main() {

	// todo 抽卡模块
	csvs.CheckLoadCsv()
	player := game.NewTestPlayer()
	RUN(player)

	//ticker := time.NewTicker(time.Second * 10)
	//
	//for {
	//	select {
	//	case <-ticker.C:
	//		fmt.Println(time.Now().Unix())
	//		//if time.Now().Unix()%3 == 0 {
	//		//	player.ModGuardian.AddExp(50)
	//		//}
	//		//if time.Now().Unix()%5 == 0 {
	//		//	player.RecvUpgrade()
	//		//}
	//	}
	//}
}

func RUN(self *game.Player) {

	for {
		fmt.Println("###############################")
		fmt.Println("# This is my test mod         #")
		fmt.Println("# Please select your function #")
		fmt.Println("###############################")
		fmt.Println("# 1. 背包模块 2. 尽请期待        #")
		fmt.Println("###############################")
		var selected int
		fmt.Scan(&selected)
		switch selected {
		case 1:
			fmt.Println(utils.Cyan("# 选择了背包模块"))
			test.RUN_BAG_TEST(self)
		default:
			fmt.Println(utils.Cyan("# 尽请期待"))
		}
		break
	}
}
