package handler

import (
	"net/http"

	"github.com/PoorMercymain/balance_api/pkg/router"
)

func reply(w http.ResponseWriter, data interface{}) error {
	return router.Reply(w, data)
}
