package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/sp_covid_impact/pkg/handler/data"
)

// Routes struct
type Routes struct {
}

// StartGin starts the router
func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", welcome)
		api.GET("/data", data.GetAllData)
	}
	r.Run(":8000")
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to the Spain Covid Impact api",
	})
	return
}
