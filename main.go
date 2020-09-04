package main

import (
	"github.com/chittaranjan-discovery/reco-infra/modules"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "test-chitta-bucket", nil)
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())

		// Setup SQS queue
		q, err := resources.SetupSQS(ctx, "test-chitta-queue", modules.DelaySeconds, modules.MaxMessageSize, modules.MessageRetentionSeconds, modules.ReceiveWaitTimeSeconds, modules.VisibilityTimeoutSeconds)
		if err != nil {
			return err
		}
		return nil
	})
}
