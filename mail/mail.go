package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

type Mail struct {
	Username string
	Password string
	Addr     string
	Port     string
}

func (m *Mail) SendMail(sendTo string, subject string, message string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Addr)

	to := []string{sendTo}
	// msg := []byte(fmt.Sprintf("To: %s\r\n", sendTo) +
	// 	fmt.Sprintf("Subject: %s\r\n", subject) +
	// 	"\r\n" +
	// 	string(message))

	addr := fmt.Sprintf("%s:%s", m.Addr, m.Port)
	err := smtp.SendMail(addr, auth, "niklas.thomas@hotmail.com", to, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
