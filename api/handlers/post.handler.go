package handlers

import (
	"fmt"
	"net/http"
	"post-api/api/models"
	"post-api/api/services"
	"post-api/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post_Handler interface {
	GetAll(c *gin.Context) error
	GetById(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type post_Handler struct {
	serv services.Post_Service
}

func New_PostHandler(serv services.Post_Service) Post_Handler {
	return &post_Handler{serv: serv}
}

func (h *post_Handler) GetAll(c *gin.Context) error {

	posts, err := h.serv.GetAll()
	if err != nil {
		fmt.Println(c.Request.Method, err)
		return err
	}

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	pkg.Response(&pkg.Respon{Code: http.StatusOK, Message: "Berhasil get data all post", Data: posts}).Send(c)
	return nil
}

func (h *post_Handler) GetById(c *gin.Context) error {
	post := new(models.Post)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	post.ID = id

	response, err := h.serv.GetById(post)
	if err != nil {
		return err
	}

	pkg.Response(&pkg.Respon{Code: http.StatusOK, Message: "Berhasil get data", Data: response}).Send(c)
	return nil
}

func (h *post_Handler) Create(c *gin.Context) error {
	postReq := new(models.PostReq)
	if err := c.ShouldBindJSON(postReq); err != nil {
		return err
	}

	post, err := h.serv.Create(postReq)
	if err != nil {
		return err
	}

	pkg.Response(&pkg.Respon{Message: "Berhasil insert", Data: post}).Send(c)
	return nil
}

func (h *post_Handler) Update(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	postReq := new(models.PostReq)
	if err := c.ShouldBindJSON(&postReq); err != nil {
		return err
	}

	err = h.serv.Update(id, postReq)
	if err != nil {
		return err
	}

	return nil
}

func (h *post_Handler) Delete(c *gin.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = h.serv.Delete(id)
	if err != nil {
		return err
	}

	pkg.Response(&pkg.Respon{Message: "Berhasil delete data"}).Send(c)
	return nil
}
