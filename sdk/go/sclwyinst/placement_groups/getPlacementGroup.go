// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package placement_groups

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPlacementGroup(ctx *pulumi.Context, args *LookupPlacementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPlacementGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPlacementGroupResult
	err := ctx.Invoke("scaleway-instances:placement_groups:getPlacementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPlacementGroupArgs struct {
	// UUID of the placement group you want to get
	Id string `pulumi:"id"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupPlacementGroupResult struct {
	PlacementGroup *ScalewayInstanceV1PlacementGroup `pulumi:"placementGroup"`
}

// Defaults sets the appropriate defaults for LookupPlacementGroupResult
func (val *LookupPlacementGroupResult) Defaults() *LookupPlacementGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PlacementGroup = tmp.PlacementGroup.Defaults()

	return &tmp
}

func LookupPlacementGroupOutput(ctx *pulumi.Context, args LookupPlacementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPlacementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlacementGroupResult, error) {
			args := v.(LookupPlacementGroupArgs)
			r, err := LookupPlacementGroup(ctx, &args, opts...)
			var s LookupPlacementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlacementGroupResultOutput)
}

type LookupPlacementGroupOutputArgs struct {
	// UUID of the placement group you want to get
	Id pulumi.StringInput `pulumi:"id"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupPlacementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementGroupArgs)(nil)).Elem()
}

type LookupPlacementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPlacementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementGroupResult)(nil)).Elem()
}

func (o LookupPlacementGroupResultOutput) ToLookupPlacementGroupResultOutput() LookupPlacementGroupResultOutput {
	return o
}

func (o LookupPlacementGroupResultOutput) ToLookupPlacementGroupResultOutputWithContext(ctx context.Context) LookupPlacementGroupResultOutput {
	return o
}

func (o LookupPlacementGroupResultOutput) PlacementGroup() ScalewayInstanceV1PlacementGroupPtrOutput {
	return o.ApplyT(func(v LookupPlacementGroupResult) *ScalewayInstanceV1PlacementGroup { return v.PlacementGroup }).(ScalewayInstanceV1PlacementGroupPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlacementGroupResultOutput{})
}
