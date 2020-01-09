package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhexiao/ksearch/search"
	"net/http"
	"strconv"
)

func main() {
	//ks := search.NewCT_Ksearch()
	//ks.Page2Url = "http://www.deyi.com/forum-912-2.html"
	//ks.Page3Url = "http://www.deyi.com/forum-912-3.html"
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
		startPageStr := c.PostForm("start_page")
		endPageStr := c.PostForm("end_page")

		startPage, _ := strconv.Atoi(startPageStr)
		endPage, _ := strconv.Atoi(endPageStr)

		ks := search.CT_Ksearch{
			Page2Url:  url1,
			Page3Url:  url2,
			Keyword:   keyword,
			StartPage: startPage,
			EndPage:   endPage,
		}

		//搜索
		if err := ks.Search(); err != nil {
			c.JSON(http.StatusOK, err)
			return
		}

		c.JSON(http.StatusOK, ks.Data)
	})

	_ = router.Run(":18888")
}
