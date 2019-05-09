package utils

import (
	"net/http"
	"io/ioutil"
	"github.com/guoruibiao/commands"
	"github.com/guoruibiao/kindlefeeder/errors"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func WebGetContents(url string)(string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
    bytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
    	return "", err
	}
    return string(bytes), nil
}

func FileGetContents(filepath string)(string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PDFGenerator(urls []string, output string) (bool, error) {
	commander := commands.New()
	urls = append(urls, output)
	ok, output := commander.GetStatusOutput("wkhtmltopdf", urls...)
	if ok == false {
		return false, &errors.CMDError{Name: output,}
	}else{
		return true, nil
	}
}


func SendPDFToKindle() {
	e := email.NewEmail()
	e.From = "spidersmall@163.com"
	e.To = []string{"8615801479216@kindle.cn"}
	e.Subject = "Awesome PDF sending test"
	e.Text = []byte("test send pdf to kindle.")
	e.HTML = []byte("<h1>test send pdf to kindle.</h1>")
	e.AttachFile("/tmp/baidu_20183817.pdf")
	e.Send("smtp.163.cpm", smtp.PlainAuth("", "spidersmall@163.com", "guoruibiao285514", "smtp.163.com"))
}