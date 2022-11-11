// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rules

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GoogleProtobufUInt32Value struct {
}

type GoogleProtobufUInt32ValueOutput struct{ *pulumi.OutputState }

func (GoogleProtobufUInt32ValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleProtobufUInt32Value)(nil)).Elem()
}

func (o GoogleProtobufUInt32ValueOutput) ToGoogleProtobufUInt32ValueOutput() GoogleProtobufUInt32ValueOutput {
	return o
}

func (o GoogleProtobufUInt32ValueOutput) ToGoogleProtobufUInt32ValueOutputWithContext(ctx context.Context) GoogleProtobufUInt32ValueOutput {
	return o
}

type GoogleProtobufUInt32ValuePtrOutput struct{ *pulumi.OutputState }

func (GoogleProtobufUInt32ValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleProtobufUInt32Value)(nil)).Elem()
}

func (o GoogleProtobufUInt32ValuePtrOutput) ToGoogleProtobufUInt32ValuePtrOutput() GoogleProtobufUInt32ValuePtrOutput {
	return o
}

func (o GoogleProtobufUInt32ValuePtrOutput) ToGoogleProtobufUInt32ValuePtrOutputWithContext(ctx context.Context) GoogleProtobufUInt32ValuePtrOutput {
	return o
}

func (o GoogleProtobufUInt32ValuePtrOutput) Elem() GoogleProtobufUInt32ValueOutput {
	return o.ApplyT(func(v *GoogleProtobufUInt32Value) GoogleProtobufUInt32Value {
		if v != nil {
			return *v
		}
		var ret GoogleProtobufUInt32Value
		return ret
	}).(GoogleProtobufUInt32ValueOutput)
}

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
	Action         *ScalewayInstanceV1SecurityGroupRuleAction    `pulumi:"action"`
	Dest_port_from *GoogleProtobufUInt32Value                    `pulumi:"dest_port_from"`
	Dest_port_to   *GoogleProtobufUInt32Value                    `pulumi:"dest_port_to"`
	Direction      *ScalewayInstanceV1SecurityGroupRuleDirection `pulumi:"direction"`
	Editable       *bool                                         `pulumi:"editable"`
	Id             *string                                       `pulumi:"id"`
	// (IP network)
	Ip_range *string                                      `pulumi:"ip_range"`
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
	if isZero(tmp.Action) {
		action_ := ScalewayInstanceV1SecurityGroupRuleAction("accept")
		tmp.Action = &action_
	}
	if isZero(tmp.Direction) {
		direction_ := ScalewayInstanceV1SecurityGroupRuleDirection("inbound")
		tmp.Direction = &direction_
	}
	if isZero(tmp.Protocol) {
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

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Dest_port_from() GoogleProtobufUInt32ValuePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *GoogleProtobufUInt32Value { return v.Dest_port_from }).(GoogleProtobufUInt32ValuePtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRuleOutput) Dest_port_to() GoogleProtobufUInt32ValuePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *GoogleProtobufUInt32Value { return v.Dest_port_to }).(GoogleProtobufUInt32ValuePtrOutput)
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
func (o ScalewayInstanceV1SecurityGroupRuleOutput) Ip_range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SecurityGroupRule) *string { return v.Ip_range }).(pulumi.StringPtrOutput)
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

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Dest_port_from() GoogleProtobufUInt32ValuePtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *GoogleProtobufUInt32Value {
		if v == nil {
			return nil
		}
		return v.Dest_port_from
	}).(GoogleProtobufUInt32ValuePtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Dest_port_to() GoogleProtobufUInt32ValuePtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *GoogleProtobufUInt32Value {
		if v == nil {
			return nil
		}
		return v.Dest_port_to
	}).(GoogleProtobufUInt32ValuePtrOutput)
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
func (o ScalewayInstanceV1SecurityGroupRulePtrOutput) Ip_range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupRule) *string {
		if v == nil {
			return nil
		}
		return v.Ip_range
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

type ScalewayInstanceV1SetSecurityGroupRulesRequestRule struct {
	// Action to apply when the rule matches a packet
	Action *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction `pulumi:"action"`
	// Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
	Dest_port_from *float64 `pulumi:"dest_port_from"`
	// End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
	Dest_port_to *float64 `pulumi:"dest_port_to"`
	// Direction the rule applies to
	Direction *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection `pulumi:"direction"`
	// Indicates if this rule is editable. Rules with the value false will be ignored
	Editable *bool `pulumi:"editable"`
	// UUID of the security rule to update. If no value is provided, a new rule will be created
	Id *string `pulumi:"id"`
	// The range of IP address this rules applies to (IP network)
	Ip_range *string `pulumi:"ip_range"`
	// Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
	Position *float64 `pulumi:"position"`
	// Protocol family this rule applies to
	Protocol *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol `pulumi:"protocol"`
	// Zone of the rule. This field is ignored
	Zone *string `pulumi:"zone"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1SetSecurityGroupRulesRequestRule
func (val *ScalewayInstanceV1SetSecurityGroupRulesRequestRule) Defaults() *ScalewayInstanceV1SetSecurityGroupRulesRequestRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction("accept")
		tmp.Action = &action_
	}
	if isZero(tmp.Direction) {
		direction_ := ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection("inbound")
		tmp.Direction = &direction_
	}
	if isZero(tmp.Protocol) {
		protocol_ := ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol("TCP")
		tmp.Protocol = &protocol_
	}
	return &tmp
}

// ScalewayInstanceV1SetSecurityGroupRulesRequestRuleInput is an input type that accepts ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs and ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput values.
// You can construct a concrete instance of `ScalewayInstanceV1SetSecurityGroupRulesRequestRuleInput` via:
//
//	ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs{...}
type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleInput interface {
	pulumi.Input

	ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput
	ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutputWithContext(context.Context) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput
}

type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs struct {
	// Action to apply when the rule matches a packet
	Action ScalewayInstanceV1SetSecurityGroupRulesRequestRuleActionPtrInput `pulumi:"action"`
	// Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
	Dest_port_from pulumi.Float64PtrInput `pulumi:"dest_port_from"`
	// End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
	Dest_port_to pulumi.Float64PtrInput `pulumi:"dest_port_to"`
	// Direction the rule applies to
	Direction ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirectionPtrInput `pulumi:"direction"`
	// Indicates if this rule is editable. Rules with the value false will be ignored
	Editable pulumi.BoolPtrInput `pulumi:"editable"`
	// UUID of the security rule to update. If no value is provided, a new rule will be created
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The range of IP address this rules applies to (IP network)
	Ip_range pulumi.StringPtrInput `pulumi:"ip_range"`
	// Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
	Position pulumi.Float64PtrInput `pulumi:"position"`
	// Protocol family this rule applies to
	Protocol ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocolPtrInput `pulumi:"protocol"`
	// Zone of the rule. This field is ignored
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs
func (val *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs) Defaults() *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		tmp.Action = ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction("accept")
	}
	if isZero(tmp.Direction) {
		tmp.Direction = ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection("inbound")
	}
	if isZero(tmp.Protocol) {
		tmp.Protocol = ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol("TCP")
	}
	return &tmp
}
func (ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SetSecurityGroupRulesRequestRule)(nil)).Elem()
}

func (i ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput {
	return i.ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutputWithContext(context.Background())
}

func (i ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutputWithContext(ctx context.Context) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput)
}

// ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayInput is an input type that accepts ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray and ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput values.
// You can construct a concrete instance of `ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayInput` via:
//
//	ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray{ ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs{...} }
type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayInput interface {
	pulumi.Input

	ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput
	ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutputWithContext(context.Context) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput
}

type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray []ScalewayInstanceV1SetSecurityGroupRulesRequestRuleInput

func (ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1SetSecurityGroupRulesRequestRule)(nil)).Elem()
}

func (i ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput {
	return i.ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutputWithContext(context.Background())
}

func (i ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput)
}

type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SetSecurityGroupRulesRequestRule)(nil)).Elem()
}

func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput {
	return o
}

func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutputWithContext(ctx context.Context) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput {
	return o
}

// Action to apply when the rule matches a packet
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Action() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleActionPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction {
		return v.Action
	}).(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleActionPtrOutput)
}

// Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Dest_port_from() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *float64 { return v.Dest_port_from }).(pulumi.Float64PtrOutput)
}

// End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Dest_port_to() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *float64 { return v.Dest_port_to }).(pulumi.Float64PtrOutput)
}

// Direction the rule applies to
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Direction() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirectionPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection {
		return v.Direction
	}).(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirectionPtrOutput)
}

// Indicates if this rule is editable. Rules with the value false will be ignored
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Editable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *bool { return v.Editable }).(pulumi.BoolPtrOutput)
}

// UUID of the security rule to update. If no value is provided, a new rule will be created
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The range of IP address this rules applies to (IP network)
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Ip_range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *string { return v.Ip_range }).(pulumi.StringPtrOutput)
}

// Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Position() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *float64 { return v.Position }).(pulumi.Float64PtrOutput)
}

// Protocol family this rule applies to
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Protocol() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocolPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol {
		return v.Protocol
	}).(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocolPtrOutput)
}

// Zone of the rule. This field is ignored
func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SetSecurityGroupRulesRequestRule) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1SetSecurityGroupRulesRequestRule)(nil)).Elem()
}

func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput() ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput {
	return o
}

func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput) ToScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput {
	return o
}

func (o ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1SetSecurityGroupRulesRequestRule {
		return vs[0].([]ScalewayInstanceV1SetSecurityGroupRulesRequestRule)[vs[1].(int)]
	}).(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalewayInstanceV1SetSecurityGroupRulesRequestRuleInput)(nil)).Elem(), ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayInput)(nil)).Elem(), ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArray{})
	pulumi.RegisterOutputType(GoogleProtobufUInt32ValueOutput{})
	pulumi.RegisterOutputType(GoogleProtobufUInt32ValuePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetSecurityGroupRuleResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListSecurityGroupRulesResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupRulePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupRuleArrayOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArrayOutput{})
}
