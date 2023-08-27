package handler

import (
	"user_segmentation/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "user_segmentation/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
	}

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("/", h.createSegment)
			segments.DELETE("/:slug", h.deleteSegment)
		}
		users := api.Group("users")
		{
			users.GET("/show_active_segments", h.showUserActiveSegments)
			users.POST("/add_to_segment", h.addUserToSegment)
			users.DELETE("/delete_from_segment/:slug", h.deleteUserFromSegment)
		}
	}

	return router
}
