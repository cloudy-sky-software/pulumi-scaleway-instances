// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package volumes

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func ListVolumesTypes(ctx *pulumi.Context, args *ListVolumesTypesArgs, opts ...pulumi.InvokeOption) (*ListVolumesTypesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListVolumesTypesResult
	err := ctx.Invoke("scaleway-instances:volumes:listVolumesTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVolumesTypesArgs struct {
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type ListVolumesTypesResult struct {
	Items ScalewayInstanceV1ListVolumesTypesResponse `pulumi:"items"`
}

func ListVolumesTypesOutput(ctx *pulumi.Context, args ListVolumesTypesOutputArgs, opts ...pulumi.InvokeOption) ListVolumesTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVolumesTypesResult, error) {
			args := v.(ListVolumesTypesArgs)
			r, err := ListVolumesTypes(ctx, &args, opts...)
			var s ListVolumesTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVolumesTypesResultOutput)
}

type ListVolumesTypesOutputArgs struct {
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (ListVolumesTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumesTypesArgs)(nil)).Elem()
}

type ListVolumesTypesResultOutput struct{ *pulumi.OutputState }

func (ListVolumesTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumesTypesResult)(nil)).Elem()
}

func (o ListVolumesTypesResultOutput) ToListVolumesTypesResultOutput() ListVolumesTypesResultOutput {
	return o
}

func (o ListVolumesTypesResultOutput) ToListVolumesTypesResultOutputWithContext(ctx context.Context) ListVolumesTypesResultOutput {
	return o
}

func (o ListVolumesTypesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListVolumesTypesResult] {
	return pulumix.Output[ListVolumesTypesResult]{
		OutputState: o.OutputState,
	}
}

func (o ListVolumesTypesResultOutput) Items() ScalewayInstanceV1ListVolumesTypesResponseOutput {
	return o.ApplyT(func(v ListVolumesTypesResult) ScalewayInstanceV1ListVolumesTypesResponse { return v.Items }).(ScalewayInstanceV1ListVolumesTypesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVolumesTypesResultOutput{})
}
