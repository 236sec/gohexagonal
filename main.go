package main

import (
	"fmt"

	"github.com/236sec/hexagonal/repository"
	"github.com/236sec/hexagonal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	db.AutoMigrate(&repository.Customer{})
	fmt.Println("Connected Database")
	custrepo := repository.CreateNewCustomerRepository(db)
	custservice := service.NewCustomerService(custrepo)
	_ = custservice

}
