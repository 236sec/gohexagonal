package adapters

import (
	"github.com/236sec/hexagonal/core"
	"gorm.io/gorm"
)

type customerRepositoryDB struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) core.CustomerRepository {
	return &customerRepositoryDB{db: db}
}

func (r *customerRepositoryDB) FindAll() ([]core.Customer, error) {

	return nil, nil
}

func (r *customerRepositoryDB) FindById() (*core.Customer, error) {
	return nil, nil
}

func (r *customerRepositoryDB) Create(customer core.CreateCustomerDto) (error) {
	newcustomer := core.Customer{
		Name: customer.Name,
		Email: customer.Email,
		Password: customer.Password,
	}
	if result := r.db.Create(&newcustomer); result.Error != nil {
		return result.Error
	}
	return nil
}
