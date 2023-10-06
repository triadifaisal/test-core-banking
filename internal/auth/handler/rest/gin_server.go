package tsthdlrest

import (
	authRestRoute "corebanking/internal/auth/handler/rest/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

// StartServer start http rest server
// Logger must be initialized
func StartServer(port int) error {
	router := gin.Default()

	authRestRoute.RegisterRouteV1(router)

	return router.Run(fmt.Sprintf(":%d", port))
}
