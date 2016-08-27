package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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
		res, err := http.Get(url)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if IsJson(string(body)) {
			err = json.Unmarshal(body, &person)
		}
		return person, err
	}
	return nil, errors.New("invalid url")
}

func IsJson(strJson string) bool {
	var sampleJson map[string]string
	return json.Unmarshal([]byte(strJson), &sampleJson) == nil
}
