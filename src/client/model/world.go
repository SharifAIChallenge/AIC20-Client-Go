package model

type World interface {
	ChooseHand(heroIds []int)
	GetMe() *Player
	GetFriend() *Player
	GetFirstEnemy() *Player
	GetSecondEnemy() *Player
	GetMap() *Map
	GetPathsCrossingCell(cell *Cell) []*Path
	GetCellUnits(cell *Cell) []*Unit
	GetShortestPathToCell(playerId int, cell *Cell) *Path
	PutUnit(typeId, pathId int)
	GetCurrentTurn() int
	GetRemainingTime() int64
	CastUnitSpell(unitId, pathId int, cell *Cell, spellId int)
	CastAreaSpell(center *Cell, spellId int)
	GetAreaSpellTargets(center *Cell, spellId int) []*Unit
	GetRemainingTurnsToUpgrade() int
	GetRemainingTurnsToGetSpell() int
	GetRangeUpgradeNumber() int
	GetDamageUpgradeNumber() int
	GetReceivedSpell() *Spell
	GetFriendReceivedSpell() *Spell
	UpgradeUnitRange(unitId int)
	UpgradeUnitDamage(unitId int)
	GetAllBaseUnits() []*BaseUnit
	GetAllSpells() []*Spell
	GetKingById(playerId int) *King
	GetSpellById(spellId int) *Spell
	GetBaseUnitById(baseUnitId int) *BaseUnit
	GetPlayerById(playerId int) *Player
	GetUnitById(unitId int) *Unit
	GetGameConstants() *GameConstants
}
