// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servers

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPlacementGroupServers(ctx *pulumi.Context, args *LookupPlacementGroupServersArgs, opts ...pulumi.InvokeOption) (*LookupPlacementGroupServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPlacementGroupServersResult
	err := ctx.Invoke("scaleway-instances:servers:getPlacementGroupServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlacementGroupServersArgs struct {
	// UUID of the placement group
	PlacementGroupId string `pulumi:"placementGroupId"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupPlacementGroupServersResult struct {
	Servers []ScalewayInstanceV1PlacementGroupServer `pulumi:"servers"`
}

func LookupPlacementGroupServersOutput(ctx *pulumi.Context, args LookupPlacementGroupServersOutputArgs, opts ...pulumi.InvokeOption) LookupPlacementGroupServersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPlacementGroupServersResultOutput, error) {
			args := v.(LookupPlacementGroupServersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway-instances:servers:getPlacementGroupServers", args, LookupPlacementGroupServersResultOutput{}, options).(LookupPlacementGroupServersResultOutput), nil
		}).(LookupPlacementGroupServersResultOutput)
}

type LookupPlacementGroupServersOutputArgs struct {
	// UUID of the placement group
	PlacementGroupId pulumi.StringInput `pulumi:"placementGroupId"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupPlacementGroupServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementGroupServersArgs)(nil)).Elem()
}

type LookupPlacementGroupServersResultOutput struct{ *pulumi.OutputState }

func (LookupPlacementGroupServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementGroupServersResult)(nil)).Elem()
}

func (o LookupPlacementGroupServersResultOutput) ToLookupPlacementGroupServersResultOutput() LookupPlacementGroupServersResultOutput {
	return o
}

func (o LookupPlacementGroupServersResultOutput) ToLookupPlacementGroupServersResultOutputWithContext(ctx context.Context) LookupPlacementGroupServersResultOutput {
	return o
}

func (o LookupPlacementGroupServersResultOutput) Servers() ScalewayInstanceV1PlacementGroupServerArrayOutput {
	return o.ApplyT(func(v LookupPlacementGroupServersResult) []ScalewayInstanceV1PlacementGroupServer { return v.Servers }).(ScalewayInstanceV1PlacementGroupServerArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlacementGroupServersResultOutput{})
}
