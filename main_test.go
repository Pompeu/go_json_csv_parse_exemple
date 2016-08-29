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
	assert.Error(t, err)
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

func TestGetUrlToCsv(t *testing.T) {
	person, err := GetUrl("http://localhost:3000/csv")
	assert.Nil(t, err, "não é possivel pegar o csv")
	assert.Equal(t, person.Email, "pompeulimp@gmail.com", "email invalido")
	assert.Equal(t, person.Name, "Pompeu", "nome invalido")
	assert.Equal(t, person.Idade, "33", "idade invalide")
	assert.Equal(t, person.Sexo, "M", "sexo invalido")
}

func TestGetUrlCsvDynamic(t *testing.T) {
	person, err := GetUrl("http://localhost:3000/csv/mult")
	assert.Nil(t, err, "url pattern is wrong")
	assert.Equal(t, person.Email, "pompeulimp@gmail.com", "email invalido")
	assert.Equal(t, person.Name, "Pompeu", "nome invalido")
	assert.Equal(t, person.Idade, "33", "idade invalide")
	assert.Equal(t, person.Sexo, "M", "sexo invalido")
	assert.Equal(t, person.Outros["childrens"], "1")
	assert.Equal(t, person.Outros["mother_name"], "Joana")
}

func TestCsvToPerson(t *testing.T) {
	in := `name,email,sexo,idade,childrens,mother_name
Pompeu,pompeulimp@gmail.com,M,33,1,Joana`

	person, err := CsvToPerson(in)
	assert.Nil(t, err)
	assert.Equal(t, person.Email, "pompeulimp@gmail.com", "email invalido")
	assert.Equal(t, person.Name, "Pompeu", "nome invalido")
	assert.Equal(t, person.Idade, "33", "idade invalide")
	assert.Equal(t, person.Sexo, "M", "sexo invalido")
	assert.Equal(t, person.Outros["childrens"], "1")
	assert.Equal(t, person.Outros["mother_name"], "Joana")
}

func TestCsvToPersonInvalid(t *testing.T) {
	in := `name,email,sexo,idade,childrens,mother_name
Pompeu,pompeulimp@gmail.com,M,33,1,Joana,invalid`

	person, err := CsvToPerson(in)
	assert.NotNil(t, err)
	assert.Nil(t, person)
}
