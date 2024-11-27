package apiConfig

import (
	"log"
	"net/http"
)

func extractEmail(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}

	email := r.Form.Get("email")
	if email == "" {
		return "", err
	}
	return email, nil
}

func (cfg *ApiConfig) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	// send a magic link to the provided email address
	email, err := extractEmail(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	log.Printf(email)

	w.WriteHeader(200)
}

func (cfg *ApiConfig) HandleSignOut(w http.ResponseWriter, r *http.Request) {}
