package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
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
			persons := CsvToPersons(string(body))
			person = &persons[0]
		}

		return person, err
	}
	return nil, errors.New("invalid url")
}

func IsJson(strJson string) bool {
	var sampleJson map[string]interface{}
	return json.Unmarshal([]byte(strJson), &sampleJson) == nil
}

func CsvToPersons(strCsv string) []Person {
	r := csv.NewReader(strings.NewReader(strCsv))
	var keys []string
	var persons []Person
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if record[0] == "name" {
			for _, v := range record {
				keys = append(keys, v)
			}
		} else {
			person := new(Person)
			person.Name = record[0]
			person.Email = record[1]
			person.Sexo = record[2]
			person.Idade = record[3]
			internalMap := make(map[string]interface{})
			for i, v := range keys[4:] {
				internalMap[v] = record[i+4]
			}
			person.Outros = internalMap
			persons = append(persons, *person)
		}
	}
	return persons
}
