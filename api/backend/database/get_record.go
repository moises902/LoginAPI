package database

import (
	"github.com/t1cg/VL/api/backend/models"
)

//RecordResponse returns data from database
func RecordResponse(id int64) (models.Record, int64) {

	var code int64
	responseData := models.Record{}
	db := Connect()
	defer db.Close()
	results, err := db.Query(`SELECT beneCount, beneCountGe65, beneCountGe65Flag, beneCountNum, 
	city, description, drugName, firstName, ge65SuppressFlag, genericName,
	autoID, lastOrgName, npi, specialty, state, total30DayFillCount,
	total30DayFillCountGe65, totalClaimCount, totalClaimCountGe65, totalClaimCountNum,
	totalDaySupply, totalDaySupplyGe65, totalDrugCost, totalDrugCostNum, totalDrugcostGe65
	FROM PartD WHERE autoID = ?`, id)
	
	defer results.Close()
	
	if err != nil {
		code = 2
		return responseData, code
	}
		
	for results.Next() {
		err := results.Scan(&responseData.BeneCount, &responseData.BeneCountGe65, &responseData.BeneCountGe65Flag, &responseData.BeneCountNum,
		&responseData.City, &responseData.Description, &responseData.DrugName, &responseData.FirstName, &responseData.Ge65SuppressFlag,
		&responseData.GenericName, &responseData.ID, &responseData.LastOrgName, &responseData.Npi, &responseData.Specialty,
		&responseData.State, &responseData.Total30DayFillCount, &responseData.Total30DayFillCountGe65, &responseData.TotalClaimCount,
		&responseData.TotalClaimCountGe65, &responseData.TotalClaimCountNum, &responseData.TotalDaySupply, &responseData.TotalDaySupplyGe65,
		&responseData.TotalDrugCost, &responseData.TotalDrugCostNum, &responseData.TotalDrugcostGe65)
			
		if err != nil {
			code = 2
			return responseData, code
		}
	}
	
	code = 1
	return responseData, code
}