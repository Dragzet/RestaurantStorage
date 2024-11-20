package handlers

import (
	"RestaurantStorage/internal/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/go-ozzo/ozzo-log"
	"net/http"
)

type Handler struct {
	Router  *gin.Engine
	service *service.Service
	logger  *log.Logger
	authMW  *jwt.GinJWTMiddleware
	ProductHandler
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.Router.ServeHTTP(writer, request)
}

type ProductHandler interface {
	AddProduct(c *gin.Context)
	GetProductsList(c *gin.Context)
	ChangeProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetProduct(c *gin.Context)
}

func (h *Handler) registerRoutes() {
	h.Router.POST("/login", h.authMW.LoginHandler)

	authGroup := h.Router.Group("/product")
	authGroup.Use(h.authMW.MiddlewareFunc())
	{
		authGroup.POST("", h.AddProduct)
		authGroup.GET("", h.GetProductsList)
		authGroup.PUT("", h.ChangeProduct)
		authGroup.DELETE("", h.DeleteProduct)
	}

	h.Router.OPTIONS("/product", h.Options)
}

func NewHandler(router *gin.Engine, serv *service.Service, logger *log.Logger, authMW *jwt.GinJWTMiddleware) *Handler {
	handler := &Handler{
		Router:  router,
		service: serv,
		logger:  logger,
		authMW:  authMW,
	}
	handler.registerRoutes()
	return handler
}

func (h *Handler) Options(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
