package model

import . "../../common/network/data"

type Game struct {
	gameConstants GameConstants
	mp            Map
	baseUnits     []BaseUnit
	spells        []Spell
	areaSpells    []AreaSpell
	unitSpells    []UnitSpell

	currentTurn    int
	playerUnits    [4][]Unit
	castAreaSpells [4][]CastAreaSpell
	castUnitSpells [4][]CastUnitSpell

	gotRangeUpgrade         bool
	gotDamageUpgrade        bool
	availableRangeUpgrades  int
	availableDamageUpgrades int
	remainingAP             int

	players []Player

	myId, friendId, firstEnemy, secondEnemy int
}

func (game *Game) HandleInitMessage(msg Message) {
	root := msg.Args[0]
	game.gameConstants = root["gameConstants"].(GameConstants)
	game.mp = root["map"].(Map)
	for _, king := range game.mp.kings {
		game.players[king.playerId] = Player{king: king, playerId: king.playerId}
	}
	game.myId = game.mp.kings[0].playerId
	game.friendId = game.mp.kings[1].playerId
	game.firstEnemy = game.mp.kings[2].playerId
	game.secondEnemy = game.mp.kings[3].playerId

	game.baseUnits = root["baseUnits"].([]BaseUnit)
	spells := root["spells"].([]interface{})
	for _, spell := range spells {
		if spell.(map[string]interface{})["isAreaSpell"].(bool) {
			game.spells = append(game.spells, spell.(AreaSpell))
			game.areaSpells = append(game.areaSpells, spell.(AreaSpell))
		} else {
			game.spells = append(game.spells, spell.(UnitSpell))
			game.unitSpells = append(game.unitSpells, spell.(UnitSpell))
		}
	}
}
func (game *Game) HandleTurnMessage(msg Message) {
	root := msg.Args[0]
	game.currentTurn = root["currentTurn"].(int)
	game.players[game.myId].deck = root["deck"].([]int)
	game.players[game.myId].hand = root["hand"].([]int)
	kings := root["kings"].([]interface{})
	for _, king := range kings {
		playerId := king.(map[string]interface{})["playerId"].(int)
		game.players[playerId].king.isAlive = king.(map[string]interface{})["isAlive"].(bool)
		game.players[playerId].king.hp = king.(map[string]interface{})["hp"].(int)
	}

	game.mp.units = []Unit{}
	for i := 0; i < 4; i++ {
		game.playerUnits[i]=[]Unit{}
		game.castUnitSpells[i]=[]CastUnitSpell{}
		game.castAreaSpells[i]=[]CastAreaSpell{}
	}

	units := root["units"].([]interface{}) // get baseUnit by typeId
	for _, unit := range units {
		typeId := unit.(map[string]interface{})["typeId"].(int)
		playerId := unit.(map[string]interface{})["playerId"].(int)
		baseUnit := game.getBaseUnitByTypeId(typeId)
		tmpUnit := unit.(Unit)
		tmpUnit.baseUnit = baseUnit
		//tmpUnit.rng = unit.(map[string]interface{})["range"].(int)
		game.playerUnits[playerId] = append(game.playerUnits[playerId], tmpUnit)
		game.mp.units = append(game.mp.units, tmpUnit)
		/*TODO*/ // add units to cells or to game
	}
	castSpells := root["castSpells"].([]interface{})
	for _, castSpell := range castSpells {
		typeId := castSpell.(map[string]interface{})["typeId"].(int)
		playerId := castSpell.(map[string]interface{})["casterId"].(int)
		if game.isUnitSpell(typeId) {
			game.castUnitSpells[playerId] = append(game.castUnitSpells[playerId], castSpell.(CastUnitSpell))
		} else {
			game.castAreaSpells[playerId] = append(game.castAreaSpells[playerId], castSpell.(CastAreaSpell))
		}
	}
	game.players[game.myId].receivedSpell = root["receivedSpell"].(int)
	game.players[game.friendId].receivedSpell = root["friendReceivedSpell"].(int)
	var tmpSpells []Spell
	mySpells := root["mySpells"].([]int)
	for _, spellId := range mySpells {
		tmpSpells = append(tmpSpells, game.getSpellById(spellId))
	}
	game.players[game.myId].spells = tmpSpells
	tmpSpells = []Spell{}
	friendSpells := root["friendSpells"].([]int)
	for _, spellId := range friendSpells {
		tmpSpells = append(tmpSpells, game.getSpellById(spellId))
	}
	game.players[game.friendId].spells = tmpSpells
	game.gotRangeUpgrade = root["gotRangeUpgrade"].(bool)
	game.gotDamageUpgrade = root["gotDamageUpgrade"].(bool)
	game.availableRangeUpgrades = root["availableRangeUpgrades"].(int)
	game.availableDamageUpgrades = root["availableDamageUpgrades"].(int)
	game.remainingAP = root["remainingAP"].(int)
}

func (game *Game) getSpellById(typeId int) Spell {
	for _, spell := range game.spells {
		if spell.GetTypeId() == typeId {
			return spell
		}
	}
	return AreaSpell{} /*TODO*/ // WTF to do
}

func (game *Game) isUnitSpell(typeId int) bool {
	for _, unitSpell := range game.unitSpells {
		if unitSpell.typeId == typeId {
			return true
		}
	}
	return false
}

func (game *Game) getBaseUnitByTypeId(typeId int) BaseUnit {
	for _, baseUnit := range game.baseUnits {
		if baseUnit.typeId == typeId {
			return baseUnit
		}
	}
	return BaseUnit{} //TODO
}

func (game *Game) ChooseDeck(heroIds []int) {
	//TODO send message
}
func (game *Game) PlayUnit(typeId, pathId int) int {
	panic("implement me")
}
func (game *Game) CastUnitSpell(unitId, pathId, index, spellId int) {
	panic("implement me")
}

func (game *Game) CastAreaSpell(center Cell, spellId int) {
	panic("implement me")
}

func (game *Game) getFriendId(playerId int) int {
	return playerId ^ 2 //TODO make sure this is valid
}

func (game *Game) GetMyId() int {
	return game.myId
}

func (game *Game) GetFriendId() int {
	return game.friendId
}

func (game *Game) GetPlayerPosition(playerId int) Cell {
	return game.players[playerId].king.center //TODO use getters?
}

func (game *Game) GetPathsFromPlayer(playerId int) []Path { //TODO friend paths
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

func (game *Game) GetPathsToFriend(playerId int) Path {
	for _, path := range game.mp.paths {
		startCell := path.cells[0]
		endCell := path.cells[len(path.cells)]
		myCell := game.GetPlayerPosition(playerId)
		friendCell := game.GetPlayerPosition(game.getFriendId(playerId))
		if (startCell.Equals(myCell) && endCell.Equals(friendCell)) ||
			(startCell.Equals(friendCell) && endCell.Equals(myCell)) {
			return path //TODO Is there such a path?
		}
	}
	return Path{} //TODO nil?
}

func (game *Game) GetMapHeight() int {
	return game.mp.height
}

func (game *Game) GetMapWidth() int {
	return game.mp.width
}

func (game *Game) GetPathsCrossing(cell Cell) []Path {
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

func (game *Game) GetPlayerUnits(playerId int) []Unit {
	units := make([]Unit, 0)
	for _, unit := range game.mp.units {
		if unit.playerId == playerId {
			units = append(units, unit)
		}
	}
	return units
}

func (game *Game) GetCellUnits(cell Cell) []Unit {
	return cell.units
}

func (game *Game) GetShortestPathToCell(playerId int, cell Cell) Path {
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

func (game *Game) GetMaxAP() int {
	return game.gameConstants.maxAP
}

func (game *Game) GetRemainingAP(playerId int) int {
	return game.players[playerId].ap
}

func (game *Game) GetHand() []int {
	return game.players[game.myId].hand
}

func (game *Game) GetDeck() []int {
	return game.players[game.myId].deck
}

func (game *Game) GetCurrentTurn() int {
	return game.currentTurn
}

func (game *Game) GetMaxTurns() int {
	return game.gameConstants.maxTurns
}

func (game *Game) GetTurnTimeout() int64 { //TODO pickTimeout?
	return game.gameConstants.turnTimeout
}

func (game *Game) GetRemainingTime() int64 {
	panic("implement me")
}

func (game *Game) GetPlayerHP(playerId int) int {
	return game.players[playerId].king.hp
}

func (game *Game) GetAreaSpellTargets(center Cell, spell Spell) []Unit {
	panic("implement me")
}

func (game *Game) GetRemainingTurnsToUpgrade() int {
	return game.gameConstants.turnsToUpgrade - game.currentTurn%game.gameConstants.turnsToUpgrade
}

func (game *Game) GetRemainingTurnsToGetSpell() int {
	return game.gameConstants.turnsToSpell - game.currentTurn%game.gameConstants.turnsToSpell
}

func (game *Game) GetCastAreaSpell(playerId int) CastAreaSpell {
	panic("implement me")
}

func (game *Game) GetCastUnitSpell(playerId int) CastUnitSpell {
	panic("implement me")
}

func (game *Game) GetDeployedUnits(playerId int) []Unit {
	panic("implement me")
}

func (game *Game) GetActiveSpellsOnCell(cell Cell) []CastAreaSpell { //TODO Delete? or active abilities on units
	panic("implement me")
}

func (game *Game) GetUpgradeTokenNumber(playerId int) int {
	panic("implement me")
}

func (game *Game) GetSpells() []Spell {
	panic("implement me")
}

func (game *Game) GetReceivedSpell() Spell {
	panic("implement me")
}

func (game *Game) GetFriendReceivedSpell() Spell {
	panic("implement me")
}
