package dig1

import (
	"reflect"
	"testing"
)

func TestErrorType(t *testing.T) {
	_errType := reflect.TypeOf((*error)(nil))
	// *error
	t.Log(_errType)
	_errT := _errType.Elem()
	// error
	t.Log(_errT)
}
