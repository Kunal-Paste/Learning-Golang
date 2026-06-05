package main

import (
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httpserver.NewRouter()

	srv := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("api running on port %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("server closed")
			return
		}
		log.Fatalf("server error : %v", err)
	}
}
