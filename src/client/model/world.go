package model

type World interface { //TODO make it have pointer arguments?
	ChooseDeck(heroIds []int) //TODO list of heroNames(enum) or list of ints
	GetMyId() int
	GetFriendId() int
	GetPlayerPosition(playerId int) Cell
	GetPathsFromPlayer(playerId int) []Path
	GetPathsToFriend(playerId int) Path
	GetMapHeight() int
	GetMapWidth() int
	GetPathsCrossing(cell Cell) []Path
	GetPlayerUnits(playerId int) []Unit
	GetCellUnits(cell Cell) []Unit
	GetShortestPathToCell(playerId int, cell Cell) Path
	GetMaxAP(playerId int) int
	GetRemainingAP(playerId int) int
	GetHand() []Unit
	GetDeck() []Unit
	PlayUnit(typeId, pathId int) int
	GetCurrentTurn() int
	GetMaxTurns() int
	GetTurnTimeout() int64
	GetRemainingTime() int64
	GetPlayerHP() int
	CastUnitSpell(unitId, pathId, index, spellId int) //Overload :( for spell?
	CastAreaSpell(center Cell, spellId int)           //Overload
	GetAreaSpellTargets(center Cell, spell Spell) []Unit
	GetRemainingTurnsToUpgrade() int
	GetRemainingTurnsToGetSpell() int
	GetCastAreaSpell(playerId int) CastAreaSpell
	GetCastUnitSpell(playerId int) CastUnitSpell
	GetDeployedUnits(playerId int) []Unit
	GetActiveSpellsOnCell(cell Cell) []CastAreaSpell
	GetUpgradeTokenNumber(playerId int) int
	GetSpells() []Spell
	GetReceivedSpell() Spell
	GetFriendReceivedSpell() Spell
}
