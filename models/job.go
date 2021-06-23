// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Job A Job represents a scheduled task to execute.
//
// swagger:model job
type Job struct {

	// Concurrency policy for the job allow/forbid
	// Example: allow
	Concurrency string `json:"concurrency,omitempty"`

	// Array containing the jobs that depends on this one
	// Example: ["dependent_job"]
	// Read Only: true
	DependentJobs []string `json:"dependent_jobs"`

	// Disabled state of the job
	Disabled bool `json:"disabled,omitempty"`

	// Nice name for the job. Optional.
	Displayname string `json:"displayname,omitempty"`

	// Number of failed executions
	// Read Only: true
	ErrorCount int64 `json:"error_count,omitempty"`

	// Executor plugin used to run the job
	// Example: shell
	Executor string `json:"executor,omitempty"`

	// Executor plugin parameters
	// Example: {"command":"echo 'Hello from Dkron'"}
	ExecutorConfig map[string]string `json:"executor_config,omitempty"`

	// Last time this job failed
	// Read Only: true
	// Format: date-time
	LastError strfmt.DateTime `json:"last_error,omitempty"`

	// Last time this job executed successfully
	// Read Only: true
	// Format: date-time
	LastSuccess strfmt.DateTime `json:"last_success,omitempty"`

	// Extra metadata tags for this job
	// Example: {"office":"Barcelona"}
	Metadata map[string]string `json:"metadata,omitempty"`

	// Name for the job. Use only lower case letters (unicode), digits, underscore and dash.
	// Example: job1
	// Required: true
	Name *string `json:"name"`

	// Time of the next job execution
	// Read Only: true
	// Format: date-time
	Next strfmt.DateTime `json:"next,omitempty"`

	// Owner of the job
	// Example: Platform Team
	Owner string `json:"owner,omitempty"`

	// Email of the owner
	// Example: platform@example.com
	OwnerEmail string `json:"owner_email,omitempty"`

	// The name/id of the job that will trigger the execution of this job
	// Example: parent_job
	ParentJob string `json:"parent_job,omitempty"`

	// processors
	Processors Processors `json:"processors,omitempty"`

	// Number of times to retry a failed job execution
	// Example: 2
	Retries int64 `json:"retries,omitempty"`

	// Cron expression for the job.
	// Example: @every 10s
	// Required: true
	Schedule *string `json:"schedule"`

	// Status of the job
	// Example: success
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Number of successful executions
	// Read Only: true
	SuccessCount int64 `json:"success_count,omitempty"`

	// Target nodes tags of this job
	// Example: {"server":"true"}
	Tags map[string]string `json:"tags,omitempty"`

	// Timezone where the job will be executed. By default and when field is set to empty string, the job will run in local time.
	// Example: Europe/Berlin
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateLastError(formats strfmt.Registry) error {
	if swag.IsZero(m.LastError) { // not required
		return nil
	}

	if err := validate.FormatOf("last_error", "body", "date-time", m.LastError.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateLastSuccess(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSuccess) { // not required
		return nil
	}

	if err := validate.FormatOf("last_success", "body", "date-time", m.LastSuccess.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("next", "body", "date-time", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateProcessors(formats strfmt.Registry) error {
	if swag.IsZero(m.Processors) { // not required
		return nil
	}

	if m.Processors != nil {
		if err := m.Processors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processors")
			}
			return err
		}
	}

	return nil
}

func (m *Job) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this job based on the context it is used
func (m *Job) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDependentJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccessCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) contextValidateDependentJobs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dependent_jobs", "body", []string(m.DependentJobs)); err != nil {
		return err
	}

	return nil
}

func (m *Job) contextValidateErrorCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error_count", "body", int64(m.ErrorCount)); err != nil {
		return err
	}

	return nil
}

func (m *Job) contextValidateLastError(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_error", "body", strfmt.DateTime(m.LastError)); err != nil {
		return err
	}

	return nil
}

func (m *Job) contextValidateLastSuccess(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_success", "body", strfmt.DateTime(m.LastSuccess)); err != nil {
		return err
	}

	return nil
}

func (m *Job) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "next", "body", strfmt.DateTime(m.Next)); err != nil {
		return err
	}

	return nil
}

func (m *Job) contextValidateProcessors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Processors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("processors")
		}
		return err
	}

	return nil
}

func (m *Job) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *Job) contextValidateSuccessCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "success_count", "body", int64(m.SuccessCount)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Job) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Job) UnmarshalBinary(b []byte) error {
	var res Job
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}