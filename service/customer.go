package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Amount     int    `json:"amount"`
}

type CustomerService interface {
	CreateUser(CustomerResponse) error
	GetAll() ([]CustomerResponse, error)
	FindUser(int) (*CustomerResponse, error)
}
