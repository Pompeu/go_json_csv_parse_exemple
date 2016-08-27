package main

import (
	"testing"
)

func TestPersonDefined(t *testing.T) {
	person := &Person{}
	if person == nil {
		t.Error("person dont defined")
	}
}
