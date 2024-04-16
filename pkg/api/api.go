package api

import (
	"encoding/json"
	"log"
	"login-jwt-example/pkg/api/login"
	"login-jwt-example/pkg/db/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB) *chi.Mux {
    //get the router
    r := chi.NewRouter()
    //add middleware
    //in this case we will store our DB to use it later
    r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

    //routes for our service
    r.Route("/user", func(r chi.Router) {
        r.Post("/", CreateUser)
        //r.Get("/", getComments)
    })

    r.Route("/createtoken", func(r chi.Router) {
        r.Post("/", api.LoginHandler)
        r.Get("/", api.ProtectedHandler)
    })

    //test route to make sure everything works
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("up and running"))
    })

    return r
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string  `json:"email"`
    LoginData *CreateLoginData `json:"login_data"`
}

type CreateLoginData struct {
	UserName	string	`json:"user_name"`
	Password 	string	`json:"password"`
}

type UserResponse struct {
    Success bool            `json:"success"`
    Error   string          `json:"error"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    //get the request body and decode it
    req := &CreateUserRequest{}
    err := json.NewDecoder(r.Body).Decode(req)

    //if there's an error with decoding the information
    //send a response with an error
    if err != nil {
        res := &UserResponse{
            Success: false,
            Error:   err.Error(),
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
    //if we can't get the db let's handle the error
    //and send an adequate response
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
    
    hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.LoginData.Password), bcrypt.DefaultCost)
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

    //if we can get the db then
    loginData := &models.LoginData{
        ID: uuid.NewString(),
        UserName: req.LoginData.UserName,
        Password: string(hashedPass),
        OldPassword: string(hashedPass),
    }
    success, err := models.CreateUser(pgdb, &models.User{
        ID: uuid.NewString(),
        FirstName: req.FirstName,
        LastName: req.LastName,
        Phone: req.Phone,
        Email: req.Email,
        LoginData: loginData,
    })
    if err != nil {
        res := &UserResponse{
            Success: false,
            Error:   err.Error(),
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
    //everything is good
    //let's return a positive response
    res := &UserResponse{
        Success: success,
        Error:   "",
    }
    err = json.NewEncoder(w).Encode(res)
    if err != nil {
        log.Printf("error encoding after creating comment %v\n", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
}