package handler

import (
	"encoding/json"
	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"
	"github.com/PoorMercymain/REST-API-work-duration-counter/pkg/router"
	"net/http"
)

type order struct {
	srv domain.OrderService
}

func NewOrder(srv domain.OrderService) *order {
	return &order{srv: srv}
}

type addService struct {
	orderId   domain.Id `json:"order_id"`
	serviceId domain.Id `json:"service_id"`
}

func (h *order) Create(w http.ResponseWriter, r *http.Request) {
	var data domain.Order

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.srv.Create(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := struct {
		ID domain.Id `json:"id"`
	}{ID: id}

	if err = reply(w, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *order) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := router.Params(r).Uint32("id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.srv.Delete(r.Context(), domain.Id(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *order) Update(w http.ResponseWriter, r *http.Request) {
	var data domain.Order

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.srv.Update(r.Context(), data.OrderId, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *order) Read(w http.ResponseWriter, r *http.Request) {
	id, err := router.Params(r).Uint32("id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.srv.Read(r.Context(), domain.Id(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = reply(w, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *order) AddService(w http.ResponseWriter, r *http.Request) {
	var data addService

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.srv.AddService(r.Context(), data.orderId, data.serviceId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*buf.ServiceIds = append(buf.ServiceIds, domain.Id(id))

	err = h.srv.Update(r.Context(), domain.Id(id), buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
}
