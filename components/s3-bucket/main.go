package s3bucket

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type S3Bucket struct {
	pulumi.ResourceState

	Bucket     *s3.Bucket          `pulumi:"bucket"`
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
}

type BucketArgs struct {
	Name string
}

func NewS3Bucket(ctx *pulumi.Context, name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*S3Bucket, error) {
	s3Bucket := &S3Bucket{}
	err := ctx.RegisterComponentResource("andel:index:S3Bucket", name, s3Bucket, opts...)
	if err != nil {
		return nil, err
	}

	bucket, err := s3.NewBucket(ctx, name, &s3.BucketArgs{
		Bucket: pulumi.String(args.Name),
	}, pulumi.Parent(s3Bucket))

	s3Bucket.Bucket = bucket
	s3Bucket.BucketName = bucket.Bucket

	err = ctx.RegisterResourceOutputs(s3Bucket, pulumi.Map{
		"bucket":     bucket,
		"bucketName": bucket.Bucket,
	})

	if err != nil {
		return nil, err
	}

	return s3Bucket, nil
}
