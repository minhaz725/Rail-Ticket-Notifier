package notifier

import (
	"Rail-Ticket-Notifier/internal/arguments"
	"Rail-Ticket-Notifier/utils/constants"
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail(messageBody string) bool {
	//Sender data.

	// Receiver email address.
	to := []string{
		arguments.RECEIVER_EMAIL_ADDRESS,
		constants.OWNER_EMAIL_ADDRESS,
	}
	//smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	mail := generateMail(messageBody, to)
	// Authentication.
	auth := smtp.PlainAuth("", constants.SENDER_EMAIL_ADDRESS, constants.SENDER_EMAIL_PASSWORD, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, constants.SENDER_EMAIL_ADDRESS, to, []byte(mail))
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent Successfully!")
	////makeCall()
	return true
}

func MakeCall() bool {

	//urlTimu := "https://e83c-103-243-82-92.ngrok-free.app/call/timu"
	//
	//// Make a GET request to the specified URL
	//_, err := http.Get(urlTimu)
	//if err != nil {
	//	fmt.Println("Error making GET request:", err)
	//	return false
	//} else {
	//	fmt.Println("call made successfully")
	//}
	//
	//urlMuna := "https://e83c-103-243-82-92.ngrok-free.app/call/muna"
	//
	//// Make a GET request to the specified URL
	//_, err = http.Get(urlMuna)
	//if err != nil {
	//	fmt.Println("Error making GET request:", err)
	//	return false
	//} else {
	//	fmt.Println("call made successfully")
	//}
	return true
}

func generateMail(messageBody string, to []string) string {
	// Message.
	msg := "From: " + constants.SENDER_EMAIL_NAME + " <" + arguments.FROM + ">\r\n"
	msg += "To: " + strings.Join(to, ";") + "\r\n"
	msg += "Subject: Available Tickets on " + arguments.DATE + "\r\n"
	msg += "\r\n" + messageBody
	return msg
}
