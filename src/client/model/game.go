package model

import (
	. "../../common/network/data"
	"time"
)

type Game struct {
	gameConstants GameConstants
	mp            Map
	baseUnits     []BaseUnit
	spells        []Spell
	areaSpells    []AreaSpell
	unitSpells    []UnitSpell

	currentTurn   int
	playerUnits   [4][]Unit
	castAreaSpell [4]CastAreaSpell
	castUnitSpell [4]CastUnitSpell

	gotRangeUpgrade         bool
	gotDamageUpgrade        bool
	availableRangeUpgrades  int
	availableDamageUpgrades int
	remainingAP             int

	startTime int64

	players []Player

	myId, friendId, firstEnemy, secondEnemy int

	sender func(message Message)
}

func NewGame(sender func(message Message)) *Game {
	return &Game{sender: sender}
}

func (game *Game) HandleInitMessage(msg Message) {
	root := msg.Args.(map[string]interface{}) //TODO make it work?
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
	game.startTime = time.Now().UnixNano()
}
func (game *Game) HandleTurnMessage(msg Message) {
	root := msg.Args.(map[string]interface{})
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
		game.playerUnits[i] = []Unit{}
		game.castUnitSpell[i] = CastUnitSpell{}
		game.castAreaSpell[i] = CastAreaSpell{}
	}
	game.mp.unitsInCell = make([][][]Unit, 0)
	for i := 0; i < game.GetMapWidth(); i++ {
		game.mp.unitsInCell = append(game.mp.unitsInCell, make([][]Unit, 0))
		for j := 0; j < game.GetMapHeight(); j++ {
			game.mp.unitsInCell[i] = append(game.mp.unitsInCell[i], make([]Unit, 0))
		}
	}
	units := root["units"].([]interface{}) // get baseUnit by typeId
	for _, unit := range units {
		typeId := unit.(map[string]interface{})["typeId"].(int)
		playerId := unit.(map[string]interface{})["playerId"].(int)
		baseUnit := game.getBaseUnitByTypeId(typeId)
		tmpUnit := unit.(Unit)
		tmpUnit.baseUnit = baseUnit
		game.playerUnits[playerId] = append(game.playerUnits[playerId], tmpUnit)
		game.mp.units = append(game.mp.units, tmpUnit)
		game.mp.unitsInCell[tmpUnit.cell.row][tmpUnit.cell.col] =
			append(game.mp.unitsInCell[tmpUnit.cell.row][tmpUnit.cell.col], tmpUnit) //TODO pointers or ID
	}
	castSpells := root["castSpells"].([]interface{})
	for i := 0; i < 4; i++ {
		game.castUnitSpell[i] = CastUnitSpell{} //TODO nil?
		game.castAreaSpell[i] = CastAreaSpell{}
	}
	for _, castSpell := range castSpells {
		typeId := castSpell.(map[string]interface{})["typeId"].(int)
		playerId := castSpell.(map[string]interface{})["casterId"].(int)
		if game.isUnitSpell(typeId) {
			game.castUnitSpell[playerId] = castSpell.(CastUnitSpell)
		} else {
			game.castAreaSpell[playerId] = castSpell.(CastAreaSpell)
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
	game.startTime = time.Now().UnixNano()
}

func (game Game) getSpellById(typeId int) Spell {
	for _, spell := range game.spells {
		if spell.GetTypeId() == typeId {
			return spell
		}
	}
	var retSpell Spell
	return retSpell
}

func (game Game) isUnitSpell(typeId int) bool {
	for _, unitSpell := range game.unitSpells {
		if unitSpell.typeId == typeId {
			return true
		}
	}
	return false
}

func (game Game) getBaseUnitByTypeId(typeId int) BaseUnit {
	for _, baseUnit := range game.baseUnits {
		if baseUnit.typeId == typeId {
			return baseUnit
		}
	}
	var baseUnit BaseUnit
	return baseUnit
}

func (game Game) ChooseDeck(heroIds []int) {
	i := make([]interface{}, 0)
	for _, v := range heroIds {
		i = append(i, v)
	}
	msg := Message{Name: "chooseDeck", Args: i} //TODO check server message format
	game.sender(msg)
}
func (game Game) PutUnit(typeId, pathId int) {
	msg := Message{Name: "putUnit", Args: []int{typeId, pathId}, Turn: game.currentTurn} //TODO named args?
	game.sender(msg)
}
func (game Game) getPathById(pathId int) Path {
	for _, path := range game.mp.paths {
		if pathId == path.pathId {
			return path
		}
	}
	var path Path
	return path
}

func (game Game) CastUnitSpell(unitId, pathId, index, spellId int) {
	path := game.getPathById(pathId)
	if len(path.cells) <= index {
		return
	}
	cell := path.cells[index]
	msg := Message{Name: "castSpell", Args: []interface{}{spellId, []int{cell.row, cell.col}, unitId, pathId}}
	game.sender(msg)
}

func (game Game) CastAreaSpell(center Cell, spellId int) {
	msg := Message{Name: "castSpell", Args: []interface{}{spellId, []int{center.row, center.col}, -1, -1}}
	game.sender(msg)
}

func (game Game) getFriendId(playerId int) int {
	return playerId ^ 2 //TODO make sure this is valid
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

func (game Game) GetPathsFromPlayer(playerId int) []Path { //TODO friend paths
	paths := make([]Path, 0)
	for _, path := range game.mp.paths {
		startCell := path.cells[0]
		endCell := path.cells[len(path.cells)-1]
		playerCell := game.GetPlayerPosition(playerId)
		if startCell == playerCell || endCell == playerCell {
			paths = append(paths, path)
		}
	}
	return paths
}

func (game Game) GetPathToFriend(playerId int) Path {
	for _, path := range game.mp.paths {
		startCell := path.cells[0]
		endCell := path.cells[len(path.cells)-1]
		myCell := game.GetPlayerPosition(playerId)
		friendCell := game.GetPlayerPosition(game.getFriendId(playerId))
		if (startCell == myCell && endCell == friendCell) ||
			(startCell == friendCell && endCell == myCell) {
			return path //TODO Is there such a path?
		}
	}
	var path Path
	return path
}

func (game Game) GetMapHeight() int {
	return game.mp.height
}

func (game Game) GetMapWidth() int {
	return game.mp.width
}

func (game Game) GetPathsCrossingCell(cell Cell) []Path {
	paths := make([]Path, 0)
	for _, path := range game.mp.paths {
		for _, cell1 := range path.cells {
			if cell1 == cell {
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
	if !game.isValid(cell) {
		return []Unit{}
	}
	return game.mp.unitsInCell[cell.row][cell.col]
}

func (game Game) GetShortestPathToCell(playerId int, cell Cell) Path {
	var ans Path
	var minAns = -1
	friendPathLen := len(game.GetPathToFriend(playerId).cells)
	for _, path := range game.mp.paths {
		startCell := path.cells[0]
		endCell := path.cells[len(path.cells)-1]
		playerCell := game.GetPlayerPosition(playerId)
		friendCell := game.GetPlayerPosition(game.getFriendId(playerId))
		if startCell == playerCell {
			for i := range path.cells {
				if path.cells[i] == cell && (minAns == -1 || i < minAns) {
					minAns = i
					ans = path
					break
				}
			}
		}
		if endCell == playerCell {
			lng := len(path.cells) - 1
			for i := range path.cells {
				if path.cells[lng-i] == cell && (minAns == -1 || i < minAns) {
					minAns = i
					ans = path
					break
				}
			}
		}
		if startCell == friendCell {
			for i := range path.cells {
				if path.cells[i] == cell && (minAns == -1 || i+friendPathLen < minAns) {
					minAns = i + friendPathLen
					ans = path
					break
				}
			}
		}
		if endCell == friendCell {
			lng := len(path.cells) - 1
			for i := range path.cells {
				if path.cells[lng-i] == cell && (minAns == -1 || i+friendPathLen < minAns) {
					minAns = i + friendPathLen
					ans = path
					break
				}
			}
		}
	}
	return ans
}

func (game Game) GetMaxAP() int {
	return game.gameConstants.maxAP
}

func (game Game) GetRemainingAP(playerId int) int {
	return game.players[playerId].ap
}

func (game Game) GetHand() []int {
	return game.players[game.myId].hand
}

func (game Game) GetDeck() []int {
	return game.players[game.myId].deck
}

func (game Game) GetCurrentTurn() int {
	return game.currentTurn
}

func (game Game) GetMaxTurns() int {
	return game.gameConstants.maxTurns
}

func (game Game) GetTurnTimeout() int64 { //TODO pickTimeout?
	return game.gameConstants.turnTimeout
}

func (game Game) GetRemainingTime() int64 {
	if game.currentTurn == 0 {
		return game.GetPickTimeout() + (time.Now().UnixNano()-game.startTime)/1e6
	} else {
		return game.GetTurnTimeout() + (time.Now().UnixNano()-game.startTime)/1e6
	}
}

func (game Game) GetPlayerHP(playerId int) int {
	return game.players[playerId].king.hp
}

func (game Game) GetRemainingTurnsToUpgrade() int {
	return game.gameConstants.turnsToUpgrade - game.currentTurn%game.gameConstants.turnsToUpgrade
}

func (game Game) GetRemainingTurnsToGetSpell() int {
	return game.gameConstants.turnsToSpell - game.currentTurn%game.gameConstants.turnsToSpell
}

func (game Game) GetCastAreaSpell(playerId int) CastAreaSpell {
	return game.castAreaSpell[playerId]
}

func (game Game) GetCastUnitSpell(playerId int) CastUnitSpell {
	return game.castUnitSpell[playerId]
}

func (game Game) GetReceivedSpell() Spell {
	return game.getSpellById(game.players[game.myId].receivedSpell)
}

func (game Game) GetFriendReceivedSpell() Spell {
	return game.getSpellById(game.players[game.friendId].receivedSpell)
}

func (game Game) GetFirstEnemy() int {
	return game.firstEnemy
}

func (game Game) GetSecondEnemy() int {
	return game.secondEnemy
}

func (game Game) GetPickTimeout() int64 {
	return game.gameConstants.pickTimeout
}

func (game Game) GetAreaSpellTargets(center Cell, spellId int) []Unit {
	units := make([]Unit, 0)
	spell := game.getSpellById(spellId).(AreaSpell)
	for i := center.row - spell.rng; i <= center.row+spell.rng; i++ {
		for j := center.col - spell.rng; j <= center.col+spell.rng; j++ {
			units = append(units, game.GetCellUnits(Cell{i, j})...)
		}
	}
	return units
}

func (game Game) GetActivePoisonsOnUnit(unitId int) int {
	return game.getUnitById(unitId).activePoisons
}

func (game Game) GetRangeUpgradeNumber() int {
	return game.availableRangeUpgrades
}

func (game Game) GetDamageUpgradeNumber() int {
	return game.availableDamageUpgrades
}

func (game Game) GetSpellsList() []Spell {
	return game.spells
}

func (game Game) GetSpells() map[Spell]int {
	spellMap := make(map[Spell]int, 0)
	for _, spell := range game.spells {
		var val, ok = spellMap[spell]
		if !ok {
			spellMap[spell] = 1
		} else {
			spellMap[spell] = val + 1
		}
	}
	return spellMap
}

func (game Game) UpgradeUnitRange(unitId int) {
	msg := Message{Name: "rangeUpgrade", Args: []interface{}{unitId}}
	game.sender(msg)
}

func (game Game) UpgradeUnitDamage(unitId int) {
	msg := Message{Name: "damageUpgrade", Args: []interface{}{unitId}}
	game.sender(msg)
}

func (game Game) GetPlayerCloneUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.units {
		if unit.isClone && unit.playerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) GetPlayerHastedUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.units {
		if unit.isHasted && unit.playerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) GetPlayerPoisonedUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.units {
		if unit.activePoisons > 0 && unit.playerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) GetPlayerPlayedUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.units {
		if unit.wasPlayedThisTurn && unit.playerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) isValid(cell Cell) bool {
	return cell.row >= 0 && cell.row < game.GetMapWidth() && cell.col >= 0 && cell.col < game.GetMapHeight()
}

func (game Game) getUnitById(unitId int) Unit {
	for _, unit := range game.mp.units {
		if unit.unitId == unitId {
			return unit
		}
	}
	var unit Unit
	return unit
}
