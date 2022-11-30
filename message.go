package main

type Message struct {
	Body string `json:"body"`
	By   string `json:"by"`
	Room string `json:"room"`
}

type OrderMessage struct {
	id          string   `json:"id"`
	client      string   `json:"client"`
	marketplace string   `json:"marketplace"`
	table       int      `json:"table"`
	products    Products `json:"products"`
	totalPrice  float32  `json:"totalPrice"`
	createdAt   string   `json:"marketplace"`
}

type Products struct {
	id          string  `json:"id"`
	name        string  `json:"name"`
	price       float32 `json:"price"`
	marketplace string  `json:"marketplace"`
}
