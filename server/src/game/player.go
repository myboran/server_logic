package game

import (
	"fmt"
	"server_logic/server/src/csvs"
	"strconv"
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
	ModBag        *ModBag
	ModWeapon     *ModWeapon
}

func NewTestPlayer() *Player {
	player := new(Player)

	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfo = make(map[int]*Icon)

	player.ModCard = new(ModCard)
	player.ModCard.CardInfo = make(map[int]*Card)

	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	//player.ModUniqueTask.Locker = new(sync.RWMutex)
	player.ModRole = new(ModRole)
	player.ModRole.RoleInfo = make(map[int]*RoleInfo)
	// *******************************
	player.ModPlayer = new(ModPlayer)
	player.ModPlayer.PlayerLevel = 1 // 初始等级为 1 级
	player.ModPlayer.WorldLevel = 1
	player.ModPlayer.Name = "旅行者"
	player.ModPlayer.WorldLevelNow = 1
	player.ModPlayer.WorldLevelCool = time.Now().Unix()
	// *******************************
	player.ModBag = new(ModBag)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)
	// *******************************
	player.ModWeapon = new(ModWeapon)
	player.ModWeapon.WeaponInfo = make(map[int]*Weapon)
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
	if len(showCard) > csvs.ShowSize {
		fmt.Println("非法操作")

		return
	}

	self.ModPlayer.SetShowCard(showCard, self)
}

func (self *Player) SetShowTeam(showRole []int) {

	if len(showRole) > csvs.ShowSize {
		fmt.Println("非法操作")

		return
	}

	self.ModPlayer.SetShowTeam(showRole, self)
}

func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}

func (self *Player) Run() {
	fmt.Println("个人学习作品: MYBORAN")
	fmt.Println("学习来源:B站 golang大海葵")
	// 监听动作
	for {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println(self.ModPlayer.Name, ",欢迎来到提瓦特大陆,请选择功能： 1基础信息  2背包  3地图(未开放)")
		var modChoose int
		fmt.Scan(&modChoose)
		switch modChoose {
		case 1:
			self.HandleBase()
		case 2:
			self.HandleBag()
		case 3:
			self.HandleMap()
		}
	}
}

func (self *Player) HandleBase() {
	for {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("当前处于基础信息界面,请选择操作：0返回  1查询信息  2设置名字  3设置签名  4头像  5名片  6设置生日")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBaseGetInfo()
		case 2:
			self.HandleBagSetName()
		case 3:
			self.HandleBagSetSign()
		case 4:
			self.HandleBagSetIcon()
		case 5:
			self.HandleBagSetCard()
		case 6:
			self.HandleBagSetBirth()
		}
	}
}

func (self *Player) HandleBag() {
	for {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("当前处于基础信息界面,请选择操作：0返回  1增加物品  2扣除物品  3查看物品")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBagAddItem()
		case 2:
			self.HandleBagRemoveItem()
		case 3:
			self.HandleBagShowItem()
		}
	}
}

func (self *Player) HandleMap() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("向着星辰与深渊,欢迎来到冒险家协会！")
	fmt.Println("当前位置:", "蒙德城")
	fmt.Println("地图模块还没写到......")
}

func (self *Player) HandleBaseGetInfo() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("名字:", self.ModPlayer.Name)
	fmt.Println("等级:", self.ModPlayer.PlayerLevel)
	fmt.Println("大世界等级:", self.ModPlayer.WorldLevelNow)
	if self.ModPlayer.Sign == "" {
		fmt.Println("签名:", "未设置")
	} else {
		fmt.Println("签名:", self.ModPlayer.Sign)
	}

	if self.ModPlayer.Icon == 0 {
		fmt.Println("头像:", "未设置")
	} else {
		fmt.Println("头像:", csvs.GetItemConfig(self.ModPlayer.Icon))
	}

	if self.ModPlayer.Card == 0 {
		fmt.Println("名片:", "未设置")
	} else {
		fmt.Println("名片:", csvs.GetItemConfig(self.ModPlayer.Card), self.ModPlayer.Card)
	}

	if self.ModPlayer.Birth == 0 {
		fmt.Println("生日:", "未设置")
	} else {
		fmt.Println("生日:", self.ModPlayer.Birth/100, "月", self.ModPlayer.Birth%100, "日")
	}
}

func (self *Player) HandleBagSetName() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("请输入名字:")
	var name string
	fmt.Scan(&name)
	self.RecvSetName(name)
}
func (self *Player) HandleBagSetSign() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("请输入签名:")
	var sign string
	fmt.Scan(&sign)
	self.RecvSetSign(sign)
}
func (self *Player) HandleBagSetIcon() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("当前处于基础信息--头像界面,请选择操作：0返回  1查询头像背包  2设置头像")
	var action int
	fmt.Scan(&action)
	switch action {
	case 0:
		return
	case 1:
		self.HandleBagSetIconGetInfo()
	case 2:
		self.HandleBagSetIconSet()
	}
}
func (self *Player) HandleBagSetCard() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("当前处于基础信息--头像界面,请选择操作：0返回  1查询名片背包  2设置名片")
	var action int
	fmt.Scan(&action)
	switch action {
	case 0:
		return
	case 1:
		self.HandleBagSetCardGetInfo()
	case 2:
		self.HandleBagSetCardSet()
	}
}

func (self *Player) HandleBagSetIconGetInfo() {
	fmt.Println("当前拥有头像如下:")
	for _, v := range self.ModIcon.IconInfo {
		config := csvs.GetItemConfig(v.Id)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (self *Player) HandleBagSetIconSet() {

	fmt.Println("请输入头像id:")
	var icon int
	fmt.Scan(&icon)
	self.RecvSetIcon(icon)
}

func (self *Player) HandleBagSetCardGetInfo() {
	fmt.Println("当前拥有名片如下:")
	for _, v := range self.ModCard.CardInfo {
		config := csvs.GetItemConfig(v.CardId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (self *Player) HandleBagSetCardSet() {
	fmt.Println("请输入名片id:")
	var card int
	fmt.Scan(&card)
	self.RecvSetCard(card)
}
func (self *Player) HandleBagSetBirth() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("请输入生日:")
	var day int
	fmt.Scan(&day)
	self.SetBirth(day)
}

func (self *Player) HandleBagAddItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	self.ModBag.AddItem(itemId, int64(itemNum), self)
}

func (self *Player) HandleBagShowItem() {

	data := "您当前有:"
	for _, v := range self.ModBag.BagInfo {
		item := csvs.GetItemConfig(v.ItemId)
		str := "\n" + item.ItemName + ":" + strconv.Itoa(int(v.ItemNum)) + "个   id:" + strconv.Itoa(v.ItemId) + ""
		data = data + str
	}
	fmt.Println(data)
}

func (self *Player) HandleBagRemoveItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	self.ModBag.RemoveItemToBag(itemId, int64(itemNum), self)
}
