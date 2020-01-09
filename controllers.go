package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SearchHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "/static/search.html", gin.H{})
}

func SearchData(c *gin.Context) {
	//ks := search.NewCT_Ksearch()
	//ks.Page2Url = "http://www.deyi.com/forum-912-2.html"
	//ks.Page3Url = "http://www.deyi.com/forum-912-3.html"
	//ks.Keyword = "光谷东"

	url1 := c.PostForm("url1")
	url2 := c.PostForm("url2")
	keyword := c.PostForm("keyword")
	startPageStr := c.PostForm("start_page")
	endPageStr := c.PostForm("end_page")

	startPage, _ := strconv.Atoi(startPageStr)
	endPage, _ := strconv.Atoi(endPageStr)

	ks := CT_Ksearch{
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
}
