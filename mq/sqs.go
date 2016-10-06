package mq

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/MiteshSharma/Sarthi/errors"
	"github.com/MiteshSharma/Sarthi/utils"
)

type SqsAgent struct {
	Region            string
	QueueUrl          string
	VisibilityTimeout int64
}

func NewSqsAgent() (*SqsAgent, error) {
	var agent *SqsAgent
	var ok bool

	var config map[string]interface{}
	if config, ok = utils.ConfigParam.ExecutorConfig["sqs"].(map[string]interface{}); !ok {
		return agent, &errors.ConfigurationNotFoundError{
			Type: "sqs",
		}
	}

	if _, ok = config["queue-url"]; !ok {
		return agent, &errors.MissingConfigurationKeyError{
			Type: "sqs",
			Key:  "queue-url",
		}
	}
	if _, ok = config["visibility-timeout"]; !ok {
		return agent, &errors.MissingConfigurationKeyError{
			Type: "sqs",
			Key:  "visibility-timeout",
		}
	}
	if _, ok = config["region"]; !ok {
		return agent, &errors.MissingConfigurationKeyError{
			Type: "sqs",
			Key:  "region",
		}
	}

	return &SqsAgent{
		Region:            config["region"].(string),
		QueueUrl:          config["queue-url"].(string),
		VisibilityTimeout: int64(config["visibility-timeout"].(float64)),
	}, nil
}

func (s *SqsAgent) Read() (*MqObject, error) {
	var m MqObject

	// create aws session
	sess, err := session.NewSession()
	if err != nil {
		return &m, &errors.MqError{
			Type:    "sqs",
			Message: "Failed to create session.",
			Err:     err,
		}
	}

	// read 1 message from sqs
	svc := sqs.New(sess, &aws.Config{Region: aws.String(s.Region)})
	params := &sqs.ReceiveMessageInput{
		QueueUrl:          aws.String(s.QueueUrl),
		VisibilityTimeout: aws.Int64(s.VisibilityTimeout),
	}
	resp, err := svc.ReceiveMessage(params)
	if err != nil {
		return &m, &errors.MqError{
			Type:    "sqs",
			Message: "Error receiving message.",
			Err:     err,
		}
	}

	// process response
	if len(resp.Messages) == 1 {
		msg := resp.Messages[0]
		fmt.Println("Found message -> ", msg)
		m.Id = *msg.ReceiptHandle
		m.Message = []byte(*msg.Body)
	}
	return &m, nil
}

func (s *SqsAgent) Write(message []byte) error {
	// create aws session
	sess, err := session.NewSession()
	if err != nil {
		return &errors.MqError{
			Type:    "sqs",
			Message: "Failed to create session.",
			Err:     err,
		}
	}

	// write message to sqs
	svc := sqs.New(sess, &aws.Config{Region: aws.String(s.Region)})
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(string(message)),
		QueueUrl:    aws.String(s.QueueUrl),
	}
	_, err = svc.SendMessage(params)
	if err != nil {
		return &errors.MqError{
			Type:    "sqs",
			Message: "Error writing message.",
			Err:     err,
		}
	}

	return nil
}

func (s *SqsAgent) Delete(MessageId string) error {
	// create aws session
	sess, err := session.NewSession()
	if err != nil {
		return &errors.MqError{
			Type:    "sqs",
			Message: "Failed to create session.",
			Err:     err,
		}
	}

	// delete message from sqs
	svc := sqs.New(sess, &aws.Config{Region: aws.String(s.Region)})
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(s.QueueUrl),
		ReceiptHandle: aws.String(MessageId),
	}
	_, err = svc.DeleteMessage(params)
	if err != nil {
		return &errors.MqError{
			Type:    "sqs",
			Message: "Error deleting message.",
			Err:     err,
		}
	}

	return nil
}
