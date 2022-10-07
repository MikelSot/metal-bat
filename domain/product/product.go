package product

import (
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/MikelSot/metal-bat/model"
)

type Storage interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Product, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Products, error)
}

type UseCase interface {
	Create(m *model.Product) (model.Product, error)
	Update(m *model.Product) (model.Product, error)
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Product, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Products, error)
}
