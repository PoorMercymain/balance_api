package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	keyParams = iota
	KeyPagination
)

var (
	ErrParamKey  = errors.New("error no key")
	ErrParamType = errors.New("error type key")
)

type Pagination struct {
	Limit         uint32
	Offset        uint32
	SortField     string
	SortDirection string
}

func NewPagination(limit string, page string, sortField string, sortDirection string) *Pagination {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	return &Pagination{Limit: uint32(limitInt), Offset: uint32((pageInt - 1) * limitInt), SortField: sortField, SortDirection: sortDirection}
}

type params struct {
	value httprouter.Params
}

func Params(r *http.Request) *params {
	return &params{value: r.Context().Value(keyParams).(httprouter.Params)}
}

func (p params) String(key string) (string, error) {
	value := p.value.ByName(key)
	if value == "" {
		return "", ErrParamKey
	}
	return value, nil
}

func (p params) Int(key string) (int, error) {
	value := p.value.ByName(key)
	if value == "" {
		return 0, ErrParamKey
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, ErrParamType
	}
	return i, nil
}

func (p params) Uint32(key string) (uint32, error) {
	i, err := p.Int(key)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

func (p params) Uint16(key string) (uint16, error) {
	i, err := p.Int(key)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}

func (p params) Bool(key string) (bool, error) {
	value := p.value.ByName(key)
	if value == "" {
		return false, ErrParamKey
	}

	res, err := strconv.ParseBool(value)
	if err != nil {
		return false, ErrParamType
	}
	return res, nil
}

func Reply(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(data)
}

func WrapHandler(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), keyParams, ps)
		ctx = context.WithValue(ctx, KeyPagination, NewPagination(r.FormValue("limit"), r.FormValue("page"), r.FormValue("sort_by"), r.FormValue("sort_dir")))

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}
