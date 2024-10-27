package handlers

import (
	"RestaurantStorage/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/go-ozzo/ozzo-log"
	"net/http"
)

type Handler struct {
	Router  *gin.Engine
	service *service.Service
	logger  *log.Logger
	ProductHandler
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.Router.ServeHTTP(writer, request)
}

type UserHandler interface {
}

type ProductHandler interface {
	AddProduct(c *gin.Context)
	GetProductsList(c *gin.Context)
	ChangeProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetProduct(c *gin.Context)
}

func (h *Handler) registerProductRoutes() {
	h.Router.POST("/product", h.AddProduct)
	h.Router.GET("/product", h.GetProductsList)
	h.Router.PUT("/product", h.ChangeProduct)
	h.Router.DELETE("/product", h.DeleteProduct)
	//h.Router.GET("/product/:name", h.GetProduct)
}

func NewHandler(router *gin.Engine, serv *service.Service, logger *log.Logger) *Handler {
	handler := &Handler{
		Router:  router,
		service: serv,
		logger:  logger,
	}
	handler.registerProductRoutes()
	return handler
}
