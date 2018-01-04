package model

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCodeOK(t *testing.T) {
	c := CodeResponse{
		Code: http.StatusOK,
	}
	if !reflect.DeepEqual(c, CodeOK()) {
		t.Error("expected equal structs")
	}
}

func TestNewBoolResponse(t *testing.T) {
	c := BooleanResponse{
		Code: http.StatusOK,
		Data: true,
	}
	if !reflect.DeepEqual(c, NewBoolResponse(true)) {
		t.Error("expected equal structs")
	}
}
