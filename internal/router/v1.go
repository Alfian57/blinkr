package router

import (
	"github.com/Alfian57/belajar-golang/internal/di"
	"github.com/Alfian57/belajar-golang/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterV1Route(router *gin.RouterGroup) {

	authHandler := di.InitializeAuthHandler()
	userHandler := di.InitializeUserHandler()
	urlHandler := di.InitializeUrlHandler()
	urlVisitorHandler := di.InitializeUrlVisitorHandler()
	bannedDomainHandler := di.InitializeBannedDomainHandler()

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
	router.POST("/refresh", middleware.AuthMiddleware(), authHandler.Refresh)
	router.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)

	admin := router.Group("admin", middleware.AuthMiddleware(), middleware.AdminMiddleware())

	users := admin.Group("users")
	{
		users.GET("/", userHandler.GetAllUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/count", userHandler.CountUsers)
		users.POST("/:id/banned", userHandler.BannedUser)
	}

	urls := admin.Group("urls")
	{
		urls.GET("/count", urlHandler.CountAllUrl)
	}

	urlsVisitor := admin.Group("urls-visitors")
	{
		urlsVisitor.GET("/count", urlVisitorHandler.CountAllUrlVisitors)
		urlsVisitor.GET(":urlID/count", urlVisitorHandler.CountUrlVisitorByID)
	}

	bannedDomain := admin.Group("banned-domains")
	{
		bannedDomain.GET("/", bannedDomainHandler.GetAllBannedDomains)
		bannedDomain.POST("/", bannedDomainHandler.CreateBannedDomain)
		bannedDomain.PUT("/:id", bannedDomainHandler.UpdateBannedDomain)
		bannedDomain.DELETE("/:id", bannedDomainHandler.DeleteBannedDomain)
	}
}
