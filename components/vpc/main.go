package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	pulumi.ResourceState

	Vpc   *ec2.Vpc        `pulumi:"vpc"`
	VpcId pulumi.IDOutput `pulumi:"vpc-id"`
}

type ModuleArgs struct {
	Cidr string
}

func NewVpc(ctx *pulumi.Context, name string, args *ModuleArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	module := &Vpc{}
	err := ctx.RegisterComponentResource("thomassjogren:index:aws-vpc", name, module, opts...)
	if err != nil {
		return nil, err
	}

	vpc, err := ec2.NewVpc(ctx, name, &ec2.VpcArgs{
		CidrBlock: pulumi.String(args.Cidr),
	}, pulumi.Parent(module))

	module.Vpc = vpc
	module.VpcId = vpc.ID()

	err = ctx.RegisterResourceOutputs(module, pulumi.Map{
		"vpc":     vpc,
		"vpc-id":  vpc.ID(),
		"vpc-arn": vpc.Arn,
	})

	if err != nil {
		return nil, err
	}

	return module, nil
}
