import React from 'react';

const SingleRecordFormat = (props) => {

    const data = props.data

    return (
        <div>
            <p>
                ID: {data.id} <br/>
                BeneCount: {data.beneCount} <br/>
                BeneCountGe65: {data.beneCountGe65} <br/>
                BeneCountGe65Flag: {data.beneCountGe65Flag} <br/>
                BeneCountNum: {data.beneCountNum} <br/>
                City: {data.city} <br/>
                Description: {data.description} <br/>
                Drug Name: {data.drugName} <br/>
                First Name: {data.firstName} <br/>
                Ge65SuppressFlag: {data.ge65SuppressFlag} <br/>
                Generic Name: {data.genericName} <br/>
                Last Org. Name: {data.lastOrgName} <br/>
                NPI: {data.npi} <br/>
                Specialty: {data.specialty} <br/>
                State: {data.state} <br/>
                Total 30 Day Fill COunt: {data.total30DayFillCount} <br/>
                Total 30 Day Fill Count Ge65: {data.total30DayFillCountGe65} <br/>
                Total Claim Count: {data.totalClaimCount} <br/>
                Total Claim Count Ge65: {data.totalClaimCountGe65} <br/>
                Total Claim Count Num: {data.totalClaimCountNum} <br/>
                Total Day Supply: {data.totalDaySupply} <br/>
                Total Day Supply Ge65: {data.totalDaySupplyGe65} <br/>
                Total Drug Cost: {data.totalDrugCost} <br/>
                Total Drug Cost Num: {data.totalDrugCostNum} <br/>
                Total Drug Cost Ge65: {data.totalDaySupplyGe65} <br/>
            </p>
        </div>
    )
    
}

export default SingleRecordFormat;

/*
{
  "beneCount": "string",
  "beneCountGe65": "string",
  "beneCountGe65Flag": "string",
  "beneCountNum": 0,
  "city": "string",
  "description": "string",
  "drugName": "string",
  "firstName": "string",
  "ge65SuppressFlag": "string",
  "genericName": "string",
  "id": 0,
  "lastOrgName": "string",
  "npi": "string",
  "specialty": "string",
  "state": "string",
  "total30DayFillCount": "string",
  "total30DayFillCountGe65": "string",
  "totalClaimCount": "string",
  "totalClaimCountGe65": "string",
  "totalClaimCountNum": 0,
  "totalDaySupply": "string",
  "totalDaySupplyGe65": "string",
  "totalDrugCost": "string",
  "totalDrugCostNum": 0,
  "totalDrugcostGe65": "string"
}
*/