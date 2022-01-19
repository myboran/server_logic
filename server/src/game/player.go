package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
}

func NewTestPlayer() *Player {
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	// *******************************
	player.ModPlayer.Icon = 0

	// *******************************
	return player
}
