package model

import (
	. "../../common/network/data"
	"encoding/json"
	"time"
)

type Game struct {
	gameConstants *GameConstants
	Map           *Map
	baseUnits     []BaseUnit
	spells        []Spell

	currentTurn             int
	castSpells              []CastSpell
	gotRangeUpgrade         bool
	gotDamageUpgrade        bool
	availableRangeUpgrades  int
	availableDamageUpgrades int
	rangeUpgradedUnit       int
	damageUpgradedUnit      int
	remainingAP             int

	startTime int64

	players [4]Player

	myId, friendId, firstEnemy, secondEnemy int
	receivedSpell, friendReceivedSpell      *Spell
	sender                                  chan Message
}

func NewGame(sender chan Message) *Game {
	return &Game{sender: sender}
}

func (game *Game) HandleInitMessage(msg Message) {
	root := msg.Args.(map[string]interface{}) //TODO make it work?
	mapToStruct(root["gameConstants"], &game.gameConstants)
	mapToStruct(root["map"], &game.Map)
	for _, king := range game.Map.Kings {
		game.players[king.PlayerId] = Player{King: &king, PlayerId: king.PlayerId}
	}
	game.myId = game.Map.Kings[0].PlayerId
	game.friendId = game.Map.Kings[1].PlayerId
	game.firstEnemy = game.Map.Kings[2].PlayerId
	game.secondEnemy = game.Map.Kings[3].PlayerId

	mapToStruct(root["baseUnits"], &game.baseUnits)
	spells := root["spells"].([]interface{})
	for _, spell := range spells {
		var tmpSpell Spell
		mapToStruct(spell, &tmpSpell)
		game.spells = append(game.spells, tmpSpell)
	}
	game.startTime = time.Now().UnixNano()
}
func (game *Game) HandleTurnMessage(msg Message) {
	root := msg.Args.(map[string]interface{})
	mapToStruct(root["currTurn"], &game.currentTurn)
	mapToStruct(root["deck"], &game.players[game.myId].Deck)
	mapToStruct(root["hand"], &game.players[game.myId].Hand)
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
			game.players[playerId].King.Target = tmpUnit
		}
	}

	game.Map.Units = make([]Unit, 0)
	for i := 0; i < 4; i++ {
		game.players[i].Units = make([]Unit, 0)
		game.players[i].DuplicateUnits = make([]Unit, 0)
		game.players[i].HastedUnits = make([]Unit, 0)
		game.players[i].PlayedUnits = make([]Unit, 0)
		game.players[i].DiedUnits = make([]Unit, 0)
	}
	game.Map.UnitsInCell = make([][][]Unit, 0)
	for i := 0; i < game.Map.RowNum; i++ {
		game.Map.UnitsInCell = append(game.Map.UnitsInCell, make([][]Unit, 0))
		for j := 0; j < game.Map.ColNum; j++ {
			game.Map.UnitsInCell[i] = append(game.Map.UnitsInCell[i], make([]Unit, 0))
		}
	}
	units := root["units"].([]interface{}) // get baseUnit by TypeId
	for _, unit := range units {
		var typeId int
		mapToStruct(unit.(map[string]interface{})["typeId"], &typeId)
		var playerId int
		mapToStruct(unit.(map[string]interface{})["playerId"], &playerId)
		var pathId int
		mapToStruct(unit.(map[string]interface{})["pathId"], &pathId)
		var targetId int
		mapToStruct(unit.(map[string]interface{})["target"], &targetId)
		var wasPlayedThisTurn bool
		mapToStruct(unit.(map[string]interface{})["wasPlayedThisTurn"], &wasPlayedThisTurn)
		var affectedSpells []int
		mapToStruct(unit.(map[string]interface{})["affectedSpells"], &affectedSpells)
		baseUnit := game.GetBaseUnitById(typeId)
		path := game.GetPathById(pathId)
		target := game.GetUnitById(targetId)
		var tmpUnit Unit
		mapToStruct(unit, &tmpUnit)
		tmpUnit.BaseUnit = baseUnit
		tmpUnit.Path = path
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
		tmpUnit.AffectedSpells = make([]CastSpell, 0)
		for _, castSpellId := range affectedSpells {
			affectedSpell := game.getCastSpellById(castSpellId)
			tmpUnit.AffectedSpells = append(tmpUnit.AffectedSpells, *affectedSpell)
		}

		game.players[playerId].Units = append(game.players[playerId].Units, tmpUnit)
		if tmpUnit.IsDuplicate {
			game.players[playerId].DuplicateUnits = append(game.players[playerId].DuplicateUnits, tmpUnit)
		}
		if tmpUnit.IsHasted {
			game.players[playerId].HastedUnits = append(game.players[playerId].HastedUnits, tmpUnit)
		}
		if wasPlayedThisTurn {
			game.players[playerId].PlayedUnits = append(game.players[playerId].PlayedUnits, tmpUnit)
		}
		game.Map.Units = append(game.Map.Units, tmpUnit)
		game.Map.UnitsInCell[tmpUnit.Cell.Row][tmpUnit.Cell.Col] =
			append(game.Map.UnitsInCell[tmpUnit.Cell.Row][tmpUnit.Cell.Col], tmpUnit) //TODO pointers or ID
	}
	deadUnits := root["units"].([]interface{}) // get baseUnit by TypeId
	for _, unit := range deadUnits {
		var typeId int
		mapToStruct(unit.(map[string]interface{})["typeId"], &typeId)
		var playerId int
		mapToStruct(unit.(map[string]interface{})["playerId"], &playerId) //TODO do dead units need something else?
		baseUnit := game.GetBaseUnitById(typeId)
		var tmpUnit Unit
		mapToStruct(unit, &tmpUnit)
		tmpUnit.BaseUnit = baseUnit
		game.players[playerId].DiedUnits = append(game.players[playerId].DiedUnits, tmpUnit)
	}
	castSpells, ok := root["castSpells"].([]interface{})
	if !ok {
		castSpells = make([]interface{}, 0)
	}
	for i := 0; i < 4; i++ {
		game.players[i].CastUnitSpell = nil
		game.players[i].CastAreaSpell = nil
	}
	game.castSpells = make([]CastSpell, 0)

	for _, castSpell := range castSpells {
		var typeId int
		mapToStruct(castSpell.(map[string]interface{})["typeId"], &typeId)
		var playerId int
		mapToStruct(castSpell.(map[string]interface{})["casterId"], &playerId)
		thisTurn := castSpell.(map[string]interface{})["wasCastThisTurn"].(bool)
		if game.isUnitSpell(typeId) {
			var tmpCastSpell CastUnitSpell
			mapToStruct(castSpell, &tmpCastSpell)
			if thisTurn {
				game.players[playerId].CastUnitSpell = &tmpCastSpell
			}
			castSpells = append(castSpells, tmpCastSpell)
		} else {
			var tmpCastSpell CastAreaSpell
			mapToStruct(castSpell, &tmpCastSpell)
			if thisTurn {
				game.players[playerId].CastAreaSpell = &tmpCastSpell
			}
			castSpells = append(castSpells, tmpCastSpell)
		}
	}
	var tmpSpellId int
	mapToStruct(root["receivedSpell"], &tmpSpellId)
	game.receivedSpell = game.GetSpellById(tmpSpellId)
	mapToStruct(root["friendReceivedSpell"], &tmpSpellId)
	var tmpSpells = make([]Spell, 0)
	var mySpells []int
	mapToStruct(root["mySpells"], &mySpells)
	for _, spellId := range mySpells {
		tmpSpells = append(tmpSpells, *game.GetSpellById(spellId))
	}
	game.players[game.myId].Spells = tmpSpells
	tmpSpells = make([]Spell, 0)
	var friendSpells []int
	mapToStruct(root["friendSpells"], &friendSpells)
	for _, spellId := range friendSpells {
		tmpSpells = append(tmpSpells, *game.GetSpellById(spellId))
	}
	game.players[game.friendId].Spells = tmpSpells
	game.gotRangeUpgrade = root["gotRangeUpgrade"].(bool)
	game.gotDamageUpgrade = root["gotDamageUpgrade"].(bool)
	mapToStruct(root["availableRangeUpgrades"], &game.availableRangeUpgrades)
	mapToStruct(root["availableDamageUpgrades"], &game.availableDamageUpgrades)
	mapToStruct(root["rangeUpgradedUnit"], &game.rangeUpgradedUnit)
	mapToStruct(root["damageUpgradedUnit"], &game.damageUpgradedUnit)
	if game.gotRangeUpgrade {
		game.players[game.myId].RangeUpgradedUnit = game.GetUnitById(game.rangeUpgradedUnit)
	}
	if game.gotDamageUpgrade {
		game.players[game.myId].DamageUpgradedUnit = game.GetUnitById(game.damageUpgradedUnit)
	}
	mapToStruct(root["remainingAP"], &game.remainingAP)
	game.startTime = time.Now().UnixNano()
}

func mapToStruct(mp interface{}, v interface{}) {
	js, _ := json.Marshal(mp)
	_ = json.Unmarshal(js, v)
}

func (game Game) ChooseDeck(heroIds []int) {
	i := make([]interface{}, 0)
	for _, v := range heroIds {
		i = append(i, v)
	}
	msg := Message{Name: "pick", Args: map[string]interface{}{"units": i}} //TODO check server message format
	game.sender <- msg
}

func (game Game) GetMe() *Player {
	return &game.players[game.myId]
}

func (game Game) GetFriend() *Player {
	return &game.players[game.friendId]
}

func (game Game) GetFirstEnemy() *Player {
	return &game.players[game.firstEnemy]
}

func (game Game) GetSecondEnemy() *Player {
	return &game.players[game.secondEnemy]
}

func (game Game) PutUnit(typeId, pathId int) {
	msg := Message{Name: "putUnit", Args: map[string]int{"typeId": typeId, "pathId": pathId}, Turn: game.currentTurn} //TODO named args?
	game.sender <- msg
}
func (game Game) CastUnitSpell(unitId, pathId int, cell Cell, spellId int) {
	msg := Message{Name: "castSpell",
		Args: map[string]interface{}{"typeId": spellId, "cell": cell, "unitId": unitId, "pathId": pathId}, Turn: game.currentTurn}
	game.sender <- msg
}

func (game Game) CastAreaSpell(center Cell, spellId int) {
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
	for _, spell := range game.spells {
		if spell.TypeId == typeId {
			return &spell
		}
	}
	return nil
}

func (game Game) GetBaseUnitById(typeId int) *BaseUnit {
	for _, baseUnit := range game.baseUnits {
		if baseUnit.TypeId == typeId {
			return &baseUnit
		}
	}
	return nil
}

func (game Game) GetPathById(pathId int) *Path {
	for _, path := range game.Map.Paths {
		if pathId == path.Id {
			return &path
		}
	}
	return nil
}

func (game Game) getFriendId(playerId int) int {
	return playerId ^ 2 //TODO make sure this is valid
}

func (game Game) GetPathsFromPlayer(playerId int) []Path { //TODO friend paths
	paths := make([]Path, 0)
	friendPath := game.GetPathToFriend(playerId)
	for _, path := range game.Map.Paths {
		if path.Id != friendPath.Id {
			startCell := path.Cells[0]
			endCell := path.Cells[len(path.Cells)-1]
			playerCell := game.players[playerId].GetPlayerPosition()
			if startCell == playerCell || endCell == playerCell {
				paths = append(paths, path)
			}
		}
	}
	return paths
}

func (game Game) GetPathToFriend(playerId int) *Path {
	for _, path := range game.Map.Paths {
		startCell := path.Cells[0]
		endCell := path.Cells[len(path.Cells)-1]
		myCell := game.players[playerId].GetPlayerPosition()
		friendCell := game.players[game.getFriendId(playerId)].GetPlayerPosition()
		if (startCell == myCell && endCell == friendCell) ||
			(startCell == friendCell && endCell == myCell) {
			return &path //TODO Is there such a path?
		}
	}
	return nil
}

func (game Game) GetPathsCrossingCell(cell Cell) []Path {
	paths := make([]Path, 0)
	for _, path := range game.Map.Paths {
		for _, cell1 := range path.Cells {
			if cell1 == cell {
				paths = append(paths, path)
				break
			}
		}
	}
	return paths
}

func (game Game) GetCellUnits(cell Cell) []Unit {
	if !game.isValid(cell) {
		return []Unit{}
	}
	return game.Map.UnitsInCell[cell.Row][cell.Col]
}

func (game Game) GetShortestPathToCell(playerId int, cell Cell) Path {
	var ans Path
	var minAns = -1
	friendPathLen := len(game.GetPathToFriend(playerId).Cells)
	for _, path := range game.Map.Paths {
		startCell := path.Cells[0]
		endCell := path.Cells[len(path.Cells)-1]
		playerCell := game.players[playerId].GetPlayerPosition()
		friendCell := game.players[game.getFriendId(playerId)].GetPlayerPosition()
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

func (game Game) GetRemainingAP(playerId int) int {
	return game.players[playerId].Ap
}

func (game Game) GetHand() []BaseUnit {
	hand := make([]BaseUnit, 0)
	for _, id := range game.players[game.myId].Hand {
		hand = append(hand, *game.GetBaseUnitById(id))
	}
	return hand
}

func (game Game) GetDeck() []BaseUnit {
	deck := make([]BaseUnit, 0)
	for _, id := range game.players[game.myId].Deck {
		deck = append(deck, *game.GetBaseUnitById(id))
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

func (game Game) GetPickTimeout() int64 {
	return game.gameConstants.PickTimeout
}

func (game Game) GetAreaSpellTargets(center Cell, spellId int) []Unit {
	units := make([]Unit, 0)
	spell := game.GetSpellById(spellId)
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
	return game.players[game.myId].Spells
}

func (game Game) GetSpells() map[Spell]int {
	spellMap := make(map[Spell]int, 0)
	for _, spell := range game.GetSpellsList() {
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
	for _, unit := range game.Map.Units {
		if unit.IsDuplicate && unit.PlayerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) GetPlayerHastedUnits(playerId int) []Unit {
	res := make([]Unit, 0)
	for _, unit := range game.Map.Units {
		if unit.IsHasted && unit.PlayerId == playerId {
			res = append(res, unit)
		}
	}
	return res
}

func (game Game) isValid(cell Cell) bool {
	return cell.Row >= 0 && cell.Row < game.GetMap().RowNum && cell.Col >= 0 && cell.Col < game.GetMap().ColNum //TODO move to Map?
}

func (game Game) GetUnitById(unitId int) *Unit {
	for _, unit := range game.Map.Units {
		if unit.UnitId == unitId {
			return &unit
		}
	}
	return nil
}
func (game Game) getCastSpellById(id int) *CastSpell {
	for _, c := range game.castSpells {
		if c.GetId() == id {
			return &c
		}
	}
	return nil
}

func (game Game) isUnitSpell(typeId int) bool {
	return game.GetSpellById(typeId).Type == "TELE" //TODO avoid hard coding
}

func (game Game) GetAllBaseUnits() []BaseUnit {
	baseUnits := make([]BaseUnit, len(game.baseUnits))
	copy(baseUnits, game.baseUnits)
	return baseUnits
}

func (game Game) GetAllSpells() []Spell {
	spells := make([]Spell, len(game.spells))
	copy(spells, game.spells)
	return spells
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
	return &game.players[playerId]
}
