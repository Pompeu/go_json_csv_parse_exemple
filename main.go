package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	if res, err := http.Get(url); err == nil {
		content := res.Header.Get("ContentType")
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		person := &Person{}
		if content == "application/json" {
			err = json.Unmarshal(body, &person)
		} else if content == "text/csv" {
			person, err = CsvToPerson(string(body))
		}
		return person, err
	} else {

		return nil, err
	}
}

func CsvToPerson(strCsv string) (*Person, error) {
	r := csv.NewReader(strings.NewReader(strCsv))
	person := new(Person)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	keys := records[0]
	for _, record := range records[1:] {
		person.Name = record[0]
		person.Email = record[1]
		person.Sexo = record[2]
		person.Idade = record[3]
		internalMap := make(map[string]interface{})
		for i, v := range keys[4:] {
			internalMap[v] = record[i+4]
		}
		person.Outros = internalMap

	}
	return person, nil
}
