package game

import (
	"fmt"
	"server_logic/server/src/csvs"
	"time"
)

const (
	TASK_STATE_INIT   = 0
	TASK_STATE_DOING  = 1
	TASK_STATE_FINISH = 2
)

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModCard       *ModCard
	ModUniqueTask *ModUniqueTask
	ModRole       *ModRole
}

func NewTestPlayer() *Player {
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	//player.ModUniqueTask.Locker = new(sync.RWMutex)
	player.ModRole = new(ModRole)
	// *******************************
	player.ModPlayer.PlayerLevel = 1 // 初始等级为 1 级
	player.ModPlayer.WorldLevel = 1
	player.ModPlayer.WorldLevelNow = 1
	player.ModPlayer.WorldLevelCool = time.Now().Unix()
	// *******************************
	return player
}

//对外接口
func (self *Player) RecvSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}

func (self *Player) RecvSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}

func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}

func (self *Player) RecvSetSign(sign string) {
	self.ModPlayer.SetSign(sign, self)
}

func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}

func (self *Player) ReturnWorldLevel() {
	self.ModPlayer.ReturnWorldLevel(self)
}

func (self *Player) SetBirth(birth int) {
	self.ModPlayer.SetBirth(birth, self)
}

func (self *Player) SetShowCard(showCard []int) {
	if len(showCard) > csvs.SHOW_SIZE {
		fmt.Println("非法操作")

		return
	}

	self.ModPlayer.SetShowCard(showCard, self)
}

func (self *Player) SetShowTeam(showRole []int) {

	if len(showRole) > csvs.SHOW_SIZE {
		fmt.Println("非法操作")

		return
	}

	self.ModPlayer.SetShowTeam(showRole, self)
}

func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}

func (self *Player) Run() {

	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
		}

	}
}
