// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cat cat
//
// swagger:model Cat
type Cat struct {

	// name
	// Example: my-cat
	Name string `json:"name,omitempty"`

	// path
	// Example: my-cat.png
	Path string `json:"path,omitempty"`

	// url
	// Example: https://raw.githubusercontent.com/Pandahoro/cats/main/arrow-cat.png
	URL string `json:"url,omitempty"`
}

// Validate validates this cat
func (m *Cat) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cat based on context it is used
func (m *Cat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cat) UnmarshalBinary(b []byte) error {
	var res Cat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}