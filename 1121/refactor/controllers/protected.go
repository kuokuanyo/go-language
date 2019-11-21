package controllers

import (
	"net/http"
	"refactor/utils"
)

func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseJSON(w, "Yes")
	}
}
