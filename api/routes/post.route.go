package routes

import (
	"net/http"
	"post-api/api/handlers"
	"post-api/api/repos"
	"post-api/api/services"
	"post-api/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostRoute(r *gin.Engine, db *gorm.DB) {
	repo := repos.New_PostRepo(db)
	serv := services.New_PostService(repo)
	post := handlers.New_PostHandler(serv)

	r.GET("/api/posts", wrap(post.GetAll))
	r.GET("/api/posts/:id", wrap(post.GetById))
	r.POST("/api/posts", wrap(post.Create))
	r.PUT("/api/posts/:id", wrap(post.Update))
	r.DELETE("/api/posts/:id", wrap(post.Delete))

}

type ginHandler func(*gin.Context) error

func wrap(f ginHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := f(ctx); err != nil {
			pkg.ErrorRes(http.StatusBadRequest, err).Send(ctx)
		}
	}
}
