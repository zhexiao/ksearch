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

func TestReadUrlTmp(t *testing.T) {
	ks := initData()
	tmp, err := ks.readUrlTmp()
	if err != nil {
		t.Errorf("读取模板失败err=%s.", err)
	}

	if strings.Compare(tmp, "http://www.deyi.com/forum-912-%d.html") != 0 {
		t.Errorf("模板匹配失败=%s.", tmp)
	}
}

func TestValidateParams(t *testing.T) {
	ks := initData()
	err := ks.validateParams()
	if err != nil {
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
