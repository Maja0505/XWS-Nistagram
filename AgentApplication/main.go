package main

import (
	"AgentApplication/model"
	"fmt"
)

func main() {

	user := model.User{FistName: "Mika",LastName: "Mikic",Email: "mika@mikic.com",Password: "12345678"}
	product := model.Product{Name: "Komp",Description: "adad",AvailableQuantity: 30,Price: 100}
	order := model.Order{Product: product,Amount: 5}
	paymentDetails := model.PaymentDetails{PhoneNumber: "+3816767987",Address: model.Address{City: "Novi Sad",StreetName: "Vojvode Supljikca",StreetNumber: "21",Latitude: "153",Longitude: "123"}}
	shoppingCart := model.ShoppingCart{Customer: user,Orders:[]model.Order{order},TotalPrice: 500,PaymentDetails: paymentDetails}

	fmt.Println(shoppingCart)

}
