package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.POST("/register", controllers.RegisterUser)
	api.POST("/login", controllers.LoginUser)

	auth := api.Group("/")
	auth.Use(middlewares.AuthenticateUser())
	auth.POST("/orders", controllers.PlaceOrder)
	auth.GET("/orders", controllers.ListOrders)
	auth.PATCH("/orders/:id/cancel", controllers.CancelOrder)

	auth.PATCH("/orders/:id/status", middlewares.AuthorizeAdmin(), controllers.UpdateOrderStatus)
}
