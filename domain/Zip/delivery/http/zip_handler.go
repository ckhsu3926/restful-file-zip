package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"restful-file-zip/entities"
	"restful-file-zip/tools"
)

type zipHttpHandler struct {
	Usecase entities.ZipUsecase
}

func NewZipHttpHandler(r *gin.RouterGroup, u entities.ZipUsecase) {
	handler := &zipHttpHandler{
		Usecase: u,
	}

	zipRouter := r.Group("zip")
	zipRouter.GET("/file", handler.GetZipFile)
}

type zipResponse struct {
	tools.GinResponse
}

// Zip godoc
// @Summary      GetZipFile
// @Description  Serving zip file
// @Tags         Zip
// @Success      200  {object}  zipResponse
// @Router       /api/zip/file [get]
func (h *zipHttpHandler) GetZipFile(c *gin.Context) {
	response := zipResponse{}

	clearErr := h.Usecase.Clear(c.Request.Context())
	if clearErr != nil {
		response.ErrorResponse(c, clearErr)
		return
	}

	filePath, createErr := h.Usecase.Create(c.Request.Context())
	if createErr != nil {
		response.ErrorResponse(c, createErr)
		return
	}

	// todo
	var fs http.FileSystem

	response.Result = 1
	c.FileFromFS(filePath, fs)
	c.AbortWithStatusJSON(200, response)
}
