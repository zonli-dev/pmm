// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListKubernetesClustersReader is a Reader for the ListKubernetesClusters structure.
type ListKubernetesClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubernetesClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubernetesClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubernetesClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubernetesClustersOK creates a ListKubernetesClustersOK with default headers values
func NewListKubernetesClustersOK() *ListKubernetesClustersOK {
	return &ListKubernetesClustersOK{}
}

/*ListKubernetesClustersOK handles this case with default header values.

A successful response.
*/
type ListKubernetesClustersOK struct {
	Payload *ListKubernetesClustersOKBody
}

func (o *ListKubernetesClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/List][%d] listKubernetesClustersOk  %+v", 200, o.Payload)
}

func (o *ListKubernetesClustersOK) GetPayload() *ListKubernetesClustersOKBody {
	return o.Payload
}

func (o *ListKubernetesClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListKubernetesClustersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubernetesClustersDefault creates a ListKubernetesClustersDefault with default headers values
func NewListKubernetesClustersDefault(code int) *ListKubernetesClustersDefault {
	return &ListKubernetesClustersDefault{
		_statusCode: code,
	}
}

/*ListKubernetesClustersDefault handles this case with default header values.

An unexpected error response.
*/
type ListKubernetesClustersDefault struct {
	_statusCode int

	Payload *ListKubernetesClustersDefaultBody
}

// Code gets the status code for the list kubernetes clusters default response
func (o *ListKubernetesClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListKubernetesClustersDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/List][%d] ListKubernetesClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubernetesClustersDefault) GetPayload() *ListKubernetesClustersDefaultBody {
	return o.Payload
}

func (o *ListKubernetesClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListKubernetesClustersDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

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

/*KubernetesClustersItems0 Cluster contains public info about kubernetes cluster.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model KubernetesClustersItems0
*/
type KubernetesClustersItems0 struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// operators
	Operators *KubernetesClustersItems0Operators `json:"operators,omitempty"`
}

// Validate validates this kubernetes clusters items0
func (o *KubernetesClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KubernetesClustersItems0) validateOperators(formats strfmt.Registry) error {

	if swag.IsZero(o.Operators) { // not required
		return nil
	}

	if o.Operators != nil {
		if err := o.Operators.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesClustersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesClustersItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesClustersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KubernetesClustersItems0Operators Operators contains list of operators installed in kubernetes cluster.
swagger:model KubernetesClustersItems0Operators
*/
type KubernetesClustersItems0Operators struct {

	// psmdb
	PSMDB *KubernetesClustersItems0OperatorsPSMDB `json:"psmdb,omitempty"`

	// pxc
	Pxc *KubernetesClustersItems0OperatorsPxc `json:"pxc,omitempty"`
}

// Validate validates this kubernetes clusters items0 operators
func (o *KubernetesClustersItems0Operators) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePSMDB(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KubernetesClustersItems0Operators) validatePSMDB(formats strfmt.Registry) error {

	if swag.IsZero(o.PSMDB) { // not required
		return nil
	}

	if o.PSMDB != nil {
		if err := o.PSMDB.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "psmdb")
			}
			return err
		}
	}

	return nil
}

func (o *KubernetesClustersItems0Operators) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(o.Pxc) { // not required
		return nil
	}

	if o.Pxc != nil {
		if err := o.Pxc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesClustersItems0Operators) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesClustersItems0Operators) UnmarshalBinary(b []byte) error {
	var res KubernetesClustersItems0Operators
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KubernetesClustersItems0OperatorsPSMDB Operator contains all information about operator installed in kubernetes cluster.
swagger:model KubernetesClustersItems0OperatorsPSMDB
*/
type KubernetesClustersItems0OperatorsPSMDB struct {

	// OperatorsStatus defines status of operators installed in kubernetes cluster.
	//
	//  - OPERATORS_STATUS_INVALID: OPERATORS_STATUS_INVALID represents unknown state.
	//  - OPERATORS_STATUS_OK: OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	//  - OPERATORS_STATUS_UNSUPPORTED: OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	//  - OPERATORS_STATUS_UNAVAILABLE: OPERATORS_STATUS_UNAVAILABLE represents that kubernetes cluster is unavailable.
	// Enum: [OPERATORS_STATUS_INVALID OPERATORS_STATUS_OK OPERATORS_STATUS_UNSUPPORTED OPERATORS_STATUS_UNAVAILABLE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this kubernetes clusters items0 operators PSMDB
func (o *KubernetesClustersItems0OperatorsPSMDB) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPERATORS_STATUS_INVALID","OPERATORS_STATUS_OK","OPERATORS_STATUS_UNSUPPORTED","OPERATORS_STATUS_UNAVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum = append(kubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum, v)
	}
}

const (

	// KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSINVALID captures enum value "OPERATORS_STATUS_INVALID"
	KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSINVALID string = "OPERATORS_STATUS_INVALID"

	// KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSOK captures enum value "OPERATORS_STATUS_OK"
	KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSOK string = "OPERATORS_STATUS_OK"

	// KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSUNSUPPORTED captures enum value "OPERATORS_STATUS_UNSUPPORTED"
	KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSUNSUPPORTED string = "OPERATORS_STATUS_UNSUPPORTED"

	// KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSUNAVAILABLE captures enum value "OPERATORS_STATUS_UNAVAILABLE"
	KubernetesClustersItems0OperatorsPSMDBStatusOPERATORSSTATUSUNAVAILABLE string = "OPERATORS_STATUS_UNAVAILABLE"
)

// prop value enum
func (o *KubernetesClustersItems0OperatorsPSMDB) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kubernetesClustersItems0OperatorsPsmdbTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *KubernetesClustersItems0OperatorsPSMDB) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("operators"+"."+"psmdb"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesClustersItems0OperatorsPSMDB) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesClustersItems0OperatorsPSMDB) UnmarshalBinary(b []byte) error {
	var res KubernetesClustersItems0OperatorsPSMDB
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KubernetesClustersItems0OperatorsPxc Operator contains all information about operator installed in kubernetes cluster.
swagger:model KubernetesClustersItems0OperatorsPxc
*/
type KubernetesClustersItems0OperatorsPxc struct {

	// OperatorsStatus defines status of operators installed in kubernetes cluster.
	//
	//  - OPERATORS_STATUS_INVALID: OPERATORS_STATUS_INVALID represents unknown state.
	//  - OPERATORS_STATUS_OK: OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	//  - OPERATORS_STATUS_UNSUPPORTED: OPERATORS_STATUS_UNSUPPORTED represents that operators are installed, but doesn't have supported API version.
	//  - OPERATORS_STATUS_UNAVAILABLE: OPERATORS_STATUS_UNAVAILABLE represents that kubernetes cluster is unavailable.
	// Enum: [OPERATORS_STATUS_INVALID OPERATORS_STATUS_OK OPERATORS_STATUS_UNSUPPORTED OPERATORS_STATUS_UNAVAILABLE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this kubernetes clusters items0 operators pxc
func (o *KubernetesClustersItems0OperatorsPxc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kubernetesClustersItems0OperatorsPxcTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPERATORS_STATUS_INVALID","OPERATORS_STATUS_OK","OPERATORS_STATUS_UNSUPPORTED","OPERATORS_STATUS_UNAVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubernetesClustersItems0OperatorsPxcTypeStatusPropEnum = append(kubernetesClustersItems0OperatorsPxcTypeStatusPropEnum, v)
	}
}

const (

	// KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSINVALID captures enum value "OPERATORS_STATUS_INVALID"
	KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSINVALID string = "OPERATORS_STATUS_INVALID"

	// KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSOK captures enum value "OPERATORS_STATUS_OK"
	KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSOK string = "OPERATORS_STATUS_OK"

	// KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSUNSUPPORTED captures enum value "OPERATORS_STATUS_UNSUPPORTED"
	KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSUNSUPPORTED string = "OPERATORS_STATUS_UNSUPPORTED"

	// KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSUNAVAILABLE captures enum value "OPERATORS_STATUS_UNAVAILABLE"
	KubernetesClustersItems0OperatorsPxcStatusOPERATORSSTATUSUNAVAILABLE string = "OPERATORS_STATUS_UNAVAILABLE"
)

// prop value enum
func (o *KubernetesClustersItems0OperatorsPxc) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kubernetesClustersItems0OperatorsPxcTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *KubernetesClustersItems0OperatorsPxc) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("operators"+"."+"pxc"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesClustersItems0OperatorsPxc) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesClustersItems0OperatorsPxc) UnmarshalBinary(b []byte) error {
	var res KubernetesClustersItems0OperatorsPxc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListKubernetesClustersDefaultBody list kubernetes clusters default body
swagger:model ListKubernetesClustersDefaultBody
*/
type ListKubernetesClustersDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list kubernetes clusters default body
func (o *ListKubernetesClustersDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListKubernetesClusters default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListKubernetesClustersOKBody list kubernetes clusters OK body
swagger:model ListKubernetesClustersOKBody
*/
type ListKubernetesClustersOKBody struct {

	// Kubernetes clusters.
	KubernetesClusters []*KubernetesClustersItems0 `json:"kubernetes_clusters"`
}

// Validate validates this list kubernetes clusters OK body
func (o *ListKubernetesClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubernetesClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersOKBody) validateKubernetesClusters(formats strfmt.Registry) error {

	if swag.IsZero(o.KubernetesClusters) { // not required
		return nil
	}

	for i := 0; i < len(o.KubernetesClusters); i++ {
		if swag.IsZero(o.KubernetesClusters[i]) { // not required
			continue
		}

		if o.KubernetesClusters[i] != nil {
			if err := o.KubernetesClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listKubernetesClustersOk" + "." + "kubernetes_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBody) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
