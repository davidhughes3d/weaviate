//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsPropertiesUpdateOKCode is the HTTP code returned for type SchemaActionsPropertiesUpdateOK
const SchemaActionsPropertiesUpdateOKCode int = 200

/*SchemaActionsPropertiesUpdateOK Changes applied.

swagger:response schemaActionsPropertiesUpdateOK
*/
type SchemaActionsPropertiesUpdateOK struct {
}

// NewSchemaActionsPropertiesUpdateOK creates SchemaActionsPropertiesUpdateOK with default headers values
func NewSchemaActionsPropertiesUpdateOK() *SchemaActionsPropertiesUpdateOK {

	return &SchemaActionsPropertiesUpdateOK{}
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SchemaActionsPropertiesUpdateUnauthorizedCode is the HTTP code returned for type SchemaActionsPropertiesUpdateUnauthorized
const SchemaActionsPropertiesUpdateUnauthorizedCode int = 401

/*SchemaActionsPropertiesUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response schemaActionsPropertiesUpdateUnauthorized
*/
type SchemaActionsPropertiesUpdateUnauthorized struct {
}

// NewSchemaActionsPropertiesUpdateUnauthorized creates SchemaActionsPropertiesUpdateUnauthorized with default headers values
func NewSchemaActionsPropertiesUpdateUnauthorized() *SchemaActionsPropertiesUpdateUnauthorized {

	return &SchemaActionsPropertiesUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaActionsPropertiesUpdateForbiddenCode is the HTTP code returned for type SchemaActionsPropertiesUpdateForbidden
const SchemaActionsPropertiesUpdateForbiddenCode int = 403

/*SchemaActionsPropertiesUpdateForbidden Forbidden

swagger:response schemaActionsPropertiesUpdateForbidden
*/
type SchemaActionsPropertiesUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesUpdateForbidden creates SchemaActionsPropertiesUpdateForbidden with default headers values
func NewSchemaActionsPropertiesUpdateForbidden() *SchemaActionsPropertiesUpdateForbidden {

	return &SchemaActionsPropertiesUpdateForbidden{}
}

// WithPayload adds the payload to the schema actions properties update forbidden response
func (o *SchemaActionsPropertiesUpdateForbidden) WithPayload(payload *models.ErrorResponse) *SchemaActionsPropertiesUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties update forbidden response
func (o *SchemaActionsPropertiesUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsPropertiesUpdateUnprocessableEntityCode is the HTTP code returned for type SchemaActionsPropertiesUpdateUnprocessableEntity
const SchemaActionsPropertiesUpdateUnprocessableEntityCode int = 422

/*SchemaActionsPropertiesUpdateUnprocessableEntity Invalid update.

swagger:response schemaActionsPropertiesUpdateUnprocessableEntity
*/
type SchemaActionsPropertiesUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesUpdateUnprocessableEntity creates SchemaActionsPropertiesUpdateUnprocessableEntity with default headers values
func NewSchemaActionsPropertiesUpdateUnprocessableEntity() *SchemaActionsPropertiesUpdateUnprocessableEntity {

	return &SchemaActionsPropertiesUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the schema actions properties update unprocessable entity response
func (o *SchemaActionsPropertiesUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaActionsPropertiesUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties update unprocessable entity response
func (o *SchemaActionsPropertiesUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsPropertiesUpdateInternalServerErrorCode is the HTTP code returned for type SchemaActionsPropertiesUpdateInternalServerError
const SchemaActionsPropertiesUpdateInternalServerErrorCode int = 500

/*SchemaActionsPropertiesUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaActionsPropertiesUpdateInternalServerError
*/
type SchemaActionsPropertiesUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesUpdateInternalServerError creates SchemaActionsPropertiesUpdateInternalServerError with default headers values
func NewSchemaActionsPropertiesUpdateInternalServerError() *SchemaActionsPropertiesUpdateInternalServerError {

	return &SchemaActionsPropertiesUpdateInternalServerError{}
}

// WithPayload adds the payload to the schema actions properties update internal server error response
func (o *SchemaActionsPropertiesUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaActionsPropertiesUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties update internal server error response
func (o *SchemaActionsPropertiesUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
