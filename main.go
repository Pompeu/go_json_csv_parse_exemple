package main

import (
	"errors"
	"regexp"
)

type Person struct {
	Name  string `name:json`
	Email string `email:json`
	Sexo  string `sexo:json`
	Idade string `idade:json`
}

func GetUrl(url string) (*Person, error) {
	valid, _ := regexp.MatchString("^(http|https)://[a-z.:0-9]+", url)
	person := &Person{}
	if valid {
		return person, nil
	}
	return nil, errors.New("invalid url")
}
