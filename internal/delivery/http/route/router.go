package route

import (
	"ecommerce/internal/config"
	"ecommerce/internal/delivery/http/handler"
	"ecommerce/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	cfg config.Config,
	authHandler *handler.AuthHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/password-reset", authHandler.RequestPasswordReset)
		}

		products := api.Group("/products")
		{
			products.GET("", productHandler.ListProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.GET("/search", productHandler.SearchProducts)
		}

		categories := api.Group("/categories")
		{
			categories.GET("", productHandler.ListCategories)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			cart := protected.Group("/cart")
			{
				cart.GET("", cartHandler.GetCart)
				cart.POST("/items", cartHandler.AddToCart)
				cart.PUT("/items/:product_id", cartHandler.UpdateCartItem)
			}

			orders := protected.Group("/orders")
			{
				orders.POST("", orderHandler.PlaceOrder)
				orders.GET("", orderHandler.GetUserOrders)
				orders.GET("/:id", orderHandler.GetOrderDetails)
			}

			admin := protected.Group("/admin")
			{
				admin.POST("/products", productHandler.CreateProduct)
				admin.POST("/categories", productHandler.CreateCategory)
			}

			protected.GET("/profile", userHandler.GetProfile)
			protected.PUT("/profile", userHandler.UpdateProfile)

			protected.GET("/me", func(c *gin.Context) {
				userID, _ := c.Get("user_id")
				c.JSON(200, gin.H{"user_id": userID})
			})
		}
	}

	return r
}
