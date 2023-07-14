package http

import (
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/api/middleware"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/api/routes"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/config"

	"github.com/gin-gonic/gin"
	_ "github.com/nikhilnarayanan623/go-socket-chat/server/cmd/api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(cfg config.Config, authHandler interfaces.AuthHandler,
	middleware middleware.Middleware, chatHandler interfaces.ChatHandler) *Server {

	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := engine.Group("/api")

	routes.SetupRoutes(api, authHandler, middleware, chatHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() error {
	return c.engine.Run(c.port)
}
