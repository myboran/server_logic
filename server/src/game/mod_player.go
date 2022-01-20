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
