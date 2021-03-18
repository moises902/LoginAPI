package database

import (
	"fmt"

	"github.com/t1cg/VL/api/backend/models"
)

//RecordsResponse queries the database and returns the records model
func RecordsResponse(city, state, specialty, drugName *string) ( models.Records, int64) {

	var code int64
	var responseData []*models.MinRecord
	db := Connect()
	defer db.Close()

	query := `SELECT city, drugName, firstName, genericName, autoID,
	lastOrgName, specialty, state FROM PartD`
	
	// Applies WHERE on sql query if there are parameters
	if (state != nil) || (specialty != nil) || (drugName != nil) || (city != nil) {
		query += ` WHERE`
		
		//Concats to query string depending on params
		if city != nil {
			var tempStr string
			if (state != nil) || (specialty != nil) || (drugName != nil) {
				tempStr = fmt.Sprintf(` city = '%s' AND`, *city)
			} else {
				tempStr = fmt.Sprintf(` city = '%s'`, *city)
			}
			query += tempStr
		}
		if state != nil {
			var tempStr string
			if (specialty != nil) || (drugName != nil) {
				tempStr = fmt.Sprintf(` state = '%s' AND`, *state)
			} else {
				tempStr = fmt.Sprintf(` state = '%s'`, *state)
			}
			query += tempStr
		}
		if specialty != nil {
			var tempStr string
			if drugName != nil {
				tempStr = fmt.Sprintf(` specialty = '%s' AND`, *specialty)
			} else {
				tempStr = fmt.Sprintf(` specialty = '%s'`, *specialty)
			}
			query += tempStr
		}
		if drugName != nil {
			tempStr := fmt.Sprintf(` drugName = '%s'`, *drugName)
			query += tempStr
		}
	}
	
	//Only show up to 10 results
	query += " LIMIT 10"

	results, err := db.Query(query)
	defer results.Close()

	if err != nil {
		code = 2
		return nil, code
	} 
	
	for results.Next() {
		data := models.MinRecord{}
		err := results.Scan(&data.City, &data.DrugName, &data.FirstName, &data.GenericName,
		&data.ID, &data.LastOrgName, &data.Specialty, &data.State)

		if err != nil {
			code = 2
			return nil, code
		} 
		
		responseData = append(responseData, &data)
	}
	

	code = 1
	return responseData, code
}
