package main

import (
	log "github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_purchase-aggregator/src/api/controller"
)

var (
	Router *gin.Engine
)
func main() {
	router := initRouter()
	MapUrl(router)
	router.Run(":8080")
}

//set up the router
func initRouter() *gin.Engine {

	var r *gin.Engine = gin.New()

	w := log.GetOut()
	r.Use(gin.LoggerWithWriter(w, "/ping"))
	r.Use(gin.RecoveryWithWriter(w))
	return r
}


//set up mapping of URLs
func MapUrl(router *gin.Engine) {

	// Add health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})


	router.GET("/purchases/:purchase_id", controller.GetByPurchaseId)


}

