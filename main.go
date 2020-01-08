package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhexiao/ksearch/search"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//ks := search.NewCT_Ksearch()
	//ks.Page2Url = "http://www.deyi.com/forum-912-2.html"
	//ks.Page3Url = "http://www.deyi.com/forum-912-3.html"
	//ks.SearchPage = 30
	//ks.Keyword = "光谷东"

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.html", gin.H{})
	})
	router.POST("/search", func(c *gin.Context) {
		url1 := c.PostForm("url1")
		url2 := c.PostForm("url2")
		keyword := c.PostForm("keyword")
		searchPage := c.PostForm("search_page")

		page, _ := strconv.Atoi(searchPage)
		ks := search.CT_Ksearch{
			Page2Url:   url1,
			Page3Url:   url2,
			Keyword:    keyword,
			SearchPage: page,
		}

		//搜索
		if err := ks.Search(); err != nil {
			log.Panicln(err)
		}

		c.JSON(http.StatusOK, ks.Data)
	})

	if err := router.Run(":8080"); err != nil {
		log.Printf("运行出错,err=%s.", err)
	}
}
