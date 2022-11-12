// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rules

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityGroupRule struct {
	pulumi.CustomResourceState

	Action ActionOutput `pulumi:"action"`
	// The beginning of the range of ports to apply this rule to (inclusive)
	Dest_port_from pulumi.Float64PtrOutput `pulumi:"dest_port_from"`
	// The end of the range of ports to apply this rule to (inclusive)
	Dest_port_to pulumi.Float64PtrOutput `pulumi:"dest_port_to"`
	Direction    DirectionOutput         `pulumi:"direction"`
	// Indicates if this rule is editable (will be ignored)
	Editable pulumi.BoolPtrOutput `pulumi:"editable"`
	// (IP network)
	Ip_range pulumi.StringOutput `pulumi:"ip_range"`
	// The position of this rule in the security group rules list
	Position pulumi.Float64PtrOutput `pulumi:"position"`
	Protocol ProtocolOutput          `pulumi:"protocol"`
}

// NewSecurityGroupRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroupRule(ctx *pulumi.Context,
	name string, args *SecurityGroupRuleArgs, opts ...pulumi.ResourceOption) (*SecurityGroupRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip_range == nil {
		return nil, errors.New("invalid value for required argument 'Ip_range'")
	}
	if isZero(args.Action) {
		args.Action = Action("accept")
	}
	if isZero(args.Direction) {
		args.Direction = Direction("inbound")
	}
	if isZero(args.Protocol) {
		args.Protocol = Protocol("TCP")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SecurityGroupRule
	err := ctx.RegisterResource("scaleway-instances:rules:SecurityGroupRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroupRule gets an existing SecurityGroupRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroupRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupRuleState, opts ...pulumi.ResourceOption) (*SecurityGroupRule, error) {
	var resource SecurityGroupRule
	err := ctx.ReadResource("scaleway-instances:rules:SecurityGroupRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroupRule resources.
type securityGroupRuleState struct {
}

type SecurityGroupRuleState struct {
}

func (SecurityGroupRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupRuleState)(nil)).Elem()
}

type securityGroupRuleArgs struct {
	Action Action `pulumi:"action"`
	// The beginning of the range of ports to apply this rule to (inclusive)
	Dest_port_from *float64 `pulumi:"dest_port_from"`
	// The end of the range of ports to apply this rule to (inclusive)
	Dest_port_to *float64  `pulumi:"dest_port_to"`
	Direction    Direction `pulumi:"direction"`
	// Indicates if this rule is editable (will be ignored)
	Editable *bool `pulumi:"editable"`
	// UUID of the security group
	Id *string `pulumi:"id"`
	// (IP network)
	Ip_range string `pulumi:"ip_range"`
	// The position of this rule in the security group rules list
	Position *float64 `pulumi:"position"`
	Protocol Protocol `pulumi:"protocol"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a SecurityGroupRule resource.
type SecurityGroupRuleArgs struct {
	Action ActionInput
	// The beginning of the range of ports to apply this rule to (inclusive)
	Dest_port_from pulumi.Float64PtrInput
	// The end of the range of ports to apply this rule to (inclusive)
	Dest_port_to pulumi.Float64PtrInput
	Direction    DirectionInput
	// Indicates if this rule is editable (will be ignored)
	Editable pulumi.BoolPtrInput
	// UUID of the security group
	Id pulumi.StringPtrInput
	// (IP network)
	Ip_range pulumi.StringInput
	// The position of this rule in the security group rules list
	Position pulumi.Float64PtrInput
	Protocol ProtocolInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (SecurityGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupRuleArgs)(nil)).Elem()
}

type SecurityGroupRuleInput interface {
	pulumi.Input

	ToSecurityGroupRuleOutput() SecurityGroupRuleOutput
	ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput
}

func (*SecurityGroupRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupRule)(nil)).Elem()
}

func (i *SecurityGroupRule) ToSecurityGroupRuleOutput() SecurityGroupRuleOutput {
	return i.ToSecurityGroupRuleOutputWithContext(context.Background())
}

func (i *SecurityGroupRule) ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRuleOutput)
}

type SecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (SecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupRule)(nil)).Elem()
}

func (o SecurityGroupRuleOutput) ToSecurityGroupRuleOutput() SecurityGroupRuleOutput {
	return o
}

func (o SecurityGroupRuleOutput) ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput {
	return o
}

func (o SecurityGroupRuleOutput) Action() ActionOutput {
	return o.ApplyT(func(v *SecurityGroupRule) ActionOutput { return v.Action }).(ActionOutput)
}

// The beginning of the range of ports to apply this rule to (inclusive)
func (o SecurityGroupRuleOutput) Dest_port_from() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.Float64PtrOutput { return v.Dest_port_from }).(pulumi.Float64PtrOutput)
}

// The end of the range of ports to apply this rule to (inclusive)
func (o SecurityGroupRuleOutput) Dest_port_to() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.Float64PtrOutput { return v.Dest_port_to }).(pulumi.Float64PtrOutput)
}

func (o SecurityGroupRuleOutput) Direction() DirectionOutput {
	return o.ApplyT(func(v *SecurityGroupRule) DirectionOutput { return v.Direction }).(DirectionOutput)
}

// Indicates if this rule is editable (will be ignored)
func (o SecurityGroupRuleOutput) Editable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.BoolPtrOutput { return v.Editable }).(pulumi.BoolPtrOutput)
}

// (IP network)
func (o SecurityGroupRuleOutput) Ip_range() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringOutput { return v.Ip_range }).(pulumi.StringOutput)
}

// The position of this rule in the security group rules list
func (o SecurityGroupRuleOutput) Position() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.Float64PtrOutput { return v.Position }).(pulumi.Float64PtrOutput)
}

func (o SecurityGroupRuleOutput) Protocol() ProtocolOutput {
	return o.ApplyT(func(v *SecurityGroupRule) ProtocolOutput { return v.Protocol }).(ProtocolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRuleInput)(nil)).Elem(), &SecurityGroupRule{})
	pulumi.RegisterOutputType(SecurityGroupRuleOutput{})
}