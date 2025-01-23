package cli

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/service"
	"log"
)

func AddPaymentMethod(paymentMethodService service.PaymentMethodService, name string) {
	addPaymentMethod := entity.PaymentMethod{
		Name: name,
	}

	err := paymentMethodService.Add(addPaymentMethod)
	if err != nil {
		log.Fatal("Failed to add payment method:", err.Error())
	}
	log.Println("Successfully added payment method:", addPaymentMethod)
}

func FindAllPaymentMethods(paymentMethodService service.PaymentMethodService) {
	paymentMethods, err := paymentMethodService.FindAll()
	if err != nil {
		log.Fatal("Failed to find all payment methods:", err.Error())
	}
	log.Println("Successfully found payment methods:", paymentMethods)
}

func FindPaymentMethodByID(paymentMethodService service.PaymentMethodService, id int) {
	paymentMethod, err := paymentMethodService.FindByID(id)
	if err != nil {
		log.Fatal("Failed to find payment method by ID:", err.Error())
	}
	log.Println("Successfully found payment method by ID:", paymentMethod)
}

func UpdatePaymentMethod(paymentMethodService service.PaymentMethodService, id int, name string) {
	updatePaymentMethod := entity.PaymentMethod{
		PaymentMethodID: id,
		Name:            name,
	}
	err := paymentMethodService.Update(updatePaymentMethod)
	if err != nil {
		log.Fatal("Failed to update payment method:", err.Error())
	}
	log.Println("Successfully updated payment method:", updatePaymentMethod)
}

func DeletePaymentMethod(paymentMethodService service.PaymentMethodService, id int) {
	err := paymentMethodService.Delete(id)
	if err != nil {
		log.Fatal("Failed to delete payment method:", err.Error())
	}
	log.Println("Successfully deleted payment method with ID:", id)
}
