package main

type Message struct {
	Body interface{} `json:"body"`
	By   string      `json:"by"`
	Room string      `json:"room"`
}

type Body struct {
	normalMessage string        `json:"normalMessage"`
	orderMessage  []interface{} `json:"orderMessage"`
}

type Order struct {
	Body OrderMessage `json:"body"`
	By   string       `json:"by"`
	Room string       `json:"room"`
}

type OrderMessage struct {
	By          string  `json:"by"`
	Room        string  `json:"room"`
	id          string  `json:"id"`
	client      string  `json:"client"`
	marketplace string  `json:"marketplace"`
	table       int     `json:"table"`
	products    Product `json:"products"`
	totalPrice  float32 `json:"totalPrice"`
	date        string  `json:"date"`
	hour        string  `json:"hour"`
}

type Product struct {
	id    string  `json:"id"`
	name  string  `json:"name"`
	price float32 `json:"price"`
}
