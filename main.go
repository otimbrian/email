package main

import (
	"emailgo/message"

	"gopkg.in/gomail.v2"
)

func main() {
	password := ""
	m := gomail.NewMessage()

	message.AddReciepientAndUserToheader(m, "otims.brian@gmail.com", "brianotims56@gmail.com")

	message.CreateMessageBody(m, "Hello Brian", "Hello Message to you my friend")

	message.SendEmail(m, "otims.brian@gmail.com", password)

}
