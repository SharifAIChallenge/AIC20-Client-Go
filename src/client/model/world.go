package model

type World interface { //TODO make it have pointer arguments?
	ChooseDeck(heroIds []int) //TODO list of baseUnits overload
	GetMe() *Player
	GetFriend() *Player
	GetFirstEnemy() *Player
	GetSecondEnemy() *Player
	GetMap() *Map
	GetPathsCrossingCell(cell Cell) []Path //TODO Overload
	GetCellUnits(cell Cell) []Unit
	GetShortestPathToCell(playerId int, cell Cell) *Path
	PutUnit(typeId, pathId int)
	GetCurrentTurn() int
	GetRemainingTime() int64
	CastUnitSpell(unitId, pathId int, cell Cell, spellId int) //TODO Overload :( for spell?
	CastAreaSpell(center Cell, spellId int)                   //TODO Overload
	GetAreaSpellTargets(center Cell, spellId int) []Unit      //TODO overload
	GetRemainingTurnsToUpgrade() int
	GetRemainingTurnsToGetSpell() int
	GetRangeUpgradeNumber() int
	GetDamageUpgradeNumber() int
	GetReceivedSpell() *Spell
	GetFriendReceivedSpell() *Spell
	UpgradeUnitRange(unitId int)  //todo unit
	UpgradeUnitDamage(unitId int) //todo unit
	GetAllBaseUnits() []BaseUnit
	GetAllSpells() []Spell
	GetKingById(playerId int) *King
	GetSpellById(spellId int) *Spell
	GetBaseUnitById(baseUnitId int) *BaseUnit
	GetPlayerById(playerId int) *Player
	GetUnitById(unitId int) *Unit
	GetGameConstants() *GameConstants
}
