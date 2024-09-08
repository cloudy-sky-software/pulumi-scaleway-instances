// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rules

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSecurityGroupRules(ctx *pulumi.Context, args *ListSecurityGroupRulesArgs, opts ...pulumi.InvokeOption) (*ListSecurityGroupRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListSecurityGroupRulesResult
	err := ctx.Invoke("scaleway-instances:rules:listSecurityGroupRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSecurityGroupRulesArgs struct {
	// UUID of the security group
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type ListSecurityGroupRulesResult struct {
	// List of security rules
	Rules []ScalewayInstanceV1SecurityGroupRule `pulumi:"rules"`
}

func ListSecurityGroupRulesOutput(ctx *pulumi.Context, args ListSecurityGroupRulesOutputArgs, opts ...pulumi.InvokeOption) ListSecurityGroupRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSecurityGroupRulesResult, error) {
			args := v.(ListSecurityGroupRulesArgs)
			r, err := ListSecurityGroupRules(ctx, &args, opts...)
			var s ListSecurityGroupRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSecurityGroupRulesResultOutput)
}

type ListSecurityGroupRulesOutputArgs struct {
	// UUID of the security group
	SecurityGroupId pulumi.StringInput `pulumi:"securityGroupId"`
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (ListSecurityGroupRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSecurityGroupRulesArgs)(nil)).Elem()
}

type ListSecurityGroupRulesResultOutput struct{ *pulumi.OutputState }

func (ListSecurityGroupRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSecurityGroupRulesResult)(nil)).Elem()
}

func (o ListSecurityGroupRulesResultOutput) ToListSecurityGroupRulesResultOutput() ListSecurityGroupRulesResultOutput {
	return o
}

func (o ListSecurityGroupRulesResultOutput) ToListSecurityGroupRulesResultOutputWithContext(ctx context.Context) ListSecurityGroupRulesResultOutput {
	return o
}

// List of security rules
func (o ListSecurityGroupRulesResultOutput) Rules() ScalewayInstanceV1SecurityGroupRuleArrayOutput {
	return o.ApplyT(func(v ListSecurityGroupRulesResult) []ScalewayInstanceV1SecurityGroupRule { return v.Rules }).(ScalewayInstanceV1SecurityGroupRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSecurityGroupRulesResultOutput{})
}
