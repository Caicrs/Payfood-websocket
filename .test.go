package main

import (
	"fmt"
	"strconv"
	"time"
)

type OrderMessage struct {
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

// -------------------
func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	t := time.Now()
	hr := t.Hour()
	min := t.Minute()
	sec := t.Second()
	hour := strconv.Itoa(hr) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)

	day := t.Day()
	year := t.Year()
	month := t.Month().String()
	date := strconv.Itoa(day) + "/" + month + "/" + strconv.Itoa(year)

	var x *OrderMessage = &OrderMessage{
		id:          "client-id",
		client:      "client-name",
		marketplace: "marketplace-id",
		table:       1,
		products: Product{
			id:    "id-product",
			name:  "name-product",
			price: 29.99,
		},
		totalPrice: 29.99,
		date:       date,
		hour:       hour,
	}

	fmt.Println(*x)
}

// -------------------
var o *WSMessageOrder = &WSMessageOrder{
	Type: "order",
	Payload: OrderMessage{
		By:          message.Payload.By,
		Room:        message.Payload.Room,
		id:          "1",
		client:      messageOrder.Payload.client,
		marketplace: messageOrder.Payload.marketplace,
		table:       messageOrder.Payload.table,
		products: Product{
			id:    "product-id-1",
			name:  messageProduct.Payload.name,
			price: messageProduct.Payload.price,
		},
		totalPrice: messageOrder.Payload.totalPrice,
		date:       messageOrder.Payload.date,
		hour:       messageOrder.Payload.hour,
	},
}

//-------------------
client: "client-name",
marketplace: "general-marketplace",
table: 1,
products: [
  {
	id: "product-id-1",
	name: "product-name-1",
	price: 12.99,
  },
  {
	id: "product-id-2",
	name: "product-name-2",
	price: 22.99,
  },
],
totalPrice: 12.99,
date: "20/20/2022",
hour: "9:10",
},
by: username,
room: room,

// ---------------

form.onsubmit = function (e) {
	e.preventDefault();
	const m = {
	  type: "message",
	  payload: {
		body: [
		  {
			client: "client-name",
		  },
		  {
			marketplace: "general-marketplace",
		  },
		  { table: 1 },
		  {
			products: [
			  {
				id: "product-id-1",
				name: "product-name-1",
				price: 12.99,
			  },
			  {
				id: "product-id-2",
				name: "product-name-2",
				price: 22.99,
			  },
			],
		  },
		  {
			totalPrice: 12.99,
		  },
		  {
			date: "20/20/2022",
		  },
		  {
			hour: "9:10",
		  },
		],
	  },
	  by: username,
	  room: room,
	};
	ws.send(JSON.stringify(m));
	messageInput.value = "";
  };