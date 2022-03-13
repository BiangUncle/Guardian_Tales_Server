package game

import (
	"fmt"
	"time"
)

const (
	BUILD_STATUS_BUILDING = 1
	BUILD_STATUS_FINISH   = 2

	BUILD_TYPE_PRODUCT  = 1
	BUILD_TYPE_LANDMARK = 1
)

type Build struct {
	Location  int
	Status    int
	KeyId     int
	Level     int
	Name      string
	BuildType int

	BuildFinishTime int64
}

func (self *Build) GetLocation() int {
	return self.Location
}

func (self *Build) GetStatus() int {
	return self.Status
}

func (self *Build) GetKeyId() int {
	return self.KeyId
}

func (self *Build) GetLevel() int {
	return self.Level
}

func (self *Build) IsFinish() bool {
	return self.Status == BUILD_STATUS_FINISH
}

func (self *Build) GetFinishTime() int64 {
	return self.BuildFinishTime
}

func (self *Build) UpdateStatus() {
	if time.Now().Unix() > self.BuildFinishTime {
		self.Status = BUILD_STATUS_FINISH
	}
}

func (self *Build) GetName() string {
	return self.Name
}

func (self *Build) GetBuildType() int {
	return self.BuildType
}

func (self *Build) GetRemainTime() string {
	now := time.Now().Unix()
	remain := self.BuildFinishTime - now
	if remain <= 0 {
		self.Status = BUILD_STATUS_FINISH
		return "Finish"
	}

	var hour = remain / (60 * 60)
	var minus = remain / 60 % 60
	var second = remain % 60
	return fmt.Sprintf("%d:%d:%d", hour, minus, second)

}

//#########################################################

type ProductBuild struct {
	Build
	Type       int
	TypeOutput int
	GoldOutput int
}

func (self *ProductBuild) GetType() int {
	return self.Type
}

func (self *ProductBuild) GetTypeOutput() int {
	return self.TypeOutput
}

func (self *ProductBuild) GetGoldOutput() int {
	return self.GoldOutput
}

//#########################################################

const (
	SHOP_MAX_LEVEL = 5

	INN_BUILD_ID = iota
	SKYGARDEN_KEY
	DEFENSETOWER_KEY
	POWERTOWER_KEY
	HEROSTATUE_KEY
	SHOP_KEY
)

type LandMarkBuild interface {
	GetFunction() string
}

//#########################################################

type Inn struct {
	Build
	BuildLevelLimit     int
	HeroNumLimit        int
	SplitPointLimitHero int
}

func (self *Inn) GetFunction() string {
	return fmt.Sprintf("当前建筑等级上限【%d】，英雄数量上限【%d】，英雄灵魂点数上限【%d】",
		self.BuildLevelLimit,
		self.HeroNumLimit,
		self.SplitPointLimitHero,
	)
}

//#########################################################

type SkyGarden struct {
	Build
	SplitPointLimitCity int
}

func (self *SkyGarden) GetFunction() string {
	return fmt.Sprintf("当前浮游城灵魂点数上限【%d】", self.SplitPointLimitCity)
}

//#########################################################

type PowerTower struct {
	Build
	AtkAdd int
}

func (self *PowerTower) GetFunction() string {
	return fmt.Sprintf("当前攻击力提升【%d%%】", self.AtkAdd)
}

//#########################################################

type DefenseTower struct {
	Build
	DefAdd int
}

func (self *DefenseTower) GetFunction() string {
	return fmt.Sprintf("当前防御力提升【%d】", self.DefAdd)
}

//#########################################################

type HeroStatue struct {
	Build
	VisitRangeMinus int
}

func (self *HeroStatue) GetFunction() string {
	return fmt.Sprintf("当前访问周期减少【%d分钟】", self.VisitRangeMinus)
}

//#########################################################

const (
	GOODS_LEVLE_LOW  = 1
	GOODS_LEVLE_MID  = 2
	GOODS_LEVLE_HIGH = 3
)

type Shop struct {
	Build
	UpdateTime int64
	Goods      [4]*Item
	GoodsLevel int
}

func (self *Shop) GetFunction() string {
	return fmt.Sprintf("当前商品等级【%d】", self.GoodsLevel)
}
