package handler

import (
	"database/sql"
	"learn-go/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed *newsfeed.Repo, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM Item")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		defer rows.Close()
		c.JSON(http.StatusOK, rows)
	}
}
