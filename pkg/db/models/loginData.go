package models

type LoginData struct {
	ID					string	`json:"id"`
	UserName			string	`json:"user_name"`
	Password 			string	`json:"password"`
	OldPassword 		string	`json:"old_password"`
	WrongLoginAttempt	int32	`json:"wrong_login_attempt"`
	TodayLoginAttempt	int32	`json:"today_login_attempt"`
	IsNowLogin			bool	`json:"is_now_login"`
}