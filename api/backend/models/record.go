// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Record record
//
// swagger:model Record
type Record struct {

	// bene count
	BeneCount string `json:"beneCount,omitempty"`

	// bene count ge65
	BeneCountGe65 string `json:"beneCountGe65,omitempty"`

	// bene count ge65 flag
	BeneCountGe65Flag string `json:"beneCountGe65Flag,omitempty"`

	// bene count num
	BeneCountNum int64 `json:"beneCountNum,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// drug name
	DrugName string `json:"drugName,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// ge65 suppress flag
	Ge65SuppressFlag string `json:"ge65SuppressFlag,omitempty"`

	// generic name
	GenericName string `json:"genericName,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// last org name
	LastOrgName string `json:"lastOrgName,omitempty"`

	// npi
	Npi string `json:"npi,omitempty"`

	// specialty
	Specialty string `json:"specialty,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// total30 day fill count
	Total30DayFillCount string `json:"total30DayFillCount,omitempty"`

	// total30 day fill count ge65
	Total30DayFillCountGe65 string `json:"total30DayFillCountGe65,omitempty"`

	// total claim count
	TotalClaimCount string `json:"totalClaimCount,omitempty"`

	// total claim count ge65
	TotalClaimCountGe65 string `json:"totalClaimCountGe65,omitempty"`

	// total claim count num
	TotalClaimCountNum string `json:"totalClaimCountNum,omitempty"`

	// total day supply
	TotalDaySupply string `json:"totalDaySupply,omitempty"`

	// total day supply ge65
	TotalDaySupplyGe65 string `json:"totalDaySupplyGe65,omitempty"`

	// total drug cost
	TotalDrugCost string `json:"totalDrugCost,omitempty"`

	// total drug cost num
	TotalDrugCostNum float64 `json:"totalDrugCostNum,omitempty"`

	// total drugcost ge65
	TotalDrugcostGe65 string `json:"totalDrugcostGe65,omitempty"`
}

// Validate validates this record
func (m *Record) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this record based on context it is used
func (m *Record) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Record) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Record) UnmarshalBinary(b []byte) error {
	var res Record
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
