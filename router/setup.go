package router

import (
	"blog.phumaster.dev-api/article"
	"github.com/labstack/echo/v4"
)

// SetupRoutes app
func SetupRoutes(app *echo.Echo) {
	api := app.Group("/api/v1")

	articleRoutes := api.Group("/article")
	aInstance := (*article.Controller)(nil)
	articleRoutes.GET("/", aInstance.GetArticleList)
	articleRoutes.POST("/", aInstance.CreateArticle)
	articleRoutes.GET("/:id", aInstance.GetArticleByID)
	articleRoutes.PUT("/:id", aInstance.UpdateArticle)
	articleRoutes.DELETE("/:id", aInstance.DeleteArticle)
}
