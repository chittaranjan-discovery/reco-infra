package modules

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SetupSQS create a SQS queue.
func SetupSQS(ctx *pulumi.Context, name string, delaySeconds int, maxMessageSize int, messageRetentionSeconds int, receiveWaitTimeSeconds int, visibilityTimeoutSeconds int) (*sqs.Queue, error) {
	//q := CreateRecoResourceName(ctx, name)
	//return sqs.NewQueue(ctx, q, &sqs.QueueArgs{
	return sqs.NewQueue(ctx{
		DelaySeconds:             pulumi.Int(delaySeconds),
		MaxMessageSize:           pulumi.Int(maxMessageSize),
		MessageRetentionSeconds:  pulumi.Int(messageRetentionSeconds),
		ReceiveWaitTimeSeconds:   pulumi.Int(receiveWaitTimeSeconds),
		VisibilityTimeoutSeconds: pulumi.Int(visibilityTimeoutSeconds),
	})
}
