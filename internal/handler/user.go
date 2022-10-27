package handler

import (
	"encoding/json"
	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"
	"github.com/PoorMercymain/REST-API-work-duration-counter/pkg/router"
	"net/http"
)

type user struct {
	srv domain.UserService
}

func NewUser(srv domain.UserService) *user {
	return &user{srv: srv}
}

type addMoney struct {
	id    domain.Id `json:"id"`
	money uint32    `json:"money"`
}

func (h *user) Create(w http.ResponseWriter, r *http.Request) {
	var data domain.User

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

func (h *user) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *user) Update(w http.ResponseWriter, r *http.Request) {
	var data domain.User

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.srv.Update(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *user) Read(w http.ResponseWriter, r *http.Request) {
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

func (h *user) ReadBalance(w http.ResponseWriter, r *http.Request) {
	id, err := router.Params(r).Uint32("id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.srv.ReadBalance(r.Context(), domain.Id(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = reply(w, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *user) ReserveMoney(w http.ResponseWriter, r *http.Request) {
	var data reserveData

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.srv.ReserveMoney(r.Context(), data.userId, data.serviceId, data.orderId, data.amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *user) AddMoney(w http.ResponseWriter, r *http.Request) {
	var data addMoney

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.srv.AddMoney(r.Context(), data.id, data.money)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
