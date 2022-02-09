package game

import (
	"fmt"
	"server_logic/server/src/csvs"
)

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelCool int
	Birth          int
	ShowTeam       []int
	ShowCard       []int

	IsProhibit int
	IsGM       int
}

//对外接口

// 设置头像
func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端，操作非法
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标：", player.ModPlayer.Icon)
}

// 设置名片
func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.IsHasCard(cardId) {
		// 通知客户端，操作非法
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("当前名片：", player.ModPlayer.Card)
}

// 设置名字
func (self *ModPlayer) SetName(name string, player *Player) {
	// 验证名字是否合法
	if GetManageBanWord().IsBanWord(name) {
		return
	}
	player.ModPlayer.Name = name
	fmt.Println("当前名字：", player.ModPlayer.Name)
}

// 设置签名
func (self *ModPlayer) SetSign(sign string, player *Player) {

	if GetManageBanWord().IsBanWord(sign) {
		return
	}

	player.ModPlayer.Sign = sign
	fmt.Println("当前签名：", player.ModPlayer.Sign)
}

func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp
	fmt.Println("当前等级: ", self.PlayerLevel, "当前经验: ", self.PlayerExp)
	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel - 1)
		if config == nil {
			break
		}
		// 60 级
		if config.PlayerExp == 0 {
			break
		}
		// 是否完成任务 TODO
		if config.ChapterId > 0 && !player.ModUniqueTask.IsTaskFinish(config.ChapterId) {
			break
		}
		// 升级
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1
			self.PlayerExp -= config.PlayerExp
			fmt.Println("当前等级: ", self.PlayerLevel, "当前经验: ", self.PlayerExp)
		} else {
			break
		}
	}
}
