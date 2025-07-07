package admin_app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusgomes28/urchin/common"
	"github.com/matheusgomes28/urchin/database"
)

func SetupRoutes(app_settings common.AppSettings, database database.Database) *gin.Engine {

	r := gin.Default()
	r.MaxMultipartMemory = 1

	r.Use(CORSMiddleware())

	r.GET("/posts/:id", getPostHandler(database))
	r.POST("/posts", postPostHandler(database))
	r.PUT("/posts", putPostHandler(database))
	r.DELETE("/posts/:id", deletePostHandler(database))

	r.POST("/images", postImageHandler(app_settings))
	r.DELETE("/images/:name", deleteImageHandler(app_settings))

	r.POST("/pages", postPageHandler(database))

	r.POST("/permalinks/:permalink/:post_id", postPermalinkHandler(database))

	// For container health purposes
	r.Any("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, PostIdResponse{Id: 0})
	})

	return r
}

// CORS middleware function
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
