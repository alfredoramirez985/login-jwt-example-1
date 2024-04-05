package api

import (
	"encoding/json"
	"fmt"
	"log"
	"login-jwt-example/pkg/db/models"
	"net/http"

	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type UserResponse struct {
    Success bool            `json:"success"`
    Error   string          `json:"error"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var loginUser LoginUser
	json.NewDecoder(r.Body).Decode(&loginUser)
	fmt.Printf("The user request value %v", loginUser)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(loginUser.Password), bcrypt.DefaultCost)
    if err != nil {
        res := &UserResponse{
            Success: false,
            Error: err.Error(),
        }
        err = json.NewEncoder(w).Encode(res)
        //if there's an error with encoding handle it
        if err != nil {
            log.Printf("error sending response %v\n", err)
        }
        //return a bad request and exist the function
        w.WriteHeader(http.StatusBadRequest)
        return
    }

	//get the db from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
        res := &UserResponse{
            Success: false,
            Error:   "could not get the DB from context",
        }
        err = json.NewEncoder(w).Encode(res)
        //if there's an error with encoding handle it
        if err != nil {
            log.Printf("error sending response %v\n", err)
        }
        //return a bad request and exist the function
        w.WriteHeader(http.StatusBadRequest)
        return
    }

	loginData, err := models.GetLoginData(pgdb, loginUser.Username)

	if loginUser.Username == loginData.UserName && string(hashedPass) == loginData.Password {
		tokenString, err := CreateToken(loginUser.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	fmt.Fprint(w, "Welcome to the the protected area")

}