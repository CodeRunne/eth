package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/infinityethback/pkg/controllers"
)

var UserRoutes = func(router *gin.Engine) {
	router.GET("/users", controllers.UIndex)
	router.GET("/users/:address", controllers.UShow)
	router.GET("/users/leaderboard", controllers.ULeaderboard)
	router.POST("/users", controllers.UStore)
}

var ReferralRoutes = func(router *gin.Engine) {
	router.GET("/referrals", controllers.RIndex)
	router.POST("/referrals/:token", controllers.Refer)
}

var QuoteRoutes = func(router *gin.Engine) {
	router.GET("/quotes/latest", controllers.QShow)
	router.POST("/quotes", controllers.QStore)
}

var SocialRoutes = func(router *gin.Engine) {
	router.POST("/quotes/socials", controllers.SStore)
}