package routes

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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

	token := jwt_auth.GenerateToken(userDB)
	data, _ := json.Marshal(map[string]string{
		"token": token,
		"admin": strconv.FormatBool(userDB.Admin),
	})

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)

}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var UserRegisterDTO models.UserRegisterDTO

	_ = json.NewDecoder(r.Body).Decode(&UserRegisterDTO)
	user.Username = UserRegisterDTO.Username
	user.Email = UserRegisterDTO.Email
	hash, _ := hashPassword(UserRegisterDTO.Password)
	user.Password = hash
	user.Admin = false

	checkUserByUsernameExists, _ := repository.GetUserByUsername(user.Username)
	if checkUserByUsernameExists.Username != "" {
		_, _ = w.Write([]byte("User already exists"))
		return
	}

	checkUserByEmailExists, _ := repository.GetUserByEmail(user.Email)
	if checkUserByEmailExists.Email != "" {
		_, _ = w.Write([]byte("Email already in use"))
		return
	}

	user, err := repository.CreateUser(user)
	if err != nil {
		_, _ = w.Write([]byte("Error creating user"))
		return
	}
	_, _ = w.Write([]byte("User created successfully"))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
