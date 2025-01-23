package repository

import "clothing-pair-project/entity"

type PaymentMethodRepository interface {
	FindAll() ([]entity.PaymentMethod, error)
	FindByID(paymentMethodID int) (entity.PaymentMethod, error)
	Add(paymentMethod entity.PaymentMethod) error
}
