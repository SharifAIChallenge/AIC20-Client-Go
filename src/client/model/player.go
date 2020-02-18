package model

type Player struct {
	PlayerId           int            `json:"playerId"`
	Deck               []*BaseUnit    `json:"deck"`
	Hand               []*BaseUnit    `json:"hand"`
	Ap                 int            `json:"ap"`
	King               *King          `json:"king"`
	PathsFromPlayer    []*Path        `json:"pathsFromPlayer"`
	PathToFriend       *Path          `json:"pathToFriend"`
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
	spellCount         map[int]int
}

func (player Player) isAlive() bool {
	return player.King.IsAlive
}

func (player Player) getHp() int {
	return player.King.Hp
}

func (player Player) GetPlayerPosition() *Cell {
	return player.King.Center
}

func (player Player) GetSpellCount(spellId int) int {
	if ret, ok := player.spellCount[spellId]; ok {
		return ret
	}
	cnt := 0
	for _, spell1 := range player.Spells {
		if spellId == spell1.TypeId {
			cnt++
		}
	}
	player.spellCount[spellId] = cnt
	return cnt
}
