package main

import (
	"encoding/json"
	"errors"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Person struct {
	Name   string `csv:"name"`
	Email  string `csv:"email"`
	Sexo   string `csv:"sexo"`
	Idade  string `csv:"idade"`
	Outros map[string]interface{}
}

func GetUrl(url string) (*Person, error) {
	valid, _ := regexp.MatchString("^(http|https)://[a-z.:0-9]+", url)
	person := &Person{}
	if valid {
		res, err := http.Get(url)
		content := res.Header.Get("ContentType")
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if IsJson(string(body)) && content == "application/json" {
			err = json.Unmarshal(body, &person)
		} else if content == "text/csv" {
			var persons = []*Person{}
			err = gocsv.UnmarshalBytes(body, &persons)
			person = persons[0]
		}

		return person, err
	}
	return nil, errors.New("invalid url")
}

func IsJson(strJson string) bool {
	var sampleJson map[string]interface{}
	return json.Unmarshal([]byte(strJson), &sampleJson) == nil
}
