package notifier

//
//import (
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/ses"
//)
//
//type SESService struct {
//	svc *ses.SES
//}
//
//// NewSESService Creates an AWS Connection and builds an SESService
//func NewSESService() (*SESService, error) {
//	sesSrv := &SESService{}
//	sess, err := session.NewSession(&aws.Config{
//		Region: aws.String("us-west-2"), // Adjust to your region
//	})
//	if err != nil {
//		return nil, err
//	}
//	sesSrv.svc = ses.New(sess)
//
//	return sesSrv, nil
//}
//
//func (ss *SESService) SendSESEmail(to string, subject string, body string) error {
//	input := &ses.SendEmailInput{
//		Destination: &ses.Destination{
//			ToAddresses: []*string{aws.String(to)},
//		},
//		Message: &ses.Message{
//			Body: &ses.Body{
//				Text: &ses.Content{
//					Charset: aws.String("UTF-8"),
//					Data:    aws.String(body),
//				},
//			},
//			Subject: &ses.Content{
//				Charset: aws.String("UTF-8"),
//				Data:    aws.String(subject),
//			},
//		},
//		Source: aws.String("mqtt-bot@mail.impermanent.faith"),
//	}
//
//	_, err := ss.svc.SendEmail(input)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
