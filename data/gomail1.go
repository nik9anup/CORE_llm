// Example 1: Sending a basic email using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Hello from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent using go-mail library.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email: %v", err)
    }
    log.Println("Email sent successfully!")
}





// Example 2: Sending an email with HTML content using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("HTML Email from Go-Mail")
    email.SetBody(mail.TextHTML, "<html><body><h1>Hello</h1><p>This is a <b>test</b> email.</p></body></html>")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email: %v", err)
    }
    log.Println("HTML Email sent successfully!")
}





// Example 3: Sending an email with attachments using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Attachment from Go-Mail")
    email.SetBody(mail.TextPlain, "This email contains an attachment.")

    // Attach a file
    file, err := os.Open("attachment.pdf")
    if err != nil {
        log.Fatalf("Error opening attachment: %v", err)
    }
    defer file.Close()

    email.Attach(file, "attachment.pdf")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with attachment: %v", err)
    }
    log.Println("Email with attachment sent successfully!")
}





// Example 4: Sending an email with inline images using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Inline Image from Go-Mail")
    email.SetBody(mail.TextHTML, `<html><body><h1>Hello</h1><p>This is an email with an inline image: <img src="cid:image1"></p></body></html>`)

    // Attach inline image
    file, err := os.Open("image.jpg")
    if err != nil {
        log.Fatalf("Error opening image: %v", err)
    }
    defer file.Close()

    email.Attach(file, "image.jpg", mail.InlineFile("image1"))

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with inline image: %v", err)
    }
    log.Println("Email with inline image sent successfully!")
}





// Example 5: Sending an email using TLS encryption with go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client with TLS
    smtp := mail.NewSMTPWithTLS("smtp.example.com", 587, "username", "password", nil)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("TLS Encrypted Email from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent using TLS encryption.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending TLS encrypted email: %v", err)
    }
    log.Println("TLS encrypted email sent successfully!")
}





// Example 6: Sending an email using SSL with go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client with SSL
    smtp := mail.NewSMTPWithSSL("smtp.example.com", 465, "username", "password", nil)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("SSL Encrypted Email from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent using SSL encryption.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending SSL encrypted email: %v", err)
    }
    log.Println("SSL encrypted email sent successfully!")
}





// Example 7: Sending an email with custom headers using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with custom headers
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Custom Headers from Go-Mail")
    email.SetBody(mail.TextPlain, "This email includes custom headers.")
    email.AddHeader("X-Custom-Header", "Custom Value")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom headers: %v", err)
    }
    log.Println("Email with custom headers sent successfully!")
}





// Example 8: Sending an email using a template with go-mail library.

package main

import (
    "html/template"
    "log"
    "bytes"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Define email template
    const emailTemplate = `
        <html>
        <body>
            <h1>Hello, {{.Name}}!</h1>
            <p>This is a test email.</p>
        </body>
        </html>`

    // Prepare data for template
    data := struct {
        Name string
    }{
        Name: "John Doe",
    }

    // Render template
    tmpl := template.Must(template.New("emailTemplate").Parse(emailTemplate))
    var tpl bytes.Buffer
    err := tmpl.Execute(&tpl, data)
    if err != nil {
        log.Fatalf("Error executing template: %v", err)
    }

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Template from Go-Mail")
    email.SetBody(mail.TextHTML, tpl.String())

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with template: %v", err)
    }
    log.Println("Email with template sent successfully!")
}





// Example 9: Sending an email to multiple recipients using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient1@example.com", "recipient2@example.com")
    email.SetSubject("Email to Multiple Recipients from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent to multiple recipients.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email to multiple recipients: %v", err)
    }
    log.Println("Email sent to multiple recipients successfully!")
}





// Example 10: Sending an email with CC and BCC recipients using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.AddCC("cc1@example.com", "cc2@example.com")
    email.AddBCC("bcc1@example.com", "bcc2@example.com")
    email.SetSubject("Email with CC and BCC from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with CC and BCC.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with CC and BCC: %v", err)
    }
    log.Println("Email with CC and BCC sent successfully!")
}





// Example 11: Handling errors when sending an email using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Handling Errors Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email demonstrating error handling.")

    // Send email with error handling
    err := smtp.Send(email)
    if err != nil {
        log.Fatalf("Error sending email: %v", err)
    }
    log.Println("Email sent successfully!")
}





// Example 12: Sending an email with custom SMTP server configuration using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Custom SMTP server configuration
    smtpConfig := mail.SMTPConfig{
        Host:     "smtp.example.com",
        Port:     587,
        Username: "username",
        Password: "password",
        TLSConfig: &mail.TLSConfig{
            InsecureSkipVerify: true, // Example: Insecure skip verify; use carefully in production.
        },
    }

    // Initialize SMTP client with custom configuration
    smtp := mail.NewCustomSMTP(smtpConfig)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom SMTP Configuration Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with custom SMTP configuration.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom SMTP configuration: %v", err)
    }
    log.Println("Email sent with custom SMTP configuration successfully!")
}





// Example 13: Sending an email using OAuth2 authentication with go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // OAuth2 authentication configuration
    oauth2Config := &mail.OAuth2Config{
        Email:      "sender@example.com",
        ClientID:   "your_client_id",
        ClientSecret: "your_client_secret",
        AccessToken: "access_token",
        RefreshToken: "refresh_token",
    }

    // Initialize SMTP client with OAuth2 authentication
    smtp := mail.NewSMTPWithOAuth2("smtp.example.com", 587, oauth2Config)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("OAuth2 Authentication Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent using OAuth2 authentication.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with OAuth2 authentication: %v", err)
    }
    log.Println("Email sent using OAuth2 authentication successfully!")
}





// Example 14: Sending an email with embedded images using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Embedded Image from Go-Mail")
    email.SetBody(mail.TextHTML, `<html><body><h1>Hello</h1><p>This email includes an embedded image: <img src="cid:image1"></p></body></html>`)

    // Attach embedded image
    file, err := os.Open("image.jpg")
    if err != nil {
        log.Fatalf("Error opening image: %v", err)
    }
    defer file.Close()

    email.Attach(file, "image.jpg", mail.InlineFile("image1"))

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with embedded image: %v", err)
    }
    log.Println("Email with embedded image sent successfully!")
}





// Example 15: Sending an email with alternative text (plaintext and HTML) using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with alternative text (plaintext and HTML)
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Alternative Text from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with alternative plain text.")
    email.AddAlternative(mail.TextHTML, "<html><body><h1>Hello</h1><p>This is a test email with alternative HTML.</p></body></html>")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with alternative text: %v", err)
    }
    log.Println("Email with alternative text sent successfully!")
}





// Example 16: Sending an email with retry mechanism using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Retry Mechanism Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with retry mechanism.")

    // Send email with retry
    var err error
    for attempt := 1; attempt <= 3; attempt++ {
        if err = smtp.Send(email); err == nil {
            log.Println("Email sent successfully!")
            break
        }
        log.Printf("Error sending email (attempt %d): %v", attempt, err)
        time.Sleep(5 * time.Second) // Wait before retrying
    }
    if err != nil {
        log.Fatalf("Failed to send email after 3 attempts: %v", err)
    }
}





// Example 17: Sending an email with custom headers and attachments using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with custom headers and attachments
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Custom Headers and Attachments from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with custom headers and attachments.")
    email.AddHeader("X-Custom-Header", "Custom Value")

    // Attach a file
    file, err := os.Open("attachment.pdf")
    if err != nil {
        log.Fatalf("Error opening attachment: %v", err)
    }
    defer file.Close()

    email.Attach(file, "attachment.pdf")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom headers and attachments: %v", err)
    }
    log.Println("Email with custom headers and attachments sent successfully!")
}





// Example 18: Sending an email with multiple attachments using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with multiple attachments
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Multiple Attachments from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with multiple attachments.")

    // Attach multiple files
    files := []string{"attachment1.pdf", "attachment2.docx"}
    for _, filename := range files {
        file, err := os.Open(filename)
        if err != nil {
            log.Printf("Error opening attachment %s: %v", filename, err)
            continue
        }
        defer file.Close()
        email.Attach(file, filename)
    }

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with multiple attachments: %v", err)
    }
    log.Println("Email with multiple attachments sent successfully!")
}





// Example 19: Sending an email with dynamic recipients from a list using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Dynamic recipients list
    recipients := []string{"recipient1@example.com", "recipient2@example.com", "recipient3@example.com"}

    // Send emails to each recipient
    for _, recipient := range recipients {
        // Compose email
        email := mail.NewMessage()
        email.SetFrom("sender@example.com")
        email.AddTo(recipient)
        email.SetSubject("Dynamic Recipients Example from Go-Mail")
        email.SetBody(mail.TextPlain, "This is a test email sent to dynamic recipients.")

        // Send email
        if err := smtp.Send(email); err != nil {
            log.Printf("Error sending email to %s: %v", recipient, err)
            continue
        }
        log.Printf("Email sent to %s successfully!", recipient)
    }
}





// Example 20: Sending an email with custom sender name and reply-to address using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with custom sender name and reply-to address
    email := mail.NewMessage()
    email.SetFromWithAlias("sender@example.com", "Sender Name")
    email.AddTo("recipient@example.com")
    email.SetReplyTo("replyto@example.com")
    email.SetSubject("Email with Custom Sender and Reply-To from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with custom sender and reply-to.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom sender and reply-to: %v", err)
    }
    log.Println("Email with custom sender and reply-to sent successfully!")
}





// Example 21: Sending an email with UTF-8 content using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")
    smtp.SetCharset("UTF-8") // Set UTF-8 encoding

    // Compose email with UTF-8 content
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("UTF-8 Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with UTF-8 content: 你好, Hello, ¡Hola!")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending UTF-8 email: %v", err)
    }
    log.Println("UTF-8 email sent successfully!")
}





// Example 22: Sending an email with inline CSS using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with inline CSS
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Inline CSS Example from Go-Mail")
    email.SetBody(mail.TextHTML, `<html><head><style>body { font-family: Arial, sans-serif; }</style></head><body><h1>Hello</h1><p style="color: blue;">This is a test email with inline CSS.</p></body></html>`)

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with inline CSS: %v", err)
    }
    log.Println("Email with inline CSS sent successfully!")
}





// Example 23: Sending an email with dynamic content using go-mail library.

package main

import (
    "log"
    "fmt"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Dynamic content
    recipient := "recipient@example.com"
    subject := "Dynamic Content Email Example from Go-Mail"
    message := "This is a test email with dynamic content."

    // Compose email with dynamic content
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo(recipient)
    email.SetSubject(subject)
    email.SetBody(mail.TextPlain, message)

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with dynamic content: %v", err)
    }
    log.Println("Email with dynamic content sent successfully!")
}





// Example 24: Sending an email using an HTML template with go-mail library.

package main

import (
    "log"
    "html/template"
    "bytes"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Define HTML email template
    const emailTemplate = `
        <html>
        <body>
            <h1>Hello, {{.Name}}!</h1>
            <p>This is a test email sent using an HTML template.</p>
        </body>
        </html>`

    // Prepare data for template
    data := struct {
        Name string
    }{
        Name: "John Doe",
    }

    // Render template
    tpl := template.Must(template.New("emailTemplate").Parse(emailTemplate))
    var tplBuffer bytes.Buffer
    if err := tpl.Execute(&tplBuffer, data); err != nil {
        log.Fatalf("Error executing template: %v", err)
    }

    // Compose email with HTML template
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("HTML Template Email Example from Go-Mail")
    email.SetBody(mail.TextHTML, tplBuffer.String())

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with HTML template: %v", err)
    }
    log.Println("Email with HTML template sent successfully!")
}





// Example 25: Sending an email with scheduled delivery using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Schedule email delivery for 1 minute later
    deliveryTime := time.Now().Add(1 * time.Minute)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Scheduled Delivery Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with scheduled delivery.")

    // Schedule delivery
    if err := smtp.Schedule(email, deliveryTime); err != nil {
        log.Fatalf("Error scheduling email delivery: %v", err)
    }
    log.Printf("Email scheduled for delivery at %s", deliveryTime.Format(time.RFC3339))
}





// Example 26: Sending an email with high priority using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with high priority
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("High Priority Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with high priority.")
    email.SetPriority(mail.PriorityHigh)

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending high priority email: %v", err)
    }
    log.Println("High priority email sent successfully!")
}





// Example 27: Sending an email with custom SMTP timeout using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Custom SMTP configuration with timeout
    smtpConfig := mail.SMTPConfig{
        Host:     "smtp.example.com",
        Port:     587,
        Username: "username",
        Password: "password",
        Timeout:  30 * time.Second, // Custom timeout
    }

    // Initialize SMTP client with custom configuration
    smtp := mail.NewCustomSMTP(smtpConfig)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom Timeout Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with custom SMTP timeout.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom timeout: %v", err)
    }
    log.Println("Email sent with custom timeout successfully!")
}





// Example 28: Sending an email with HTML content and attachments using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with HTML content and attachments
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("HTML Email with Attachments Example from Go-Mail")
    email.SetBody(mail.TextHTML, `<html><body><h1>Hello</h1><p>This is a test email with HTML content and attachments.</p></body></html>`)

    // Attach a file
    file, err := os.Open("attachment.pdf")
    if err != nil {
        log.Fatalf("Error opening attachment: %v", err)
    }
    defer file.Close()

    email.Attach(file, "attachment.pdf")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with HTML content and attachments: %v", err)
    }
    log.Println("Email with HTML content and attachments sent successfully!")
}





// Example 29: Sending an email with read receipt request using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with read receipt request
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Read Receipt Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with read receipt request.")
    email.SetReadReceipt("readreceipt@example.com")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with read receipt request: %v", err)
    }
    log.Println("Email with read receipt request sent successfully!")
}





// Example 30: Sending an email with custom SMTP headers using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with custom SMTP headers
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom Headers Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with custom SMTP headers.")
    email.AddHeader("X-Custom-Header", "Custom Value")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom headers: %v", err)
    }
    log.Println("Email with custom headers sent successfully!")
}





// Example 31: Sending an email using an HTML template with inline CSS using go-mail library.

package main

import (
    "log"
    "html/template"
    "bytes"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Define HTML email template with inline CSS
    const emailTemplate = `
        <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    color: #333333;
                }
                h1 {
                    color: #0099ff;
                }
            </style>
        </head>
        <body>
            <h1>Hello, {{.Name}}!</h1>
            <p>This is a test email sent using an HTML template with inline CSS.</p>
        </body>
        </html>`

    // Prepare data for template
    data := struct {
        Name string
    }{
        Name: "Jane Doe",
    }

    // Render template
    tpl := template.Must(template.New("emailTemplate").Parse(emailTemplate))
    var tplBuffer bytes.Buffer
    if err := tpl.Execute(&tplBuffer, data); err != nil {
        log.Fatalf("Error executing template: %v", err)
    }

    // Compose email with HTML template and inline CSS
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("HTML Template with Inline CSS Email Example from Go-Mail")
    email.SetBody(mail.TextHTML, tplBuffer.String())

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with HTML template and inline CSS: %v", err)
    }
    log.Println("Email with HTML template and inline CSS sent successfully!")
}





// Example 32: Sending an email to multiple recipients using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Multiple recipients
    recipients := []string{"recipient1@example.com", "recipient2@example.com", "recipient3@example.com"}

    // Compose email for each recipient
    for _, recipient := range recipients {
        email := mail.NewMessage()
        email.SetFrom("sender@example.com")
        email.AddTo(recipient)
        email.SetSubject("Multiple Recipients Email Example from Go-Mail")
        email.SetBody(mail.TextPlain, "This is a test email sent to multiple recipients.")

        // Send email
        if err := smtp.Send(email); err != nil {
            log.Printf("Error sending email to %s: %v", recipient, err)
            continue
        }
        log.Printf("Email sent to %s successfully!", recipient)
    }
}





// Example 33: Sending an email with different encodings using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")
    smtp.SetCharset("ISO-8859-1") // Set encoding to ISO-8859-1 (Latin-1)

    // Compose email with different encoding
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Different Encoding Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with different encoding: café")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with different encoding: %v", err)
    }
    log.Println("Email with different encoding sent successfully!")
}





// Example 34: Sending an email with delayed delivery using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Schedule email delivery after 5 minutes
    deliveryTime := time.Now().Add(5 * time.Minute)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Delayed Delivery Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with delayed delivery.")

    // Schedule delivery
    if err := smtp.Schedule(email, deliveryTime); err != nil {
        log.Fatalf("Error scheduling email delivery: %v", err)
    }
    log.Printf("Email scheduled for delivery at %s", deliveryTime.Format(time.RFC3339))
}





// Example 35: Sending an email with a custom SMTP client using go-mail library.

package main

import (
    "log"
    "net/smtp"

    "github.com/go-mail/mail"
)

func main() {
    // Custom SMTP client configuration
    smtpClient := smtp.NewClient(nil, "smtp.example.com")
    if err := smtpClient.Auth(smtp.PlainAuth("", "username", "password", "smtp.example.com")); err != nil {
        log.Fatalf("Error authenticating SMTP client: %v", err)
    }

    // Initialize custom SMTP client
    smtp := mail.NewCustomSMTPClient(smtpClient)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom SMTP Client Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with a custom SMTP client.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom SMTP client: %v", err)
    }
    log.Println("Email sent with custom SMTP client successfully!")
}





// Example 36: Sending an email using secure SMTP (TLS/SSL) with go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Secure SMTP configuration (using TLS)
    smtpConfig := mail.SMTPConfig{
        Host:     "smtp.example.com",
        Port:     587,
        Username: "username",
        Password: "password",
        TLSConfig: &mail.TLSConfig{
            InsecureSkipVerify: true, // Example: Insecure skip verify; use carefully in production.
        },
    }

    // Initialize secure SMTP client
    smtp := mail.NewCustomSMTP(smtpConfig)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Secure SMTP Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent using secure SMTP (TLS).")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with secure SMTP: %v", err)
    }
    log.Println("Email sent using secure SMTP successfully!")
}





// Example 37: Sending an email with CC and BCC using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with CC and BCC
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.AddCC("cc1@example.com")
    email.AddCC("cc2@example.com")
    email.AddBCC("bcc1@example.com")
    email.AddBCC("bcc2@example.com")
    email.SetSubject("CC and BCC Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with CC and BCC recipients.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with CC and BCC: %v", err)
    }
    log.Println("Email with CC and BCC sent successfully!")
}





// Example 38: Sending an email with a custom Message-ID using go-mail library.

package main

import (
    "log"
    "github.com/go-mail/mail"
    "fmt"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with custom Message-ID
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom Message-ID Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with a custom Message-ID.")

    // Set custom Message-ID
    customMessageID := fmt.Sprintf("<custom-%d@example.com>", time.Now().UnixNano())
    email.SetMessageID(customMessageID)

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom Message-ID: %v", err)
    }
    log.Println("Email with custom Message-ID sent successfully!")
}





// Example 39: Sending an email with retry mechanism using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Retry Mechanism Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with retry mechanism.")

    // Retry sending email up to 3 times with exponential backoff
    retryCount := 3
    for i := 0; i < retryCount; i++ {
        if err := smtp.Send(email); err != nil {
            log.Printf("Error sending email (attempt %d): %v", i+1, err)
            // Exponential backoff before retrying
            time.Sleep(time.Duration(i*2) * time.Second)
            continue
        }
        log.Println("Email sent successfully!")
        break
    }
}





// Example 40: Sending an email with HTML content and embedded images using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with HTML content and embedded images
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("HTML Email with Embedded Images Example from Go-Mail")
    email.SetBody(mail.TextHTML, `<html><body><h1>Hello</h1><p>This is a test email with HTML content and embedded image:<br/><img src="cid:logo"></p></body></html>`)

    // Attach embedded image
    imgFile, err := os.Open("logo.png")
    if err != nil {
        log.Fatalf("Error opening image file: %v", err)
    }
    defer imgFile.Close()

    email.Attach(imgFile, "logo.png", mail.InlineFile("logo"))

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with embedded images: %v", err)
    }
    log.Println("Email with HTML content and embedded images sent successfully!")
}





// Example 41: Sending an email with low priority using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with low priority
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Low Priority Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with low priority.")
    email.SetPriority(mail.PriorityLow)

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending low priority email: %v", err)
    }
    log.Println("Low priority email sent successfully!")
}





// Example 42: Sending an email with a Reply-To address using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with Reply-To address
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Reply-To Address Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with a Reply-To address.")
    email.SetReplyTo("reply-to@example.com")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with Reply-To address: %v", err)
    }
    log.Println("Email with Reply-To address sent successfully!")
}





// Example 43: Sending an email with a persistent SMTP connection using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client with persistent connection
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")
    smtp.SetKeepAlive(30 * time.Second) // Keep the connection alive for 30 seconds

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Persistent Connection Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with a persistent SMTP connection.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with persistent connection: %v", err)
    }
    log.Println("Email sent with persistent connection successfully!")
}





// Example 44: Sending an email with a custom retry strategy using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Custom retry strategy
    retryAttempts := 5
    retryInterval := 5 * time.Second

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom Retry Strategy Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with a custom retry strategy.")

    // Retry sending email
    for i := 1; i <= retryAttempts; i++ {
        if err := smtp.Send(email); err != nil {
            log.Printf("Attempt %d failed: %v", i, err)
            time.Sleep(retryInterval)
            continue
        }
        log.Println("Email sent successfully!")
        break
    }
}





// Example 45: Sending an email with custom SMTP timeout using go-mail library.

package main

import (
    "log"
    "time"

    "github.com/go-mail/mail"
)

func main() {
    // Custom SMTP configuration with timeout
    smtpConfig := mail.SMTPConfig{
        Host:     "smtp.example.com",
        Port:     587,
        Username: "username",
        Password: "password",
        Timeout:  30 * time.Second, // Custom timeout
    }

    // Initialize SMTP client with custom configuration
    smtp := mail.NewCustomSMTP(smtpConfig)

    // Compose email
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom Timeout Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email sent with custom SMTP timeout.")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom timeout: %v", err)
    }
    log.Println("Email sent with custom timeout successfully!")
}





// Example 46: Sending an email with an attachment from a file using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with attachment from file
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Attachment Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with an attachment.")

    // Attach a file
    file, err := os.Open("document.pdf")
    if err != nil {
        log.Fatalf("Error opening attachment file: %v", err)
    }
    defer file.Close()

    email.Attach(file, "document.pdf")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with attachment: %v", err)
    }
    log.Println("Email with attachment sent successfully!")
}





// Example 47: Sending an email with a custom header and UTF-8 encoding using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")
    smtp.SetCharset("UTF-8") // Set encoding to UTF-8

    // Compose email with custom header and UTF-8 encoding
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Custom Header and UTF-8 Encoding Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with custom header and UTF-8 encoding.")
    email.AddHeader("X-Custom-Header", "Custom Value")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with custom header and UTF-8 encoding: %v", err)
    }
    log.Println("Email with custom header and UTF-8 encoding sent successfully!")
}





// Example 48: Sending an email with inline content (inline attachments) using go-mail library.

package main

import (
    "log"
    "os"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with inline content
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Inline Content Email Example from Go-Mail")
    email.SetBody(mail.TextHTML, `<html><body><h1>Hello</h1><p>This is a test email with inline image: <img src="cid:image1"></p></body></html>`)

    // Attach inline content (image)
    imgFile, err := os.Open("image.jpg")
    if err != nil {
        log.Fatalf("Error opening image file: %v", err)
    }
    defer imgFile.Close()

    email.Attach(imgFile, "image.jpg", mail.InlineFile("image1"))

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with inline content: %v", err)
    }
    log.Println("Email with inline content sent successfully!")
}





// Example 49: Sending an email with delivery status notification (DSN) using go-mail library.

package main

import (
    "log"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Compose email with delivery status notification (DSN)
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("DSN Email Example from Go-Mail")
    email.SetBody(mail.TextPlain, "This is a test email with delivery status notification (DSN).")
    email.SetDSN("never")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with delivery status notification (DSN): %v", err)
    }
    log.Println("Email with delivery status notification (DSN) sent successfully!")
}





// Example 50: Sending an email with template data loaded from a JSON file using go-mail library.

package main

import (
    "log"
    "io/ioutil"
    "encoding/json"

    "github.com/go-mail/mail"
)

func main() {
    // Initialize SMTP client
    smtp := mail.NewSMTP("smtp.example.com", 587, "username", "password")

    // Load template data from JSON file
    templateData, err := ioutil.ReadFile("email_data.json")
    if err != nil {
        log.Fatalf("Error reading template data from JSON file: %v", err)
    }

    // Unmarshal JSON data
    var data map[string]interface{}
    if err := json.Unmarshal(templateData, &data); err != nil {
        log.Fatalf("Error unmarshaling JSON data: %v", err)
    }

    // Compose email with template data
    email := mail.NewMessage()
    email.SetFrom("sender@example.com")
    email.AddTo("recipient@example.com")
    email.SetSubject("Email with Template Data Example from Go-Mail")

    // Example assumes email_data.json contains key "name" with value "John Doe"
    email.SetBody(mail.TextPlain, "Hello, "+data["name"].(string)+"!")

    // Send email
    if err := smtp.Send(email); err != nil {
        log.Fatalf("Error sending email with template data: %v", err)
    }
    log.Println("Email with template data sent successfully!")
}
