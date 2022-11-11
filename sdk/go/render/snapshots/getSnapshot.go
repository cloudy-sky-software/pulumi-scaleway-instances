// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snapshots

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotResult
	err := ctx.Invoke("scaleway-instances:snapshots:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSnapshotArgs struct {
	// UUID of the snapshot you want to get
	Id string `pulumi:"id"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupSnapshotResult struct {
	Items ScalewayInstanceV1GetSnapshotResponse `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupSnapshotResult
func (val *LookupSnapshotResult) Defaults() *LookupSnapshotResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

type LookupSnapshotOutputArgs struct {
	// UUID of the snapshot you want to get
	Id pulumi.StringInput `pulumi:"id"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) Items() ScalewayInstanceV1GetSnapshotResponseOutput {
	return o.ApplyT(func(v LookupSnapshotResult) ScalewayInstanceV1GetSnapshotResponse { return v.Items }).(ScalewayInstanceV1GetSnapshotResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
