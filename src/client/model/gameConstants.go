package model

type GameConstants struct {
	maxAP                 int
	maxTurns              int
	turnTimeout           int
	pickTimeout           int
	turnsToUpgrade        int
	turnsToSpell          int
	damageUpgradeAddition int
	rangeUpgradeAddition  int
}

func (g *GameConstants) TurnsToUpgrade() int {
	return g.turnsToUpgrade
}

func (g *GameConstants) SetTurnsToUpgrade(turnsToUpgrade int) {
	g.turnsToUpgrade = turnsToUpgrade
}

func (g *GameConstants) PickTimeout() int {
	return g.pickTimeout
}

func (g *GameConstants) SetPickTimeout(pickTimeout int) {
	g.pickTimeout = pickTimeout
}

func (g *GameConstants) TurnTimeout() int {
	return g.turnTimeout
}

func (g *GameConstants) SetTurnTimeout(turnTimeout int) {
	g.turnTimeout = turnTimeout
}

func (g *GameConstants) MaxTurns() int {
	return g.maxTurns
}

func (g *GameConstants) SetMaxTurns(maxTurns int) {
	g.maxTurns = maxTurns
}

func (g *GameConstants) MaxAP() int {
	return g.maxAP
}

func (g *GameConstants) SetMaxAP(maxAP int) {
	g.maxAP = maxAP
}
