package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

func NewTestPlayer() *Player {
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	//player.ModIcon = new(ModIcon)
	//player.ModCard = new(ModCard)
	// *******************************

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
