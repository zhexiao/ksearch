package main

import (
	"html/template"
	"io/ioutil"
	"strings"
)

//把templates到go文件，参考如下
//https://github.com/gin-gonic/examples/tree/master/assets-in-binary
func LoadTemplate() (*template.Template, error) {
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
