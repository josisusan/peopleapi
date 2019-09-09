package people

import (
	"context"
	"fmt"

	"github.com/graniticio/granitic/v2/ws"
)

// PersonCreateIFace Writable Interface
type PersonCreateIFace interface {
	Write(string) (string, error)
}

// PersonCreate struct
type PersonCreate struct {
	Store PersonCreateIFace
}

// ProcessPayload Create Action for people endpoint
func (pC *PersonCreate) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, p *AddPersonRequest) {
	d := fmt.Sprintf("%s,%s,%s\n", p.Name, p.Age, p.Gender)
	if _, err := pC.Store.Write(d); err != nil {
		res.HTTPStatus = 500
		res.Body = "Something went wrong"
	} else {
		res.HTTPStatus = 201
		res.Body = "Person info was succesfully registered"
	}
}
