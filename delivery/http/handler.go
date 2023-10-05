package http

import (
	"net/http"

	"ton-ton/domain"
	"ton-ton/utils"

	"github.com/gin-gonic/gin"
)

// handler represent the httphandler
type handler struct {
	usecase domain.Usecase
}

// NewHandler will initialize the endpoint
func NewHandler(g *gin.RouterGroup, us domain.Usecase) {
	handler := &handler{usecase: us}

	article := g.Group("/article")
	article.GET("/:id", handler.GetArticle)
	article.POST("", handler.AddArticle)
	article.PUT("", handler.EditArticle)
	article.DELETE("", handler.DeleteArticle)
}

func (a *handler) GetArticle(ctx *gin.Context) {
	article, err := a.usecase.GetArticle(ctx, utils.Pointer(ctx.Param("id")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func (a *handler) AddArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (a *handler) EditArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (a *handler) DeleteArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
