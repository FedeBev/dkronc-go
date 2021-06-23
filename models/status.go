// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Status Status represents details about the node.
//
// swagger:model status
type Status struct {

	// Node basic details
	// Read Only: true
	Agent map[string]interface{} `json:"agent,omitempty"`

	// Serf status
	// Read Only: true
	Serf map[string]interface{} `json:"serf,omitempty"`

	// Tags asociated with this node
	// Read Only: true
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this status based on the context it is used
func (m *Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Status) contextValidateSerf(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Status) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}