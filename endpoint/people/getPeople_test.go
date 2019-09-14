package people

import (
	"context"
	"errors"
	"testing"

	"github.com/graniticio/granitic/v2/test"
	"github.com/graniticio/granitic/v2/ws"
)

type MockStoreResponse struct {
	people [][]string
	err    error
}

func (m MockStoreResponse) Read() ([][]string, error) {
	return m.people, m.err
}

func TestPeopleIndex_Process(t *testing.T) {
	t.Log("When the Get API for people succeeds")
	{
		t.Run("returns all person records", func(t *testing.T) {
			people := make([][]string, 2)
			people[0] = []string{"uid1", "John Doe", "20", "Male"}
			people[1] = []string{"uid2", "Jane Doe", "20", "Female"}
			mock := MockStoreResponse{people: people, err: nil}
			pI := PeopleIndex{Store: mock}
			req := &ws.Request{}
			res := &ws.Response{}
			pI.Process(context.TODO(), req, res)
			test.ExpectString(t, res.Body.([]Person)[0].Name, "John Doe")
			test.ExpectString(t, res.Body.([]Person)[1].Name, "Jane Doe")
		})
	}

	t.Log("When the Get API for people fails")
	{
		t.Run("returns error message", func(t *testing.T) {
			mock := MockStoreResponse{people: [][]string{}, err: errors.New("Db Error")}
			pI := PeopleIndex{Store: mock}
			req := &ws.Request{}
			res := &ws.Response{}
			pI.Process(context.TODO(), req, res)
			test.ExpectInt(t, res.HTTPStatus, 500)
			// test.ExpectNil(t, res.Body)
		})
	}
}
