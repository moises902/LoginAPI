package login

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/t1cg/VL/api/backend/database"
	"github.com/t1cg/VL/api/backend/models"
)

//PostFunc will get the params and insert it into LoginResponse
//returned value from LoginResponse will determine payload
func PostFunc(body *models.Login) (response middleware.Responder) {

	code := database.LoginResponse(body)
	message := models.User{}

	token, err := GenerateJWT(*body.Email)
	if err != nil {
		code = 3
	}

	fmt.Println(token)
	payloadToken, cookieToken := SplitToken(token)

	switch code {
	case 1:
		message.Email = *body.Email
		message.Token = payloadToken
		response = NewCustomResponder(NewPostLoginOK().WithPayload(&message), cookieToken)
	case 2:
		message := models.ReturnCode{}
		message.Code = 400
		message.Message = "Incorrect password for email"
		response = NewPostLoginBadRequest().WithPayload(&message)
	case 3:
		message := models.ReturnCode{}
		message.Code = 400
		message.Message = "Email does not exists in the system"
		response = NewPostLoginInternalServerError().WithPayload(&message)
	}

	return
}

//GenerateJWT will create the full JWT token
func GenerateJWT(email string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err = token.SignedString([]byte("some_secret_key_val_123123"))
	if err != nil {
		tokenString = ""
		fmt.Println(err)
		return
	}

	return
}

//SplitToken takes in a token and splits it in half to be a payload
//or a cookie
func SplitToken(token string) (payloadToken, cookieToken string) {
	payloadToken = token[:len((token))/2]
	cookieToken = token[len((token))/2:]

	return
}

//CustomResponder struct that includes a cookie
type CustomResponder struct {
	responder   middleware.Responder
	cookieToken string
}

//NewCustomResponder formats the new responder for cookie
func NewCustomResponder(response middleware.Responder, token string) *CustomResponder {

	return &CustomResponder{
		responder:   response,
		cookieToken: token,
	}
}

//WriteResponse appends a cookie to the responder
func (payload *CustomResponder) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
	cookie := http.Cookie{
		Name: "cookie", 
		Value: payload.cookieToken, 
		SameSite: http.SameSiteStrictMode}
	http.SetCookie(rw, &cookie)
	payload.responder.WriteResponse(rw, p)
}
