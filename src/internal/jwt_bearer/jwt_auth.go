package jwt_bearer

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
	"wiselink-challenge/src/models"
)

const secretJWTKey = "e18924e1982e4wqa4sd89asd"

func GenerateToken(user models.User) string {
	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Username,
		"admin":    user.Admin,
	}
	jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*60))
	jwtauth.SetIssuedAt(claims, time.Now())
	_, token, _ := jwtauth.New("HS512", []byte(secretJWTKey), nil).Encode(claims)

	return token
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(secretJWTKey)
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

func TokenGetClaims(token string) jwt.MapClaims {
	claims, _ := ExtractClaims(token)

	if claims == nil {
		return nil
	}
	return claims
}
