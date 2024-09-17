package main

import (
	"chinmay-sawant/Spider-Query/controllers"
	"chinmay-sawant/Spider-Query/internal"
	"chinmay-sawant/Spider-Query/middleware"
	"chinmay-sawant/Spider-Query/model"
	"chinmay-sawant/Spider-Query/services"

	"github.com/gin-gonic/gin"
)

func main() {
	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:9999"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           3600, // 1 hour
	}
	router := gin.Default()
	router.Use(middleware.CORSMiddleware(corsConfig))
	router.LoadHTMLGlob("templates/*")
	dbs := model.Dbs{Host: "localhost", User: "sc", Password: "sc", Dbname: "sc", Port: "5432"}
	db := internal.InitDB(dbs)
	if db == nil {
		// error connecting to DB
	}

	dbservice := services.Dblist{}
	dbservice.InitService(db)
	dbservice.DefaultRow()
	cachedDbList := dbservice.GetAllDbList()
	router.Use(middleware.BasicAuthMiddleware("u", "p"))

	// Controllers
	LoginController := &controllers.LoginController{}
	ViewController := &controllers.ViewsController{}
	SpiderWebController := &controllers.SpiderWebController{}

	// Init Controllers
	LoginController.InitLoginControllerRouter(router)
	ViewController.InitViewController(router, dbservice)
	SpiderWebController.InitSpiderWebControllerRouter(router, dbservice, cachedDbList)
	router.Run(":9999")
}
