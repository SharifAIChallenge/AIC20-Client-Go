package main

import (
	. "../common/model"
	. "../common/network/data"
	. "./model"
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


}

func (controller Controller) handleMessage(msg Message) {

}

func (controller Controller) handleInitMessage(msg Message) {

}

func (controller Controller) handlePickMessage(msg Message) {

}

func (controller Controller) handleTurnMessage(msg Message) {

}

func (controller Controller) handleShutdownMessage(msg Message) {

}

func (controller Controller) pick(game Game, endEvent Event) {

}

func (controller Controller) turn(game Game, endEvent Event) {

}
