// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecognizeResult recognize result
//
// swagger:model RecognizeResult
type RecognizeResult struct {

	// text
	Text string `json:"text,omitempty"`

	// UUID
	UUID string `json:"UUID,omitempty"`

	// words
	Words []*WordResult `json:"result"`
}

// Validate validates this recognize result
func (m *RecognizeResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecognizeResult) validateWords(formats strfmt.Registry) error {
	if swag.IsZero(m.Words) { // not required
		return nil
	}

	for i := 0; i < len(m.Words); i++ {
		if swag.IsZero(m.Words[i]) { // not required
			continue
		}

		if m.Words[i] != nil {
			if err := m.Words[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this recognize result based on the context it is used
func (m *RecognizeResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecognizeResult) contextValidateWords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Words); i++ {

		if m.Words[i] != nil {
			if err := m.Words[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecognizeResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecognizeResult) UnmarshalBinary(b []byte) error {
	var res RecognizeResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}