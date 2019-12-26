package network

//TODO implement errors and error handling
import (
	. "./data"
	"bufio"
	"encoding/json"
	"net"
)

const tag = "JsonSocket"

type JsonSocket struct {
	socket net.Conn
	reader *bufio.Reader
}

func NewJsonSocket(host string, port int) *JsonSocket {
	socket, _ := net.Dial("tcp", host+":"+string(port))
	return &JsonSocket{socket: socket}
}

func (jsonSocket JsonSocket) Close() {
	jsonSocket.socket.Close()
}

func (jsonSocket JsonSocket) Send(msg Message) {
	js, _ := json.Marshal(msg) //TODO should we convert to utf8 or is it okay already?
	jsonSocket.socket.Write(append(js, byte('\000')))
}

func (jsonSocket JsonSocket) Get() interface{} {
	js := make([]byte, 1000)
	if jsonSocket.reader == nil {
		jsonSocket.reader = bufio.NewReader(jsonSocket.socket)
	}
	for {
		char, err := jsonSocket.reader.ReadByte()
		if err != nil || char == 0 { //TODO error
			break
		}
		js = append(js, char)
	}
	var result interface{}
	json.Unmarshal(js, &result)
	return result
}
