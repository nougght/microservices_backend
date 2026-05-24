package main

import (
	"fmt"
	"store-server/config"
	"store-server/database"
	"store-server/internal/auth"
	"store-server/internal/auth/middleware"
	"store-server/internal/auth/tools"
	"store-server/internal/cart"
	"store-server/internal/category"
	"store-server/internal/minio"
	"store-server/internal/od"
	"store-server/internal/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// TODO: переделать сомнительные моменты

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env file not found")
		return
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Configuration loading error: %v", err)
		return
	}

	db, err := database.NewPostgresDB(cfg.Postgres)
	if err != nil {
		fmt.Printf("Database initialization error: %v", err)
		return
	}
	defer db.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // или "*" для всех
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.RedirectTrailingSlash = false

	jwtTools := tools.NewJwtTools(cfg.Jwt.SecretKey)
	authModule := auth.NewAuthModule(cfg, db, jwtTools)
	productModule := product.NewProductModule(db)
	cartModule := cart.NewCartModule(db)
	categoryModule := category.NewCategoryModule(db)
	odModule := od.NewODModule(db)
	minioModule, err := minio.NewMinioModule(cfg.Minio, db)
	if err != nil {
		fmt.Printf("Minio initialization error: %v", err)
		return
	}

	public := r.Group("/")
	private := r.Group("/")
	private.Use(middleware.AccessTokenAuthMiddleware(jwtTools))
	// маршруты без аутентификации
	authModule.RegisterPublicRoutes(public)

	// все остальные c jwt
	authModule.RegisterPrivateRoutes(private)
	productModule.RegisterRoutes(private)
	cartModule.RegisterRoutes(private)
	categoryModule.RegisterRoutes(private)
	minioModule.RegisterRoutes(private)
	odModule.RegisterRoutes(private)

	public.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	err = r.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
