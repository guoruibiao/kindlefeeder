package utils

import "testing"

// 测试 网页链接
func Test_WebGetContents(t *testing.T) {
	url := "http://www.douban.com"
	content, err := WebGetContents(url)
	if err != nil {
		t.Error(err)
	}else{
		t.Log(content)
	}
}

// 测试获取 存在的文件  的内容
func Test_FileGetContents_Exists(t *testing.T) {
	filepath := "/tmp/hello"
	content, err := FileGetContents(filepath)
	if err != nil {
		t.Error(err)
	}else{
		t.Log(content)
	}
}

// 测试 获取不存在的文件内容
func Test_FileGetContents_NotExists(t *testing.T) {
	filepath := "/tmp/notexists"
	content, err := FileGetContents(filepath)
	if err != nil {
		t.Log(err.Error())
	}else{
		t.Log(content)
	}
}


// 测试调用外部工具生成PDF
func TestPDFGenerator(t *testing.T) {
	urls := []string{
		"http://www.baidu.com",
		"http://www.douban.com",
		"http://github.com",
	}
	ok, errmsg := PDFGenerator(urls, "/tmp/demo.pdf")
	if ok == false {
		t.Error(errmsg)
	}else{
		t.Log(errmsg)
	}
}