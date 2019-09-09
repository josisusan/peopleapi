package people

import "github.com/graniticio/granitic/v2/types"

type Person struct {
	UUID   string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type AddPersonRequest struct {
	Name   *types.NilableString `json:"name"`
	Age    *types.NilableString `json:"age"`
	Gender *types.NilableString `json:"gender"`
}

type UpdatePersonRequest struct {
	UUID   string
	Name   *types.NilableString `json:"name"`
	Age    *types.NilableString `json:"age"`
	Gender *types.NilableString `json:"gender"`
}
