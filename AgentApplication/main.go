package main

import (
	"XWS-Nistagram/AgentApplication/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {


	user := model.User{FirstName: "Mika",LastName: "Mikic",Email: "mika@mikic.com",Password: "12345678"}
	address := model.Address{City: "Novi Sad",StreetName: "Vojvode Supljikca",StreetNumber: "21",Latitude: "153",Longitude: "123"}
	product := model.Product{Name: "Komp",Description: "adad",AvailableQuantity: 30,Price: 100}
	product2 := model.Product{Name: "AAAA",Description: "aAAAAAdad",AvailableQuantity: 20,Price: 170}

	order := model.Order{Product: product,Amount: 5}
	order2 := model.Order{Product: product2,Amount: 10}

	paymentDetails := model.PaymentDetails{PhoneNumber: "+3816767987",Address:address}
	shoppingCart := model.ShoppingCart{Orders:[]model.Order{order,order2},TotalPrice: 500,User: user,PaymentDetails: paymentDetails}

	host := os.Getenv("DBHOST")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")
	user1 := os.Getenv("USER")


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user1, password, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.Address{})
	db.Migrator().DropTable(&model.ShoppingCart{})
	db.Migrator().DropTable(&model.PaymentDetails{})
	db.Migrator().DropTable(&model.Order{})
	db.Migrator().DropTable(&model.Product{})


	db.AutoMigrate(&model.User{},&model.Address{},&model.Product{},&model.PaymentDetails{},&model.Order{},&model.ShoppingCart{})

	db.Create(&shoppingCart)

	fmt.Println("Successfully connected!")


}
