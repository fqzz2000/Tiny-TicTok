package main

import (
	"github.com/fqzz2000/tiny-tictok/controller"
	"github.com/fqzz2000/tiny-tictok/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.JwtMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", middleware.SHA1Middleware(), controller.Register) // need a middleware hash the passcode
	apiRouter.POST("/user/login/", middleware.SHA1Middleware(), controller.Login)
	apiRouter.POST("/publish/action/", middleware.JwtMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.JwtMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JwtMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JwtMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.JwtMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
