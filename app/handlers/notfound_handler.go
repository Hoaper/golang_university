package handlers

import (
	"github.com/Hoaper/golang_university/app/utils"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	utils.RespondWithError(w, 404, "Page not found!")
}
