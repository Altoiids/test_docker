package controller

import (
	"net/http"
	"strings"
	"strconv"
	"mvc/pkg/models"
)

func WithdrawReturnRequest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	tokenString := strings.TrimSpace(cookie.Value)
	claims, err := models.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	username := claims.Username

	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	
	error := models.WithdrawReturnRequest(username,bookId)
	if error != nil {
		http.Redirect(w, r, "/client/serverError", http.StatusFound)
		return
	}
	
	http.Redirect(w, r, "/client/userReturn", http.StatusSeeOther)

}

