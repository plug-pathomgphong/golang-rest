package models

import "github.com/plug-pathomgphong/golang-rest/db"

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Salt     string
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password, salt) VALUES(?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password, u.Salt)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.Id = userId
	return err
}
