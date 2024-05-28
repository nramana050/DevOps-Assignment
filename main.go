package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func provisionHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	err := cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Infrastructure provisioned")
}

func destroyHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	err := cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Infrastructure destroyed")
}

func main() {
	http.HandleFunc("/provision", provisionHandler)
	http.HandleFunc("/destroy", destroyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
