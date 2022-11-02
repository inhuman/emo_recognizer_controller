// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClearSoundResult clear sound result
//
// swagger:model ClearSoundResult
type ClearSoundResult struct {

	// UUID
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this clear sound result
func (m *ClearSoundResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clear sound result based on context it is used
func (m *ClearSoundResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClearSoundResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClearSoundResult) UnmarshalBinary(b []byte) error {
	var res ClearSoundResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
