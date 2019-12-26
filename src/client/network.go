package main

import (
	. "../common/network/data"
	"container/list"
	"net"
)

const tag = "Network"

type Network struct {
	messageHandler func(msg Message)
	messagesToSend list.List
	port           int
	host           string
	token          string
	socket         net.Conn //TODO implement JsonSocket?
}

func (network Network) connect() {
	isConnected := false
	for {
	}
}
