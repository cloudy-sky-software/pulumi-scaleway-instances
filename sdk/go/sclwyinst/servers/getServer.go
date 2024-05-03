// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servers

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerResult
	err := ctx.Invoke("scaleway-instances:servers:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServerArgs struct {
	// UUID of the server you want to get
	Id string `pulumi:"id"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupServerResult struct {
	Items ScalewayInstanceV1GetServerResponse `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupServerResult
func (val *LookupServerResult) Defaults() *LookupServerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	// UUID of the server you want to get
	Id pulumi.StringInput `pulumi:"id"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) Items() ScalewayInstanceV1GetServerResponseOutput {
	return o.ApplyT(func(v LookupServerResult) ScalewayInstanceV1GetServerResponse { return v.Items }).(ScalewayInstanceV1GetServerResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
