// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rules

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityGroupRule(ctx *pulumi.Context, args *LookupSecurityGroupRuleArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityGroupRuleResult
	err := ctx.Invoke("scaleway-instances:rules:getSecurityGroupRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSecurityGroupRuleArgs struct {
	Id              string `pulumi:"id"`
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type LookupSecurityGroupRuleResult struct {
	Items ScalewayInstanceV1GetSecurityGroupRuleResponse `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupSecurityGroupRuleResult
func (val *LookupSecurityGroupRuleResult) Defaults() *LookupSecurityGroupRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupSecurityGroupRuleOutput(ctx *pulumi.Context, args LookupSecurityGroupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupRuleResult, error) {
			args := v.(LookupSecurityGroupRuleArgs)
			r, err := LookupSecurityGroupRule(ctx, &args, opts...)
			var s LookupSecurityGroupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityGroupRuleResultOutput)
}

type LookupSecurityGroupRuleOutputArgs struct {
	Id              pulumi.StringInput `pulumi:"id"`
	SecurityGroupId pulumi.StringInput `pulumi:"securityGroupId"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupSecurityGroupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupRuleArgs)(nil)).Elem()
}

type LookupSecurityGroupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupRuleResult)(nil)).Elem()
}

func (o LookupSecurityGroupRuleResultOutput) ToLookupSecurityGroupRuleResultOutput() LookupSecurityGroupRuleResultOutput {
	return o
}

func (o LookupSecurityGroupRuleResultOutput) ToLookupSecurityGroupRuleResultOutputWithContext(ctx context.Context) LookupSecurityGroupRuleResultOutput {
	return o
}

func (o LookupSecurityGroupRuleResultOutput) Items() ScalewayInstanceV1GetSecurityGroupRuleResponseOutput {
	return o.ApplyT(func(v LookupSecurityGroupRuleResult) ScalewayInstanceV1GetSecurityGroupRuleResponse { return v.Items }).(ScalewayInstanceV1GetSecurityGroupRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupRuleResultOutput{})
}
