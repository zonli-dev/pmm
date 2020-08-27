// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteBackupReader is a Reader for the DeleteBackup structure.
type DeleteBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBackupOK creates a DeleteBackupOK with default headers values
func NewDeleteBackupOK() *DeleteBackupOK {
	return &DeleteBackupOK{}
}

/*DeleteBackupOK handles this case with default header values.

A successful response.
*/
type DeleteBackupOK struct {
	Payload interface{}
}

func (o *DeleteBackupOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Backups/Delete][%d] deleteBackupOk  %+v", 200, o.Payload)
}

func (o *DeleteBackupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBackupDefault creates a DeleteBackupDefault with default headers values
func NewDeleteBackupDefault(code int) *DeleteBackupDefault {
	return &DeleteBackupDefault{
		_statusCode: code,
	}
}

/*DeleteBackupDefault handles this case with default header values.

An unexpected error response
*/
type DeleteBackupDefault struct {
	_statusCode int

	Payload *DeleteBackupDefaultBody
}

// Code gets the status code for the delete backup default response
func (o *DeleteBackupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBackupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Backups/Delete][%d] DeleteBackup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBackupDefault) GetPayload() *DeleteBackupDefaultBody {
	return o.Payload
}

func (o *DeleteBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteBackupDefaultBody delete backup default body
swagger:model DeleteBackupDefaultBody
*/
type DeleteBackupDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this delete backup default body
func (o *DeleteBackupDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteBackupDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeleteBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBackupDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBackupDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteBackupDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
