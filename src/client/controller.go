package main

import (
	"../common/log"
	. "../common/network/data"
	. "./model"
	"os"
	"time"
)

const Tag = "Controller"

type Controller struct {
	port       int
	host       string
	token      string
	retryDelay int64
	game       Game
	network    Network
	sender     func(message Message)
}

func (controller Controller) Start() {
	controller.network = Network{messageHandler: controller.handleMessage,
		host: controller.host, port: controller.port, token: controller.token}
	controller.sender = controller.network.send
	controller.game = *NewGame(controller.sender)
	for !controller.network.isConnected {
		controller.network.connect()
		time.Sleep(time.Duration(controller.retryDelay) * time.Millisecond)
	}
}

func (controller Controller) handleMessage(msg Message) {
	switch msg.Name {
	case "init":
		controller.handleInitMessage(msg)
	case "turn":
		controller.handleTurnMessage(msg)
	case "shutdown":
		controller.handleShutdownMessage(msg)
	default:
		log.W(Tag, "Undefined message received "+msg.Name)
	}
}

func (controller Controller) handleInitMessage(msg Message) {
	//TODO make new Game
	controller.game.HandleInitMessage(msg)
	controller.pick(controller.game.GetCurrentTurn())
}

func (controller Controller) handleTurnMessage(msg Message) {
	//TODO make new game
	controller.game.HandleTurnMessage(msg)
	controller.turn(controller.game.GetCurrentTurn())
}

func (controller Controller) handleShutdownMessage(msg Message) {
	controller.network.terminate()
	os.Exit(0)
}

func (controller Controller) pick(turnNumber int) {
	go func() {
		pick(controller.game)
		controller.sender(Message{Name: "pick-end", Args: []interface{}{turnNumber}}) //TODO message format
	}()
}

func (controller Controller) turn(turnNumber int) {
	go func() {
		turn(controller.game)
		controller.sender(Message{Name: "turn-end", Args: []interface{}{turnNumber}}) //TODO message format
	}()
}
