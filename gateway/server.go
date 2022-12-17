package main

import (
	log "github.com/sirupsen/logrus"
	"server/src/game"
	"sync"
)

type Server struct {
}

var server *Server
var serverOnce sync.Once

func ServerInstance() *Server {
	serverOnce.Do(func() {
		server = &Server{}
	})
	return server
}

// Hero 英雄结构体
type Hero struct {
	HeroName string `json:"hero_name"` // 英雄名称
}

func (s *Server) Summon(typ string) ([]Hero, error) {
	// 进行抽奖
	heroList, err := game.SummonTenTimes()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	// 转成 HTTP 专用格式
	heroListHTTP := make([]Hero, 10)
	for i, hero := range heroList {
		heroListHTTP[i] = Hero{
			HeroName: hero,
		}
	}

	return heroListHTTP, nil
}
