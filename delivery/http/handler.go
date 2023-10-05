package http

import (
	"net/http"

	"ton-ton/domain"
	"ton-ton/utils"

	"github.com/labstack/echo"
)

// handler represent the httphandler
type handler struct {
	usecase domain.Usecase
}

// NewHandler will initialize the endpoint
func NewHandler(g *echo.Group, us domain.Usecase) {
	handler := &handler{usecase: us}

	article := g.Group("/article")
	article.GET("/:id", handler.GetArticle)
	article.POST("", handler.AddArticle)
	article.PUT("", handler.EditArticle)
	article.DELETE("", handler.DeleteArticle)
}

func (a *handler) GetArticle(ctx echo.Context) error {
	article, err := a.usecase.GetArticle(ctx.Request().Context(), utils.Pointer(ctx.Param("id")))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, article)
}

func (a *handler) AddArticle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

func (a *handler) EditArticle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

func (a *handler) DeleteArticle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
