package http

import (
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/middleware"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/routes"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/config"

	"github.com/gin-gonic/gin"
	docs "github.com/nikhilnarayanan623/react-go-messenger/server/cmd/api/docs"
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

	engine.Use(middleware.Cors())
	engine.Use(gin.Logger())

	docs.SwaggerInfo.BasePath = routes.BaseUrl

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.SetupRoutes(engine, authHandler, middleware, userHandler, chatHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() error {
	return c.engine.Run(c.port)
}
