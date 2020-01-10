package main

import (
	"strings"
	"testing"
)

func initData() *CT_Ksearch {
	ks := CT_Ksearch{
		Page2Url:  "http://www.deyi.com/forum-912-2.html",
		Page3Url:  "http://www.deyi.com/forum-912-3.html",
		StartPage: 1,
		EndPage:   2,
		Keyword:   "得意",
	}

	return &ks
}

func initData1() *CT_Ksearch {
	ks := NewCT_Ksearch()
	ks.Page2Url = "http://gaoloumi.cc/forum.php?mod=forumdisplay&fid=40&page=2"
	ks.Page3Url = "http://gaoloumi.cc/forum.php?mod=forumdisplay&fid=40&page=3"
	ks.StartPage = 1
	ks.EndPage = 2
	ks.Keyword = "武汉"

	return ks
}

func TestReadUrlTmp(t *testing.T) {
	ks := initData()
	tmp, err := ks.readUrlTmp()
	if err != nil {
		t.Errorf("读取模板失败err=%s.", err)
	}

	if strings.Compare(tmp, "http://www.deyi.com/forum-912-%d.html") != 0 {
		t.Errorf("模板匹配失败=%s.", tmp)
	}

	ks.Page2Url = "abc.com"
	ks.Page3Url = "bcd.com"
	tmp, err = ks.readUrlTmp()
	if err == nil {
		t.Errorf("数据检验失败 err=%s.", err)
	}

	ks.Page2Url = "abcd.com"
	ks.Page3Url = "abcdefg.com"
	tmp, err = ks.readUrlTmp()
	if err == nil {
		t.Errorf("数据检验失败 err=%s.", err)
	}
}

func TestValidateParams(t *testing.T) {
	ks := initData()
	err := ks.validateParams()
	if err != nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	ks.StartPage = 0
	ks.EndPage = 0
	err = ks.validateParams()
	if err != nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	if ks.StartPage != 1 || ks.EndPage != 2 {
		t.Errorf("参数值初始化验证失败 err=%s.", err)
	}

	ks.StartPage = 2
	ks.EndPage = 1
	err = ks.validateParams()
	if err == nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	ks.StartPage = 2
	ks.EndPage = 100
	err = ks.validateParams()
	if err == nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	ks.Page2Url = ""
	ks.Page3Url = ""
	err = ks.validateParams()
	if err == nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	ks.Keyword = ""
	err = ks.validateParams()
	if err == nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}
}

func TestReadUrlData(t *testing.T) {
	ks := initData()
	err := ks.readUrlData(ks.Page3Url)
	if err != nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	if len(ks.Data) <= 0 {
		t.Errorf("没有搜索到数据，有疑问 err=%s.", ks.Keyword)
	}
}

func TestSearch(t *testing.T) {
	ks := initData()
	err := ks.Search()
	if err != nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	if len(ks.Data) <= 0 {
		t.Errorf("没有搜索到数据，有疑问 err=%s.", ks.Keyword)
	}
}

func TestUriAddDomain(t *testing.T) {
	respUrl := "forum.php?mod=viewthread&tid=885674&extra=page%3D1"

	ks := initData1()
	realLink, err := ks.fillLink(respUrl)
	if err != nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	if strings.Compare(respUrl, realLink) == 0 {
		t.Errorf("地址应该被转换，请检查")
	}
}

func TestUriNotAddDomain(t *testing.T) {
	respUrl := "http://www.deyi.com/thread-10363276-1-1.html"

	ks := initData()
	realLink, err := ks.fillLink(respUrl)
	if err != nil {
		t.Errorf("参数验证失败 err=%s.", err)
	}

	if strings.Compare(respUrl, realLink) != 0 {
		t.Errorf("地址不应该被转换，请检查")
	}
}

func TestKsearchData(t *testing.T) {
	dt := NewCT_KSearchData()
	dt.Url = ""
	dt.Title = ""
}
