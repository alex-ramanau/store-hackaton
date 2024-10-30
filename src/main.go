package main

import (
	"github.com/gin-gonic/gin"
	"inventory-service/handlers"
	"inventory-service/models"
	"log"
	//"net/http"
)

func main() {
	// Initialize the database.
	models.InitDB()

	// Set up the Gin router
	r := gin.Default()

	//App health check
    r.GET("/ping", handlers.CheckHealth)

	// Define app routes
	v1 := r.Group("/v1/inventory")
	{
		v1.POST("/get", handlers.GetInventory)
		v1.POST("/grant", handlers.GrantItem)
		v1.POST("/consume", handlers.ConsumeItem)
		v1.POST("/update", handlers.UpdateInventory)

		v1.POST("/catalog/get", handlers.GetCatalog)
		v1.POST("/catalog/create", handlers.CreateCatalogEntry)
		v1.POST("/catalog/update", handlers.UpdateCatalogEntry)
		v1.POST("/catalog/delete", handlers.DeleteCatalogEntry)
	}

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
