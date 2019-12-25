package model

import . "../../common/network/data"

type Game struct {
	gameConstants GameConstants
	mp            Map
	baseUnits     []BaseUnit
	spells        []Spell
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

func (game *Game) HandleInitMessage(msg Message) {
	root := msg.Args[0]
	game.gameConstants = root["gameConstants"].(GameConstants)
	game.mp = root["map"].(Map)
	game.baseUnits = root["baseUnits"].([]BaseUnit)
	spells := root["spells"].([]interface{})
	for _, spell := range spells {
		if spell.(map[string]interface{})["isAreaSpell"].(bool) {
			game.spells = append(game.spells, spell.(AreaSpell))
		} else {
			game.spells = append(game.spells, spell.(UnitSpell))
		}
	}
}
func HandleTurnMessage(msg Message) {

}
func (game Game) ChooseDeck(heroIds []int) {
	//TODO send message
}

func (game Game) GetMyId() int {
	return game.myId
}

func (game Game) GetFriendId() int {
	return game.friendId
}

func (game Game) GetPlayerPosition(playerId int) Cell {
	return game.players[playerId].king.center //TODO use getters?
}

func (game Game) GetPathsFromPlayer(playerId int) []Path {
	paths := make([]Path, 0)
	for _, path := range game.mp.paths {
		startCell := path.cells[0]
		endCell := path.cells[len(path.cells)]
		playerCell := game.GetPlayerPosition(playerId)
		if startCell.Equals(playerCell) || endCell.Equals(playerCell) {
			paths = append(paths, path)
		}
	}
	return paths
}

func (game Game) GetPathsToFriend(playerId int) Path {
	for _, path := range game.mp.paths {
		startCell := path.cells[0]
		endCell := path.cells[len(path.cells)]
		myCell := game.GetPlayerPosition(game.myId)
		friendCell := game.GetPlayerPosition(game.friendId)
		if (startCell.Equals(myCell) && endCell.Equals(friendCell)) ||
			(startCell.Equals(friendCell) && endCell.Equals(myCell)) {
			return path //Is there such a path?
		}
	}
	return Path{} //TODO nil?
}

func (game Game) GetMapHeight() int {
	return game.mp.height
}

func (game Game) GetMapWidth() int {
	return game.mp.width
}

func (game Game) GetPathsCrossing(cell Cell) []Path {
	paths := make([]Path, 0)
	for _, path := range game.mp.paths {
		for _, cell1 := range path.cells {
			if cell1.Equals(cell) {
				paths = append(paths, path)
				break
			}
		}
	}
	return paths
}

func (game Game) GetPlayerUnits(playerId int) []Unit {
	units := make([]Unit, 0)
	for _, unit := range game.mp.units {
		if unit.playerId == playerId {
			units = append(units, unit)
		}
	}
	return units
}

func (game Game) GetCellUnits(cell Cell) []Unit {
	return cell.units
}

func (game Game) GetShortestPathToCell(playerId int, cell Cell) Path {
	/*	ans := -1
		for _, path := range game.mp.paths {
			startCell := path.cells[0]
			endCell := path.cells[len(path.cells)]
			playerCell := game.GetPlayerPosition(playerId)
			if startCell.Equals(playerCell) {
				paths = append(paths,path)
			}
		}
	*/
	panic("fill me up")
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
	return game.currentTurn
}

func (game Game) GetMaxTurns() int {
	return game.gameConstants.maxTurns
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
