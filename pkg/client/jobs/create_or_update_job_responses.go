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

// CreateOrUpdateJobReader is a Reader for the CreateOrUpdateJob structure.
type CreateOrUpdateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrUpdateJobCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrUpdateJobCreated creates a CreateOrUpdateJobCreated with default headers values
func NewCreateOrUpdateJobCreated() *CreateOrUpdateJobCreated {
	return &CreateOrUpdateJobCreated{}
}

/* CreateOrUpdateJobCreated describes a response with status code 201, with default header values.

Successful response
*/
type CreateOrUpdateJobCreated struct {
	Payload *models.Job
}

func (o *CreateOrUpdateJobCreated) Error() string {
	return fmt.Sprintf("[POST /jobs][%d] createOrUpdateJobCreated  %+v", 201, o.Payload)
}
func (o *CreateOrUpdateJobCreated) GetPayload() *models.Job {
	return o.Payload
}

func (o *CreateOrUpdateJobCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
