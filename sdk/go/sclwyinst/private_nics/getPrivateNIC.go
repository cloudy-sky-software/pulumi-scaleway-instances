// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_nics

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateNIC(ctx *pulumi.Context, args *LookupPrivateNICArgs, opts ...pulumi.InvokeOption) (*LookupPrivateNICResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateNICResult
	err := ctx.Invoke("scaleway-instances:private_nics:getPrivateNIC", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateNICArgs struct {
	Id       string `pulumi:"id"`
	ServerId string `pulumi:"serverId"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupPrivateNICResult struct {
	PrivateNic *ScalewayInstanceV1PrivateNIC `pulumi:"privateNic"`
}

// Defaults sets the appropriate defaults for LookupPrivateNICResult
func (val *LookupPrivateNICResult) Defaults() *LookupPrivateNICResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateNic = tmp.PrivateNic.Defaults()

	return &tmp
}
func LookupPrivateNICOutput(ctx *pulumi.Context, args LookupPrivateNICOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateNICResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPrivateNICResultOutput, error) {
			args := v.(LookupPrivateNICArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway-instances:private_nics:getPrivateNIC", args, LookupPrivateNICResultOutput{}, options).(LookupPrivateNICResultOutput), nil
		}).(LookupPrivateNICResultOutput)
}

type LookupPrivateNICOutputArgs struct {
	Id       pulumi.StringInput `pulumi:"id"`
	ServerId pulumi.StringInput `pulumi:"serverId"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupPrivateNICOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateNICArgs)(nil)).Elem()
}

type LookupPrivateNICResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateNICResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateNICResult)(nil)).Elem()
}

func (o LookupPrivateNICResultOutput) ToLookupPrivateNICResultOutput() LookupPrivateNICResultOutput {
	return o
}

func (o LookupPrivateNICResultOutput) ToLookupPrivateNICResultOutputWithContext(ctx context.Context) LookupPrivateNICResultOutput {
	return o
}

func (o LookupPrivateNICResultOutput) PrivateNic() ScalewayInstanceV1PrivateNICPtrOutput {
	return o.ApplyT(func(v LookupPrivateNICResult) *ScalewayInstanceV1PrivateNIC { return v.PrivateNic }).(ScalewayInstanceV1PrivateNICPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateNICResultOutput{})
}
