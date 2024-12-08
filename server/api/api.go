package api

import (
	"log"
	"net/http"
	"os"
	"path"
	"server/api/routes"
	"server/extra"
	"strings"

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
	server.Use(Cors())
	server.Use(gzip.Gzip(
		gzip.DefaultCompression,
		gzip.WithExcludedPaths([]string{"/api/"}),
	))

	jwtToken := os.Getenv("JWT_TOKEN")
	if jwtToken == "" {
		jwtToken = extra.RandomString(16)
		log.Println("JWT_TOKEN not provided, using random secret")
	}
	routes.SetJwtToken(jwtToken)

	webFiles := os.Getenv("WEB_FILES_PATH")
	if webFiles == "" {
		webFiles = "./www"
		log.Println("WEB_FILES_PATH not provided, using ./www")
	}
	// it'll serve a ReactJS application statically
	server.Use(static.Serve("/", static.LocalFile(webFiles, true)))
	server.NoRoute(func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
			return
		}
		c.File(path.Join(webFiles, "index.html"))
	})

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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func registerAPIRoutes(g *gin.RouterGroup) {
	// this can be used for testing porpuses
	g.GET("/ping", routes.AuthMiddleware(), routes.Ping)

	g.POST("/login", routes.Login)

	g.POST("/locations", routes.AuthMiddleware(), routes.CreateLocation)
	g.GET("/locations", routes.GetAllLocations)
	g.DELETE("/locations/:id", routes.DeleteLocation)

	g.GET("/clients", routes.AuthMiddleware(), routes.GetAllClients)
	g.POST("/clients", routes.AuthMiddleware(), routes.CreateClient)
	g.GET("/clients/:phone", routes.GetClient)
	g.DELETE("/clients/:phone", routes.DeleteClient)

	g.GET("/products", routes.GetAllProducts)
	g.POST("/products", routes.AuthMiddleware(), routes.CreateProduct)
	g.GET("/products/:id", routes.GetProduct)
	g.DELETE("/products/:id", routes.AuthMiddleware(), routes.DeleteProduct)
	g.POST("/products/:id/sizes", routes.AuthMiddleware(), routes.AddProductSize)
	g.DELETE("/products/:id/sizes/:sid", routes.AuthMiddleware(), routes.DeleteProductSize)

	g.POST("/purchases", routes.CreatePurchase)
	g.GET("/purchases", routes.AuthMiddleware(), routes.GetAllPurchases)
	g.GET("/purchases/:id", routes.AuthMiddleware(), routes.GetPurchase)
	g.PUT("/purchases/:id", routes.AuthMiddleware(), routes.SetPurchaseStage)

	g.GET("/tokens/:id", routes.GetTokenUser)
}
