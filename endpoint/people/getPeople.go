package people

import (
	"context"
	"strconv"

	"github.com/graniticio/granitic/v2/ws"
)

type PeopleIndexIFace interface {
	Read() ([][]string, error)
}

type PeopleIndex struct {
	Store PeopleIndexIFace
}

// Process Index handler for People
func (pI *PeopleIndex) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	data, err := pI.Store.Read()
	if err != nil {
		res.HTTPStatus = 500
		return
	}
	people := make([]Person, len(data))
	for i, d := range data {
		people[i].UUID = d[0]
		people[i].Name = d[1]
		people[i].Age, _ = strconv.Atoi(d[2])
		people[i].Gender = d[3]
	}

	res.Body = people
	res.HTTPStatus = 200
}
