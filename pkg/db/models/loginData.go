package models

import "github.com/go-pg/pg/v10"

type LoginData struct {
	ID					string	`json:"id"`
	UserName			string	`json:"user_name"`
	Password 			string	`json:"password"`
	OldPassword 		string	`json:"old_password"`
	WrongLoginAttempt	int32	`json:"wrong_login_attempt"`
	TodayLoginAttempt	int32	`json:"today_login_attempt"`
	IsNowLogin			bool	`json:"is_now_login"`
}

func GetLoginData(db *pg.DB, username string) (*LoginData, error) {
    loginData := &LoginData{}

    err := db.Model(loginData).
        Relation("User").
        Where("loginData.user_name = ?", username).
        Select()

    return loginData, err
}