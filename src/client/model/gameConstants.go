package model

type GameConstants struct {
	MaxAP                 int   `json:"maxAp"`
	MaxTurns              int   `json:"maxTurns"`
	TurnTimeout           int64 `json:"turnTimeout"`
	PickTimeout           int64 `json:"pickTimeout"`
	TurnsToUpgrade        int   `json:"turnsToUpgrade"`
	TurnsToSpell          int   `json:"turnsToSpell"`
	DamageUpgradeAddition int   `json:"damageUpgradeAddition"`
	RangeUpgradeAddition  int   `json:"rangeUpgradeAddition"`
	DeckSize              int   `json:"deckSize"`
	HandSize              int   `json:"handSize"`
	APAddition            int   `json:"apAddition"`
}
