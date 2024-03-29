package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"server/src/core/test"
	"server/src/csvs"
	"server/src/game"
	"server/src/utils"
)

// Init 设置工作路径
func Init() {
	err := os.Chdir("/Users/biang/Desktop/github.com/BiangUncle/Guardian_Tales_Server")
	if err != nil {
		panic(err)
	}
	log.Infoln("设置工作路径")
}

func main() {

	// todo 抽卡模块
	Init()
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
	fmt.Println("###############################")
	fmt.Println("# This is my test mod         #")
	fmt.Println("# Please select your function #")

	for {
		fmt.Println("###############################")
		fmt.Println("# 1. 背包模块 2. 抽奖模块       #")
		fmt.Println("# 3. 英雄模块 4. 浮游城         #")
		fmt.Println("# 0. 退出                     #")
		fmt.Println("###############################")
		var selected int
		fmt.Scan(&selected)
		switch selected {
		case 1:
			fmt.Println(utils.Cyan("# 选择了背包模块"))
			test.RUN_BAG_TEST(self)
		case 2:
			fmt.Println(utils.Cyan("# 选择了抽奖模块"))
			test.RUN_SUMMON_TEST(self)
		case 3:
			fmt.Println(utils.Cyan("# 选择了英雄模块"))
			test.RUN_HERO_TEST(self)
		case 4:
			fmt.Println(utils.Cyan("# 进入浮游城"))
			test.RUN_FLOAT_CITY_TEST(self)
		case 0:
			return
		default:
			fmt.Println(utils.Cyan("# 尽请期待"))
		}
	}
}
