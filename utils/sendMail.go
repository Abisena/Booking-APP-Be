package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type EmailData struct {
	Subject string
	Body   string
	To     []string
}


func SendMail(data EmailData) error {
    auth := smtp.PlainAuth(
        "",
        "abisena1123@gmail.com",
        "swojswzpimauiihl",
        "smtp.gmail.com",
    )

    t, err := template.ParseFiles("./email.html")
    if err != nil {
        return err
    }

    var buf bytes.Buffer
    err = t.Execute(&buf, data)
    if err != nil {
        return err
    }

    plainText := convertHTMLToPlainText(buf.String())

    msg := fmt.Sprintf("Subject: %s\nContent-Type: text/plain; charset=UTF-8\n\n%s", data.Subject, plainText)

    err = smtp.SendMail(
        "smtp.gmail.com:587",
        auth,
        "abisena1123@gmail.com",
        data.To,
        []byte(msg),
    )

    return err
}

func convertHTMLToPlainText(html string) string {
	p := bluemonday.NewPolicy()
	p.AllowAttrs("style").OnElements("span")
	p.AllowStandardURLs()
	p.AllowElements("a", "br", "p", "ul", "ol", "li")
	p.AllowAttrs("rel").Matching(regexp.MustCompile("nofollow")).OnElements("a")
	
	plainText := p.Sanitize(html)

	// Remove HTML tags and replace line breaks with spaces
	plainText = strings.ReplaceAll(plainText, "<br>", "\n")
	plainText = strings.ReplaceAll(plainText, "</p>", "\n\n")
	plainText = strings.ReplaceAll(plainText, "</li>", "\n")
	plainText = strings.ReplaceAll(plainText, "</ul>", "\n")
	plainText = strings.ReplaceAll(plainText, "</ol>", "\n")
	plainText = strings.ReplaceAll(plainText, "<.*?>", "")

	return plainText
}