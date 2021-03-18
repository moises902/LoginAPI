// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MinRecord min record
//
// swagger:model MinRecord
type MinRecord struct {

	// city
	City string `json:"city,omitempty"`

	// drug name
	DrugName string `json:"drugName,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// generic name
	GenericName string `json:"genericName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last org name
	LastOrgName string `json:"lastOrgName,omitempty"`

	// specialty
	Specialty string `json:"specialty,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this min record
func (m *MinRecord) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this min record based on context it is used
func (m *MinRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MinRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MinRecord) UnmarshalBinary(b []byte) error {
	var res MinRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
