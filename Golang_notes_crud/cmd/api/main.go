package main

import (
	"fmt"
	"log"
	"note-api/internal/config"
	"note-api/internal/db"
	"note-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error")
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db error")
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Fatalf("mongo disconnect error: %v", err)
		}
	}()

	router := server.Newrouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}

}
