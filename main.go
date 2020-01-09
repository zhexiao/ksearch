package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhexiao/ksearch/search"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	//ks := search.NewCT_Ksearch()
	//ks.Page2Url = "http://www.deyi.com/forum-912-2.html"
	//ks.Page3Url = "http://www.deyi.com/forum-912-3.html"
	//ks.Keyword = "光谷东"

	router := gin.Default()

	t, _ := loadTemplate()
	router.SetHTMLTemplate(t)

	router.Static("/static", "./static")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/templates/search.html", gin.H{})
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

//把templates到go文件，参考如下
//https://github.com/gin-gonic/examples/tree/master/assets-in-binary
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
