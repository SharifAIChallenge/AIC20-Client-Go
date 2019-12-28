package model

type World interface { //TODO make it have pointer arguments?
	ChooseDeck(heroIds []int) //TODO list of heroNames(enum) or list of ints
	GetMyId() int
	GetFriendId() int
	GetFirstEnemy() int
	GetSecondEnemy() int
	GetPlayerPosition(playerId int) Cell
	GetPathsFromPlayer(playerId int) []Path
	GetPathToFriend(playerId int) Path
	GetMapHeight() int
	GetMapWidth() int
	GetPathsCrossingCell(cell Cell) []Path
	GetPlayerUnits(playerId int) []Unit
	GetCellUnits(cell Cell) []Unit
	GetShortestPathToCell(playerId int, cell Cell) Path
	GetMaxAP() int
	GetRemainingAP(playerId int) int
	GetHand() []int
	GetDeck() []int
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
	GetActivePoisonsOnUnit(unitId int) int
	GetRangeUpgradeNumber() int
	GetDamageUpgradeNumber() int
	GetSpellsList() []Spell
	GetSpells() map[Spell]int
	GetReceivedSpell() Spell
	GetFriendReceivedSpell() Spell
	UpgradeUnitRange(unitId int)
	UpgradeUnitDamage(unitId int)
	GetPlayerCloneUnits(playerId int) []Unit
	GetPlayerHastedUnits(playerId int) []Unit
	GetPlayerPoisonedUnits(playerId int) []Unit
	GetPlayerPlayedUnits(playerId int) []Unit
}
