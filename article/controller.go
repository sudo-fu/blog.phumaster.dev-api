package article

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller nil
type Controller struct{}

var service = &Service{}

// GetArticleByID nil
func (c *Controller) GetArticleByID(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, service.Find(id))
}

// CreateArticle nil
func (c *Controller) CreateArticle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, service.Create())
}

// UpdateArticle nil
func (c *Controller) UpdateArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, service.Update(id))
}

// DeleteArticle nil
func (c *Controller) DeleteArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, service.Delete(id))
}

// GetArticleList nil
func (c *Controller) GetArticleList(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, service.All(id))
}
