package spider

import "testing"


func TestSpider_FetchCSDN(t *testing.T) {
	spider, err := New()
	if err != nil {
		t.Error(err)
	}
	url := "https://guoruibiao.blog.csdn.net"
	link, err := spider.FetchCSDNPage(url)
	if err != nil {
		t.Error(err)
	}else{
		t.Log(link)
	}
}
