package main

import (
	"github.com/guoruibiao/commands"
	"fmt"
	"github.com/guoruibiao/kindlefeeder/spider"
	"github.com/guoruibiao/kindlefeeder/utils"
)

func main() {
	spider, _ := spider.New()
	//links, err:=spider.FetchCSDNAuthor("marksinoberg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	urls := []string{
		"http://www.voidcn.com/article/p-zgwmiipi-bay.html",
		"https://www.jianshu.com/p/ba7852fbcc80",
	}
	links,_ := spider.FetchURLs(urls)
	fmt.Println(len(links))
	utils.PDFGenerator(links, "/tmp/baidu_20183817.pdf")
	//
	//utils.SendPDFToKindle()
	commander := commands.New()
	status, output := commander.GetStatusOutput("python", "./sendmail.py", "-f", "./.mailerrc", "-r", "8615801479216_20efef@kindle.cn", "-a", "/tmp/baidu_20183817.pdf", "-t", "content", "-s", "Awesome kindle feeder")
	if status == false{
		fmt.Println(output)
	}else{
		fmt.Println("发送附件成功", output)
	}
}
