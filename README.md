# kindle feeder

1. download web blogs
2. translate to PDF (based on `wkhtmltopdf`)
3. feed my kindle by sending mails

CSDN加了个"阅读更多"真的是，对CSDN实现了个大概，结果发现因为这个按钮的存在，没什么价值了。

```
 oooO ↘┏━┓ ↙ Oooo  ( 踩)→┃你┃ ←(死 )   \ ( →┃√┃ ← ) / 　 \_)↗┗━┛ ↖(_/
```

测试了下简书和其他几个网站，生成结果还不错。

go 发送email不是很方便，这里还是用之前写好的python脚本来发邮件。
```
python sendmail.py --help
usage: sendmail.py [-h] [-f FILE] [-r RECEIVERS [RECEIVERS ...]]
                   [-a ATTACHMENT] [-t TEXT] [-s SUBJECT]

命令行邮件发送工具

optional arguments:
  -h, --help            show this help message and exit
  -f FILE, --file FILE  .mailerrc配置文件的位置
  -r RECEIVERS [RECEIVERS ...], --receivers RECEIVERS [RECEIVERS ...]
                        收件人列表，注意不能少于一个，而且多
                        个收件人要用空格隔开
  -a ATTACHMENT, --attachment ATTACHMENT
                        附件的全路径，注意windows和linux上会稍有
                        不同~,
                        多个附件的时候需要重复指定此参数~
  -t TEXT, --text TEXT  邮件正文部分，相当于纯文本模式
  -s SUBJECT, --subject SUBJECT
                        邮件主题
```

其中的`-f` 参数存储的是邮箱的设置，内容参考`./.mailerrc` 作为脚本运行时的动态读取。

本来想加点UI的，但是CSDN这个"阅读更多"让我瞬间没了做下去的兴趣了，后面搞不搞再看时间吧。