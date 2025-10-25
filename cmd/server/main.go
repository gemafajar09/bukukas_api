package main

import (
	"go-project/internal/config"
	"go-project/internal/db/mysql"
	"go-project/internal/delivery/http"
	"log"

	_ "go-project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title BUKU KAS API
// @version 1.0
// @description Ini adalah contoh API dengan Otentikasi JWT di Go menggunakan Swagger
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()

	db, err := mysql.NewPgConnection(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBName,
		cfg.DBPort,
	)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := mysql.AutoMigrate(db); err != nil {
		log.Fatalf("Auto-migrate failed: %v", err)
	}

	port := cfg.ServerPort
	r := http.NewRouter(cfg, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if len(port) > 0 && port[0] != ':' {
		port = ":" + port
	}
	log.Fatal(r.Run(port))
}
