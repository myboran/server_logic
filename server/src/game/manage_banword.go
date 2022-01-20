package game

import (
	"fmt"
	"regexp"
	"server_logic/server/src/csvs"
	"time"
)

// 单例模型 （节省空间）
var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string // 配置生成
	BanWordExtra []string // 更新
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordExtra = []string{"原神"}
	}
	return manageBanWord
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, "---", v)
		if match {
			return match
		}
	}

	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, "---", v)
		if match {
			return match
		}
	}
	return false
}

func (self *ManageBanWord) Run() {
	self.BanWordBase = csvs.GetBanWordBase()
	// 基础词库的更新
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%10 == 0 {
				fmt.Println("更新词库")
			} else {
				fmt.Println("待机")
			}
		}
	}
}
