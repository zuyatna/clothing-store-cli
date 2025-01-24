package handler

import (
	"clothing-pair-project/entity"

	"github.com/jmoiron/sqlx"
)

type PaymentMethodHandler struct {
	db *sqlx.DB
}

func NewPaymentMethodHandler(db *sqlx.DB) *PaymentMethodHandler {
	return &PaymentMethodHandler{db: db}
}

func (h *PaymentMethodHandler) FindAll() ([]entity.PaymentMethod, error) {
	var paymentMethods []entity.PaymentMethod
	query := `SELECT payment_method_id, name FROM payment_methods`
	err := h.db.Select(&paymentMethods, query)
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

func (h *PaymentMethodHandler) FindByID(paymentMethodID int) (entity.PaymentMethod, error) {
	var paymentMethod entity.PaymentMethod
	query := `SELECT payment_method_id, name FROM payment_methods WHERE payment_method_id = $1`
	err := h.db.Get(&paymentMethod, query, paymentMethodID)
	if err != nil {
		return entity.PaymentMethod{}, err
	}
	return paymentMethod, nil
}

func (h *PaymentMethodHandler) Add(paymentMethod entity.PaymentMethod) error {
	query := `INSERT INTO payment_methods (name) VALUES ($1)`
	_, err := h.db.Exec(query, paymentMethod.Name)
	return err
}

func (h *PaymentMethodHandler) Delete(paymentMethodID int) error {
	query := `DELETE FROM payment_methods WHERE payment_method_id = $1`
	_, err := h.db.Exec(query, paymentMethodID)
	return err
}

func (h *PaymentMethodHandler) Update(paymentMethod entity.PaymentMethod) error {
	query := `UPDATE payment_methods SET name = $1 WHERE payment_method_id = $2`
	_, err := h.db.Exec(query, paymentMethod.Name, paymentMethod.PaymentMethodID)
	return err
}

func (h *PaymentMethodHandler) ResetIncrement() error {
	query := `ALTER SEQUENCE "PaymentMethod_PaymentMethodID_seq" RESTART WITH 1`
	_, err := h.db.Exec(query)
	return err
}
