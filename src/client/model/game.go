package model

type Game struct {
	mp             Map
	players        []Player
	myId, friendId int
}

func (g Game) ChooseDeck(...interface{}) {
	//TODO send message
}

func (g Game) GetMyId() int {
	panic("implement me")
}

func (g Game) GetFriendId() int {
	panic("implement me")
}

func (g Game) GetPlayerPosition(playerId int) Cell {
	panic("implement me")
}

func (g Game) GetPathsFromPlayer(playerId int) []Path {
	panic("implement me")
}

func (g Game) GetPathsToFriend(playerId int) Path {
	panic("implement me")
}

func (g Game) GetMapHeight() int {
	panic("implement me")
}

func (g Game) GetMapWidth() int {
	panic("implement me")
}

func (g Game) GetPathsCrossing(cell Cell) []Path {
	panic("implement me")
}

func (g Game) GetPlayerUnits(playerId int) []Unit {
	panic("implement me")
}

func (g Game) GetCellUnits(cell Cell) []Unit {
	panic("implement me")
}

func (g Game) GetShortestPathToCell(playerId int, cell Cell) Path {
	panic("implement me")
}

func (g Game) GetMaxAP(playerId int) int {
	panic("implement me")
}

func (g Game) GetRemainingAP(playerId int) int {
	panic("implement me")
}

func (g Game) GetHand() []Unit {
	panic("implement me")
}

func (g Game) GetDeck() []Unit {
	panic("implement me")
}

func (g Game) PlayUnit(typeId, pathId int) int {
	panic("implement me")
}

func (g Game) GetCurrentTurn() int {
	panic("implement me")
}

func (g Game) GetMaxTurns() int {
	panic("implement me")
}

func (g Game) GetTurnTimeout() int64 {
	panic("implement me")
}

func (g Game) GetRemainingTime() int64 {
	panic("implement me")
}

func (g Game) GetPlayerHP() int {
	panic("implement me")
}
