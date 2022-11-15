package http

import (
	"errors"

	"github.com/gin-gonic/gin"

	"restful-file-zip/entities"
	"restful-file-zip/tools"
)

type fileHttpHandler struct {
	Usecase entities.FileUsecase
}

func NewFileHttpHandler(r *gin.RouterGroup, u entities.FileUsecase) {
	handler := &fileHttpHandler{
		Usecase: u,
	}

	fileRouter := r.Group("file")
	fileRouter.POST("/", handler.SaveFile)
	fileRouter.DELETE("/", handler.Delete)
	fileRouter.DELETE("/clear", handler.DeleteClear)
	fileRouter.GET("/list", handler.GetFileList)
}

type response struct {
	tools.GinResponse
}
type fileResponse struct {
	tools.GinResponse
	Data []entities.FileObject `json:"data"`
}

// File godoc
// @Summary      SaveFile
// @Description  Save Post File
// @Description
// @Description  **required** Content-Type: multipart/form-data
// @Tags         File
// @param        file  formData  file  true  "reource files"
// @Success      200   {object}  response
// @Router       /api/file/ [post]
func (h *fileHttpHandler) SaveFile(c *gin.Context) {
	response := response{}

	file, getFileErr := c.FormFile("file")
	if getFileErr != nil {
		response.ErrorResponse(c, getFileErr)
		return
	}

	saveErr := h.Usecase.Save(c.Request.Context(), file)
	if saveErr != nil {
		response.ErrorResponse(c, saveErr)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// File godoc
// @Summary      Delete
// @Description  Delete File
// @Tags         File
// @Param        filename  query     string  true  "File name"
// @Success      200       {object}  response
// @Router       /api/file/ [delete]
func (h *fileHttpHandler) Delete(c *gin.Context) {
	response := response{}

	filename := c.Query("filename")
	if filename == "" {
		response.ErrorResponse(c, errors.New(`filename is empty`))
		return
	}

	deleteErr := h.Usecase.Delete(c.Request.Context(), filename)
	if deleteErr != nil {
		response.ErrorResponse(c, deleteErr)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// File godoc
// @Summary      DeleteClear
// @Description  Clear Source Dir
// @Tags         File
// @Success      200  {object}  response
// @Router       /api/file/clear [delete]
func (h *fileHttpHandler) DeleteClear(c *gin.Context) {
	response := response{}

	clearErr := h.Usecase.Clear(c.Request.Context())
	if clearErr != nil {
		response.ErrorResponse(c, clearErr)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// File godoc
// @Summary      GetFileList
// @Description  Get Source File List
// @Tags         File
// @Success      200  {object}  fileResponse
// @Router       /api/file/list [get]
func (h *fileHttpHandler) GetFileList(c *gin.Context) {
	response := fileResponse{}

	list, getListErr := h.Usecase.GetList(c.Request.Context())
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
