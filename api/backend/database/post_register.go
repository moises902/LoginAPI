package database

import (
	"fmt"

	"github.com/t1cg/VL/api/backend/models"
	"golang.org/x/crypto/bcrypt"
)

//RegisterResponse takes in the parameters and returns an int64
// as the response code
func RegisterResponse(body *models.Register) (code int64) {

	db := ConnectWrite()
	defer db.Close()

	var flag bool
	row := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM Users WHERE email = ?)`, body.Email)
	if err := row.Scan(&flag); err != nil {
		code = 5
		return
	} else if flag { //If email exists error
		code = 1
		return
	} else if *body.Password1 != *body.Password2 { //If passwords do not match error
		code = 2
		return
	} else if (len(*body.Password1) < 6) || (len(*body.Password2) < 6) { //If password length is less than 6
		code = 3
		return 
	} else {

		pass := []byte(*body.Password2)
		hashPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
		if err != nil {
			code = 5
			return
		}

		results, err := db.Exec(`INSERT INTO Users (email, pass) VALUES (?,?)`, body.Email, hashPass)
		if err != nil {
			fmt.Println(err)
			code = 5
			return
		}
		rows, err := results.RowsAffected()
		if err != nil {
			code = 5
			return
		}
		if rows != 1 {
			code = 5
			return
		}

		code = 4
		return 
	}
}