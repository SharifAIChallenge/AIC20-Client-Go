package main

import (
	. "../common/network"
	. "../common/network/data"
)

const tag = "Network"

type Network struct {
	messageHandler chan Message
	port           int
	host           string
	token          string
	socket         JsonSocket
	isConnected    bool
	terminateFlag  bool
	messagesToSend chan Message
}

func (network Network) connect() {
	network.isConnected = false
	var client JsonSocket
	client = *NewJsonSocket(network.host, network.port)
	client.Send(Message{Name: "token", Args: map[string]string{"token": network.token}})
	init := client.Get()
	if init.Name != "init" {
		client.Close()
	}
	network.isConnected = true
	network.socket = client
	network.messageHandler <- init
	go network.startReceiving()
	go network.startSending()
}

func (network Network) startReceiving() {
	for !network.terminateFlag {
		network.doReceive()
	}
}

func (network Network) doReceive() {
	msg := network.socket.Get()
	network.messageHandler <- msg
}

func (network Network) startSending() {
	for !network.terminateFlag {
		msg := <-network.messagesToSend
		network.socket.Send(msg)
	}
}
func (network Network) terminate() {
	network.terminateFlag = true
}

