package helpers

import "fmt"

func ErrParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}
