package util

import (
	"crypto/tls"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"gopkg.in/gomail.v2"
)

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

type Email struct {
	*SMTPInfo
}

func New(info *SMTPInfo) *Email {
	return &Email{info}
}

func (e *Email) Send(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: !e.IsSSL}
	var err error
	if err = dialer.DialAndSend(m); err != nil {
		logging.Errorf("邮件发送失败：%v; %#v", err, e.SMTPInfo)
	}
	return err
}

func _default() *Email {
	return &Email{&SMTPInfo{
		Host:     global.Setting.Email.Host,
		Port:     global.Setting.Email.Port,
		IsSSL:    global.Setting.Email.IsSSL,
		UserName: global.Setting.Email.UserName,
		Password: global.Setting.Email.Password,
		From:     global.Setting.Email.From,
	}}
}

func SendEmail(to []string, subject, body string) error {
	return _default().Send(to, subject, body)
}
