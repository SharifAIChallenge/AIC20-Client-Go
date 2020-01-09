package main

import (
	. "../common/network"
	. "../common/network/data"
	"fmt"
)

const tag = "Network"

type Network struct {
	messageHandler chan Message
	port           int
	host           string
	token          string
	socket         JsonSocket
	isConnected    bool //TODO should i add number of exceptions?
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
		//TODO throw error? exception? use try catch?
	}
	network.isConnected = true
	network.socket = client
	network.messageHandler <- init
	go network.startReceiving() //TODO is this okay?
	go network.startSending()   //TODO should this be in another place?
}

func (network Network) startReceiving() {
	for !network.terminateFlag {
		network.doReceive()
	}
}

func (network Network) doReceive() {
	msg := network.socket.Get() //TODO errors and error handling?
	network.messageHandler <- msg
}

func (network Network) startSending() {
	for !network.terminateFlag {
		msg := <-network.messagesToSend
		fmt.Println(msg)
		network.socket.Send(msg)
	}
}
func (network Network) terminate() {
	network.terminateFlag = true
	network.socket.Close() //TODO errors?
}

//TODO handleIOE?
