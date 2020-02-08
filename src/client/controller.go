package main

import (
	"../common/log"
	. "../common/network/data"
	. "./model"
	"encoding/json"
	"os"
	"time"
)

const Tag = "Controller"

type Controller struct {
	port           int
	host           string
	token          string
	retryDelay     int64
	game           *Game
	network        *Network
	sender         chan Message
	messageHandler chan Message
}

func (controller Controller) Start() {
	controller.messageHandler = make(chan Message)
	controller.network = &Network{messageHandler: controller.messageHandler,
		host: controller.host, port: controller.port, token: controller.token, messagesToSend: make(chan Message)}
	controller.sender = controller.network.messagesToSend
	controller.game = NewGame(controller.sender)
	go controller.handleMessages()
	for !controller.network.isConnected {
		controller.network.connect()
		time.Sleep(time.Duration(controller.retryDelay) * time.Millisecond)
	}
}

func (controller Controller) handleMessages() {
	for {
		msg := <-controller.messageHandler
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
}

func (controller Controller) handleInitMessage(msg Message) {
	//TODO make new Game
	controller.game.HandleInitMessage(msg)
	controller.pick(controller.game)
}

func (controller Controller) handleTurnMessage(msg Message) {
	//TODO make new game
	controller.game.HandleTurnMessage(msg)
	controller.turn(controller.game)
}

func (controller Controller) handleShutdownMessage(msg Message) {
	info := msg.Args.(map[string]interface{})
	controller.game.HandleTurnMessage(Message{Name: msg.Name, Args: info["turnMessage"], Turn: msg.Turn})
	var scoresList []map[string]int
	mapToStruct(info["scores"], &scoresList)
	scores := make(map[int]int)
	for _, score := range scoresList {
		scores[score["playerId"]] = score["score"]
	}
	controller.end(controller.game, scores)
	controller.network.terminate()
	os.Exit(0)
}

func mapToStruct(mp interface{}, v interface{}) {
	js, _ := json.Marshal(mp)
	_ = json.Unmarshal(js, v)
}

func (controller Controller) pick(game *Game) {
	go func() {
		pick(game)
		controller.sender <- Message{Name: "endTurn", Args: map[string]interface{}{}, Turn: 0}
	}()
}

func (controller Controller) turn(game *Game) {
	go func() {
		turn(game)
		controller.sender <- Message{Name: "endTurn", Args: map[string]interface{}{}, Turn: game.GetCurrentTurn()}
	}()
}

func (controller Controller) end(game *Game, scores map[int]int) {
	end(game, scores)
}
