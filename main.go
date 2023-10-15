package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	go func() {
		resume := GetResume()
		data, _ := json.Marshal(resume)
		ioutil.WriteFile("resume.json", data, 0644)
	}()
	http.HandleFunc("/", HandleResume)
	http.ListenAndServe(":8080", nil)
}

func HandleResume(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-type")
	switch contentType {
	case "application/json":
		HandleResumeJSON(w, r)
	default:
		HandleResumeHTML(w, r)
	}
}

func HandleResumeJSON(w http.ResponseWriter, r *http.Request) {
	resume := GetResume()
	data, _ := json.Marshal(resume)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandleResumeHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Austin Crane</h1>"))
}
