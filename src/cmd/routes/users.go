package routes

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"wiselink-challenge/src/cmd/service"
	jwt_auth "wiselink-challenge/src/internal/jwt"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))

	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}
	data := service.GetProfileService(claims)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
