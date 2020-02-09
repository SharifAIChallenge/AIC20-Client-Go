package main

import (
	"./model"
	"fmt"
)

func pick(world model.World) {
	units := make([]int, 0)
	for i := 0; i < 6; i++ {
		units = append(units, world.GetAllBaseUnits()[i].TypeId)
	}
	world.ChooseDeck(units)
}

func turn(world model.World) {
	var me=world.GetMe()
	if len(world.GetAllSpells()) > 0 {

		world.CastAreaSpell(me.PathsFromPlayer[0].Cells[0], world.GetAllSpells()[0].TypeId)
	}
	world.PutUnit(me.Hand[0].TypeId, me.PathsFromPlayer[0].Id)
	fmt.Println("|||||||||||||||||| currTurn",world.GetCurrentTurn())
	var cell=world.GetMap().GetCell(6,6)
	var paths=world.GetPathsCrossingCell(cell)
	for _,path := range paths {
		fmt.Println("|||||")
		for _,cell := range path.Cells {
			fmt.Println(*cell)
		}
	}
	fmt.Println("|||||||||||||||||||||||| King",world.GetMe().King.Center.Row," ",world.GetMe().King.Center.Col)
	var path=world.GetShortestPathToCell(world.GetMe().PlayerId,&model.Cell{Row: 6,Col:6})
	fmt.Println("|||||")
	for _,cell := range path.Cells {
		fmt.Println(*cell)
	}
}

func end(world model.World, scores map[int]int) {

}
