package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	t, _ := LoadTemplate()
	router.SetHTMLTemplate(t)
	router.Static("/static", "./static")

	router.GET("", SearchHtml)
	router.POST("/search", SearchData)

	_ = router.Run(":18888")
}
