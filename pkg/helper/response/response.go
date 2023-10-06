package response

import (
	dtobase "corebanking/pkg/dto/base"

	"github.com/gin-gonic/gin"
)

// ErrorResponse returns an error response
func ErrorResponse(c *gin.Context, err error, httpCode int, data interface{}) {
	if data != nil {
		c.AbortWithStatusJSON(httpCode, data)
	} else {
		c.AbortWithStatusJSON(httpCode, dtobase.BaseRes{
			Code:    httpCode,
			Message: err.Error(),
		})
	}
}
