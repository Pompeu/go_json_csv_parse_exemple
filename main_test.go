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

func TestGetUrlDontAcceptInvaidUrl(t *testing.T) {
	_, err := GetUrl("")
	if err == nil {
		t.Error("url pattern is wrong")
	}
}
