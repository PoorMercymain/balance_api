package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PoorMercymain/balance_api/internal/domain"
	"github.com/PoorMercymain/balance_api/pkg/router"
	"net/http"
)

type user struct {
	srv domain.UserService
}

func NewUser(srv domain.UserService) *user {
	return &user{srv: srv}
}

type addMoney struct {
	Id      domain.Id `json:"id"`
	Money   uint32    `json:"money"`
	WhoMade string    `json:"who_made"`
	Reason  string    `json:"reason"`
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

	fmt.Println(data)

	err := h.srv.ReserveMoney(r.Context(), data.UserId, data.ServiceId, data.OrderId, data.Amount, data.WhoMade)
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

	err := h.srv.AddMoney(r.Context(), data.Id, data.Money, data.WhoMade, data.Reason)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *user) TransactionList(w http.ResponseWriter, r *http.Request) {
	userId, err := router.Params(r).Uint32("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	page, err := h.srv.TransactionList(r.Context(), domain.Id(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = reply(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *user) MakeReport(w http.ResponseWriter, r *http.Request) {
	var data domain.DateForReport

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := h.srv.MakeReport(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = reply(w, link); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
