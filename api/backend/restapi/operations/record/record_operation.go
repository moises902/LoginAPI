package record

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/t1cg/VL/api/backend/database"
	"github.com/t1cg/VL/api/backend/models"
)

//GetFunc formats the response for the endpoint
func GetFunc(id int64) (response middleware.Responder) {

	data, code := database.RecordResponse(id)
	
	switch code {
	case 1:
		response = NewGetRecordOK().WithPayload(&data)
	case 2:
		message := models.ReturnCode{}
		message.Code = 500
		message.Message = "Internal Server Error"
		response = NewGetRecordInternalServerError().WithPayload(&message)
	}
	
	return 
}
 