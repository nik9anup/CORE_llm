/*
Package main demonstrates how to send an email using SMTP with authentication and TLS encryption in Go.

This program prompts the user for SMTP credentials (username and password), establishes a connection to Gmail's SMTP server,
and sends a simple email message using the provided credentials.
*/
package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func main() {
	// Prompt the user to enter SMTP username (email address).
	var email string
	fmt.Println("Enter username for smtp: ")
	fmt.Scanln(&email)

	// Prompt the user to enter SMTP password.
	var pass string
	fmt.Println("Enter password for smtp: ")
	fmt.Scanln(&pass)

	// Authenticate using PlainAuth with the provided credentials.
	auth := smtp.PlainAuth("",
		email,
		pass,
		"smtp.gmail.com")

	// Connect to Gmail's SMTP server on port 587 (TLS encryption).
	c, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Enable TLS encryption on the SMTP connection.
	config := &tls.Config{ServerName: "smtp.gmail.com"}
	if err = c.StartTLS(config); err != nil {
		panic(err)
	}

	// Authenticate with the server using the provided credentials.
	if err = c.Auth(auth); err != nil {
		panic(err)
	}

	// Set the sender's email address.
	if err = c.Mail(email); err != nil {
		panic(err)
	}

	// Set the recipient's email address (same as sender in this example).
	if err = c.Rcpt(email); err != nil {
		panic(err)
	}

	// Open a data connection to send the email content.
	w, err := c.Data()
	if err != nil {
		panic(err)
	}

	// Define the email message content (in this case, a simple text message).
	msg := []byte("Hello, this is the email content")

	// Write the message content to the data connection.
	if _, err := w.Write(msg); err != nil {
		panic(err)
	}

	// Close the data connection.
	err = w.Close()
	if err != nil {
		panic(err)
	}

	// Quit the SMTP session.
	err = c.Quit()
	if err != nil {
		panic(err)
	}
}