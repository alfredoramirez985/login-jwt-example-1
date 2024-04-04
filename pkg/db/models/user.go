package models

import "github.com/go-pg/pg/v10"

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     int32  `json:"email"`
	LoginData *LoginData `pg:"rel:has-one" json:"login_data"`
}

func CreateUser(db *pg.DB, req *User) (bool, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return false, err
	}

	user := &User{}

	err = db.Model(user).
        Relation("LoginData").
        Where("user.id = ?", req.ID).
        Select()
	
	return err != nil, err
}
