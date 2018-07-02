// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// cars HTTP client types
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package client

import (
	carssvc "goa.design/goa/examples/streaming/gen/cars"
	carssvcviews "goa.design/goa/examples/streaming/gen/cars/views"
)

// ListResponseBody is the type of the "cars" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// The make of the car
	Make *string `form:"make,omitempty" json:"make,omitempty" xml:"make,omitempty"`
	// The car model
	Model *string `form:"model,omitempty" json:"model,omitempty" xml:"model,omitempty"`
	// The car body style
	BodyStyle *string `form:"body_style,omitempty" json:"body_style,omitempty" xml:"body_style,omitempty"`
}

// LoginUnauthorizedResponseBody is the type of the "cars" service "login"
// endpoint HTTP response body for the "unauthorized" error.
type LoginUnauthorizedResponseBody string

// ListInvalidScopesResponseBody is the type of the "cars" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody string

// ListUnauthorizedResponseBody is the type of the "cars" service "list"
// endpoint HTTP response body for the "unauthorized" error.
type ListUnauthorizedResponseBody string

// NewLoginUnauthorized builds a cars service login endpoint unauthorized error.
func NewLoginUnauthorized(body LoginUnauthorizedResponseBody) carssvc.Unauthorized {
	v := carssvc.Unauthorized(body)
	return v
}

// NewListCarOK builds a "cars" service "list" endpoint result from a HTTP "OK"
// response.
func NewListCarOK(body *ListResponseBody) *carssvcviews.CarView {
	v := &carssvcviews.CarView{
		Make:      body.Make,
		Model:     body.Model,
		BodyStyle: body.BodyStyle,
	}
	return v
}

// NewListInvalidScopes builds a cars service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body ListInvalidScopesResponseBody) carssvc.InvalidScopes {
	v := carssvc.InvalidScopes(body)
	return v
}

// NewListUnauthorized builds a cars service list endpoint unauthorized error.
func NewListUnauthorized(body ListUnauthorizedResponseBody) carssvc.Unauthorized {
	v := carssvc.Unauthorized(body)
	return v
}