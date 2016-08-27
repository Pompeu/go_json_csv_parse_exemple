package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func resJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	js := map[string]string{
		"name":  "Pompeu",
		"email": "pompeulimp@gmail.com",
		"sexo":  "M",
		"idade": "33"}

	res, _ := json.Marshal(js)

	io.WriteString(w, string(res))
}

func main() {
	http.HandleFunc("/", resJson)
	http.ListenAndServe(":3000", nil)
}
