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

func ListVolumes(ctx *pulumi.Context, args *ListVolumesArgs, opts ...pulumi.InvokeOption) (*ListVolumesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListVolumesResult
	err := ctx.Invoke("scaleway-instances:volumes:listVolumes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVolumesArgs struct {
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type ListVolumesResult struct {
	Items ScalewayInstanceV1ListVolumesResponse `pulumi:"items"`
}

func ListVolumesOutput(ctx *pulumi.Context, args ListVolumesOutputArgs, opts ...pulumi.InvokeOption) ListVolumesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVolumesResult, error) {
			args := v.(ListVolumesArgs)
			r, err := ListVolumes(ctx, &args, opts...)
			var s ListVolumesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVolumesResultOutput)
}

type ListVolumesOutputArgs struct {
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (ListVolumesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumesArgs)(nil)).Elem()
}

type ListVolumesResultOutput struct{ *pulumi.OutputState }

func (ListVolumesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumesResult)(nil)).Elem()
}

func (o ListVolumesResultOutput) ToListVolumesResultOutput() ListVolumesResultOutput {
	return o
}

func (o ListVolumesResultOutput) ToListVolumesResultOutputWithContext(ctx context.Context) ListVolumesResultOutput {
	return o
}

func (o ListVolumesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListVolumesResult] {
	return pulumix.Output[ListVolumesResult]{
		OutputState: o.OutputState,
	}
}

func (o ListVolumesResultOutput) Items() ScalewayInstanceV1ListVolumesResponseOutput {
	return o.ApplyT(func(v ListVolumesResult) ScalewayInstanceV1ListVolumesResponse { return v.Items }).(ScalewayInstanceV1ListVolumesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVolumesResultOutput{})
}
