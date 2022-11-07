package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

// golang自己的net/smtp其实也可以发邮件，但是比较麻烦，使用一个第三方库email来发送，这个本就是对net/smtp的封装

const (
	USERNAME_QQ = "1824517828@qq.com"
	PASSWORD_QQ = "mathtfnstavleggi"
	HOST_QQ     = "smtp.qq.com"
	PORT_QQ     = 587
)

// SendEmail 发送邮件
func SendEmail(subject string, from string, to []string, cc []string, content string) {
	sendText(subject, from, to, cc, nil, content)
}

// SendEmailHtml 发送邮件
func SendEmailHtml(subject string, from string, to []string, cc []string, htmlStr string) {
	sendHtml(subject, from, to, cc, nil, htmlStr)
}

func sendText(subject string, from string, to []string, cc []string, bcc []string, content string) {
	log.Print("开始发送邮箱")
	e := email.NewEmail()
	if from != "" {
		e.From = from
	}

	if to != nil && len(to) > 0 {
		e.To = to
	}

	if bcc != nil && len(bcc) > 0 {
		e.Bcc = bcc
	}

	if cc != nil && len(cc) > 0 {
		e.Cc = cc
	}

	if subject != "" {
		e.Subject = subject
	}

	if content != "" {
		e.Text = []byte(content)
	}
	//e.HTML = []byte("<h1> html 邮件内容！</h1>")
	err := e.Send(fmt.Sprintf("%v:%v", HOST_QQ, PORT_QQ), smtp.PlainAuth("", USERNAME_QQ, PASSWORD_QQ, HOST_QQ))
	if err != nil {
		log.Print(fmt.Sprint(err.Error()))
		return
	}
	log.Print("结束发送邮箱")
}

func sendHtml(subject string, from string, to []string, cc []string, bcc []string, htmlStr string) {
	log.Print("开始发送邮箱")
	e := email.NewEmail()
	if from != "" {
		e.From = from
	}

	if to != nil && len(to) > 0 {
		e.To = to
	}

	if bcc != nil && len(bcc) > 0 {
		e.Bcc = bcc
	}

	if cc != nil && len(cc) > 0 {
		e.Cc = cc
	}

	if subject != "" {
		e.Subject = subject
	}

	if htmlStr != "" {
		e.HTML = []byte(htmlStr)
	}
	err := e.Send(fmt.Sprintf("%v:%v", HOST_QQ, PORT_QQ), smtp.PlainAuth("", USERNAME_QQ, PASSWORD_QQ, HOST_QQ))
	if err != nil {
		log.Print(fmt.Sprint(err.Error()))
		return
	}
	log.Print("结束发送邮箱")
}
