package main

import (
	"time"

	"github.com/gin-gonic/gin"

	_ "restful-file-zip/docs"

	"restful-file-zip/config"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_middleware "restful-file-zip/middleware"

	_fileDeliveryHttp "restful-file-zip/domain/File/delivery/http"
	_fileRepository "restful-file-zip/domain/File/repository/file"
	_fileUsecase "restful-file-zip/domain/File/usecase"

	_zipDeliveryHttp "restful-file-zip/domain/Zip/delivery/http"
	_zipRepository "restful-file-zip/domain/Zip/repository/file"
	_zipUsecase "restful-file-zip/domain/Zip/usecase"
)

var timeContext = time.Duration(2) * time.Second

func init() {
	// env
	if err := config.InitialEnvConfiguration(); err != nil {
		panic(err)
	}

	if config.EnvConfig.SourcePath[len(config.EnvConfig.SourcePath)-1] == '/' {
		config.EnvConfig.SourcePath = config.EnvConfig.SourcePath[0 : len(config.EnvConfig.SourcePath)-1]
	}
	if config.EnvConfig.ArchivePath[len(config.EnvConfig.ArchivePath)-1] == '/' {
		config.EnvConfig.ArchivePath = config.EnvConfig.ArchivePath[0 : len(config.EnvConfig.ArchivePath)-1]
	}
}

// @Title  restful-file-zip
// @Description
// @Version  0.1
// @Host     localhost:8081
func main() {
	// gin
	if config.EnvConfig.Env != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// middleware
	corsMiddleware := _middleware.NewCORSMiddleware()
	r.Use(corsMiddleware.CORS())

	// api
	apiRouter := r.Group("api")

	fileRepo := _fileRepository.NewFileRepository(config.EnvConfig.SourcePath)
	fileUsecase := _fileUsecase.NewFileUsecase(fileRepo, timeContext)
	_fileDeliveryHttp.NewFileHttpHandler(apiRouter, fileUsecase)

	zipRepo := _zipRepository.NewFileZipRepository(config.EnvConfig.SourcePath+"/", config.EnvConfig.ArchivePath)
	zipUsecase := _zipUsecase.NewZipUsecase(zipRepo, timeContext)
	_zipDeliveryHttp.NewZipHttpHandler(apiRouter, zipUsecase)

	// gin swagger
	if config.EnvConfig.Env == "local" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	if err := r.Run(config.EnvConfig.Host + ":" + config.EnvConfig.Port); err != nil {
		panic(err)
	}
}
