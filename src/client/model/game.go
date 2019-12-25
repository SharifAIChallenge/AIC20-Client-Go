package model

type Game struct {
	gameConstants GameConstants
	mp            Map
	baseUnits     []BaseUnit
	areaSpells    []AreaSpell
	unitSpells    []UnitSpell

	currentTurn       int
	playerUnits       [4][]Unit
	myCastAreaSpells  []CastAreaSpell
	myCastUnitSpells  []CastUnitSpell
	oppCastAreaSpells []CastAreaSpell
	oppCastUnitSpells []CastUnitSpell

	gotRangedUpgrade        bool
	gotDamageUpgrade        bool
	availableRangeUpgrades  int
	availableDamageUpgrades int

	players        []Player
	myId, friendId int
}

func (game Game) ChooseDeck(heroIds []int) {
	//TODO send message
}

func (game Game) GetMyId() int {
	panic("implement me")
}

func (game Game) GetFriendId() int {
	panic("implement me")
}

func (game Game) GetPlayerPosition(playerId int) Cell {
	panic("implement me")
}

func (game Game) GetPathsFromPlayer(playerId int) []Path {
	panic("implement me")
}

func (game Game) GetPathsToFriend(playerId int) Path {
	panic("implement me")
}

func (game Game) GetMapHeight() int {
	panic("implement me")
}

func (game Game) GetMapWidth() int {
	panic("implement me")
}

func (game Game) GetPathsCrossing(cell Cell) []Path {
	panic("implement me")
}

func (game Game) GetPlayerUnits(playerId int) []Unit {
	panic("implement me")
}

func (game Game) GetCellUnits(cell Cell) []Unit {
	panic("implement me")
}

func (game Game) GetShortestPathToCell(playerId int, cell Cell) Path {
	panic("implement me")
}

func (game Game) GetMaxAP(playerId int) int {
	panic("implement me")
}

func (game Game) GetRemainingAP(playerId int) int {
	panic("implement me")
}

func (game Game) GetHand() []Unit {
	panic("implement me")
}

func (game Game) GetDeck() []Unit {
	panic("implement me")
}

func (game Game) PlayUnit(typeId, pathId int) int {
	panic("implement me")
}

func (game Game) GetCurrentTurn() int {
	panic("implement me")
}

func (game Game) GetMaxTurns() int {
	panic("implement me")
}

func (game Game) GetTurnTimeout() int64 {
	panic("implement me")
}

func (game Game) GetRemainingTime() int64 {
	panic("implement me")
}

func (game Game) GetPlayerHP() int {
	panic("implement me")
}
