package main

import(
	"github.com/gin-gonic/gin"
	_ "github.com/infinityethback/pkg/config"
	"github.com/infinityethback/pkg/routes"
)

func main() {
	router := gin.Default()
	
	// Import routes
	routes.UserRoutes(router)
	routes.ReferralRoutes(router)
	routes.QuoteRoutes(router)
	routes.SocialRoutes(router)

	// Start the gin server
	router.Run()
}