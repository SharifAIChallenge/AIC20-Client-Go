package model

type World interface { //TODO make it have pointer arguments?
	ChooseDeck(heroIds []int) //TODO list of baseUnits overload
	GetMyId() int
	GetFriendId() int
	GetFirstEnemy() int
	GetSecondEnemy() int
	GetPlayerPosition(playerId int) Cell
	GetPathsFromPlayer(playerId int) []Path
	GetPathToFriend(playerId int) Path
	GetMapRowNum() int
	GetMapColNum() int
	GetPathsCrossingCell(cell Cell) []Path //TODO Overload
	GetPlayerUnits(playerId int) []Unit
	GetCellUnits(cell Cell) []Unit
	GetShortestPathToCell(playerId int, cell Cell) Path
	GetMaxAP() int
	GetRemainingAP(playerId int) int
	GetHand() []BaseUnit
	GetDeck() []BaseUnit
	PutUnit(typeId, pathId int)
	GetCurrentTurn() int
	GetMaxTurns() int
	GetPickTimeout() int64
	GetTurnTimeout() int64
	GetRemainingTime() int64
	GetPlayerHP(playerId int) int
	CastUnitSpell(unitId, pathId, index, spellId int)    //TODO Overload :( for spell?
	CastAreaSpell(center Cell, spellId int)              //TODO Overload
	GetAreaSpellTargets(center Cell, spellId int) []Unit //TODO overload
	GetRemainingTurnsToUpgrade() int
	GetRemainingTurnsToGetSpell() int
	GetCastAreaSpell(playerId int) CastAreaSpell
	GetCastUnitSpell(playerId int) CastUnitSpell
	GetCastSpellsOnUnit(unitId int) []CastSpell
	GetRangeUpgradeNumber() int
	GetDamageUpgradeNumber() int
	GetSpellsList() []Spell
	GetSpells() map[Spell]int
	GetReceivedSpell() Spell
	GetFriendReceivedSpell() Spell
	UpgradeUnitRange(unitId int)
	UpgradeUnitDamage(unitId int)
	GetPlayerDuplicateUnits(playerId int) []Unit
	GetPlayerHastedUnits(playerId int) []Unit
	GetPlayerPlayedUnits(playerId int) []Unit
	GetUnitTarget(unitId int) Unit //TODO overload
	GetUnitTargetCell(unitId int) Cell
	GetKingTarget(playerId int) Unit
	GetKingTargetCell(playerId int) Cell
	GetKingUnitIsAttacking(unitId int) int
	GetAllBaseUnits() []BaseUnit
	GetAllSpells() []Spell
	GetKingById(playerId int) King
}
