package api

import (
    "encoding/json"
    "login-jwt-example/pkg/db/models"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-pg/pg/v10"
)

//start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB) *chi.Mux {
    //get the router
    r := chi.NewRouter()
    //add middleware
    //in this case we will store our DB to use it later
    r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

    //routes for our service
    r.Route("/comments", func(r chi.Router) {
        r.Post("/", CreateUser)
        //r.Get("/", getComments)
    })

    //test route to make sure everything works
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("up and running"))
    })

    return r
}

type CreateUserRequest struct {
    ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     int32  `json:"email"`
}

type LoginData struct {
	
}

type UserResponse struct {
    Success bool            `json:"success"`
    Error   string          `json:"error"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}