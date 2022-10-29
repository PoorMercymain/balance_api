package handler

import "github.com/PoorMercymain/balance_api/internal/domain"

type reserveData struct {
	UserId    domain.Id `json:"user_id"`
	ServiceId domain.Id `json:"service_id"`
	OrderId   domain.Id `json:"order_id"`
	Amount    uint32    `json:"amount"`
	WhoMade   string    `json:"who_made"`
}
