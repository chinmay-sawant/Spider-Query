package controllers

import (
	"chinmay-sawant/Spider-Query/model"
	"chinmay-sawant/Spider-Query/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SpiderWebController struct {
	dblist    services.Dblist
	cacheList []model.Dbs
}

func (sw *SpiderWebController) InitSpiderWebControllerRouter(router *gin.Engine, dblist services.Dblist, cacheList []model.Dbs) {
	sw.dblist = dblist
	sw.cacheList = cacheList
	loginRoutes := router.Group("/sw")
	loginRoutes.GET("/", sw.DbList())
	loginRoutes.POST("/", sw.fireQuery())
}

func (sw *SpiderWebController) DbList() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sw.cacheList == nil {
			sw.cacheList = sw.dblist.GetAllDbList()
			for _, v := range sw.cacheList {
				fmt.Println(v)
			}
		} else {
			for _, v := range sw.cacheList {
				fmt.Println(v)
			}
		}
		c.JSON(200, gin.H{
			"results": sw.cacheList,
		})
	}
}

func (sw *SpiderWebController) fireQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqDbs model.ReqDbs
		c.BindJSON(&reqDbs)
		c.JSON(200, gin.H{
			"results": sw.dblist.DynamicQuery(reqDbs.GetQuery()),
		})
	}
}
