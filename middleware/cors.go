package middleware

import (
	"restful-file-zip/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CORSMiddleware struct{}

func (m *CORSMiddleware) CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: config.EnvConfig.Cors.AllowAllOrigins,
		AllowOrigins:    config.EnvConfig.Cors.AllowOrigins,
		// AllowOriginFunc: func(origin string) bool {},
		AllowMethods:     config.EnvConfig.Cors.AllowMethods,
		AllowHeaders:     config.EnvConfig.Cors.AllowHeaders,
		AllowCredentials: config.EnvConfig.Cors.AllowCredentials,
		ExposeHeaders:    config.EnvConfig.Cors.ExposeHeaders,
		MaxAge:           0,
		// AllowWildcard:          false,
		// AllowBrowserExtensions: false,
		// AllowWebSockets:        false,
		// AllowFiles:             false,
	})
}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}
