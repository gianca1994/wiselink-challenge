package jwt

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"wiselink-challenge/src/models"
)

func GenerateToken(user models.User) string {
	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Username,
		"admin":    user.Admin,
	}
	jwt_expire, _ := strconv.Atoi(os.Getenv("JWT_TIME_EXPIRE_MINUTES"))
	jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*time.Duration(jwt_expire)))
	jwtauth.SetIssuedAt(claims, time.Now())
	_, token, _ := jwtauth.New("HS512", []byte(os.Getenv("JWT_SECRET")), nil).Encode(claims)
	return token
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func TokenGetClaims(token string, w http.ResponseWriter) jwt.MapClaims {
	claims, _ := ExtractClaims(token)
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return nil
	}
	return claims
}
