package packageName

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/module-name"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Module struct {
	pulumi.ResourceState

	SomeModule *AWS.TYPE           `pulumi:"output-name"`
	SomeOutput pulumi.StringOutput `pulumi:"output-name"`
}

type ModuleArgs struct {
	Prop string
}

func ModuleName(ctx *pulumi.Context, name string, args *ModuleArgs, opts ...pulumi.ResourceOption) (*AWSResource, error) {
	someModule := &Module{}
	err := ctx.RegisterComponentResource("org:index:Module", name, someModule, opts...)
	if err != nil {
		return nil, err
	}

	returnValue, err := aws.resource(ctx, name, &aws.resourceArgs{
		Prop: pulumi.String(args.Prop),
	}, pulumi.Parent(someModule))

	someModule.SomeModule = returnValue
	someModule.SomeOutput = returnValue.someProp

	err = ctx.RegisterResourceOutputs(someModule, pulumi.Map{
		"some-output":      returnValue,
		"some-output-more": returnValue.someProp,
	})

	if err != nil {
		return nil, err
	}

	return someModule, nil
}
