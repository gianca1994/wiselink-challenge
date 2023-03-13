package routes

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"wiselink-challenge/src/cmd/repository"
	jwt_auth "wiselink-challenge/src/internal/jwt"
	"wiselink-challenge/src/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var UserLoginDTO models.UserLoginDTO
	var userDB models.User

	_ = json.NewDecoder(r.Body).Decode(&UserLoginDTO)
	userDB, err := repository.GetUserByUsername(UserLoginDTO.Username)
	if err != nil {
		_, _ = w.Write([]byte("User not found"))
		return
	}
	if !checkPasswordHash(userDB.Password, UserLoginDTO.Password) {
		_, _ = w.Write([]byte("Invalid password"))
		return
	}

	data := jwt_auth.GenerateToken(userDB)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)

}

func Register(w http.ResponseWriter, r *http.Request) {
	var UserRegisterDTO models.UserRegisterDTO

	_ = json.NewDecoder(r.Body).Decode(&UserRegisterDTO)

	checkUserByUsernameExists, _ := repository.GetUserByUsername(UserRegisterDTO.Username)
	if checkUserByUsernameExists.Username != "" {
		_, _ = w.Write([]byte("User already exists"))
		return
	}
	checkUserByEmailExists, _ := repository.GetUserByEmail(UserRegisterDTO.Email)
	if checkUserByEmailExists.Email != "" {
		_, _ = w.Write([]byte("Email already in use"))
		return
	}

	_, err := repository.CreateUser(models.User{
		Username: UserRegisterDTO.Username,
		Email:    UserRegisterDTO.Email,
		Password: hashPassword(UserRegisterDTO.Password),
		Admin:    UserRegisterDTO.Username == "admin",
	})
	if err != nil {
		_, _ = w.Write([]byte("Error creating user"))
		return
	}
	_, _ = w.Write([]byte("User created successfully"))
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
