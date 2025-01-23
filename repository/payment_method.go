package repository

import "clothing-pair-project/entity"

type PaymentMethodRepository interface {
	GetAll() ([]entity.PaymentMethod, error)
	GetByID(paymentMethodID int) (entity.PaymentMethod, error)
}

type paymentMethodRepository struct {
	paymentMethods []entity.PaymentMethod
}
