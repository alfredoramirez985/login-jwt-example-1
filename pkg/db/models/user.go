package models

import "github.com/go-pg/pg/v10"

type User struct {
	ID        string 	`json:"id"`
	FirstName string 	`json:"first_name"`
	LastName  string 	`json:"last_name"`
	Phone     string 	`json:"phone"`
	Email     string  	`json:"email"`
	IDLoginData string 	`json:"id_login_data"`
	LoginData *LoginData `pg:"rel:has-one" json:"login_data"`
}

func CreateUser(db *pg.DB, req *User) (bool, error) {
	// Check if LoginData already exists
    existingLoginData, err := GetLoginData(db, req.Email)
    if err != nil && err != pg.ErrNoRows {
        return false, err
    }

    // If LoginData doesn't exist, create it
    if existingLoginData == nil {
        reqLoginData := &LoginData{
            UserName: req.Email,
            Password: req.LoginData.Password,
            OldPassword: req.LoginData.Password,
        }

        _, err := db.Model(reqLoginData).Insert()
        if err != nil {
            return false, err
        }
        req.IDLoginData = reqLoginData.ID
    } else {
        req.IDLoginData = existingLoginData.ID
    }

    reqUser := &User{
        ID: req.ID,
        FirstName: req.FirstName,
        LastName: req.LastName,
        Phone: req.Phone,
        Email: req.Email,
        IDLoginData: req.IDLoginData,
    }

    _, err = db.Model(reqUser).Insert()
    if err != nil {
        return false, err
    }

    return true, nil
}
