package entity

type PaymentMethod struct {
	PaymentMethodID int    `db:"payment_method_id"`
	Name            string `db:"name"`
}
