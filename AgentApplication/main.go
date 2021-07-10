package main

import (
	"XWS-Nistagram/AgentApplication/handler"
	"XWS-Nistagram/AgentApplication/model"
	"XWS-Nistagram/AgentApplication/repository"
	"XWS-Nistagram/AgentApplication/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func seedData(db *gorm.DB){
	user := model.User{FirstName: "Mika",LastName: "Mikic",Email: "mika",Password: "mika",Role:"User"}
	user1 := model.User{FirstName: "Pera",LastName: "Vlahovic",Email: "pera",Password: "pera",Role:"Agent"}
	user2 := model.User{FirstName: "Zoki",LastName: "Lazic",Email: "zoran",Password: "zoran",Role:"User"}
	address := model.Address{City: "Novi Sad",Address: "Vojvode Supljikca 22",ZipCode: "21000",Country: "Srbija"}
	product := model.Product{Name: "Lenovo thinkpad T14 G1",Description: "Lenovo ThinkPad T14 G1 (20S0000SCX) laptop Intel® Quad Core™ i7 10510U 14\" UHD 16GB 512GB SSD Intel® UHD Graphics Win10 Pro crni",AvailableQuantity: 30,Price: 100,Image: "thinkpad1.jpg"}
	product2 := model.Product{Name: "Computer",Description: "Gandiva Economical C2D Desktop Computer(Core2Duo CPU/4GB DDR3 RAM/250GB HDD/15.6 inch Monitor/WiFi)Windows 10&MS Office(Trial Version)& Antivirus(Free Version)",AvailableQuantity: 20,Price: 170,Image: "computer.jpg"}
	product3 := model.Product{Name: "Mouse",Description: "The Level 20 RGB is a high-performance gaming mouse equipped with a powerful gaming grade 16,000 DPI optical sensor and durable OMRON switches rated up to 50 million clicks for endless hours of gameplay.",AvailableQuantity: 20,Price: 170,Image: "mouse.jpg"}
	product4 := model.Product{Name: "Thinkpad T4",Description: "Lagan, tanak i građen da nastupa bilo gdje. Trajanje baterije do 15 sati, moćni procesor i dovoljno memorije za sve vrste uživanja. Dostupna touch screen opcija.",AvailableQuantity: 20,Price: 1750,Image: "thinkpadt4.jpg"}


	order := model.Order{Product: product,Amount: 5}
//	order2 := model.Order{Product: product2,Amount: 10}
//	order3 := model.Order{Product: product3,Amount: 10}
//	order4 := model.Order{Product: product4,Amount: 150}


	shoppingCart := model.ShoppingCart{Orders:[]model.Order{order},TotalPrice: 500,User: user,Address: address}
	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.Address{})
	db.Migrator().DropTable(&model.Product{})
	db.Migrator().DropTable(&model.Order{})
	db.Migrator().DropTable(&model.ShoppingCart{})
	db.Migrator().DropTable(&model.Purchase{})


	db.AutoMigrate(&model.User{},&model.Address{},&model.Product{},&model.Order{},&model.ShoppingCart{},&model.Purchase{})
	db.Create(&user1)
	db.Create(&user2)
	//db.Create(&product)
	db.Create(&product3)
	db.Create(&product4)
	db.Create(&product2)
	db.Create(&shoppingCart)
}
func initDB() *gorm.DB {
/*	host := os.Getenv("DBHOST")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	user1 := os.Getenv("USER")

*/
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", "localhost", "postgres", "root", "postgres",  "5432")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func initUserRepo(database *gorm.DB) *repository.UserRepository {
	return &repository.UserRepository{Database: database}
}

func initUserService(repo *repository.UserRepository) *service.UserService {
	return &service.UserService{Repository: repo}
}

func initUserHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler{Service: service}
}

func initProductRepo(database *gorm.DB) *repository.ProductRepository {
	return &repository.ProductRepository{Database: database}
}

func initProductService(repo *repository.ProductRepository) *service.ProductService {
	return &service.ProductService{Repository: repo}
}

func initProductHandler(service *service.ProductService) *handler.ProductHandler{
	return &handler.ProductHandler{Service: service}
}

func initMediaStorageHandler() *handler.MediaStorageHandler{
	return &handler.MediaStorageHandler{}
}

func initShoppingCartRepo(database *gorm.DB) *repository.ShoppingCartRepository {
	return &repository.ShoppingCartRepository{Database: database}
}

func initShoppingCartService(repo *repository.ShoppingCartRepository,productRepo *repository.ProductRepository) *service.ShoppingCartService {
	return &service.ShoppingCartService{Repository: repo,ProductRepository: productRepo}
}

func initShoppingCartHandler(service *service.ShoppingCartService) *handler.ShoppingCartHandler{
	return &handler.ShoppingCartHandler{Service: service}
}

func initAddressRepo(database *gorm.DB) *repository.AddressRepository {
	return &repository.AddressRepository{Database: database}
}

func initAddressService(repo *repository.AddressRepository) *service.AddressService {
	return &service.AddressService{Repository: repo}
}

func initAddressHandler(service *service.AddressService) *handler.AddressHandler{
	return &handler.AddressHandler{Service: service}
}


func handleFunctions(userHandler *handler.UserHandler,productHandler *handler.ProductHandler,mediaStorageHandler *handler.MediaStorageHandler,shoppingCartHandler *handler.ShoppingCartHandler,addressHandler *handler.AddressHandler) {
	router := gin.Default()
	router.POST("/users/registerUser", userHandler.RegisterUser)
	router.POST("/users/login", userHandler.Login)
	router.POST("/products/create", productHandler.CreateProduct)
	router.POST("/products/update", productHandler.UpdateProduct)
	router.POST("/products/delete", productHandler.DeleteProduct)
	router.GET("/products/findAll", productHandler.FindAll)
	router.POST("/products/findById", productHandler.FindById)
	router.GET("/media/getImage", mediaStorageHandler.GetMediaImage)
	router.POST("/media/uploadImage", mediaStorageHandler.UploadMediaImage)
	router.POST("/shoppingCart/create", shoppingCartHandler.CreateShoppingCart)
	router.POST("/shoppingCart/delete", shoppingCartHandler.DeleteShoppingCart)
	router.POST("/shoppingCart/findById", shoppingCartHandler.FindById)
	router.POST("/shoppingCart/addOrderToCart", shoppingCartHandler.AddOrderToShoppingCart)
	router.POST("/shoppingCart/deleteOrderFromCart", shoppingCartHandler.DeleteOrderFromShoppingCart)
	router.POST("/shoppingCart/findByUser", shoppingCartHandler.FindByUser)
	router.POST("/shoppingCart/updateOrderQuantity", shoppingCartHandler.UpdateOrderQuantity)
	router.POST("/address/createAddress", addressHandler.CreateAddress)
	router.POST("/purchase/createPurchase", shoppingCartHandler.CreatePurchase)
	router.POST("/shoppingCart/emptyShoppingCart", shoppingCartHandler.EmptyShoppingCart)
	log.Fatal(router.Run(":8070"))

}
func main() {
	database := initDB()
	seedData(database)
	userRepo := initUserRepo(database)
	userService := initUserService(userRepo)
	userHandler :=initUserHandler(userService)
	productRepo := initProductRepo(database)
	productService  := initProductService(productRepo)
	productHandler := initProductHandler(productService)
	shoppingCartRepo := initShoppingCartRepo(database)
	shoppingCartService  := initShoppingCartService(shoppingCartRepo,productRepo)
	shoppingCartHandler := initShoppingCartHandler(shoppingCartService)
	addressRepo := initAddressRepo(database)
	addressService  := initAddressService(addressRepo)
	addressHandler := initAddressHandler(addressService)
	mediaStorageHandler := initMediaStorageHandler()
	handleFunctions(userHandler,productHandler,mediaStorageHandler,shoppingCartHandler,addressHandler)
}
