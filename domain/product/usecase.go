package product

import (
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/metal-bat/model"
)

var allowedFieldsForQuery = []string{"id"}

type Product struct {
	storage Storage
}

func New(s Storage) Product {
	return Product{s}
}

func (p Product) Create(m *model.Product) (model.Product, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.Product{}, fmt.Errorf("product: %w", model.ErrNilPointer)
	}

	if err := m.ValidateFields(); err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	err := p.storage.Create(m)
	if err != nil {
		return model.Product{}, handleStorageErr(err)
	}

	return model.Product{}, nil
}

func (p Product) Update(m *model.Product) (model.Product, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.Product{}, fmt.Errorf("product: %w", model.ErrNilPointer)
	}

	if err := m.ValidateFields(); err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	err := p.storage.Update(m)
	if err != nil {
		return model.Product{}, handleStorageErr(err)
	}

	return model.Product{}, nil
}

func (p Product) Delete(ID uint) error {
	err := p.storage.Delete(ID)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

func (p Product) GetWhere(specification models.FieldsSpecification) (model.Product, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	group, err := p.storage.GetWhere(specification)
	if err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	return group, nil
}

func (p Product) GetAllWhere(specification models.FieldsSpecification) (model.Products, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("product: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("product: %w", err)
	}

	products, err := p.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("product: %w", err)
	}

	return products, nil
}

func handleStorageErr(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
