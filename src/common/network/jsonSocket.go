package network

import (
	. "./data"
	"bufio"
	"encoding/json"
	"net"
	"strconv"
)

const tag = "JsonSocket"

type JsonSocket struct {
	socket net.Conn
	reader *bufio.Reader
}

func NewJsonSocket(host string, port int) *JsonSocket {
	socket, _ := net.Dial("tcp", host+":"+strconv.Itoa(port))
	return &JsonSocket{socket: socket}
}

func (jsonSocket JsonSocket) Close() {
	_ = jsonSocket.socket.Close()
}

func (jsonSocket JsonSocket) Send(msg Message) {
	js, _ := json.Marshal(msg)
	_, _ = jsonSocket.socket.Write(append(js, byte('\000')))
}

func (jsonSocket JsonSocket) Get() Message {
	js := make([]byte, 0)
	if jsonSocket.reader == nil {
		jsonSocket.reader = bufio.NewReader(jsonSocket.socket)
	}
	for {
		char, _ := jsonSocket.reader.ReadByte()
		if char == 0 {
			break
		}
		js = append(js, char)
	}
	var result Message
	_ = json.Unmarshal(js, &result)
	return result
}
