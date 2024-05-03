// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rules

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1GetSecurityGroupRuleResponse struct {
	Rule *ScalewayInstanceV1SecurityGroupRule `pulumi:"rule"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1GetSecurityGroupRuleResponse
func (val *ScalewayInstanceV1GetSecurityGroupRuleResponse) Defaults() *ScalewayInstanceV1GetSecurityGroupRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Rule = tmp.Rule.Defaults()

	return &tmp
}

type ScalewayInstanceV1GetSecurityGroupRuleResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetSecurityGroupRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetSecurityGroupRuleResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetSecurityGroupRuleResponseOutput) ToScalewayInstanceV1GetSecurityGroupRuleResponseOutput() ScalewayInstanceV1GetSecurityGroupRuleResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetSecurityGroupRuleResponseOutput) ToScalewayInstanceV1GetSecurityGroupRuleResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetSecurityGroupRuleResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetSecurityGroupRuleResponseOutput) Rule() ScalewayInstanceV1SecurityGroupRulePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetSecurityGroupRuleResponse) *ScalewayInstanceV1SecurityGroupRule {
		return v.Rule
	}).(ScalewayInstanceV1SecurityGroupRulePtrOutput)
}

type ScalewayInstanceV1ListSecurityGroupRulesResponse struct {
	// List of security rules
	Rules []ScalewayInstanceV1SecurityGroupRule `pulumi:"rules"`
}

type ScalewayInstanceV1ListSecurityGroupRulesResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListSecurityGroupRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListSecurityGroupRulesResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1ListSecurityGroupRulesResponseOutput) ToScalewayInstanceV1ListSecurityGroupRulesResponseOutput() ScalewayInstanceV1ListSecurityGroupRulesResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListSecurityGroupRulesResponseOutput) ToScalewayInstanceV1ListSecurityGroupRulesResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1ListSecurityGroupRulesResponseOutput {
	return o
}

// List of security rules
func (o ScalewayInstanceV1ListSecurityGroupRulesResponseOutput) Rules() ScalewayInstanceV1SecurityGroupRuleArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ListSecurityGroupRulesResponse) []ScalewayInstanceV1SecurityGroupRule {
		return v.Rules
	}).(ScalewayInstanceV1SecurityGroupRuleArrayOutput)
}

type ScalewayInstanceV1SecurityGroupRule struct {
	Action       *ScalewayInstanceV1SecurityGroupRuleAction    `pulumi:"action"`
	DestPortFrom *float64                                      `pulumi:"destPortFrom"`
	DestPortTo   *float64                                      `pulumi:"destPortTo"`
	Direction    *ScalewayInstanceV1SecurityGroupRuleDirection `pulumi:"direction"`
	Editable     *bool                                         `pulumi:"editable"`
	Id           *string                                       `pulumi:"id"`
	// (IP network)
	IpRange  *string                                      `pulumi:"ipRange"`
	Position *float64                                     `pulumi:"position"`
	Protocol *ScalewayInstanceV1SecurityGroupRuleProtocol `pulumi:"protocol"`
	Zone     *string                                      `pulumi:"zone"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1SecurityGroupRule
func (val *ScalewayInstanceV1SecurityGroupRule) Defaults() *ScalewayInstanceV1SecurityGroupRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Action == nil {
		action_ := ScalewayInstanceV1SecurityGroupRuleAction("accept")
		tmp.Action = &action_
	}
	if tmp.Direction == nil {
		direction_ := ScalewayInstanceV1SecurityGroupRuleDirection("inbound")
		tmp.Direction = &direction_
	}
	if tmp.Protocol == nil {
		protocol_ := ScalewayInstanceV1SecurityGroupRuleProtocol("TCP")
		tmp.Protocol = &protocol_
	}
	return &tmp
}

type ScalewayInstanceV1SecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SecurityGroupRule)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) ToScalewayInstanceV1SecurityGroupRuleOutput() ScalewayInstanceV1SecurityGroupRuleOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) ToScalewayInstanceV1SecurityGroupRuleOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupRuleOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Action() ScalewayInstanceV1SecurityGroupRuleActionPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *ScalewayInstanceV1SecurityGroupRuleAction {
		return v.Action
	}).(ScalewayInstanceV1SecurityGroupRuleActionPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) DestPortFrom() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *float64 { return v.DestPortFrom }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) DestPortTo() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *float64 { return v.DestPortTo }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Direction() ScalewayInstanceV1SecurityGroupRuleDirectionPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *ScalewayInstanceV1SecurityGroupRuleDirection {
		return v.Direction
	}).(ScalewayInstanceV1SecurityGroupRuleDirectionPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Editable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *bool { return v.Editable }).(pulumi.BoolPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (IP network)
func (o ScalewayInstanceV1SecurityGroupRuleOutput) IpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *string { return v.IpRange }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Position() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *float64 { return v.Position }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Protocol() ScalewayInstanceV1SecurityGroupRuleProtocolPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *ScalewayInstanceV1SecurityGroupRuleProtocol {
		return v.Protocol
	}).(ScalewayInstanceV1SecurityGroupRuleProtocolPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SecurityGroupRulePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SecurityGroupRule)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) ToScalewayInstanceV1SecurityGroupRulePtrOutput() ScalewayInstanceV1SecurityGroupRulePtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) ToScalewayInstanceV1SecurityGroupRulePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupRulePtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Elem() ScalewayInstanceV1SecurityGroupRuleOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) ScalewayInstanceV1SecurityGroupRule {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SecurityGroupRule
		return ret
	}).(ScalewayInstanceV1SecurityGroupRuleOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Action() ScalewayInstanceV1SecurityGroupRuleActionPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *ScalewayInstanceV1SecurityGroupRuleAction {
		if v == nil {
			return nil
		}
		return v.Action
	}).(ScalewayInstanceV1SecurityGroupRuleActionPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) DestPortFrom() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *float64 {
		if v == nil {
			return nil
		}
		return v.DestPortFrom
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) DestPortTo() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *float64 {
		if v == nil {
			return nil
		}
		return v.DestPortTo
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Direction() ScalewayInstanceV1SecurityGroupRuleDirectionPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *ScalewayInstanceV1SecurityGroupRuleDirection {
		if v == nil {
			return nil
		}
		return v.Direction
	}).(ScalewayInstanceV1SecurityGroupRuleDirectionPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Editable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *bool {
		if v == nil {
			return nil
		}
		return v.Editable
	}).(pulumi.BoolPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// (IP network)
func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) IpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *string {
		if v == nil {
			return nil
		}
		return v.IpRange
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Position() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *float64 {
		if v == nil {
			return nil
		}
		return v.Position
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Protocol() ScalewayInstanceV1SecurityGroupRuleProtocolPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *ScalewayInstanceV1SecurityGroupRuleProtocol {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(ScalewayInstanceV1SecurityGroupRuleProtocolPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *string {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SecurityGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1SecurityGroupRule)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupRuleArrayOutput) ToScalewayInstanceV1SecurityGroupRuleArrayOutput() ScalewayInstanceV1SecurityGroupRuleArrayOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupRuleArrayOutput) ToScalewayInstanceV1SecurityGroupRuleArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupRuleArrayOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupRuleArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1SecurityGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1SecurityGroupRule {
		return vs[0].([]ScalewayInstanceV1SecurityGroupRule)[vs[1].(int)]
	}).(ScalewayInstanceV1SecurityGroupRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1GetSecurityGroupRuleResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListSecurityGroupRulesResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupRulePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupRuleArrayOutput{})
}
