package typeconvertutil_test

import (
	"fmt"
	"testing"

	typeconvertutil "github.com/synzofficial/nsd-synz-sharelib-api/pkg/util/type-convert-util"
)

func TestToBodyReader(t *testing.T) {
	type X struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	x := X{
		Name:  "test",
		Email: "email",
	}
	res, err := typeconvertutil.ToBodyReader(x)
	fmt.Printf("\n%+v\n", res)
	fmt.Printf("\n%+v\n", err)
}
