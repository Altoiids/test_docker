package controller

import (
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"strings"
)

func ProfilePage(w http.ResponseWriter, r *http.Request) {
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
	
	booksList,err := models.ProfileBooks(username)
	if err != nil {
		http.Redirect(w, r, "/client/serverError", http.StatusFound)
		return
	}
	file := views.FileNames()
	t := views.ViewClientPages(file.Profile)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, booksList)
}

