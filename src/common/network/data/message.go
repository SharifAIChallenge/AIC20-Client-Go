package data

import "encoding/json"

type Message struct {
	Name string
	Args []map[string]interface{}
}

func NewMessage(name string, args ...interface{}) *Message {
	var args1 = make([]map[string]interface{}, 0)
	for _, arg := range args {
		var tmp interface{}
		b, _ := json.Marshal(arg)
		_ = json.Unmarshal(b, &tmp)
		args1 = append(args1, tmp.(map[string]interface{}))
	}
	return &Message{name, args1}
}