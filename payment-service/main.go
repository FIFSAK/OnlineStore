package main

import (
	"OnlineStore/payment-service/services"
	"fmt"
)

func main() {
	response, err := services.CreatePayment()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(response)
}
