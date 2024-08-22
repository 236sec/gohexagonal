package repository

import (
	"time"

	"github.com/236sec/hexagonal/service"
)

type Customer struct {
	CustomerID uint      `db:"customer_id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Amount     int       `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type CustomerRepository interface {
	Create(service.CustomerResponse) error
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
