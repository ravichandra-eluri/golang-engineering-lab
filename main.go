package main

import (
	"log/slog"
	"net/http"
	"os"

	"golang-engineering-lab/config"
	"golang-engineering-lab/db"
	"golang-engineering-lab/handler"
	"golang-engineering-lab/middleware"
	"golang-engineering-lab/service"
)

func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg.DBConn)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	authSvc := service.NewAuthService(cfg.JWTSecret)
	userSvc := service.NewUserService(database)

	authHandler := &handler.AuthHandler{Secret: cfg.JWTSecret}
	userHandler := &handler.UserHandler{Users: userSvc}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", authHandler.Login)
	mux.Handle("GET /users", middleware.Auth(authSvc)(http.HandlerFunc(userHandler.List)))

	slog.Info("starting server", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, middleware.Logging(mux)); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
