package main

import (
	"log"

	"ecommerce/internal/config"
	"ecommerce/internal/delivery/http/handler"
	"ecommerce/internal/delivery/http/route"
	"ecommerce/internal/domain/entity"
	"ecommerce/internal/domain/service"
	"ecommerce/internal/infrastructure/cache"
	"ecommerce/internal/infrastructure/database"
	"ecommerce/internal/infrastructure/database/repository"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Printf("Starting with default config or environment variables")
	}

	database.ConnectDB(cfg)
	cache.ConnectRedis(cfg)

	err = database.DB.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Product{},
		&entity.Cart{},
		&entity.CartItem{},
		&entity.Order{},
		&entity.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	userRepo := repository.NewUserRepository(database.DB)
	productRepo := repository.NewProductRepository(database.DB)
	categoryRepo := repository.NewCategoryRepository(database.DB)
	cartRepo := repository.NewCartRepository(database.DB)
	orderRepo := repository.NewOrderRepository(database.DB)

	authService := service.NewAuthService(userRepo, cfg)
	productService := service.NewProductService(productRepo, categoryRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	orderService := service.NewOrderService(orderRepo, cartRepo, productRepo)
	userService := service.NewUserService(userRepo)

	authHandler := handler.NewAuthHandler(authService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)
	orderHandler := handler.NewOrderHandler(orderService)
	userHandler := handler.NewUserHandler(userService)

	r := route.SetupRouter(cfg, authHandler, productHandler, cartHandler, orderHandler, userHandler)

	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
