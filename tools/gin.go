package tools

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type GinResponse struct {
	Result int `json:"result" example:"1"`
}
type ErrorResponse struct {
	GinResponse
	Error string `json:"error" example:""`
}

// Print error debug and abort with 400 response
func (res *GinResponse) ErrorResponse(c *gin.Context, err error) {
	response := ErrorResponse{
		GinResponse: GinResponse{
			Result: 0,
		},
		Error: err.Error(),
	}

	fmt.Println(err.Error())
	if _, ok := errorList[err.Error()]; !ok {
		debug.PrintStack()
	}

	c.AbortWithStatusJSON(400, response)
}
