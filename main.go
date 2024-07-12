package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/AustinMCrane/resume/pkg/model/memory"
	"github.com/AustinMCrane/resume/pkg/server"
)

func main() {
	flag.Parse()
	resumeStore := memory.GetResumeStore()
	resumeServer := server.GetResumeServer(resumeStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/", resumeServer.HandleResume)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portStr := ":" + port

	log.Fatal(http.ListenAndServe(portStr, mux))
}
