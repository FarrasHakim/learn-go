package handler

import (
	"database/sql"
	"fmt"
	"learn-go/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// type errorResponse struct {
// 	err
// }

func NewsfeedPost(feed *newsfeed.Repo, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)

		result, err := db.Exec("INSERT INTO Item (title, post) VALUES (?, ?)", requestBody.Title, requestBody.Post)
		fmt.Println(result)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		// item := newsfeed.Item{
		// 	Title: requestBody.Title,
		// 	Post: requestBody.Post,
		// }
		// feed.Add(item)

		c.JSON(http.StatusOK, result)
	}
}
