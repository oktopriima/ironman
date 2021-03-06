/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package helper

import (
	"fmt"

	"github.com/KulinaID/kulina-go-libraries/kuconfig"
	"github.com/keighl/mandrill"
)

// Sends email via mandrill.
// Accepts mailTo as destination, mailSubject as subject, content as
// the rendered content (usually html).
// Returns email, status, reject reason, and id from Mandrill.
func SendEmail(mailTo, mailSubject, content string) []*mandrill.Response {
	cfg := kuconfig.NewConfig()

	client := mandrill.ClientWithKey(cfg.GetString("mandrill.key"))
	message := new(mandrill.Message)
	message.AddRecipient(mailTo, "", "to")
	message.FromEmail = cfg.GetString("mandrill.from.email")
	message.FromName = cfg.GetString("mandrill.from.name")
	message.Subject = mailSubject
	message.HTML = content
	message.Text = content

	responses, err := client.MessagesSend(message)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return responses
}
