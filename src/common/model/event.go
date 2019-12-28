package model

const EVENT = "event"

type Event struct {
	Typ  string
	Args []interface{}
}
