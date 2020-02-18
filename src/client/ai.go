package main

import (
	"./model"
	"fmt"
	"math/rand"
)

var rows, cols int
var pathForMyUnits *model.Path

func pick(world model.World) {
	fmt.Println("pick started")

	//preprocess
	mp := world.GetMap()
	rows = mp.RowNum
	cols = mp.ColNum

	allBaseUnits := world.GetAllBaseUnits()
	myDeck := make([]int, 0)

	//choosing all flying units
	for _, baseUnit := range allBaseUnits {
		if baseUnit.IsFlying {
			myDeck = append(myDeck, baseUnit.TypeId)
		}
	}

	// picking the chosen deck - rest of the deck will automatically be filled with random baseUnits
	world.ChooseHandById(myDeck)

	//other preprocess
	pathForMyUnits = world.GetFriend().PathsFromPlayer[0]
}

func turn(world model.World) {
	fmt.Println("turn started: ", world.GetCurrentTurn())

	myself := world.GetMe()
	maxAp := world.GetGameConstants().MaxAP

	// play all of hand once your ap reaches maximum. if ap runs out, putUnit doesn't do anything
	if myself.Ap == maxAp {
		for _, baseUnit := range world.GetMe().Hand {
			world.PutUnit(baseUnit.TypeId, pathForMyUnits.Id)
		}
	}

	// this code tries to cast the received spell
	receivedSpell := world.GetReceivedSpell()
	if receivedSpell != nil {
		if receivedSpell.IsAreaSpell() {
			switch receivedSpell.Target {
			case "ENEMY":
				enemyUnits := world.GetFirstEnemy().Units
				if len(enemyUnits) > 0 {
					world.CastAreaSpell(enemyUnits[0].Cell, receivedSpell.TypeId)
				}
			case "ALLIED":
				friendUnits := world.GetFriend().Units
				if len(friendUnits) > 0 {
					world.CastAreaSpell(friendUnits[0].Cell, receivedSpell.TypeId)
				}
			case "SELF":
				myUnits := myself.Units
				if len(myUnits) > 0 {
					world.CastAreaSpell(myUnits[0].Cell, receivedSpell.TypeId)
				}
			}
		} else {
			myUnits := myself.Units
			if len(myUnits) > 0 {
				unit := myUnits[0]
				myPaths := myself.PathsFromPlayer
				path := myPaths[rand.Intn(len(myPaths))]
				size := len(path.Cells)
				cell := path.Cells[size/2]

				world.CastUnitSpell(unit.UnitId, path.Id, cell, receivedSpell.TypeId)
			}
		}

		// this code tries to upgrade damage of first unit. in case there's no damage token, it tries to upgrade range
		if len(myself.Units) > 0 {
			unit := myself.Units[0]
			world.UpgradeUnitDamage(unit.PlayerId)
			world.UpgradeUnitRange(unit.PlayerId)
		}
	}
}

func end(world model.World, scores map[int]int) {
	fmt.Println("end started")
	fmt.Println("My Score: ", scores[world.GetMe().PlayerId])
}
