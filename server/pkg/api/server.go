package http

import (
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/middleware"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/routes"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/config"

	"github.com/gin-gonic/gin"
	_ "github.com/nikhilnarayanan623/server/react-go-messenger/cmd/api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(cfg config.Config, authHandler interfaces.AuthHandler,
	middleware middleware.Middleware, userHandler interfaces.UserHandler,
	chatHandler interfaces.ChatHandler) *Server {

	engine := gin.New()
	// engine.Use(cors.Default())
	engine.Use(middleware.Cors())
	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := engine.Group("/api")

	routes.SetupRoutes(api, authHandler, middleware, userHandler, chatHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() error {
	return c.engine.Run(c.port)
}
