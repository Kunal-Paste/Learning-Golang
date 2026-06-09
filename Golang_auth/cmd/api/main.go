package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("startup failed : %v", err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Fatalf("shutdown warning %s", err)
		}
	}()

	router := httpserver.NewRouter(a)

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
