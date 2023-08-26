package handler

import (
	"user_segmentation/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

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
			segments := users.Group("/segments")
			{
				segments.GET("/show_active", h.showUserActiveSegments)
				segments.POST("/add_segment/:slug", h.addUserToSegment)
				segments.DELETE("/delete_segment/:slug", h.deleteUserFromSegment)
			}
		}
	}

	return router
}
