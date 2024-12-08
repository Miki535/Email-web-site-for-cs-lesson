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

	r.LoadHTMLGlob("../templates/htmls/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/SMTPtest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "send-info.html", gin.H{})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{})
	})

	r.POST("/SMTPtest", SMTPtest)
	r.Run() //run on localhost:8080
}

func SMTPtest(c *gin.Context) {
	data := Data{
		SenderEmail: c.PostForm("senderemail"),
		Email:       c.PostForm("email"),
		Message:     c.PostForm("message"),
		SecretCode:  c.PostForm("secretcode"),
	}
	auth := smtp.PlainAuth("", data.SenderEmail, data.SecretCode, "smtp.gmail.com")

	to := []string{data.Email}
	msg := data.Message
	// send message on email
	err := smtp.SendMail("smtp.gmail.com:587", auth, data.SenderEmail, to, []byte(fmt.Sprint(msg)))

	if err != nil {
		c.Redirect(http.StatusFound, "/send-error.html")
	} else {
		c.Redirect(http.StatusSeeOther, "/SMTPtest")
	}
}
