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

func TestGetUrlFromServer(t *testing.T) {
	person, err := GetUrl("http://localhost:3000")
	if err != nil {
		t.Error("url pattern is wrong")
	}

	if person.Email != "pompeulimp@gmail.com" {
		t.Error("email invalido")
	}

	if person.Name != "Pompeu" {
		t.Error("nome invalido")
	}

	if person.Idade != "33" {
		t.Error("idade invalide")
	}

	if person.Sexo != "M" {
		t.Error("sexo invalido")
	}
}
