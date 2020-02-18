package model

import (
	"encoding/json"
	"fmt"
	"time"

	. "../../common/network/data"
)

type Game struct {
	gameConstants *GameConstants
	Map           *Map
	baseUnits     []*BaseUnit
	spells        []*Spell

	currentTurn             int
	castSpells              []*CastSpell
	gotRangeUpgrade         bool
	gotDamageUpgrade        bool
	availableRangeUpgrades  int
	availableDamageUpgrades int
	rangeUpgradedUnit       int
	damageUpgradedUnit      int

	startTime int64

	players [4]*Player

	myId, friendId, firstEnemy, secondEnemy int
	receivedSpell, friendReceivedSpell      *Spell

	ShortestPaths [4]map[Cell]int
	pathsCrossing map[Cell][]*Path

	unitById      map[int]*Unit
	pathById      map[int]*Path
	spellById     map[int]*Spell
	castSpellById map[int]*CastSpell
	baseUnitById  map[int]*BaseUnit

	sender chan Message
}

func NewGame(sender chan Message) *Game {
	return &Game{sender: sender}
}

func (game *Game) HandleInitMessage(msg Message) {
	root := msg.Args.(map[string]interface{})
	mapToStruct(root["gameConstants"], &game.gameConstants)
	mapToStruct(root["map"], &game.Map)
	game.Map.Cells = make([][]*Cell, 0)
	for i := 0; i < game.Map.RowNum; i++ {
		game.Map.Cells = append(game.Map.Cells, make([]*Cell, 0))
		for j := 0; j < game.Map.ColNum; j++ {
			game.Map.Cells[i] = append(game.Map.Cells[i], &Cell{Row: i, Col: j})
		}
	}
	for _, king := range game.Map.Kings {
		game.players[king.PlayerId] = &Player{King: king, PlayerId: king.PlayerId}
	}
	game.myId = game.Map.Kings[0].PlayerId
	game.friendId = game.Map.Kings[1].PlayerId
	game.firstEnemy = game.Map.Kings[2].PlayerId
	game.secondEnemy = game.Map.Kings[3].PlayerId

	mapToStruct(root["baseUnits"], &game.baseUnits)
	spells := root["spells"].([]interface{})
	game.pathsCrossing = make(map[Cell][]*Path)
	game.baseUnitById = make(map[int]*BaseUnit)
	game.spellById = make(map[int]*Spell)
	game.pathById = make(map[int]*Path)
	for _, spell := range spells {
		var tmpSpell Spell
		mapToStruct(spell, &tmpSpell)
		game.spells = append(game.spells, &tmpSpell)
	}
	for i := 0; i < 4; i++ {
		game.players[i].PathToFriend = game.getPathToFriend(i)
		game.players[i].PathsFromPlayer = game.getPathsFromPlayer(i)
		game.ShortestPaths[i] = make(map[Cell]int)
	}
	game.startTime = time.Now().UnixNano()
}
func (game *Game) HandleTurnMessage(msg Message) {
	fmt.Println("Received Turn Message!")

	game.unitById = make(map[int]*Unit)
	game.castSpellById = make(map[int]*CastSpell)
	for i := 0; i < 4; i++ {
		game.players[i].spellCount = make(map[int]int)
	}
	root := msg.Args.(map[string]interface{})
	mapToStruct(root["currTurn"], &game.currentTurn)

	var tmpCards []int
	game.players[game.myId].Deck = make([]*BaseUnit, 0)
	mapToStruct(root["deck"], &tmpCards)
	for _, unitId := range tmpCards {
		game.players[game.myId].Deck = append(game.players[game.myId].Deck, game.GetBaseUnitById(unitId))
	}

	game.players[game.myId].Hand = make([]*BaseUnit, 0)
	mapToStruct(root["hand"], &tmpCards)
	for _, unitId := range tmpCards {
		game.players[game.myId].Hand = append(game.players[game.myId].Hand, game.GetBaseUnitById(unitId))
	}

	game.Map.Units = make([]*Unit, 0)
	for i := 0; i < 4; i++ {
		game.players[i].Units = make([]*Unit, 0)
		game.players[i].DuplicateUnits = make([]*Unit, 0)
		game.players[i].HastedUnits = make([]*Unit, 0)
		game.players[i].PlayedUnits = make([]*Unit, 0)
		game.players[i].DiedUnits = make([]*Unit, 0)
	}
	game.Map.UnitsInCell = make([][][]*Unit, 0)
	for i := 0; i < game.Map.RowNum; i++ {
		game.Map.UnitsInCell = append(game.Map.UnitsInCell, make([][]*Unit, 0))
		for j := 0; j < game.Map.ColNum; j++ {
			game.Map.UnitsInCell[i] = append(game.Map.UnitsInCell[i], make([]*Unit, 0))
		}
	}
	units := root["units"].([]interface{})
	for _, unit := range units {
		var typeId int
		mapToStruct(unit.(map[string]interface{})["typeId"], &typeId)
		var playerId int
		mapToStruct(unit.(map[string]interface{})["playerId"], &playerId)
		var pathId int
		mapToStruct(unit.(map[string]interface{})["pathId"], &pathId)
		var wasPlayedThisTurn bool
		mapToStruct(unit.(map[string]interface{})["wasPlayedThisTurn"], &wasPlayedThisTurn)
		var wasDamageUpgraded bool
		mapToStruct(unit.(map[string]interface{})["wasDamageUpgraded"], &wasDamageUpgraded)
		var wasRangeUpgraded bool
		mapToStruct(unit.(map[string]interface{})["wasRangeUpgraded"], &wasRangeUpgraded)
		var affectedSpells []int
		mapToStruct(unit.(map[string]interface{})["affectedSpells"], &affectedSpells)
		baseUnit := game.GetBaseUnitById(typeId)
		path := game.getPathById(pathId)
		var tmpUnit Unit
		delete(unit.(map[string]interface{}), "target")
		delete(unit.(map[string]interface{}), "affectedSpells")
		delete(unit.(map[string]interface{}), "typeId")
		delete(unit.(map[string]interface{}), "wasDamageUpgraded")
		delete(unit.(map[string]interface{}), "wasRangeUpgraded")

		mapToStruct(unit, &tmpUnit)
		if wasDamageUpgraded {
			game.GetPlayerById(playerId).DamageUpgradedUnit = &tmpUnit
		}
		if wasRangeUpgraded {
			game.GetPlayerById(playerId).RangeUpgradedUnit = &tmpUnit
		}
		tmpUnit.BaseUnit = baseUnit
		tmpUnit.Path = path
		game.players[playerId].Units = append(game.players[playerId].Units, &tmpUnit)
		if tmpUnit.IsDuplicate {
			game.players[playerId].DuplicateUnits = append(game.players[playerId].DuplicateUnits, &tmpUnit)
		}
		if tmpUnit.IsHasted {
			game.players[playerId].HastedUnits = append(game.players[playerId].HastedUnits, &tmpUnit)
		}
		if wasPlayedThisTurn {
			game.players[playerId].PlayedUnits = append(game.players[playerId].PlayedUnits, &tmpUnit)
		}
		game.Map.Units = append(game.Map.Units, &tmpUnit)
		game.Map.UnitsInCell[tmpUnit.Cell.Row][tmpUnit.Cell.Col] =
			append(game.Map.UnitsInCell[tmpUnit.Cell.Row][tmpUnit.Cell.Col], &tmpUnit)
	}

	for _, unit := range units {
		var unitId int
		mapToStruct(unit.(map[string]interface{})["unitId"], &unitId)
		var targetId int
		mapToStruct(unit.(map[string]interface{})["target"], &targetId)
		target := game.GetUnitById(targetId)
		tmpUnit := game.GetUnitById(unitId)
		targetIsKing := false
		var targetIfKing *King = nil
		for i := 0; i < 4; i++ {
			if targetId == game.players[i].PlayerId {
				targetIsKing = true
				targetIfKing = game.players[i].King
			}
		}
		if targetId != -1 && !targetIsKing {
			tmpUnit.Target = target
		} else if targetIsKing {
			tmpUnit.TargetIfKing = targetIfKing
		}
	}

	deadUnits := root["diedUnits"].([]interface{})
	for _, unit := range deadUnits {
		var typeId int
		mapToStruct(unit.(map[string]interface{})["typeId"], &typeId)
		var playerId int
		mapToStruct(unit.(map[string]interface{})["playerId"], &playerId)
		var pathId int
		mapToStruct(unit.(map[string]interface{})["pathId"], &pathId)
		baseUnit := game.GetBaseUnitById(typeId)
		path := game.getPathById(pathId)
		var tmpUnit Unit
		delete(unit.(map[string]interface{}), "target")
		delete(unit.(map[string]interface{}), "affectedSpells")
		delete(unit.(map[string]interface{}), "typeId")
		delete(unit.(map[string]interface{}), "wasDamageUpgraded")
		delete(unit.(map[string]interface{}), "wasRangeUpgraded")

		mapToStruct(unit, &tmpUnit)
		tmpUnit.BaseUnit = baseUnit
		tmpUnit.Path = path
		game.players[playerId].DiedUnits = append(game.players[playerId].DiedUnits, &tmpUnit)
	}
	kings := root["kings"].([]interface{})
	for _, king := range kings {
		var playerId int
		mapToStruct(king.(map[string]interface{})["playerId"], &playerId)
		game.players[playerId].King.IsAlive = king.(map[string]interface{})["isAlive"].(bool)
		mapToStruct(king.(map[string]interface{})["hp"], &game.players[playerId].King.Hp)
		var targetId int
		mapToStruct(king.(map[string]interface{})["target"], &targetId)
		if targetId != -1 {
			tmpUnit := game.GetUnitById(targetId)
			game.players[playerId].King.Target = tmpUnit
			game.players[playerId].King.TargetCell = tmpUnit.Cell
		}
	}
	castSpells, ok := root["castSpells"].([]interface{})
	if !ok {
		castSpells = make([]interface{}, 0)
	}
	for i := 0; i < 4; i++ {
		game.players[i].CastUnitSpell = nil
		game.players[i].CastAreaSpell = nil
	}
	game.castSpells = make([]*CastSpell, 0)

	for _, castSpell := range castSpells {
		tmpMP := castSpell.(map[string]interface{})
		var typeId int
		mapToStruct(tmpMP["typeId"], &typeId)
		var playerId int
		mapToStruct(tmpMP["casterId"], &playerId)
		thisTurn := tmpMP["wasCastThisTurn"].(bool)
		var affectedUnits []int
		mapToStruct(tmpMP["affectedUnits"], &affectedUnits)
		if game.isUnitSpell(typeId) {
			var unitId int
			mapToStruct(tmpMP["unit"], &unitId)
			var pathId int
			mapToStruct(tmpMP["path"], &pathId)
			delete(tmpMP, "typeId")
			delete(tmpMP, "affectedUnits")
			delete(tmpMP, "unit")
			delete(tmpMP, "path")
			delete(tmpMP, "wasCastThisTurn")
			var tmpCastSpell CastUnitSpell
			mapToStruct(tmpMP, &tmpCastSpell)
			tmpCastSpell.Spell = game.GetSpellById(typeId)
			tmpCastSpell.Path = game.getPathById(pathId)
			tmpCastSpell.Unit = game.GetUnitById(unitId)
			tmpCastSpell.AffectedUnits = make([]*Unit, 0)
			for _, unitId = range affectedUnits {
				tmpCastSpell.AffectedUnits = append(tmpCastSpell.AffectedUnits, game.GetUnitById(unitId))
			}
			if thisTurn {
				game.players[playerId].CastUnitSpell = &tmpCastSpell
			}
			castSpells = append(castSpells, tmpCastSpell)
		} else {
			delete(tmpMP, "typeId")
			delete(tmpMP, "affectedUnits")
			delete(tmpMP, "wasCastThisTurn")
			var tmpCastSpell CastAreaSpell
			mapToStruct(tmpMP, &tmpCastSpell)
			tmpCastSpell.Spell = game.GetSpellById(typeId)
			tmpCastSpell.AffectedUnits = make([]*Unit, 0)
			for _, unitId := range affectedUnits {
				tmpCastSpell.AffectedUnits = append(tmpCastSpell.AffectedUnits, game.GetUnitById(unitId))
			}
			if thisTurn {
				game.players[playerId].CastAreaSpell = &tmpCastSpell
			}
			castSpells = append(castSpells, tmpCastSpell)
		}
	}

	for _, unit := range units {
		var unitId int
		mapToStruct(unit.(map[string]interface{})["unitId"], &unitId)
		var affectedSpells []int
		mapToStruct(unit.(map[string]interface{})["affectedSpells"], &affectedSpells)
		var tmpUnit = game.GetUnitById(unitId)
		tmpUnit.AffectedSpells = make([]*CastSpell, 0)
		for _, castSpellId := range affectedSpells {
			affectedSpell := game.getCastSpellById(castSpellId)
			tmpUnit.AffectedSpells = append(tmpUnit.AffectedSpells, affectedSpell)
		}
	}

	var tmpSpellId int
	mapToStruct(root["receivedSpell"], &tmpSpellId)
	game.receivedSpell = game.GetSpellById(tmpSpellId)
	mapToStruct(root["friendReceivedSpell"], &tmpSpellId)
	var tmpSpells = make([]*Spell, 0)
	var mySpells []int
	mapToStruct(root["mySpells"], &mySpells)
	for _, spellId := range mySpells {
		tmpSpells = append(tmpSpells, game.GetSpellById(spellId))
	}
	game.players[game.myId].Spells = tmpSpells
	tmpSpells = make([]*Spell, 0)
	var friendSpells []int
	mapToStruct(root["friendSpells"], &friendSpells)
	for _, spellId := range friendSpells {
		tmpSpells = append(tmpSpells, game.GetSpellById(spellId))
	}
	game.players[game.friendId].Spells = tmpSpells
	game.gotRangeUpgrade = root["gotRangeUpgrade"].(bool)
	game.gotDamageUpgrade = root["gotDamageUpgrade"].(bool)
	mapToStruct(root["availableRangeUpgrades"], &game.availableRangeUpgrades)
	mapToStruct(root["availableDamageUpgrades"], &game.availableDamageUpgrades)
	mapToStruct(root["remainingAP"], &game.GetMe().Ap)
	game.startTime = time.Now().UnixNano()
}

func mapToStruct(mp interface{}, v interface{}) {
	js, _ := json.Marshal(mp)
	err := json.Unmarshal(js, v)
	if err != nil {
		fmt.Println(mp)
	}
}

func (game Game) ChooseHand(heroIds []int) {
	i := make([]interface{}, 0)
	for _, v := range heroIds {
		i = append(i, v)
	}
	msg := Message{Name: "pick", Args: map[string]interface{}{"units": i}}
	game.sender <- msg
}

func (game Game) GetMe() *Player {
	return game.players[game.myId]
}

func (game Game) GetFriend() *Player {
	return game.players[game.friendId]
}

func (game Game) GetFirstEnemy() *Player {
	return game.players[game.firstEnemy]
}

func (game Game) GetSecondEnemy() *Player {
	return game.players[game.secondEnemy]
}

func (game Game) PutUnit(typeId, pathId int) {
	msg := Message{Name: "putUnit", Args: map[string]int{"typeId": typeId, "pathId": pathId}, Turn: game.currentTurn}
	game.sender <- msg
}
func (game Game) CastUnitSpell(unitId, pathId int, cell *Cell, spellId int) {
	msg := Message{Name: "castSpell",
		Args: map[string]interface{}{"typeId": spellId, "cell": cell, "unitId": unitId, "pathId": pathId}, Turn: game.currentTurn}
	game.sender <- msg
}

func (game Game) CastAreaSpell(center *Cell, spellId int) {
	msg := Message{Name: "castSpell",
		Args: map[string]interface{}{"typeId": spellId, "cell": center}, Turn: game.currentTurn}
	game.sender <- msg
}

func (game Game) UpgradeUnitRange(unitId int) {
	msg := Message{Name: "rangeUpgrade", Args: map[string]int{"unitId": unitId}, Turn: game.currentTurn}
	game.sender <- msg
}

func (game Game) UpgradeUnitDamage(unitId int) {
	msg := Message{Name: "damageUpgrade", Args: map[string]int{"unitId": unitId}, Turn: game.currentTurn}
	game.sender <- msg
}

func (game Game) GetSpellById(typeId int) *Spell {
	if ret, ok := game.spellById[typeId]; ok {
		return ret
	}
	for _, spell := range game.spells {
		if spell.TypeId == typeId {
			game.spellById[typeId] = spell
			return spell
		}
	}
	game.spellById[typeId] = nil
	return nil
}

func (game Game) GetBaseUnitById(typeId int) *BaseUnit {
	if ret, ok := game.baseUnitById[typeId]; ok {
		return ret
	}
	for _, baseUnit := range game.baseUnits {
		if baseUnit.TypeId == typeId {
			game.baseUnitById[typeId] = baseUnit
			return baseUnit
		}
	}
	game.baseUnitById[typeId] = nil
	return nil
}

func (game Game) getPathById(pathId int) *Path {
	if ret, ok := game.pathById[pathId]; ok {
		return ret
	}
	for _, path := range game.Map.Paths {
		if pathId == path.Id {
			game.pathById[pathId] = path
			return path
		}
	}
	game.pathById[pathId] = nil
	return nil
}

func (game Game) getFriendId(playerId int) int {
	return playerId ^ 2
}

func (game Game) getPathsFromPlayer(playerId int) []*Path {
	paths := make([]*Path, 0)
	friendPath := game.getPathToFriend(playerId)
	for _, path := range game.Map.Paths {
		if path.Id != friendPath.Id {
			startCell := *path.Cells[0]
			endCell := *path.Cells[len(path.Cells)-1]
			playerCell := game.players[playerId].GetPlayerPosition()
			if startCell == playerCell || endCell == playerCell {
				paths = append(paths, path)
			}
		}
	}
	return paths
}

func (game Game) getPathToFriend(playerId int) *Path {
	for i, path := range game.Map.Paths {
		startCell := *path.Cells[0]
		endCell := *path.Cells[len(path.Cells)-1]
		myCell := game.players[playerId].GetPlayerPosition()
		friendCell := game.players[game.getFriendId(playerId)].GetPlayerPosition()
		if (startCell == myCell && endCell == friendCell) ||
			(startCell == friendCell && endCell == myCell) {
			return game.Map.Paths[i]
		}
	}
	return nil
}

func (game Game) GetPathsCrossingCell(cell *Cell) []*Path {
	if ret, ok := game.pathsCrossing[*cell]; ok {
		return ret
	}
	paths := make([]*Path, 0)
	for _, path := range game.Map.Paths {
		for _, cell1 := range path.Cells {
			if *cell1 == *cell {
				paths = append(paths, path)
				break
			}
		}
	}
	game.pathsCrossing[*cell] = paths
	return paths
}

func (game Game) GetCellUnits(cell *Cell) []*Unit {
	if !game.isValid(cell) {
		return []*Unit{}
	}
	return game.Map.UnitsInCell[cell.Row][cell.Col]
}

func (game Game) GetShortestPathToCell(playerId int, cell *Cell) *Path {
	if ret, ok := game.ShortestPaths[playerId][*cell]; ok {
		return game.getPathById(ret)
	}

	var ans *Path
	var minAns = -1
	friendPathLen := len(game.getPathToFriend(playerId).Cells)
	for _, path := range game.Map.Paths {
		startCell := *path.Cells[0]
		endCell := *path.Cells[len(path.Cells)-1]
		playerCell := game.players[playerId].GetPlayerPosition()
		friendCell := game.players[game.getFriendId(playerId)].GetPlayerPosition()
		if startCell == playerCell {
			for i := range path.Cells {
				if *path.Cells[i] == *cell && (minAns == -1 || i < minAns) {
					minAns = i
					ans = path
					break
				}
			}
		}
		if endCell == playerCell {
			lng := len(path.Cells) - 1
			for i := range path.Cells {
				if *path.Cells[lng-i] == *cell && (minAns == -1 || i < minAns) {
					minAns = i
					ans = path
					break
				}
			}
		}
		if startCell == friendCell {
			for i := range path.Cells {
				if *path.Cells[i] == *cell && (minAns == -1 || i+friendPathLen < minAns) {
					minAns = i + friendPathLen
					ans = path
					break
				}
			}
		}
		if endCell == friendCell {
			lng := len(path.Cells) - 1
			for i := range path.Cells {
				if *path.Cells[lng-i] == *cell && (minAns == -1 || i+friendPathLen < minAns) {
					minAns = i + friendPathLen
					ans = path
					break
				}
			}
		}
	}
	game.ShortestPaths[playerId][*cell] = ans.Id
	return ans
}

func (game Game) GetCurrentTurn() int {
	return game.currentTurn
}

func (game Game) GetRemainingTime() int64 {
	if game.currentTurn == 0 {
		return game.gameConstants.PickTimeout - (time.Now().UnixNano()-game.startTime)/1e6
	} else {
		return game.gameConstants.TurnTimeout - (time.Now().UnixNano()-game.startTime)/1e6
	}
}

func (game Game) GetRemainingTurnsToUpgrade() int {
	return game.gameConstants.TurnsToUpgrade - game.currentTurn%game.gameConstants.TurnsToUpgrade
}

func (game Game) GetRemainingTurnsToGetSpell() int {
	return game.gameConstants.TurnsToSpell - game.currentTurn%game.gameConstants.TurnsToSpell
}

func (game Game) GetReceivedSpell() *Spell {
	return game.receivedSpell
}

func (game Game) GetFriendReceivedSpell() *Spell {
	return game.friendReceivedSpell
}

func (game Game) GetAreaSpellTargets(center *Cell, spellId int) []*Unit {
	units := make([]*Unit, 0)
	spell := game.GetSpellById(spellId)
	for i := center.Row - spell.Range; i <= center.Row+spell.Range; i++ {
		for j := center.Col - spell.Range; j <= center.Col+spell.Range; j++ {
			units = append(units, game.GetCellUnits(&Cell{i, j})...)
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

func (game Game) isValid(cell *Cell) bool {
	return cell.Row >= 0 && cell.Row < game.GetMap().RowNum && cell.Col >= 0 && cell.Col < game.GetMap().ColNum
}

func (game Game) GetUnitById(unitId int) *Unit {
	if ret, ok := game.unitById[unitId]; ok {
		return ret
	}
	for _, unit := range game.Map.Units {
		if unit.UnitId == unitId {
			game.unitById[unitId] = unit
			return unit
		}
	}
	game.unitById[unitId] = nil
	return nil
}
func (game Game) getCastSpellById(id int) *CastSpell {
	if ret, ok := game.castSpellById[id]; ok {
		return ret
	}
	for _, c := range game.castSpells {
		if (*c).GetId() == id {
			game.castSpellById[id] = c
			return c
		}
	}
	game.castSpellById[id] = nil
	return nil
}

func (game Game) isUnitSpell(typeId int) bool {
	return game.GetSpellById(typeId).Type == "TELE"
}

func (game Game) GetAllBaseUnits() []*BaseUnit {
	return game.baseUnits
}

func (game Game) GetAllSpells() []*Spell {
	return game.spells
}

func (game Game) GetKingById(playerId int) *King {
	return game.players[playerId].King
}

func (game Game) GetGameConstants() *GameConstants {
	return game.gameConstants
}

func (game Game) GetMap() *Map {
	return game.Map
}

func (game Game) GetPlayerById(playerId int) *Player {
	return game.players[playerId]
}
