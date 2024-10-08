package main

import (
	"api/database"
	"api/routes"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// first we need to load some variables to start our system

	godotenv.Load(".env", "../.env")

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":8080"
	}

	webFiles := os.Getenv("WEB_FILES")
	dbPath := os.Getenv("DB_PATH")
	if webFiles == "" {
		panic("WEB_FILES environment variable not defined!\nIt's needed to show the webpage at /")
	} else if dbPath == "" {
		panic("DB_PATH not defined!")
	} else if err := database.Init(dbPath); err != nil { // initialize database and this if it's ok
		panic("failed to init database! " + err.Error())
	}

	server := gin.Default()

	// this will serve a ReactJS application statically
	server.Use(static.Serve("/", static.LocalFile(webFiles, true)))

	// this route can be used for testing porpuses
	server.GET("/ping", routes.Ping)

	// from there we will define all routes of the api endpoint
	apiRoute := server.Group("/api")
	{ // TODO: Add an authentication for some routes
		apiRoute.POST("/locations", routes.CreateLocation)
		apiRoute.GET("/locations", routes.GetAllLocations)
		apiRoute.DELETE("/locations/:id", routes.DeleteLocation)

		apiRoute.GET("/clients", routes.GetAllClients)
		apiRoute.POST("/clients", routes.CreateClient)
		apiRoute.GET("/clients/:phone", routes.GetClient)
		apiRoute.DELETE("/clients/:phone", routes.DeleteClient)

		apiRoute.GET("/products", routes.GetAllProducts)
		apiRoute.POST("/products", routes.CreateProduct)
		apiRoute.GET("/products/:id", routes.GetProduct)
		apiRoute.DELETE("/products/:id", routes.DeleteProduct)
		apiRoute.POST("/products/:id/sizes", routes.AddProductSize)
		apiRoute.DELETE("/products/:id/sizes/:sid", routes.DeleteProductSize)

		apiRoute.POST("/purchases", routes.CreatePurchase)
		apiRoute.GET("/purchases", routes.GetAllPurchases)
		apiRoute.GET("/purchases/:id", routes.GetPurchase)
	}

	// put this shit to run and see what it will do
	if err := server.Run(addr); err != nil {
		panic(err)
	}
}
