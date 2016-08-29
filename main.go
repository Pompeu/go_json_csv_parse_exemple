package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Person struct {
	Nome   string `csv:"nome"`
	Email  string `csv:"email"`
	Sexo   string `csv:"sexo"`
	Idade  string `csv:"idade"`
	Outros map[string]interface{}
}

func (p *Person) GetUrl(url string) (*Person, error) {
	if res, err := http.Get(url); err == nil {
		content := res.Header.Get("ContentType")
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		if content == "application/json" {
			err = json.Unmarshal(body, &p)
		} else if content == "text/csv" {
			err = p.CsvUnmarshal(string(body))
		}
		return p, err
	} else {

		return nil, err
	}
}

func (p *Person) CsvUnmarshal(strCsv string) error {
	r := csv.NewReader(strings.NewReader(strCsv))

	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	keys := records[0]
	for _, record := range records[1:] {
		p.Nome = record[0]
		p.Email = record[1]
		p.Sexo = record[2]
		p.Idade = record[3]
		internalMap := make(map[string]interface{})
		for i, v := range keys[4:] {
			internalMap[v] = record[i+4]
		}
		p.Outros = internalMap

	}
	return nil
}
