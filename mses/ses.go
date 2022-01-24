package mses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const (
	CharSet = "UTF-8"
)

var (
	_fromEmailAddress = ""
	_client           *sesv2.Client
)

func Init(cfg aws.Config, fromEmailAddress string) {
	_client = sesv2.NewFromConfig(cfg)
	_fromEmailAddress = fromEmailAddress
}

type Email struct {
	Subject          string
	Text             string
	HtmlBody         string
	ToAddresses      []string
	BccAddresses     []string
	CcAddresses      []string
	ReplyToAddresses []string
}

func NewEmail() *Email {
	return &Email{}
}

func NewEmailAll(subject, text, htmlBody string, toAddresses ...string) *Email {
	return &Email{
		Subject:     subject,
		Text:        text,
		HtmlBody:    htmlBody,
		ToAddresses: toAddresses,
	}
}

func (item *Email) SendEmail() (*sesv2.SendEmailOutput, error) {
	content := &types.EmailContent{
		Simple: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data:    aws.String(item.HtmlBody),
					Charset: aws.String(CharSet),
				},
				Text: &types.Content{ //不支持html的客户端走这个
					Data:    aws.String(item.Text),
					Charset: aws.String(CharSet),
				},
			},
			Subject: &types.Content{
				Data:    aws.String(item.Subject),
				Charset: aws.String(CharSet),
			},
		},
	}
	params := &sesv2.SendEmailInput{
		Content: content,
		Destination: &types.Destination{
			BccAddresses: item.BccAddresses, //密送
			CcAddresses:  item.CcAddresses,  //抄送
			ToAddresses:  item.ToAddresses,  //收件人
		},
		FromEmailAddress: aws.String(_fromEmailAddress),
		ReplyToAddresses: item.ReplyToAddresses,
	}
	res, err := _client.SendEmail(context.TODO(), params)
	return res, err
}

func (item *Email) SetAddress(toAddresses ...string) {
	item.ToAddresses = toAddresses
}
