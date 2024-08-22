package main

import (
	"fmt"

	"github.com/236sec/hexagonal/adapters"
	"github.com/236sec/hexagonal/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	db.AutoMigrate(&core.Customer{})
	fmt.Println("Connected Database")
	custrepo := adapters.NewCustomerRepository(db)
	custservice := core.NewCustomerService(custrepo)
	custhandler := adapters.NewHttpCustomerHandler(custservice)


	app.Post("/customers",custhandler.CreateUser)

	app.Listen(":8000")

}
