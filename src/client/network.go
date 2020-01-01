package main

import (
	. "../common/network"
	. "../common/network/data"
)

const tag = "Network"

type Network struct {
	messageHandler func(msg Message)
	messagesToSend []Message
	port           int
	host           string
	token          string
	socket         JsonSocket
	isConnected    bool //TODO should i add number of exceptions?
	terminateFlag  bool
}

func (network Network) connect() {
	network.isConnected = false
	var client JsonSocket
	client = *NewJsonSocket(network.host, network.port)
	client.Send(Message{Name: "token", Args: []interface{}{network.token}})
	init := client.Get().(Message)
	if init.Name != "init" {
		client.Close()
		//TODO throw error? exception? use try catch?
	}
	network.isConnected = true
	network.socket = client
	network.messageHandler(init)
	go network.startReceiving() //TODO is this okay?
	go network.startSending()   //TODO should this be in another place?
}

func (network Network) startReceiving() {
	for !network.terminateFlag {
		network.doReceive()
	}
}

func (network Network) doReceive() {
	msg := network.socket.Get().(Message) //TODO errors and error handling?
	network.messageHandler(msg)
}

func (network Network) startSending() {
	for !network.terminateFlag {
		if len(network.messagesToSend) > 0 {
			msg := network.messagesToSend[0]
			network.messagesToSend = network.messagesToSend[1:] //TODO check memory problems?
			network.socket.Send(msg)                            //TODO handle exceptions?
		}
	}
}
func (network Network) send(msg Message) {
	network.messagesToSend = append(network.messagesToSend, msg)
}

func (network Network) terminate() {
	network.terminateFlag = true
	network.socket.Close() //TODO errors?
}

//TODO handleIOE?
