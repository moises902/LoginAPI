package register

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/t1cg/VL/api/backend/database"
	"github.com/t1cg/VL/api/backend/models"
)

//PostFunc returns the response depending on the result from RegisterResponse
func PostFunc(body *models.Register) (response middleware.Responder) {

	code := database.RegisterResponse(body)
	message := models.ReturnCode{}

	switch code {
	case 1:
		message.Code = 400
		message.Message = "Email Already Exists"
		response = NewPostRegisterBadRequest().WithPayload(&message)
	case 2:
		message.Code = 400
		message.Message = "Passwords Are Not The Same"
		response = NewPostRegisterBadRequest().WithPayload(&message)

	case 3:
		message.Code = 400
		message.Message = "Password Length Is Less Than 6"
		response = NewPostRegisterBadRequest().WithPayload(&message)

	case 4:
		message.Code = 200
		message.Message = "Registration Success"
		response = NewPostRegisterOK().WithPayload(&message)
	case 5:
		message.Code = 500
		message.Message = "Internal Server Error"
		response = NewPostRegisterInternalServerError().WithPayload(&message)
	} 

	return  
}
