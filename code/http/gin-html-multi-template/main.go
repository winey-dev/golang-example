package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"

	"gin-html-multi-template/data"
)

func loadTemplates(baseDir string) multitemplate.Render {

	r := multitemplate.New()

	layouts, err := filepath.Glob(baseDir + "/layouts/*tmpl")
	if err != nil {
		fmt.Printf("layouts templates load failed. err=%v\n", err)
		os.Exit(1)
	}

	pages, err := filepath.Glob(baseDir + "/pages/*tmpl")
	if err != nil {
		fmt.Printf("pages templates load failed. err=%v\n", err)
		os.Exit(1)
	}

	for _, page := range pages {
		// 만들어진 하나의 page에 layout에 읽혀진 데이터를 입히는 구조
		/*
		   layout1
		   layout2
		   page
		   layout3

		   이런 형태로 반복문을 돌리기 위해
		*/
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, page)
		fmt.Println(files, filepath.Base(page))
		r.AddFromFiles(filepath.Base(page), files...)
	}

	return r

}

func main() {
	r := gin.Default()

	// Static 파일 Read 여기에 css, js, image 등이 올라간다.
	r.Static("/static", "./static/")
	r.HTMLRender = loadTemplates("./templates")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "HOME"})
		return
	})

	r.GET("/user", func(c *gin.Context) {
		users, err := data.User().GetUserList()
		if err != nil {
			c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"title": "정보 조회 실패", "error": ""})
		} else {
			c.HTML(http.StatusOK, "users.tmpl", gin.H{"title": "유저 정보 조회", "users": users})
		}
		return
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := data.User().GetUser(id)
		if err != nil {
			c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"title": id + "정보 조회 실패", "error": ""})
		} else {
			c.HTML(http.StatusOK, "user.tmpl", gin.H{"title": id + "정보 조회", "user": user})
		}
		return
	})

	r.Run()
}
