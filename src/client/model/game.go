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

	currentTurn             int
	playerUnits             [4][]Unit
	castAreaSpell           [4]CastAreaSpell
	castUnitSpell           [4]CastUnitSpell
	castSpells              []CastSpell
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
	for _, king := range game.mp.Kings {
		game.players[king.PlayerId] = Player{King: king, PlayerId: king.PlayerId}
	}
	game.myId = game.mp.Kings[0].PlayerId
	game.friendId = game.mp.Kings[1].PlayerId
	game.firstEnemy = game.mp.Kings[2].PlayerId
	game.secondEnemy = game.mp.Kings[3].PlayerId

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
	game.players[game.myId].Deck = root["deck"].([]int)
	game.players[game.myId].Hand = root["hand"].([]int)
	kings := root["kings"].([]interface{})
	for _, king := range kings {
		playerId := king.(map[string]interface{})["playerId"].(int)
		game.players[playerId].King.IsAlive = king.(map[string]interface{})["isAlive"].(bool)
		game.players[playerId].King.Hp = king.(map[string]interface{})["hp"].(int)
	}

	game.mp.Units = make([]Unit, 0)
	for i := 0; i < 4; i++ {
		game.playerUnits[i] = make([]Unit, 0)
		game.castUnitSpell[i] = CastUnitSpell{}
		game.castAreaSpell[i] = CastAreaSpell{}
	}
	game.mp.UnitsInCell = make([][][]Unit, 0)
	for i := 0; i < game.GetMapWidth(); i++ {
		game.mp.UnitsInCell = append(game.mp.UnitsInCell, make([][]Unit, 0))
		for j := 0; j < game.GetMapHeight(); j++ {
			game.mp.UnitsInCell[i] = append(game.mp.UnitsInCell[i], make([]Unit, 0))
		}
	}
	units := root["units"].([]interface{}) // get baseUnit by TypeId
	for _, unit := range units {
		typeId := unit.(map[string]interface{})["TypeId"].(int)
		playerId := unit.(map[string]interface{})["playerId"].(int)
		baseUnit := game.getBaseUnitByTypeId(typeId)
		tmpUnit := unit.(Unit)
		tmpUnit.BaseUnit = baseUnit
		game.playerUnits[playerId] = append(game.playerUnits[playerId], tmpUnit)
		game.mp.Units = append(game.mp.Units, tmpUnit)
		game.mp.UnitsInCell[tmpUnit.Cell.Row][tmpUnit.Cell.Col] =
			append(game.mp.UnitsInCell[tmpUnit.Cell.Row][tmpUnit.Cell.Col], tmpUnit) //TODO pointers or ID
	}
	castSpells := root["castSpells"].([]interface{})
	for i := 0; i < 4; i++ {
		game.castUnitSpell[i] = CastUnitSpell{} //TODO nil?
		game.castAreaSpell[i] = CastAreaSpell{}
	}
	game.castSpells = make([]CastSpell, 0)

	for _, castSpell := range castSpells {
		typeId := castSpell.(map[string]interface{})["TypeId"].(int)
		playerId := castSpell.(map[string]interface{})["casterId"].(int)
		thisTurn := castSpell.(map[string]interface{})["wasCastThisTurn"].(bool)
		if game.isUnitSpell(typeId) {
			if thisTurn {
				game.castUnitSpell[playerId] = castSpell.(CastUnitSpell)
			}
			castSpells = append(castSpells, castSpell.(CastUnitSpell))
		} else {
			if thisTurn {
				game.castAreaSpell[playerId] = castSpell.(CastAreaSpell)
			}
			castSpells = append(castSpells, castSpell.(CastAreaSpell))
		}
	}
	game.players[game.myId].ReceivedSpell = root["receivedSpell"].(int)
	game.players[game.friendId].ReceivedSpell = root["friendReceivedSpell"].(int)
	var tmpSpells = make([]Spell, 0)
	mySpells := root["mySpells"].([]int)
	for _, spellId := range mySpells {
		tmpSpells = append(tmpSpells, game.getSpellById(spellId))
	}
	game.players[game.myId].Spells = tmpSpells
	tmpSpells = make([]Spell, 0)
	friendSpells := root["friendSpells"].([]int)
	for _, spellId := range friendSpells {
		tmpSpells = append(tmpSpells, game.getSpellById(spellId))
	}
	game.players[game.friendId].Spells = tmpSpells
	game.gotRangeUpgrade = root["gotRangeUpgrade"].(bool)
	game.gotDamageUpgrade = root["gotDamageUpgrade"].(bool)
	game.availableRangeUpgrades = root["availableRangeUpgrades"].(int)
	game.availableDamageUpgrades = root["availableDamageUpgrades"].(int)
	game.remainingAP = root["remainingAP"].(int)
	game.startTime = time.Now().UnixNano()
}

func (game Game) ChooseDeck(heroIds []int) {
	i := make([]interface{}, 0)
	for _, v := range heroIds {
		i = append(i, v)
	}
	msg := Message{Name: "pick", Args: i} //TODO check server message format
	game.sender(msg)
}
func (game Game) PutUnit(typeId, pathId int) {
	msg := Message{Name: "putUnit", Args: []int{typeId, pathId}, Turn: game.currentTurn} //TODO named args?
	game.sender(msg)
}
func (game Game) CastUnitSpell(unitId, pathId, index, spellId int) {
	path := game.getPathById(pathId)
	if len(path.Cells) <= index {
		return
	}
	cell := path.Cells[index]
	msg := Message{Name: "castSpell",
		Args: []interface{}{spellId, []int{cell.Row, cell.Col}, unitId, pathId}, Turn: game.currentTurn}
	game.sender(msg)
}

func (game Game) CastAreaSpell(center Cell, spellId int) {
	msg := Message{Name: "castSpell",
		Args: []interface{}{spellId, []int{center.Row, center.Col}, -1, -1}, Turn: game.currentTurn}
	game.sender(msg)
}

func (game Game) UpgradeUnitRange(unitId int) {
	msg := Message{Name: "rangeUpgrade", Args: []interface{}{unitId}, Turn: game.currentTurn}
	game.sender(msg)
}

func (game Game) UpgradeUnitDamage(unitId int) {
	msg := Message{Name: "damageUpgrade", Args: []interface{}{unitId}, Turn: game.currentTurn}
	game.sender(msg)
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
		if unitSpell.TypeId == typeId {
			return true
		}
	}
	return false
}

func (game Game) getBaseUnitByTypeId(typeId int) BaseUnit {
	for _, baseUnit := range game.baseUnits {
		if baseUnit.TypeId == typeId {
			return baseUnit
		}
	}
	var baseUnit BaseUnit
	return baseUnit
}

func (game Game) getPathById(pathId int) Path {
	for _, path := range game.mp.Paths {
		if pathId == path.PathId {
			return path
		}
	}
	var path Path
	return path
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
	return game.players[playerId].King.Center //TODO use getters?
}

func (game Game) GetPathsFromPlayer(playerId int) []Path { //TODO friend paths
	paths := make([]Path, 0)
	for _, path := range game.mp.Paths {
		startCell := path.Cells[0]
		endCell := path.Cells[len(path.Cells)-1]
		playerCell := game.GetPlayerPosition(playerId)
		if startCell == playerCell || endCell == playerCell {
			paths = append(paths, path)
		}
	}
	return paths
}

func (game Game) GetPathToFriend(playerId int) Path {
	for _, path := range game.mp.Paths {
		startCell := path.Cells[0]
		endCell := path.Cells[len(path.Cells)-1]
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
	return game.mp.Height
}

func (game Game) GetMapWidth() int {
	return game.mp.Width
}

func (game Game) GetPathsCrossingCell(cell Cell) []Path {
	paths := make([]Path, 0)
	for _, path := range game.mp.Paths {
		for _, cell1 := range path.Cells {
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
	for _, unit := range game.mp.Units {
		if unit.PlayerId == playerId {
			units = append(units, unit)
		}
	}
	return units
}

func (game Game) GetCellUnits(cell Cell) []Unit {
	if !game.isValid(cell) {
		return []Unit{}
	}
	return game.mp.UnitsInCell[cell.Row][cell.Col]
}

func (game Game) GetShortestPathToCell(playerId int, cell Cell) Path {
	var ans Path
	var minAns = -1
	friendPathLen := len(game.GetPathToFriend(playerId).Cells)
	for _, path := range game.mp.Paths {
		startCell := path.Cells[0]
		endCell := path.Cells[len(path.Cells)-1]
		playerCell := game.GetPlayerPosition(playerId)
		friendCell := game.GetPlayerPosition(game.getFriendId(playerId))
		if startCell == playerCell {
			for i := range path.Cells {
				if path.Cells[i] == cell && (minAns == -1 || i < minAns) {
					minAns = i
					ans = path
					break
				}
			}
		}
		if endCell == playerCell {
			lng := len(path.Cells) - 1
			for i := range path.Cells {
				if path.Cells[lng-i] == cell && (minAns == -1 || i < minAns) {
					minAns = i
					ans = path
					break
				}
			}
		}
		if startCell == friendCell {
			for i := range path.Cells {
				if path.Cells[i] == cell && (minAns == -1 || i+friendPathLen < minAns) {
					minAns = i + friendPathLen
					ans = path
					break
				}
			}
		}
		if endCell == friendCell {
			lng := len(path.Cells) - 1
			for i := range path.Cells {
				if path.Cells[lng-i] == cell && (minAns == -1 || i+friendPathLen < minAns) {
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
	return game.gameConstants.MaxAP
}

func (game Game) GetRemainingAP(playerId int) int {
	return game.players[playerId].Ap
}

func (game Game) GetHand() []BaseUnit {
	hand := make([]BaseUnit, 0)
	for _, id := range game.players[game.myId].Hand {
		hand = append(hand, game.getBaseUnitByTypeId(id))
	}
	return hand
}

func (game Game) GetDeck() []BaseUnit {
	deck := make([]BaseUnit, 0)
	for _, id := range game.players[game.myId].Deck {
		deck = append(deck, game.getBaseUnitByTypeId(id))
	}
	return deck
}

func (game Game) GetCurrentTurn() int {
	return game.currentTurn
}

func (game Game) GetMaxTurns() int {
	return game.gameConstants.MaxTurns
}

func (game Game) GetTurnTimeout() int64 {
	return game.gameConstants.TurnTimeout
}

func (game Game) GetRemainingTime() int64 {
	if game.currentTurn == 0 {
		return game.GetPickTimeout() + (time.Now().UnixNano()-game.startTime)/1e6
	} else {
		return game.GetTurnTimeout() + (time.Now().UnixNano()-game.startTime)/1e6
	}
}

func (game Game) GetPlayerHP(playerId int) int {
	return game.players[playerId].King.Hp
}

func (game Game) GetRemainingTurnsToUpgrade() int {
	return game.gameConstants.TurnsToUpgrade - game.currentTurn%game.gameConstants.TurnsToUpgrade
}

func (game Game) GetRemainingTurnsToGetSpell() int {
	return game.gameConstants.TurnsToSpell - game.currentTurn%game.gameConstants.TurnsToSpell
}

func (game Game) GetCastAreaSpell(playerId int) CastAreaSpell {
	return game.castAreaSpell[playerId]
}

func (game Game) GetCastUnitSpell(playerId int) CastUnitSpell {
	return game.castUnitSpell[playerId]
}

func (game Game) GetReceivedSpell() Spell {
	return game.getSpellById(game.players[game.myId].ReceivedSpell)
}

func (game Game) GetFriendReceivedSpell() Spell {
	return game.getSpellById(game.players[game.friendId].ReceivedSpell)
}

func (game Game) GetFirstEnemy() int {
	return game.firstEnemy
}

func (game Game) GetSecondEnemy() int {
	return game.secondEnemy
}

func (game Game) GetPickTimeout() int64 {
	return game.gameConstants.PickTimeout
}

func (game Game) GetAreaSpellTargets(center Cell, spellId int) []Unit {
	units := make([]Unit, 0)
	spell := game.getSpellById(spellId).(AreaSpell)
	for i := center.Row - spell.Range; i <= center.Row+spell.Range; i++ {
		for j := center.Col - spell.Range; j <= center.Col+spell.Range; j++ {
			units = append(units, game.GetCellUnits(Cell{i, j})...)
		}
	}
	return units
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

func (game Game) GetPlayerDuplicateUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.Units {
		if unit.IsDuplicate && unit.PlayerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) GetPlayerHastedUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.Units {
		if unit.IsHasted && unit.PlayerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) GetPlayerPlayedUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.mp.Units {
		if unit.WasPlayedThisTurn && unit.PlayerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) isValid(cell Cell) bool {
	return cell.Row >= 0 && cell.Row < game.GetMapWidth() && cell.Col >= 0 && cell.Col < game.GetMapHeight()
}

func (game Game) getUnitById(unitId int) Unit {
	for _, unit := range game.mp.Units {
		if unit.UnitId == unitId {
			return unit
		}
	}
	var unit Unit
	return unit
}
func (game Game) getCastSpellById(id int) CastSpell {
	for _, c := range game.castSpells {
		if c.GetId() == id {
			return c
		}
	}
	return nil
}
func (game Game) GetCastSpellsOnUnit(unitId int) []CastSpell {
	castSpells := make([]CastSpell, 0)
	for _, c := range game.getUnitById(unitId).AffectedSpells {
		castSpells = append(castSpells, game.getCastSpellById(c))
	}
	return castSpells
}

func (game Game) GetUnitTarget(unitId int) Unit {
	return game.getUnitById(game.getUnitById(unitId).Target)
}

func (game Game) GetUnitTargetCell(unitId int) Cell {
	return game.getUnitById(unitId).TargetCell
}

func (game Game) GetKingTarget(playerId int) Unit {
	return game.getUnitById(game.players[playerId].King.Target)
}

func (game Game) GetKingTargetCell(playerId int) Cell {
	return game.GetKingTarget(playerId).Cell
}

func (game Game) GetKingUnitIsAttacking(unitId int) int {
	unit := game.getUnitById(unitId)
	if unit.Target > 4 || unit.Target < 0 {
		return -1
	}
	return unit.Target
}
