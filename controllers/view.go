package controllers

import (
	"chinmay-sawant/Spider-Query/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewsController struct {
	dblist services.Dblist
}

func (vc *ViewsController) InitViewController(router *gin.Engine, dblist services.Dblist) {
	vc.dblist = dblist
	loginRoutes := router.Group("/view")
	loginRoutes.GET("/", vc.DbListView())

}

func (vc *ViewsController) DbListView() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prepare data to be passed to the template (replace with your actual data)
		// We are passing the actual data using FETCHAPI in frontend only by caling /sw endpint
		data := map[string]interface{}{
			"Title":   "Database List",
			"DBNames": []string{"DB1", "DB2", "DB3"},
		}

		// Render the template with the data
		c.HTML(http.StatusOK, "dblist.html", data)
	}
}
