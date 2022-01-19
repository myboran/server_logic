package game

import "fmt"

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

func (self *Player) RecvSetIcon(iconId int) {
	if !self.ModIcon.IsHasIcon(iconId) {
		// 通知客户端，操作非法
		return
	}

	self.ModPlayer.Icon = iconId
	fmt.Println("当前图标：", self.ModPlayer.Icon)
}
