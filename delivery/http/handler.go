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
	article.PUT("/:id", handler.EditArticle)
	article.DELETE("/:id", handler.DeleteArticle)
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
	article := &domain.Article{}
	if err := ctx.ShouldBindJSON(article); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	if err := a.usecase.AddArticle(ctx, article); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (a *handler) EditArticle(ctx *gin.Context) {
	article := &domain.Article{ID: utils.Pointer(ctx.Param("id"))}
	if err := ctx.ShouldBindJSON(article); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	if err := a.usecase.EditArticle(ctx, article); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (a *handler) DeleteArticle(ctx *gin.Context) {
	if err := a.usecase.DeleteArticle(ctx, utils.Pointer(ctx.Param("id"))); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
