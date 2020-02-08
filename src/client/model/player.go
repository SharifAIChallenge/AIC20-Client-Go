package model

type Player struct {
	PlayerId           int            `json:"playerId"`
	Deck               []int          `json:"deck"`
	Hand               []int          `json:"hand"`
	Ap                 int            `json:"ap"`
	King               *King          `json:"king"`
	PathsFromPlayer    []*Path        `json:"pathsFromPlayer"` //TODO
	PathToFriend       *Path          `json:"pathToFriend"`    //TODO
	Units              []*Unit        `json:"units"`
	CastAreaSpell      *CastAreaSpell `json:"castAreaSpell"`
	CastUnitSpell      *CastUnitSpell `json:"castUnitSpell"`
	DuplicateUnits     []*Unit        `json:"duplicateUnits"`
	HastedUnits        []*Unit        `json:"hastedUnits"`
	PlayedUnits        []*Unit        `json:"playedUnits"`
	DiedUnits          []*Unit        `json:"diedUnits"`
	RangeUpgradedUnit  *Unit          `json:"rangeUpgradedUnit"`
	DamageUpgradedUnit *Unit          `json:"damageUpgradedUnit"`
	UpgradeTokens      int            `json:"upgradeTokens"`
	Spells             []*Spell       `json:"spells"`
}

func (player Player) isAlive() bool {
	return player.King.IsAlive
}

func (player Player) getHp() int {
	return player.King.Hp
}

func (player Player) GetPlayerPosition() Cell {
	return *player.King.Center //TODO use getters?
}

func (player Player) GetSpellCount(spell Spell) int {
	cnt := 0
	for _, spell1 := range player.Spells {
		if spell.TypeId == spell1.TypeId {
			cnt++
		}
	}
	return cnt
}
