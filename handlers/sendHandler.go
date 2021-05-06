package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wouter173/mailer/structs"
	gomail "gopkg.in/mail.v2"
)

func SendHandler(c *fiber.Ctx) error {

	if c.Method() != "POST" {
		c.Status(405)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Method not allowed."))
	}

	if c.Get("Content-Type") != "application/json" {
		c.Status(415)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Unsupported media type."))
	}

	var data structs.Email
	err := json.Unmarshal(c.Body(), &data)

	if err != nil {
		c.Status(422)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Bruh, that is not json."))
	}

	if data.Body == "" {
		c.Status(422)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Invalid body."))
	}

	if data.Subject == "" {
		c.Status(422)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Invalid subject."))
	}

	footer := fmt.Sprintf("<br><br><p style=\"font-size: 8pt; font-weight: bold\">%s</p>", c.Locals("Service"))

	mail := gomail.NewMessage()
	mail.SetHeader("From", fmt.Sprintf("%s <%s>", os.Getenv("EMAIL_NAME"), os.Getenv("EMAIL_ADDRESS")))
	mail.SetHeader("To", data.Target)
	mail.SetHeader("Subject", data.Subject)
	mail.SetBody("text/html", fmt.Sprintf("%s %s", data.Body, footer))

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	dialer := gomail.NewDialer(os.Getenv("SMTP_URI"), port, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))
	err = dialer.DialAndSend(mail)

	if err != nil {
		c.Status(422)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Invalid target address."))
	}

	c.Status(200)
	c.Set("Content-Type", "application/json")
	return c.Send(structs.NewResult(""))
}
