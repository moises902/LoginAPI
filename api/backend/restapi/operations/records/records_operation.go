package records

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/t1cg/VL/api/backend/database"
	"github.com/t1cg/VL/api/backend/models"
)

//GetFunc gets the parameters and calls RecordsResponse with the parameters
//to get the response back
func GetFunc(city, state, specialty, drugName *string) (response middleware.Responder) {
	
	data, code := database.RecordsResponse(city, state, specialty, drugName)
	
	switch code {
	case 1:
		response = NewGetRecordsOK().WithPayload(data)
	case 2:
		message := models.ReturnCode{}
		message.Code = 500
		message.Message = "Internal Server Error"
		response = NewGetRecordsInternalServerError().WithPayload(&message)
	}

	return
}