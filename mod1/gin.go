package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")

	group := r.Group("/api")
	{ // http://localhost:6066/api/hello
		group.GET("/hell", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world!")
		})
	}
	viwGroup := r.Group("/view")
	{ // http://localhost:6066/view/index
		viwGroup.GET("/index", func(c *gin.Context) {
			h := gin.H{
				"title": "test title",
				"body":  "body test",
			}
			c.HTML(http.StatusOK, "index.html", h)
		})
		// http://localhost:6066/view/json
		viwGroup.GET("/json", func(c *gin.Context) {
			c.JSON(http.StatusOK, user{ID: 11, Name: "test", Age: 23})
		})
	}

	r.Run(":6066")
}
