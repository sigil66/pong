package pong

type Message struct {
	Id     string `json:"id,omitempty"`
	To     string `json:"to"`
	Data   map[string]interface{} `json:"data,omitempty"`
}