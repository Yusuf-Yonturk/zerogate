package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerogate/payment/internal/config"
	"github.com/zerogate/payment/internal/database"
	"github.com/zerogate/payment/internal/payment"
	"github.com/zerogate/payment/internal/redis"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	locker := redis.NewLocker(cfg.RedisURL)
	repo := payment.NewRepository(db)
	svc := payment.NewService(repo, locker)
	handler := payment.NewHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/api/pay", handler.HandlePayment).Methods("POST")

	log.Printf("Server listening on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
