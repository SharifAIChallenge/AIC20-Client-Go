package data

type Message struct {
	Name string      `json:"type"`
	Args interface{} `json:"info"`
}
