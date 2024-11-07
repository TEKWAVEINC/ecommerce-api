package main

import (
	"ecommerce-api/config"
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"
	"ecommerce-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)

func main() {
	config.InitDB()
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	
	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.Authenticate(), middleware.AdminOnly())
	productRoutes.POST("", controllers.CreateProduct)
	productRoutes.GET("", controllers.GetProducts)
	productRoutes.GET("/:id", controllers.GetProduct)
	productRoutes.PUT("/:id", controllers.UpdateProduct)
	productRoutes.DELETE("/:id", controllers.DeleteProduct)

	orderRoutes := r.Group("/orders")
	orderRoutes.Use(middleware.Authenticate())
	orderRoutes.POST("", controllers.PlaceOrder)
	orderRoutes.GET("", controllers.GetUserOrders)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
