package service

import (
	"os"
	"time"

	"github.com/AuthSystemJWT/deyki/v2/database"
	"github.com/AuthSystemJWT/deyki/v2/util"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignUpResponse struct {
	CreatedAt 	time.Time
	Username 	string
}


type SignInResponse struct {
	JWToken		string
}


func SignUp(user *database.User) (*SignUpResponse, *util.ErrorMessage) {

	db, err := database.ConnectDB()
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}


	hash, errorMessage := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errorMessage != nil {
		return nil, util.ErrorMessage{}.FailedToCreateHashFromPassword()
	}

	user.Password = string(hash)

	db.Create(&user)

	return &SignUpResponse{user.CreatedAt, user.Username}, nil
}


func SignIn(requestBodyUser *database.User) (*SignInResponse, *util.ErrorMessage) {

	db, err := database.ConnectDB()
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}

	var user database.User

	errorMessage := db.First(&user, "username = ?", requestBodyUser.Username).Error
	if errorMessage != nil {
		return nil, util.ErrorMessage{}.UserNotFound()
	}

	compareHashAndPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBodyUser.Password))
	if compareHashAndPass != nil {
		return nil, util.ErrorMessage{}.UserNotFound()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, errorMessage := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if errorMessage != nil {
		return nil, util.ErrorMessage{}.FailedToCreateJWToken()
	}

	return &SignInResponse{tokenString}, nil
}