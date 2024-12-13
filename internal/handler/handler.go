package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/go-market-app/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ursulgwopp/go-market-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/profile", h.getProfile)
			users.GET("/:id", h.getUserProfile)
			users.GET("/", h.listUsers)
		}

		products := api.Group("/products")
		{
			products.POST("/", h.addProduct)
			products.GET("/", h.listProducts)
			products.GET("/:id", h.getProduct)
			products.PUT("/:id", h.updateProduct)
			products.DELETE("/:id", h.deleteProduct)
		}

		purchases := api.Group("/purchases")
		{
			purchases.POST("/:id", h.makePurchase)
			purchases.GET("/user/:id", h.getUserPurchases)
			purchases.GET("/product/:id", h.getProductPurhases)
		}
	}

	return router
}
