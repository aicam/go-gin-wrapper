package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// UserID user ID
// swagger:model UserID
type UserID struct {

	// return code.
	Code int32 `json:"code,omitempty"`

	// New user ID.
	ID int32 `json:"id,omitempty"`

	// This is not used anymore.
	Message string `json:"message,omitempty"`
}

// Validate validates this user ID
func (m *UserID) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}