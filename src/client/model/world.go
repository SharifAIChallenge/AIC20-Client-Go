package model

type World interface { //TODO make it have pointer arguments?
	ChooseDeck(heroIds []int) //TODO list of baseUnits overload
	GetMe() Player
	GetFriend() Player
	GetFirstEnemy() Player
	GetSecondEnemy() Player
	GetMap() Player
	GetPathsCrossingCell(cell Cell) []Path //TODO Overload
	GetPlayerUnits(playerId int) []Unit
	GetCellUnits(cell Cell) []Unit
	GetShortestPathToCell(playerId int, cell Cell) Path
	GetHand() []BaseUnit
	GetDeck() []BaseUnit
	PutUnit(typeId, pathId int)
	GetCurrentTurn() int
	GetRemainingTime() int64
	CastUnitSpell(unitId, pathId, index, spellId int)    //TODO Overload :( for spell?
	CastAreaSpell(center Cell, spellId int)              //TODO Overload
	GetAreaSpellTargets(center Cell, spellId int) []Unit //TODO overload
	GetRemainingTurnsToUpgrade() int
	GetRemainingTurnsToGetSpell() int
	GetRangeUpgradeNumber() int
	GetDamageUpgradeNumber() int
	GetSpellsList() []Spell
	GetSpells() map[Spell]int
	GetReceivedSpell() Spell
	GetFriendReceivedSpell() Spell
	UpgradeUnitRange(unitId int)
	UpgradeUnitDamage(unitId int)
	GetAllBaseUnits() []BaseUnit
	GetAllSpells() []Spell
	GetKingById(playerId int) King
	GetSpellById(spellId int) Spell
	GetBaseUnitById(baseUnitId int) BaseUnit
	GetPlayerById(playerId int) Player
	GetUnitById(unitId int) Unit
	GetGameConstants() GameConstants
}
