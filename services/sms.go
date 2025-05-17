package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

var (
	from     = "MatchP@0xsol.org"
	subject  = "MatchP Verification Code"
	textBody = "Your verification code,don't share to anyone."
)

func Send(key, secret, code string, to string) (success bool, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")))
	if err != nil {
		return false, err
	}
	client := sesv2.NewFromConfig(cfg)
	htmlBody := fmt.Sprintf("<h1>Code:%s</h1><p>%s</p>", code, textBody)
	// build email content
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from),
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String(subject),
				},
				Body: &types.Body{
					Html: &types.Content{
						Data: aws.String(htmlBody),
					},
					//Text: &types.Content{
					//	Data: aws.String(textBody),
					//},
				},
			},
		},
	}

	resp, err := client.SendEmail(context.TODO(), input)
	if err != nil {
		return false, errors.New("failed to send email: " + err.Error())
	}

	fmt.Printf("mail: %s send success，MessageId: %v \n", to, *resp.MessageId)

	return true, nil

}
