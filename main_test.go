package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonDefined(t *testing.T) {
	person := &Person{}
	assert.NotNil(t, person, "person dont defined")
}

func TestGetUrlDontAcceptInvaidUrl(t *testing.T) {
	_, err := GetUrl("")
	assert.Error(t, err, "url pattern is wrong")
	assert.Equal(t, err.Error(), "invalid url", "url pattern is wrong")
}

func TestGetUrlFromServer(t *testing.T) {
	person, err := GetUrl("http://localhost:3000/json")

	assert.Nil(t, err, "url pattern is wrong")
	assert.Equal(t, person.Email, "pompeulimp@gmail.com", "email invalido")
	assert.Equal(t, person.Name, "Pompeu", "nome invalido")
	assert.Equal(t, person.Idade, "33", "idade invalide")
	assert.Equal(t, person.Sexo, "M", "sexo invalido")
}

func TestGetUrlJsonDynamic(t *testing.T) {
	person, err := GetUrl("http://localhost:3000/json/mult")
	assert.Nil(t, err, "url pattern is wrong")
	assert.Equal(t, person.Email, "pompeulimp@gmail.com", "email invalido")
	assert.Equal(t, person.Name, "Pompeu", "nome invalido")
	assert.Equal(t, person.Idade, "33", "idade invalide")
	assert.Equal(t, person.Sexo, "M", "sexo invalido")
	assert.Equal(t, person.Outros["childrens"], "1")
	assert.Equal(t, person.Outros["mother_name"], "Joana")
}

func TestIsJson(t *testing.T) {
	isJson := IsJson(`{"name":"pompeu"}`)
	assert.True(t, isJson, "não é um json invalido")
}

func TestInvalidJson(t *testing.T) {
	isInvalid := IsJson(`{"email": "password", "jose"}`)
	assert.False(t, isInvalid, "esse json é valido")
}

func TestGetUrlToCsv(t *testing.T) {
	person, err := GetUrl("http://localhost:3000/csv")
	assert.Nil(t, err, "não é possivel pegar o csv")
	assert.Equal(t, person.Email, "pompeulimp@gmail.com", "email invalido")
	assert.Equal(t, person.Name, "Pompeu", "nome invalido")
	assert.Equal(t, person.Idade, "33", "idade invalide")
	assert.Equal(t, person.Sexo, "M", "sexo invalido")
}
