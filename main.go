package main

import (
	"fmt"
	"ksearch/search"
)

func main() {
	ks := search.NewCT_Ksearch()
	ks.Page2Url = "http://www.deyi.com/forum-912-2.html"
	ks.Page3Url = "http://www.deyi.com/forum-912-3.html"
	ks.SearchPage = 10
	ks.Keyword = "光谷东"

	ks.Search()

	for _, d := range ks.Data {
		fmt.Println(*d)
	}
}
