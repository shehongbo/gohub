package mail

import (
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"gohub/pkg/logger"
	"net/smtp"
)

// STMP 实现 email.Driver infterface
type STMP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s STMP) Send(email Email, config map[string]string) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发件详情", e)

	err := e.Send(fmt.Sprintf("%v,%v", config["host"], config["port"]),
		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)
	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
	}
	logger.DebugString("发送邮件", "发送成功", "")
	return true
}
