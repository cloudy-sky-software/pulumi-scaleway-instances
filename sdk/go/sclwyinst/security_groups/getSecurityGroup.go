// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security_groups

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityGroup(ctx *pulumi.Context, args *LookupSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityGroupResult
	err := ctx.Invoke("scaleway-instances:security_groups:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSecurityGroupArgs struct {
	// UUID of the security group you want to get
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupSecurityGroupResult struct {
	SecurityGroup *ScalewayInstanceV1SecurityGroup `pulumi:"securityGroup"`
}

// Defaults sets the appropriate defaults for LookupSecurityGroupResult
func (val *LookupSecurityGroupResult) Defaults() *LookupSecurityGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SecurityGroup = tmp.SecurityGroup.Defaults()

	return &tmp
}
func LookupSecurityGroupOutput(ctx *pulumi.Context, args LookupSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupResultOutput, error) {
			args := v.(LookupSecurityGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway-instances:security_groups:getSecurityGroup", args, LookupSecurityGroupResultOutput{}, options).(LookupSecurityGroupResultOutput), nil
		}).(LookupSecurityGroupResultOutput)
}

type LookupSecurityGroupOutputArgs struct {
	// UUID of the security group you want to get
	SecurityGroupId pulumi.StringInput `pulumi:"securityGroupId"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupArgs)(nil)).Elem()
}

type LookupSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupResult)(nil)).Elem()
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutput() LookupSecurityGroupResultOutput {
	return o
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutputWithContext(ctx context.Context) LookupSecurityGroupResultOutput {
	return o
}

func (o LookupSecurityGroupResultOutput) SecurityGroup() ScalewayInstanceV1SecurityGroupPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) *ScalewayInstanceV1SecurityGroup { return v.SecurityGroup }).(ScalewayInstanceV1SecurityGroupPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupResultOutput{})
}
