package main

import (
	"fmt"
	"net/http"
)

func resJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(200)

	js := `{"nome":"Pompeu","email":"pompeulimp@gmail.com","sexo":"M","idade":"33"}`

	fmt.Fprintln(w, js)
}

func resJsonMult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(200)

	js := `{"nome":"Pompeu","email":"pompeulimp@gmail.com","sexo":"M","idade":"33","outros":{"filhos":"1","mae":"Joana"}}`

	fmt.Fprintln(w, js)
}

func resCsv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "text/csv")
	w.WriteHeader(200)

	csv := `nome,email,sexo,idade
Pompeu,pompeulimp@gmail.com,M,33`

	fmt.Fprintln(w, csv)
}

func resCsvMult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "text/csv")
	w.WriteHeader(200)

	csv := `nome,email,sexo,idade,filhos,mae
Pompeu,pompeulimp@gmail.com,M,33,1,Joana`

	fmt.Fprintln(w, csv)
}

func main() {
	http.HandleFunc("/json", resJson)
	http.HandleFunc("/csv", resCsv)
	http.HandleFunc("/csv/mult", resCsvMult)
	http.HandleFunc("/json/mult", resJsonMult)
	http.ListenAndServe(":3000", nil)
}
