package core

import (
	"time"
)

type Customer struct {
	ID uint      `db:"customer_id" gorm:"primaryKey"`
	Name       string    `db:"name"`
	Email      string    `db:"email" gorm:"unique"`
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