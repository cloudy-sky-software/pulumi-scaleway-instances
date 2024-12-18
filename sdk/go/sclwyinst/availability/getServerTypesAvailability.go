// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package availability

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetServerTypesAvailability(ctx *pulumi.Context, args *GetServerTypesAvailabilityArgs, opts ...pulumi.InvokeOption) (*GetServerTypesAvailabilityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerTypesAvailabilityResult
	err := ctx.Invoke("scaleway-instances:availability:getServerTypesAvailability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetServerTypesAvailabilityArgs struct {
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type GetServerTypesAvailabilityResult struct {
	Servers map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability `pulumi:"servers"`
}

func GetServerTypesAvailabilityOutput(ctx *pulumi.Context, args GetServerTypesAvailabilityOutputArgs, opts ...pulumi.InvokeOption) GetServerTypesAvailabilityResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServerTypesAvailabilityResultOutput, error) {
			args := v.(GetServerTypesAvailabilityArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway-instances:availability:getServerTypesAvailability", args, GetServerTypesAvailabilityResultOutput{}, options).(GetServerTypesAvailabilityResultOutput), nil
		}).(GetServerTypesAvailabilityResultOutput)
}

type GetServerTypesAvailabilityOutputArgs struct {
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (GetServerTypesAvailabilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerTypesAvailabilityArgs)(nil)).Elem()
}

type GetServerTypesAvailabilityResultOutput struct{ *pulumi.OutputState }

func (GetServerTypesAvailabilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerTypesAvailabilityResult)(nil)).Elem()
}

func (o GetServerTypesAvailabilityResultOutput) ToGetServerTypesAvailabilityResultOutput() GetServerTypesAvailabilityResultOutput {
	return o
}

func (o GetServerTypesAvailabilityResultOutput) ToGetServerTypesAvailabilityResultOutputWithContext(ctx context.Context) GetServerTypesAvailabilityResultOutput {
	return o
}

func (o GetServerTypesAvailabilityResultOutput) Servers() ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput {
	return o.ApplyT(func(v GetServerTypesAvailabilityResult) map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability {
		return v.Servers
	}).(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerTypesAvailabilityResultOutput{})
}
