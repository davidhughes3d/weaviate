//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

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

	"github.com/semi-technologies/weaviate/entities/models"
)

// NewObjectsClassPatchParams creates a new ObjectsClassPatchParams object
// with the default values initialized.
func NewObjectsClassPatchParams() *ObjectsClassPatchParams {
	var ()
	return &ObjectsClassPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsClassPatchParamsWithTimeout creates a new ObjectsClassPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObjectsClassPatchParamsWithTimeout(timeout time.Duration) *ObjectsClassPatchParams {
	var ()
	return &ObjectsClassPatchParams{

		timeout: timeout,
	}
}

// NewObjectsClassPatchParamsWithContext creates a new ObjectsClassPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewObjectsClassPatchParamsWithContext(ctx context.Context) *ObjectsClassPatchParams {
	var ()
	return &ObjectsClassPatchParams{

		Context: ctx,
	}
}

// NewObjectsClassPatchParamsWithHTTPClient creates a new ObjectsClassPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObjectsClassPatchParamsWithHTTPClient(client *http.Client) *ObjectsClassPatchParams {
	var ()
	return &ObjectsClassPatchParams{
		HTTPClient: client,
	}
}

/*
ObjectsClassPatchParams contains all the parameters to send to the API endpoint
for the objects class patch operation typically these are written to a http.Request
*/
type ObjectsClassPatchParams struct {

	/*Body
	  RFC 7396-style patch, the body contains the object to merge into the existing object.

	*/
	Body *models.Object
	/*ClassName
	  The class name as defined in the schema

	*/
	ClassName string
	/*ID
	  The uuid of the data object to update.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the objects class patch params
func (o *ObjectsClassPatchParams) WithTimeout(timeout time.Duration) *ObjectsClassPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects class patch params
func (o *ObjectsClassPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects class patch params
func (o *ObjectsClassPatchParams) WithContext(ctx context.Context) *ObjectsClassPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects class patch params
func (o *ObjectsClassPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects class patch params
func (o *ObjectsClassPatchParams) WithHTTPClient(client *http.Client) *ObjectsClassPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects class patch params
func (o *ObjectsClassPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the objects class patch params
func (o *ObjectsClassPatchParams) WithBody(body *models.Object) *ObjectsClassPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the objects class patch params
func (o *ObjectsClassPatchParams) SetBody(body *models.Object) {
	o.Body = body
}

// WithClassName adds the className to the objects class patch params
func (o *ObjectsClassPatchParams) WithClassName(className string) *ObjectsClassPatchParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the objects class patch params
func (o *ObjectsClassPatchParams) SetClassName(className string) {
	o.ClassName = className
}

// WithID adds the id to the objects class patch params
func (o *ObjectsClassPatchParams) WithID(id strfmt.UUID) *ObjectsClassPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the objects class patch params
func (o *ObjectsClassPatchParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsClassPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
