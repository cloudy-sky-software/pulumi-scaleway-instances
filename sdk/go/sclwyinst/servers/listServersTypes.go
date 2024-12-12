// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servers

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServersTypes(ctx *pulumi.Context, args *ListServersTypesArgs, opts ...pulumi.InvokeOption) (*ListServersTypesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListServersTypesResult
	err := ctx.Invoke("scaleway-instances:servers:listServersTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServersTypesArgs struct {
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type ListServersTypesResult struct {
	// List of server types
	Servers map[string]ScalewayInstanceV1ServerType `pulumi:"servers"`
}

func ListServersTypesOutput(ctx *pulumi.Context, args ListServersTypesOutputArgs, opts ...pulumi.InvokeOption) ListServersTypesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListServersTypesResultOutput, error) {
			args := v.(ListServersTypesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway-instances:servers:listServersTypes", args, ListServersTypesResultOutput{}, options).(ListServersTypesResultOutput), nil
		}).(ListServersTypesResultOutput)
}

type ListServersTypesOutputArgs struct {
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (ListServersTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServersTypesArgs)(nil)).Elem()
}

type ListServersTypesResultOutput struct{ *pulumi.OutputState }

func (ListServersTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServersTypesResult)(nil)).Elem()
}

func (o ListServersTypesResultOutput) ToListServersTypesResultOutput() ListServersTypesResultOutput {
	return o
}

func (o ListServersTypesResultOutput) ToListServersTypesResultOutputWithContext(ctx context.Context) ListServersTypesResultOutput {
	return o
}

// List of server types
func (o ListServersTypesResultOutput) Servers() ScalewayInstanceV1ServerTypeMapOutput {
	return o.ApplyT(func(v ListServersTypesResult) map[string]ScalewayInstanceV1ServerType { return v.Servers }).(ScalewayInstanceV1ServerTypeMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServersTypesResultOutput{})
}
