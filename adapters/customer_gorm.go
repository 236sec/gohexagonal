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

func (r *customerRepositoryDB) Create() (*core.Customer, error) {
	return nil,nil
}
