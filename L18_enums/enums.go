package main

import "fmt"

// enumerated types

type MyType string // custom types can be made in go

type OrderStatus string

// type OrderStatus int

// const (
// 	Received  OrderStatus = iota
// 	Confirmed             
// 	Prepared              
// 	Delivered             
// )

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)
}