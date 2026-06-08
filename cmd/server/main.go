package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/anomalyco/task-management-api/internal/adapters/outbound/persistence"
	"github.com/anomalyco/task-management-api/internal/adapters/outbound/security"
	"github.com/anomalyco/task-management-api/internal/application"
	httphandler "github.com/anomalyco/task-management-api/internal/adapters/inbound/http"
	"github.com/anomalyco/task-management-api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error al cargar configuracion: %v", err)
	}

	dbCfg := persistence.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.DBSSLMode,
	}

	db, err := persistence.NewConnection(dbCfg)
	if err != nil {
		log.Fatalf("error al conectar a base de datos: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error al obtener conexion sql: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := persistence.RunMigrations(db); err != nil {
		log.Fatalf("error al ejecutar AutoMigrate: %v", err)
	}

	userRepo := persistence.NewGormUserRepository(db)
	sessionRepo := persistence.NewGormSessionRepository(db)
	taskRepo := persistence.NewGormTaskRepository(db)
	hasher := security.NewBcryptHasher()
	tokenSvc := security.NewJWTTokenService(cfg.JWTSecret)

	deps := application.Dependencies{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		TaskRepo:    taskRepo,
		Hasher:      hasher,
		TokenSvc:    tokenSvc,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	httphandler.SetupAuthRoutes(r, deps)
	httphandler.SetupUserRoutes(r, deps)
	httphandler.SetupTaskRoutes(r, deps)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("servidor iniciado en :%s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error al iniciar servidor: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("servidor deteniendose...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error al detener servidor: %v", err)
	}

	sqlDB.Close()
	log.Println("servidor detenido")
}
