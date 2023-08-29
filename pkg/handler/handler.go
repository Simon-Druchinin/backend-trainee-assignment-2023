package handler

import (
	"net/http"
	"user_segmentation/pkg/service"

	"github.com/gin-contrib/cors"
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

	// Do not use in production - explicitly set allowed origins
	router.Use(cors.Default())

	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
	}

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("", h.createSegment)
			segments.DELETE("/:slug", h.deleteSegment)
		}
		users := api.Group("users")
		{
			users.GET("/:id/show_active_segments", h.showUserActiveSegments)
			users.POST("/:id/add_to_segment", h.addUserToSegment)
			users.DELETE("/:id/delete_from_segment", h.deleteUserFromSegment)
		}
	}

	return router
}
