package people

import (
	"context"
	"fmt"

	"github.com/graniticio/granitic/v2/ws"
)

type PersonUpdateIFace interface {
	Update(string, map[string]string) error
}

type PersonUpdate struct {
	Store PersonUpdateIFace
}

// ProcessPayload Update for the person
func (pU *PersonUpdate) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, p *UpdatePersonRequest) {
	u := map[string]string{"Name": p.Name.String(), "Age": p.Age.String(), "Gender": p.Gender.String()}

	err := pU.Store.Update(p.UUID, u)

	if err != nil {
		fmt.Println(err)
		res.HTTPStatus = 500
		return
	}

	res.HTTPStatus = 201
	res.Body = map[string]string{"message": "Record has been updated"}
}
