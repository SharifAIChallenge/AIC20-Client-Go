package main

import (
	"./model"
)

func pick(world model.World) {
	units := make([]int, 0)
	for i := 0; i < 6; i++ {
		units = append(units, world.GetAllBaseUnits()[i].TypeId)
	}
	world.ChooseDeck(units)
}

func turn(world model.World) {
	//fmt.Println(world.GetDeck())
	if len(world.GetSpells()) > 0 {
		//	world.CastAreaSpell(world.GetPathsFromPlayer(world.GetMyId())[0].Cells[0], world.GetSpellsList()[0].TypeId)
	}
	//fmt.Println(world.GetCastAreaSpell(world.GetMyId()))
	world.PutUnit(world.GetHand()[0].TypeId, world.GetPathsFromPlayer(world.GetMyId())[0].PathId)
	//fmt.Println(world.GetPlayerUnits(world.GetMyId()))
}
