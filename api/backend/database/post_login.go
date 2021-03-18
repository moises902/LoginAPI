package database

import (
	"fmt"

	"github.com/t1cg/VL/api/backend/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginResponse takes in the parameters and authenticates credentials
// Returns int for return code
func LoginResponse(body *models.Login) (code int64) {

	db := Connect()
	defer db.Close()

	var pass string
	result := db.QueryRow(`SELECT pass FROM Users WHERE email = ?`, *body.Email)

	if err := result.Scan(&pass); err != nil {
		fmt.Println(err)
		code = 3
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(*body.Password))
	if err != nil {
		fmt.Println(err)
		code = 2
		return
	}

	code = 1
	return
}
