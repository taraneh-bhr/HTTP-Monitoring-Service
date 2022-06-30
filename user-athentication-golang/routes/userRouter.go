package routes

import (
    controller "user-athentication-golang/controllers"
    middleware "httpmon.com/first/middleware"

    "github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/users/signup", controller.SignUp())
    incomingRoutes.POST("/users/login", controller.Login())
    incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users/history", controller.GetHistory())
	incomingRoutes.GET("/users/alerts", controller.GetAlerts())
}
