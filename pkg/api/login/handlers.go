package api

import (
	"encoding/json"
	"fmt"
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
    
    var loginUser LoginUser
    err := json.NewDecoder(r.Body).Decode(&loginUser)
    if err != nil {
        handleError(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    // Get the database connection from the context
    pgdb, ok := r.Context().Value("DB").(*pg.DB)
    if !ok {
        handleError(w, http.StatusInternalServerError, "Failed to get database connection from context")
        return
    }

    // Get login data from the database
    loginData, err := models.GetLoginData(pgdb, loginUser.Username)
    if err != nil {
        handleError(w, http.StatusInternalServerError, "Failed to retrieve user data")
        return
    }

    // Compare the provided password with the hashed password from the database
    err = bcrypt.CompareHashAndPassword([]byte(loginData.Password), []byte(loginUser.Password))
    if err != nil {
        handleError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    // If passwords match, generate and send token
    tokenString, err := CreateToken(loginUser.Username)
    if err != nil {
        handleError(w, http.StatusInternalServerError, "Failed to generate token")
        return
    }

    // Respond with token
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, tokenString)
}

func handleError(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    fmt.Fprintf(w, `{"error": "%s"}`, message)
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

	err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	fmt.Fprint(w, "Welcome to the the protected area")

}