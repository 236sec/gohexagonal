package core

import (
	"time"
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


type CreateCustomerDto struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
}