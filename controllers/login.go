package controllers

import "github.com/gin-gonic/gin"

type LoginController struct{}

func (lc *LoginController) InitLoginControllerRouter(router *gin.Engine) {
	loginRoutes := router.Group("/login")
	loginRoutes.GET("/", lc.loginCheck())
	loginRoutes.POST("/new", lc.newUser())
}

func (lc *LoginController) loginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "secretKey",
		})
	}
}

func (lc *LoginController) newUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "msgPost",
		})
	}
}
