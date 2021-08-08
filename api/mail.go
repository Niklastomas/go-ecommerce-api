package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/niklastomas/go-ecommerce-api/mail"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

type MailBody struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (s *Server) SendMail(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("SENDGRID_USERNAME")
	password := os.Getenv("SENDGRID_API_KEY")
	server := os.Getenv("SENDGRID_SERVER")
	port := os.Getenv("SENDGRID_PORT")

	m := mail.Mail{Username: username, Password: password, Addr: server, Port: port}

	var newMail *MailBody
	json.NewDecoder(r.Body).Decode(&newMail)

	dir, _ := os.Getwd()
	file := fmt.Sprintf("%s/templates/email_test.html", dir)
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
		return
	}

	var body bytes.Buffer

	// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))
	err = t.Execute(&body, newMail)
	if err != nil {
		log.Println(err)
		return
	}

	header := make(map[string]string)
	header["From"] = "niklas.thomas@hotmail.com"
	header["To"] = newMail.To
	header["Subject"] = "test"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(body.Bytes())

	err = m.SendMail(newMail.To, newMail.Subject, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.JSON(w, r, "Success", http.StatusOK)
}
