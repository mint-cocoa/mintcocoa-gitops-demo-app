package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type response struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Env     string `json:"env"`
	Time    string `json:"time"`
}

func main() {
	version := getenv("APP_VERSION", "dev")
	env := getenv("APP_ENV", "local")
	port := getenv("PORT", "8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{
			Service: "mintcocoa-gitops-demo-app",
			Version: version,
			Env:     env,
			Time:    time.Now().UTC().Format(time.RFC3339),
		})
	})

	log.Printf("listening on :%s version=%s env=%s", port, version, env)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
