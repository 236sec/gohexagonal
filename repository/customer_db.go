package repository

import (
	"gorm.io/gorm"
)

type customerRepositoryDB struct {
	db *gorm.DB
}

func CreateNewCustomerRepository(db *gorm.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (c customerRepositoryDB) GetAll() ([]Customer, error) {
	return nil, nil
}

func (c customerRepositoryDB) GetById(id int) (*Customer, error) {
	return nil, nil
}
