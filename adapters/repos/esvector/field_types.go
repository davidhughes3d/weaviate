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

package esvector

// FieldType in Elasticsearch
// These should be provided by the official es go client, but couldn't be found
// in there
type FieldType string

const (

	// Text indexes a block of text, such as the body of an email
	Text FieldType = "text"

	// Keyword indexes an exact string, such as an email address
	Keyword FieldType = "keyword"

	// Integer indexes an integer
	Integer FieldType = "integer"

	// Float indexes a float64
	Float FieldType = "float"

	// Long indexes a long (which can be used for uint64)
	Long FieldType = "long"

	// Boolean indexes a boolean
	Boolean FieldType = "boolean"

	// Date indexes a date
	Date FieldType = "date"

	// GeoPoint indexes a geo point
	GeoPoint FieldType = "geo_point"
)
