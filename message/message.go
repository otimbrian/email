package message

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(message *gomail.Message, senderEmail string) (err error) {
	messageDialer := gomail.NewDialer("smtp.gmail.com", 465, senderEmail, "cayjcnjxggixwzwq")

	err = messageDialer.DialAndSend(message)
	if err != nil {
		return err
	}
	return
}

func AttachFile(messg *gomail.Message, filepath, newName string) *gomail.Message {
	if newName != "" {
		messg.Attach(filepath, gomail.Rename(newName))
	}
	messg.Attach(filepath)
	return messg
}
func CreateMessageBody(messg *gomail.Message, subject, body string) *gomail.Message {
	messg.SetHeader("subject", subject)
	messg.SetBody("text/html", body)

	return messg
}

func ReciepientAndUser(messg *gomail.Message, senderEmail string, ReceivingEmails ...string) {

	messg.SetHeader("From", senderEmail)
	messg.SetHeader("To", ReceivingEmails...)
}
