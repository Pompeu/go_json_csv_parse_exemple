package main

type Person struct {
	Name  string `name:json`
	Email string `email:json`
	Sexo  string `sexo:json`
	Idade string `idade:json`
}
