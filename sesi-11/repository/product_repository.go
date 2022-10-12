package repository

import "go_unit_test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}