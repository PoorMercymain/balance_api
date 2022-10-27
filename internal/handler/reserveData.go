package handler

import "github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"

type reserveData struct {
	userId    domain.Id `json:"user_id"`
	serviceId domain.Id `json:"service_id"`
	orderId   domain.Id `json:"order_id"`
	amount    uint32    `json:"amount"`
}
