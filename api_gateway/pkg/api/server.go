package api

import (
	"api-gateway/pkg/api/handler"
	"api-gateway/pkg/routes"
	"api-gateway/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServeHTTP(c *config.Config, videoHandler handler.VideoHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("template/*.html")
	engine.Static("/static", "./static")
	routes.RegisterVideoRoutes(engine.Group("/"), videoHandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})
	return &Server{
		Engine: engine,
		Port:   c.Port,
	}, nil
}

func (c *Server) Start() {
	c.Engine.Run(c.Port)
}