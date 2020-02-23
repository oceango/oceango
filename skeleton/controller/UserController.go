package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/oceango/skeleton/model"
	"github.com/oceango/web/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{db: db.GetDb()}
}

func (c UserController) Welcome(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintf(w, "welcome")
}

func (c UserController) Info(w http.ResponseWriter, r *http.Request)  {
	user := r.Context().Value("user").(model.User)

	json.NewEncoder(w).Encode(user)
}

func (c UserController) Login(w http.ResponseWriter, r *http.Request)  {
	type Credentials struct {
		Password string `json:"password"`
		Telephone string `json:"telephone"`
	}
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Print(err)
	}
	log.Print(creds)
	telephone := creds.Telephone
	password := creds.Password


	// query user
	var user model.User
	c.db.First(&user, "telephone = ?", telephone)

	if user.Password != password {
		panic("password error")
	}

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	 jid := strconv.Itoa(user.Id)
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
		Issuer:    "test",
		Subject: "test",
		Id: jid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokens, err := token.SignedString(mySigningKey)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
		User model.User `json:"user"`
	}{
		Token:tokens,
		User: user,
	})
}

func (c UserController) Register(w http.ResponseWriter, r *http.Request)  {
	type Credentials struct {
		Password string `json:"password"`
		Telephone string `json:"telephone"`
	}
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		panic(err)
	}
	// save user
	user := &model.User{
		Telephone: creds.Telephone,
		Password:  creds.Password,
	}


	c.db.Create(user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}


