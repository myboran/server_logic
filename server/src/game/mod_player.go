package game

import (
	"encoding/json"
	"fmt"
	"server_logic/server/src/csvs"
	"time"
)

type ShowRole struct {
	RoleId    int
	RoleLevel int
}

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int   // 大世界等级
	WorldLevelNow  int   // 大世界等级(当前)
	WorldLevelCool int64 // 操作大世界等级冷却时间
	Birth          int
	ShowTeam       []*ShowRole
	HideShowTeam   int // 隐藏开关
	ShowCard       []int

	Prohibit int // 封禁状态
	IsGM     int // GM账号标志
}

//对外接口

// 设置头像
func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端，操作非法
		fmt.Println("当前还没有该头像")
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("设置图标：", player.ModPlayer.Icon)
}

// 设置名片
func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.IsHasCard(cardId) {
		// 通知客户端，操作非法
		fmt.Println("当前还没有该头像")
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("设置名片：", player.ModPlayer.Card)
}

// 设置名字
func (self *ModPlayer) SetName(name string, player *Player) {
	// 验证名字是否合法
	if GetManageBanWord().IsBanWord(name) {
		fmt.Println("名字非法")
		return
	}
	player.ModPlayer.Name = name
	fmt.Println("设置名字：", player.ModPlayer.Name)
}

// 设置签名
func (self *ModPlayer) SetSign(sign string, player *Player) {

	if GetManageBanWord().IsBanWord(sign) {
		return
	}

	player.ModPlayer.Sign = sign
	fmt.Println("设置签名：", player.ModPlayer.Sign)
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

			// 升世界等级
			if self.PlayerLevel >= 25 && self.PlayerLevel%5 == 0 {
				self.WorldLevelNow += 1
				self.WorldLevel += 1
			}
			fmt.Println("当前等级: ", self.PlayerLevel, "当前经验: ", self.PlayerExp, "当前世界等级: ", self.WorldLevelNow)
		} else {
			break
		}
	}
}

func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	if self.WorldLevel < csvs.ReduceWorldLevelStart {
		fmt.Println("降低世界等级失败  当前世界等级---", self.WorldLevel)
		return
	}
	if self.WorldLevel-self.WorldLevelNow >= csvs.ReduceWorldLevelMax {
		fmt.Println("降低世界等级失败 当前世界等级: ", self.WorldLevel, "真实世界等级: ", self.WorldLevelNow)
		return
	}
	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("降低世界等级失败 冷却中")
		return
	}
	self.WorldLevelNow -= 1
	self.WorldLevelCool = time.Now().Unix() + csvs.ReduceWorldLevelCoolTime
	fmt.Println("降低世界等级成功 当前世界等级: ", self.WorldLevel, "真实世界等级: ", self.WorldLevelNow)
	return
}

func (self *ModPlayer) ReturnWorldLevel(player *Player) {
	if self.WorldLevelNow == self.WorldLevel {
		fmt.Println("恢复世界等级失败 当前世界等级: ", self.WorldLevel, "真实世界等级: ", self.WorldLevelNow)
		return
	}
	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("恢复世界等级失败 冷却中")
		return
	}
	self.WorldLevelNow += 1
	self.WorldLevelCool = time.Now().Unix() + csvs.ReduceWorldLevelCoolTime
	fmt.Println("恢复世界等级成功 当前世界等级: ", self.WorldLevel, "真实世界等级: ", self.WorldLevelNow)
	return
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {

	if self.Birth != 0 {
		fmt.Println("已经设置过生日了")
		return
	}

	month := birth / 100
	day := birth % 100

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println("设置生日非法操作: month-", month, " day-", day)

			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println("设置生日非法操作: month-", month, " day-", day)

			return
		}
	case 2:
		if day <= 0 || day > 30 {
			fmt.Println("设置生日非法操作: month-", month, " day-", day)

			return
		}
	default:
		fmt.Println("设置生日非法操作: month-", month, " day-", day)

		return
	}
	self.Birth = birth
	fmt.Println("设置成功, 生日为: ", birth)

	if self.IsBirthDay() {
		fmt.Println("今天是你的生日: ", birth)
	}
}

func (self *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()
	if int(month) == self.Birth/100 && day == self.Birth%100 {
		return true
	}
	return false
}

func (self *ModPlayer) SetShowCard(showCard []int, player *Player) {

	// 需要验证
	cardExist := make(map[int]int)
	newList := make([]int, 0)
	for _, cardId := range showCard {
		_, ok := cardExist[cardId]
		if ok {
			continue
		}
		if !player.ModCard.IsHasCard(cardId) {
			// 通知客户端，操作非法
			continue
		}
		newList = append(newList, cardId)
		cardExist[cardId] = 1
	}
	self.ShowCard = newList
	fmt.Println("玩家名片: ", self.ShowCard)
}

func (self *ModPlayer) SetShowTeam(showTeam []int, player *Player) {

	roleExist := make(map[int]int)
	newList := make([]*ShowRole, 0)
	for _, roleId := range showTeam {
		_, ok := roleExist[roleId]
		if ok {
			continue
		}
		if !player.ModRole.IsHasRole(roleId) {
			fmt.Println("没有此角色 id: ", roleId)
			continue
		}

		showRole := new(ShowRole)
		showRole.RoleId = roleId
		showRole.RoleLevel = player.ModRole.GetRoleLevel(roleId)

		newList = append(newList, showRole)
		roleExist[roleId] = 1
	}
	self.ShowTeam = newList
	data, _ := json.Marshal(self.ShowTeam)
	fmt.Println("展示阵容: ", string(data))
}

func (self *ModPlayer) SetHideShowTeam(isHide int, player *Player) {
	if isHide != csvs.LogicTrue && isHide != csvs.LogicFalse {
		fmt.Println("设置隐藏阵容非法")
		return
	}
	self.HideShowTeam = isHide
}

func (self *ModPlayer) SetProhibit(prohibit int) {
	self.Prohibit = prohibit
}

func (self *ModPlayer) SetIsGM(isGM int) {
	self.IsGM = isGM
}

func (self *ModPlayer) IsCanEnter() bool {
	return int64(self.Prohibit) < time.Now().Unix()
}
