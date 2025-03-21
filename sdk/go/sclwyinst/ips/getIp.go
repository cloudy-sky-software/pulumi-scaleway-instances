// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ips

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIp(ctx *pulumi.Context, args *LookupIpArgs, opts ...pulumi.InvokeOption) (*LookupIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpResult
	err := ctx.Invoke("scaleway-instances:ips:getIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpArgs struct {
	// The IP ID or address to get
	Id string `pulumi:"id"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupIpResult struct {
	Ip *ScalewayInstanceV1Ip `pulumi:"ip"`
}

func LookupIpOutput(ctx *pulumi.Context, args LookupIpOutputArgs, opts ...pulumi.InvokeOption) LookupIpResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIpResultOutput, error) {
			args := v.(LookupIpArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway-instances:ips:getIp", args, LookupIpResultOutput{}, options).(LookupIpResultOutput), nil
		}).(LookupIpResultOutput)
}

type LookupIpOutputArgs struct {
	// The IP ID or address to get
	Id pulumi.StringInput `pulumi:"id"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpArgs)(nil)).Elem()
}

type LookupIpResultOutput struct{ *pulumi.OutputState }

func (LookupIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpResult)(nil)).Elem()
}

func (o LookupIpResultOutput) ToLookupIpResultOutput() LookupIpResultOutput {
	return o
}

func (o LookupIpResultOutput) ToLookupIpResultOutputWithContext(ctx context.Context) LookupIpResultOutput {
	return o
}

func (o LookupIpResultOutput) Ip() ScalewayInstanceV1IpPtrOutput {
	return o.ApplyT(func(v LookupIpResult) *ScalewayInstanceV1Ip { return v.Ip }).(ScalewayInstanceV1IpPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpResultOutput{})
}
