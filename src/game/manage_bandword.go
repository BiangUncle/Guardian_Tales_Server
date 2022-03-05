package game

import (
	"fmt"
	"regexp"
	"server/src/csvs"
	"time"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"自闭"}
		manageBanWord.BanWordExtra = []string{"坎公"}
	}
	return manageBanWord
}

func (self *ManageBanWord) IsBanWord(txt string) bool {

	for _, word := range self.BanWordBase {
		match, _ := regexp.MatchString(word, txt)
		fmt.Println(SYSTEM_INFO, match, word)
		if match {
			return match
		}
	}

	for _, word := range self.BanWordExtra {
		match, _ := regexp.MatchString(word, txt)
		fmt.Println(SYSTEM_INFO, match, word)
		if match {
			return match
		}
	}
	return false
}

func (self *ManageBanWord) RUN() {

	self.BanWordBase = csvs.GetBandWordBase()

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
		}
	}
}
