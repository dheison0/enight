package api

import (
	"log"
	"os"
	"server/api/routes"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Start(debug bool) error {
	log.Println("Creating server...")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.Use(gzip.Gzip(
		gzip.DefaultCompression,
		gzip.WithExcludedPaths([]string{"/api/"}),
	))

	webFiles := os.Getenv("WEB_FILES_PATH")
	if webFiles == "" {
		webFiles = "./www"
		log.Println("WEB_FILES_PATH not provided, using ./www")
	}
	// it'll serve a ReactJS application statically
	server.NoRoute(static.Serve("/", static.LocalFile(webFiles, true)))

	// this can be used for testing porpuses
	server.GET("/ping", routes.Ping)

	apiRoute := server.Group("/api")
	registerAPIRoutes(apiRoute)
	log.Println("Starting server...")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
		log.Println("SERVER_PORT not provided, using 8080")
	}
	return server.Run(":" + port)
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
