package main

import (
	"database/sql"
	"fmt"
	"learn-go/handler"
	"learn-go/platform/newsfeed"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello There")
	feed := newsfeed.New()
	r := gin.Default()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed, db))
	r.POST("/newsfeed", handler.NewsfeedPost(feed, db))
	defer db.Close()
	r.Run()

}
