package api

import (
	"log"
	"server/api/routes"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Start(addr, webFiles string, debug bool) error {
	log.Println("Creating server...")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	// this will serve a ReactJS application statically
	server.Use(static.Serve("/", static.LocalFile(webFiles, true)))

	// this route can be used for testing porpuses
	server.GET("/ping", routes.Ping)

	// from there we will define all routes of the api endpoint
	apiRoute := server.Group("/api")
	registerAPIRoutes(apiRoute)
	log.Print("Starting server...")
	return server.Run(addr)
}

func registerAPIRoutes(g *gin.RouterGroup) {
	g.POST("/locations", routes.CreateLocation)
	g.GET("/locations", routes.GetAllLocations)
	g.DELETE("/locations/:id", routes.DeleteLocation)

	g.GET("/clients", routes.GetAllClients)
	g.POST("/clients", routes.CreateClient)
	g.GET("/clients/:phone", routes.GetClient)
	g.DELETE("/clients/:phone", routes.DeleteClient)

	g.GET("/products", routes.GetAllProducts)
	g.POST("/products", routes.CreateProduct)
	g.GET("/products/:id", routes.GetProduct)
	g.DELETE("/products/:id", routes.DeleteProduct)
	g.POST("/products/:id/sizes", routes.AddProductSize)
	g.DELETE("/products/:id/sizes/:sid", routes.DeleteProductSize)

	g.POST("/purchases", routes.CreatePurchase)
	g.GET("/purchases", routes.GetAllPurchases)
	g.GET("/purchases/:id", routes.GetPurchase)
	g.PUT("/purchases/:id", routes.SetPurchaseStage)

	g.POST("/tokens", routes.CreateToken)
	g.GET("/tokens/:id", routes.GetTokenUser)
}
