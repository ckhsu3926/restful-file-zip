package http

import (
	"io"
	"os"

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
	zipRouter.GET("/", handler.GetZipFile)
}

type zipResponse struct {
	tools.GinResponse
}

// Zip godoc
// @Summary      GetZipFile
// @Description  Serving zip file
// @Tags         Zip
// @Success      200  {object}  zipResponse
// @Router       /api/zip/ [get]
func (h *zipHttpHandler) GetZipFile(c *gin.Context) {
	response := zipResponse{}

	fileName, filePath, getErr := h.Usecase.Get(c.Request.Context())
	if getErr != nil {
		response.ErrorResponse(c, getErr)
		return
	}

	file, dataErr := os.Open(filePath)
	if dataErr != nil {
		response.ErrorResponse(c, dataErr)
		return
	}
	defer file.Close()

	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
	_, copyErr := io.Copy(c.Writer, file)
	if copyErr != nil {
		response.ErrorResponse(c, copyErr)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
