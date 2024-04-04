package models

type User struct {
	ID			string	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Phone 		string	`json:"phone"`
	Email		int32	`json:"email"`
}