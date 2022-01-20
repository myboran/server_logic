package game

import (
	"fmt"
	"regexp"
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
		manageBanWord.BanWordBase = []string{"外挂", "外挂工具"}
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
