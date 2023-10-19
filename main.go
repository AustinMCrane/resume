package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

var (
	localDev = flag.Bool("local", false, "Run in local development mode")
)

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleResumeHTML)
	mux.HandleFunc("/api/resume", HandleResumeJSON)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	if *localDev {
		log.Fatal(http.ListenAndServe(":8080", mux))
		return
	}
	log.Fatal(http.Serve(autocert.NewListener("austinmcrane.com"), mux))
}

func HandleResumeJSON(w http.ResponseWriter, r *http.Request) {
	resume := GetResume()
	data, _ := json.Marshal(resume)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandleResumeHTML(w http.ResponseWriter, r *http.Request) {
	str, err := GetHTMLResume()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(str))
}
