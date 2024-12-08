package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
)

type Data struct {
	SenderEmail string
	Email       string
	Message     string
	SecretCode  string
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/SMTPtest", func(c *gin.Context) {})
	r.POST("/SMTPtest", SMTPtest)
	r.Run()
}

func SMTPtest(c *gin.Context) {
	data := Data{
		SenderEmail: c.PostForm("SenderEmail"),
		Email:       c.PostForm("Email"),
		Message:     c.PostForm("Message"),
		SecretCode:  c.PostForm("SecretCode"),
	}
	auth := smtp.PlainAuth("", data.SenderEmail, data.SecretCode, "smtp.gmail.com")

	to := []string{data.Email}
	msg := data.Message
	// send message on email
	err := smtp.SendMail("smtp.gmail.com:587", auth, data.SenderEmail, to, []byte(fmt.Sprint(msg)))

	if err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
	}
}
