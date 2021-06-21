// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListExecutionsByJobParams creates a new ListExecutionsByJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListExecutionsByJobParams() *ListExecutionsByJobParams {
	return &ListExecutionsByJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListExecutionsByJobParamsWithTimeout creates a new ListExecutionsByJobParams object
// with the ability to set a timeout on a request.
func NewListExecutionsByJobParamsWithTimeout(timeout time.Duration) *ListExecutionsByJobParams {
	return &ListExecutionsByJobParams{
		timeout: timeout,
	}
}

// NewListExecutionsByJobParamsWithContext creates a new ListExecutionsByJobParams object
// with the ability to set a context for a request.
func NewListExecutionsByJobParamsWithContext(ctx context.Context) *ListExecutionsByJobParams {
	return &ListExecutionsByJobParams{
		Context: ctx,
	}
}

// NewListExecutionsByJobParamsWithHTTPClient creates a new ListExecutionsByJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewListExecutionsByJobParamsWithHTTPClient(client *http.Client) *ListExecutionsByJobParams {
	return &ListExecutionsByJobParams{
		HTTPClient: client,
	}
}

/* ListExecutionsByJobParams contains all the parameters to send to the API endpoint
   for the list executions by job operation.

   Typically these are written to a http.Request.
*/
type ListExecutionsByJobParams struct {

	/* JobName.

	   The job that owns the executions to be fetched.
	*/
	JobName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list executions by job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListExecutionsByJobParams) WithDefaults() *ListExecutionsByJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list executions by job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListExecutionsByJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list executions by job params
func (o *ListExecutionsByJobParams) WithTimeout(timeout time.Duration) *ListExecutionsByJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list executions by job params
func (o *ListExecutionsByJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list executions by job params
func (o *ListExecutionsByJobParams) WithContext(ctx context.Context) *ListExecutionsByJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list executions by job params
func (o *ListExecutionsByJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list executions by job params
func (o *ListExecutionsByJobParams) WithHTTPClient(client *http.Client) *ListExecutionsByJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list executions by job params
func (o *ListExecutionsByJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobName adds the jobName to the list executions by job params
func (o *ListExecutionsByJobParams) WithJobName(jobName string) *ListExecutionsByJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the list executions by job params
func (o *ListExecutionsByJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WriteToRequest writes these params to a swagger request
func (o *ListExecutionsByJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param job_name
	if err := r.SetPathParam("job_name", o.JobName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
