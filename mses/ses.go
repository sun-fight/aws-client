package mses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

var (
	_client  *sesv2.Client
	_charSet = "UTF-8"
)

func Init(cfg aws.Config) {
	_client = sesv2.NewFromConfig(cfg)
}

type ReqSendEmail struct {
	FromEmailAddress *string
	Subject          *string
	Text             *string
	HtmlBody         *string
	ToAddresses      []string
	BccAddresses     []string
	CcAddresses      []string
	ReplyToAddresses []string
}

type Email struct {
}

func NewEmail() *Email {
	return &Email{}
}

func NewReqSendEmail(subject, text, htmlBody string, toAddresses ...string) ReqSendEmail {
	return ReqSendEmail{
		Subject:     aws.String(subject),
		Text:        aws.String(text),
		HtmlBody:    aws.String(htmlBody),
		ToAddresses: toAddresses,
	}
}

func (item *Email) SendEmail(req ReqSendEmail) (*sesv2.SendEmailOutput, error) {
	content := &types.EmailContent{
		Simple: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data:    (req.HtmlBody),
					Charset: aws.String(_charSet),
				},
				Text: &types.Content{ //不支持html的客户端走这个
					Data:    (req.Text),
					Charset: aws.String(_charSet),
				},
			},
			Subject: &types.Content{
				Data:    (req.Subject),
				Charset: aws.String(_charSet),
			},
		},
	}
	params := &sesv2.SendEmailInput{
		Content: content,
		Destination: &types.Destination{
			BccAddresses: req.BccAddresses, //密送
			CcAddresses:  req.CcAddresses,  //抄送
			ToAddresses:  req.ToAddresses,  //收件人
		},
		FromEmailAddress: req.FromEmailAddress,
		ReplyToAddresses: req.ReplyToAddresses,
	}
	res, err := _client.SendEmail(context.TODO(), params)
	return res, err
}

// Set custom charSet.
// Default charSet is "UTF-8"
func (item *Email) SetCharSet(charSet string) {
	_charSet = charSet
}
