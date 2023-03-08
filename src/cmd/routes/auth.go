package routes

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"wiselink-challenge/src/internal/database"
	jwt_auth "wiselink-challenge/src/internal/jwt_bearer"
	"wiselink-challenge/src/models"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.DbConnection()

	var UserLoginDTO models.UserLoginDTO
	_ = json.NewDecoder(r.Body).Decode(&UserLoginDTO)

	var userDB models.User
	db.Where("username = ?", UserLoginDTO.Username).First(&userDB)

	if CheckPasswordHash(userDB.Password, UserLoginDTO.Password) {
		token := jwt_auth.GenerateToken(userDB)

		data, _ := json.Marshal(map[string]string{
			"token": token,
		})

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	} else {
		_, _ = w.Write([]byte("Invalid username or password"))
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := database.DbConnection()

	var user models.User
	var UserRegisterDTO models.UserRegisterDTO

	_ = json.NewDecoder(r.Body).Decode(&UserRegisterDTO)

	user.Username = UserRegisterDTO.Username
	user.Email = UserRegisterDTO.Email
	hash, _ := HashPassword(UserRegisterDTO.Password)
	user.Password = hash
	user.Admin = false

	db.Create(&user)

	_, _ = w.Write([]byte("User created successfully"))
}
