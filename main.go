package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"strings"
)

func main() {
	router := gin.Default()

	t, _ := loadTemplate()
	router.SetHTMLTemplate(t)
	router.Static("/static", "./static")

	router.GET("", SearchHtml)
	router.POST("/search", SearchData)

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
