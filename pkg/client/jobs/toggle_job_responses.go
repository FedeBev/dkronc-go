// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/FedeBev/dkronc-go/pkg/models"
)

// ToggleJobReader is a Reader for the ToggleJob structure.
type ToggleJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ToggleJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewToggleJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewToggleJobOK creates a ToggleJobOK with default headers values
func NewToggleJobOK() *ToggleJobOK {
	return &ToggleJobOK{}
}

/* ToggleJobOK describes a response with status code 200, with default header values.

Successful response
*/
type ToggleJobOK struct {
	Payload *models.Job
}

func (o *ToggleJobOK) Error() string {
	return fmt.Sprintf("[POST /jobs/{job_name}/toggle][%d] toggleJobOK  %+v", 200, o.Payload)
}
func (o *ToggleJobOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *ToggleJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
