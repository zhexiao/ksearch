package main

import (
	"fmt"
	"github.com/zhexiao/ksearch/search"
	"log"
)

func main() {
	ks := search.NewCT_Ksearch()
	ks.Page2Url = "http://www.deyi.com/forum-912-2.html"
	ks.Page3Url = "http://www.deyi.com/forum-912-3.html"
	ks.SearchPage = 30
	ks.Keyword = "光谷东"

	//搜索
	if err := ks.Search(); err != nil {
		log.Panicln(err)
	}

	for _, d := range ks.Data {
		fmt.Println(*d)
	}
}
