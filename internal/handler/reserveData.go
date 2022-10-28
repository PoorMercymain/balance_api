package handler

import "github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"

type reserveData struct {
	UserId    domain.Id `json:"user_id"`
	ServiceId domain.Id `json:"service_id"`
	OrderId   domain.Id `json:"order_id"`
	Amount    uint32    `json:"amount"`
}
