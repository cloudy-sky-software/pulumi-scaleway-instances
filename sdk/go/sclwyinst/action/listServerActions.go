// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package action

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServerActions(ctx *pulumi.Context, args *ListServerActionsArgs, opts ...pulumi.InvokeOption) (*ListServerActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListServerActionsResult
	err := ctx.Invoke("scaleway-instances:action:listServerActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServerActionsArgs struct {
	Server_id string `pulumi:"server_id"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type ListServerActionsResult struct {
	Items ScalewayInstanceV1ListServerActionsResponse `pulumi:"items"`
}

func ListServerActionsOutput(ctx *pulumi.Context, args ListServerActionsOutputArgs, opts ...pulumi.InvokeOption) ListServerActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServerActionsResult, error) {
			args := v.(ListServerActionsArgs)
			r, err := ListServerActions(ctx, &args, opts...)
			var s ListServerActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServerActionsResultOutput)
}

type ListServerActionsOutputArgs struct {
	Server_id pulumi.StringInput `pulumi:"server_id"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (ListServerActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerActionsArgs)(nil)).Elem()
}

type ListServerActionsResultOutput struct{ *pulumi.OutputState }

func (ListServerActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerActionsResult)(nil)).Elem()
}

func (o ListServerActionsResultOutput) ToListServerActionsResultOutput() ListServerActionsResultOutput {
	return o
}

func (o ListServerActionsResultOutput) ToListServerActionsResultOutputWithContext(ctx context.Context) ListServerActionsResultOutput {
	return o
}

func (o ListServerActionsResultOutput) Items() ScalewayInstanceV1ListServerActionsResponseOutput {
	return o.ApplyT(func(v ListServerActionsResult) ScalewayInstanceV1ListServerActionsResponse { return v.Items }).(ScalewayInstanceV1ListServerActionsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServerActionsResultOutput{})
}
