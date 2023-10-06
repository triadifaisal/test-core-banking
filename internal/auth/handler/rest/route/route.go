package route

import (
	"corebanking/pkg/helper/response"
	"corebanking/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// To load swaggo dependencies
	_ = response.ErrorResponse
}

// RegisterRouteV1 for GIN Server
func RegisterRouteV1(router *gin.Engine) {
	if router == nil {
		panic("router is not instantiated")
	}

	router.GET("/nnd-aruok", healthCheck)
	v1 := router.Group("/v1", middleware.CORSMiddleware())
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// enquiryController := leadfactory.GetEnquiryController()
		// v1.GET("/enquiries", enquiryController.GetEnquiriesPaginated)
		// v1.POST("/enquiries", enquiryController.PostEnquiry)
		// v1.GET("/enquiries/:uuid", enquiryController.GetEnquirySingular)
	}
}

// HealthCheck	Health check
// @Summary     Health check
// @Tags        Base
// @Accept      json
// @Produce     json
// @Success     200  {object}  					map[string]any
// @Failure     400  {object}  					map[string]any
// @Failure     500  {object}  					map[string]any
// @Router      /nnd-aruok		[get]
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
