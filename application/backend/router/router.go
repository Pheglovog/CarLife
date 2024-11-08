package router

import (
	controller "carlife-backend/controllers"
	"carlife-backend/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/Car", controller.GetCar)
		api.GET("/CarList", controller.GetCarList)
		api.POST("/SetCarTires", controller.SetCarTires)
		api.POST("/SetCarBody", controller.SetCarBody)
		api.POST("/SetCarInterior", controller.SetCarInterior)
		api.POST("/SetCarManu", controller.SetCarManu)
		api.POST("/SetCarStore", controller.SetCarStore)
		api.POST("/SetCarInsure", controller.SetCarInsure)
		api.POST("/SetCarMaint", controller.SetCarMaint)
		api.POST("/TransferCar", controller.TransferCar)
	}
	return r
}
