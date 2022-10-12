// Code generated by goa v3.10.1, DO NOT EDIT.
//
// calc HTTP client types
//
// Command:
// $ goa gen goa.design/plugins/v3/docs/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/docs/examples/calc

package client

import (
	calc "goa.design/plugins/v3/docs/examples/calc/gen/calc"
)

// AddStreamingBody is the type of the "calc" service "add" endpoint HTTP
// request body.
type AddStreamingBody struct {
	// Left operand
	A int `form:"a" json:"a" xml:"a"`
	// Right operand
	B int `form:"b" json:"b" xml:"b"`
}

// NewAddStreamingBody builds the HTTP request body from the payload of the
// "add" endpoint of the "calc" service.
func NewAddStreamingBody(p *calc.AddStreamingPayload) *AddStreamingBody {
	body := &AddStreamingBody{
		A: p.A,
		B: p.B,
	}
	return body
}
