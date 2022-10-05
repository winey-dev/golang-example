package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"html/data"
)

func main() {
	r := gin.Default()

	// Static 파일 Read 여기에 css, js, image 등이 올라간다.
	r.Static("/static", "./static/")

	// html templates 모음집이다.
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "HOME"})
		return
	})

	r.GET("/user", func(c *gin.Context) {
		users, err := data.User().GetUserList()
		if err != nil {
			c.HTML(http.StatusNotFound, "error.html", gin.H{"title": "정보 조회 실패", "error": ""})
		} else {
			c.HTML(http.StatusOK, "users.html", gin.H{"title": "유저 정보 조회", "users": users})
		}
		return
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := data.User().GetUser(id)
		if err != nil {
			c.HTML(http.StatusNotFound, "error.html", gin.H{"title": id + "정보 조회 실패", "error": ""})
		} else {
			c.HTML(http.StatusOK, "user.html", gin.H{"title": id + "정보 조회", "user": user})
		}
		return
	})

	r.Run()
}
