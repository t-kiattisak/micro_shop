package kafka

import (
	"encoding/json"
	"log"
	"payment-service/internal/usecase"
)

type ConsumerHandler struct {
	paymentUseCase *usecase.PaymentUseCase
}

func NewConsumerHandler(paymentUseCase *usecase.PaymentUseCase) *ConsumerHandler {
	return &ConsumerHandler{paymentUseCase: paymentUseCase}
}

func (h *ConsumerHandler) ProcessPaymentMessage(msg []byte) {
	var event struct {
		OrderID uint    `json:"order_id"`
		Amount  float64 `json:"amount"`
		Status  string  `json:"status"`
	}

	if err := json.Unmarshal(msg, &event); err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	log.Printf("Processing payment for order %d, amount: %.2f", event.OrderID, event.Amount)
	err := h.paymentUseCase.ProcessPayment(event.OrderID, event.Amount)
	if err != nil {
		log.Printf("Payment processing failed: %v", err)
	} else {
		log.Printf("Payment for order %d completed successfully!", event.OrderID)
	}
}
