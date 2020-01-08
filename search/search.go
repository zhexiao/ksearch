package search

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"sync"
)

type CT_Ksearch struct {
	//第二页与第三页的地址
	Page2Url string
	Page3Url string

	//搜索的关键字
	Keyword string

	//最大支持搜索的页数和需要搜索的页数
	SearchPage int

	//存储的数据
	Data []*CT_KSearchData

	//去重相同的地址
	duplicateUri []string
}

func NewCT_Ksearch() *CT_Ksearch {
	return &CT_Ksearch{}
}

// 搜索数据
func (k *CT_Ksearch) Search() {
	//验证参数
	if err := k.validateParams(); err != nil {
		log.Panicf("参数验证有误,err=%s。", err)
	}

	//读取模板
	urlTmp, err := k.readUrlTmp()
	if err != nil {
		log.Panicf("读取模板失败,err=%s。", err)
	}

	//开始搜索
	var wg sync.WaitGroup
	for i := 1; i < k.SearchPage; i++ {
		wg.Add(1)

		//生成真实的搜索地址
		realUrl := fmt.Sprintf(urlTmp, i)
		go func(url string) {
			defer wg.Done()

			if err := k.readUrlData(url); err != nil {
				log.Panicf("读取数据失败,err=%s。", err)
			}
		}(realUrl)
	}

	wg.Wait()
}

//根据第二页和第三页，找到此地址的分页模板
func (k *CT_Ksearch) readUrlTmp() (string, error) {
	var posList []int

	url := k.Page2Url
	url2 := k.Page3Url

	if len(url) != len(url2) {
		return "", NewKsError(1, "第二页与第三页的地址长度不一致，无法找到分页模板。")
	}

	for pos := 0; pos < len(url); pos++ {
		if url[pos] != url2[pos] {
			posList = append(posList, pos)
		}
	}

	if len(posList) != 1 && len(posList) > 0 {
		return "", NewKsError(1, "地址有多处不一致，无法定位分页位置。")
	}

	//找到分页位置替换为 %d
	urlTmp := fmt.Sprintf("%s%s%s", url[:posList[0]], "%d", url[(posList[0]+1):])

	return urlTmp, nil
}

//使用爬虫读取数据
func (k *CT_Ksearch) readUrlData(url string) error {
	//实例化爬虫
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"),
		colly.MaxDepth(1),
	)

	//读取href数据
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		text := strings.Trim(e.Text, " ")

		//保存数据
		k.saveData(link, text)
	})

	//请求数据
	if err := c.Visit(url); err != nil {
		return err
	}

	return nil
}

//保存数据
func (k *CT_Ksearch) saveData(url string, title string) {
	for _, u := range k.duplicateUri {
		//如果地址已经存在，则返回
		if strings.Compare(u, url) == 0 {
			return
		}
	}

	//如果关键字存在
	if strings.Contains(title, k.Keyword) {
		tmpData := &CT_KSearchData{
			Url:   url,
			Title: title,
		}
		k.Data = append(k.Data, tmpData)
		k.duplicateUri = append(k.duplicateUri, url)
	}
}

//验证搜索参数是否合法
func (k *CT_Ksearch) validateParams() error {
	//验证搜索的页数
	if k.SearchPage <= 0 {
		k.SearchPage = 10
	}

	if k.SearchPage > 20 {
		return NewKsError(2, "最大支持一次性搜索20页的数据。")
	}

	k.Page2Url = strings.Trim(k.Page2Url, " ")
	k.Page3Url = strings.Trim(k.Page3Url, " ")
	if k.Page2Url == "" || k.Page3Url == "" {
		return NewKsError(2, "第二页与第三页的链接地址不能为空。")
	}

	k.Keyword = strings.Trim(k.Keyword, " ")
	if k.Keyword == "" {
		return NewKsError(2, "搜索的关键字不能为空。")
	}

	return nil
}
