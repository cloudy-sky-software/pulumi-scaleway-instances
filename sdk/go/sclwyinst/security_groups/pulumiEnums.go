// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security_groups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The default inbound policy
type InboundDefaultPolicy string

const (
	InboundDefaultPolicyAccept = InboundDefaultPolicy("accept")
	InboundDefaultPolicyDrop   = InboundDefaultPolicy("drop")
)

func (InboundDefaultPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundDefaultPolicy)(nil)).Elem()
}

func (e InboundDefaultPolicy) ToInboundDefaultPolicyOutput() InboundDefaultPolicyOutput {
	return pulumi.ToOutput(e).(InboundDefaultPolicyOutput)
}

func (e InboundDefaultPolicy) ToInboundDefaultPolicyOutputWithContext(ctx context.Context) InboundDefaultPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InboundDefaultPolicyOutput)
}

func (e InboundDefaultPolicy) ToInboundDefaultPolicyPtrOutput() InboundDefaultPolicyPtrOutput {
	return e.ToInboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (e InboundDefaultPolicy) ToInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) InboundDefaultPolicyPtrOutput {
	return InboundDefaultPolicy(e).ToInboundDefaultPolicyOutputWithContext(ctx).ToInboundDefaultPolicyPtrOutputWithContext(ctx)
}

func (e InboundDefaultPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InboundDefaultPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InboundDefaultPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InboundDefaultPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InboundDefaultPolicyOutput struct{ *pulumi.OutputState }

func (InboundDefaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundDefaultPolicy)(nil)).Elem()
}

func (o InboundDefaultPolicyOutput) ToInboundDefaultPolicyOutput() InboundDefaultPolicyOutput {
	return o
}

func (o InboundDefaultPolicyOutput) ToInboundDefaultPolicyOutputWithContext(ctx context.Context) InboundDefaultPolicyOutput {
	return o
}

func (o InboundDefaultPolicyOutput) ToInboundDefaultPolicyPtrOutput() InboundDefaultPolicyPtrOutput {
	return o.ToInboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (o InboundDefaultPolicyOutput) ToInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) InboundDefaultPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InboundDefaultPolicy) *InboundDefaultPolicy {
		return &v
	}).(InboundDefaultPolicyPtrOutput)
}

func (o InboundDefaultPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[InboundDefaultPolicy] {
	return pulumix.Output[InboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o InboundDefaultPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InboundDefaultPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InboundDefaultPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InboundDefaultPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InboundDefaultPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InboundDefaultPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InboundDefaultPolicyPtrOutput struct{ *pulumi.OutputState }

func (InboundDefaultPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundDefaultPolicy)(nil)).Elem()
}

func (o InboundDefaultPolicyPtrOutput) ToInboundDefaultPolicyPtrOutput() InboundDefaultPolicyPtrOutput {
	return o
}

func (o InboundDefaultPolicyPtrOutput) ToInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) InboundDefaultPolicyPtrOutput {
	return o
}

func (o InboundDefaultPolicyPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*InboundDefaultPolicy] {
	return pulumix.Output[*InboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o InboundDefaultPolicyPtrOutput) Elem() InboundDefaultPolicyOutput {
	return o.ApplyT(func(v *InboundDefaultPolicy) InboundDefaultPolicy {
		if v != nil {
			return *v
		}
		var ret InboundDefaultPolicy
		return ret
	}).(InboundDefaultPolicyOutput)
}

func (o InboundDefaultPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InboundDefaultPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InboundDefaultPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InboundDefaultPolicyInput is an input type that accepts InboundDefaultPolicyArgs and InboundDefaultPolicyOutput values.
// You can construct a concrete instance of `InboundDefaultPolicyInput` via:
//
//	InboundDefaultPolicyArgs{...}
type InboundDefaultPolicyInput interface {
	pulumi.Input

	ToInboundDefaultPolicyOutput() InboundDefaultPolicyOutput
	ToInboundDefaultPolicyOutputWithContext(context.Context) InboundDefaultPolicyOutput
}

var inboundDefaultPolicyPtrType = reflect.TypeOf((**InboundDefaultPolicy)(nil)).Elem()

type InboundDefaultPolicyPtrInput interface {
	pulumi.Input

	ToInboundDefaultPolicyPtrOutput() InboundDefaultPolicyPtrOutput
	ToInboundDefaultPolicyPtrOutputWithContext(context.Context) InboundDefaultPolicyPtrOutput
}

type inboundDefaultPolicyPtr string

func InboundDefaultPolicyPtr(v string) InboundDefaultPolicyPtrInput {
	return (*inboundDefaultPolicyPtr)(&v)
}

func (*inboundDefaultPolicyPtr) ElementType() reflect.Type {
	return inboundDefaultPolicyPtrType
}

func (in *inboundDefaultPolicyPtr) ToInboundDefaultPolicyPtrOutput() InboundDefaultPolicyPtrOutput {
	return pulumi.ToOutput(in).(InboundDefaultPolicyPtrOutput)
}

func (in *inboundDefaultPolicyPtr) ToInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) InboundDefaultPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InboundDefaultPolicyPtrOutput)
}

func (in *inboundDefaultPolicyPtr) ToOutput(ctx context.Context) pulumix.Output[*InboundDefaultPolicy] {
	return pulumix.Output[*InboundDefaultPolicy]{
		OutputState: in.ToInboundDefaultPolicyPtrOutputWithContext(ctx).OutputState,
	}
}

// The default outbound policy
type OutboundDefaultPolicy string

const (
	OutboundDefaultPolicyAccept = OutboundDefaultPolicy("accept")
	OutboundDefaultPolicyDrop   = OutboundDefaultPolicy("drop")
)

func (OutboundDefaultPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundDefaultPolicy)(nil)).Elem()
}

func (e OutboundDefaultPolicy) ToOutboundDefaultPolicyOutput() OutboundDefaultPolicyOutput {
	return pulumi.ToOutput(e).(OutboundDefaultPolicyOutput)
}

func (e OutboundDefaultPolicy) ToOutboundDefaultPolicyOutputWithContext(ctx context.Context) OutboundDefaultPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OutboundDefaultPolicyOutput)
}

func (e OutboundDefaultPolicy) ToOutboundDefaultPolicyPtrOutput() OutboundDefaultPolicyPtrOutput {
	return e.ToOutboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (e OutboundDefaultPolicy) ToOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) OutboundDefaultPolicyPtrOutput {
	return OutboundDefaultPolicy(e).ToOutboundDefaultPolicyOutputWithContext(ctx).ToOutboundDefaultPolicyPtrOutputWithContext(ctx)
}

func (e OutboundDefaultPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutboundDefaultPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutboundDefaultPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OutboundDefaultPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OutboundDefaultPolicyOutput struct{ *pulumi.OutputState }

func (OutboundDefaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundDefaultPolicy)(nil)).Elem()
}

func (o OutboundDefaultPolicyOutput) ToOutboundDefaultPolicyOutput() OutboundDefaultPolicyOutput {
	return o
}

func (o OutboundDefaultPolicyOutput) ToOutboundDefaultPolicyOutputWithContext(ctx context.Context) OutboundDefaultPolicyOutput {
	return o
}

func (o OutboundDefaultPolicyOutput) ToOutboundDefaultPolicyPtrOutput() OutboundDefaultPolicyPtrOutput {
	return o.ToOutboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (o OutboundDefaultPolicyOutput) ToOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) OutboundDefaultPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutboundDefaultPolicy) *OutboundDefaultPolicy {
		return &v
	}).(OutboundDefaultPolicyPtrOutput)
}

func (o OutboundDefaultPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[OutboundDefaultPolicy] {
	return pulumix.Output[OutboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o OutboundDefaultPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutboundDefaultPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutboundDefaultPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutboundDefaultPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutboundDefaultPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutboundDefaultPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutboundDefaultPolicyPtrOutput struct{ *pulumi.OutputState }

func (OutboundDefaultPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundDefaultPolicy)(nil)).Elem()
}

func (o OutboundDefaultPolicyPtrOutput) ToOutboundDefaultPolicyPtrOutput() OutboundDefaultPolicyPtrOutput {
	return o
}

func (o OutboundDefaultPolicyPtrOutput) ToOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) OutboundDefaultPolicyPtrOutput {
	return o
}

func (o OutboundDefaultPolicyPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*OutboundDefaultPolicy] {
	return pulumix.Output[*OutboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o OutboundDefaultPolicyPtrOutput) Elem() OutboundDefaultPolicyOutput {
	return o.ApplyT(func(v *OutboundDefaultPolicy) OutboundDefaultPolicy {
		if v != nil {
			return *v
		}
		var ret OutboundDefaultPolicy
		return ret
	}).(OutboundDefaultPolicyOutput)
}

func (o OutboundDefaultPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutboundDefaultPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutboundDefaultPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OutboundDefaultPolicyInput is an input type that accepts OutboundDefaultPolicyArgs and OutboundDefaultPolicyOutput values.
// You can construct a concrete instance of `OutboundDefaultPolicyInput` via:
//
//	OutboundDefaultPolicyArgs{...}
type OutboundDefaultPolicyInput interface {
	pulumi.Input

	ToOutboundDefaultPolicyOutput() OutboundDefaultPolicyOutput
	ToOutboundDefaultPolicyOutputWithContext(context.Context) OutboundDefaultPolicyOutput
}

var outboundDefaultPolicyPtrType = reflect.TypeOf((**OutboundDefaultPolicy)(nil)).Elem()

type OutboundDefaultPolicyPtrInput interface {
	pulumi.Input

	ToOutboundDefaultPolicyPtrOutput() OutboundDefaultPolicyPtrOutput
	ToOutboundDefaultPolicyPtrOutputWithContext(context.Context) OutboundDefaultPolicyPtrOutput
}

type outboundDefaultPolicyPtr string

func OutboundDefaultPolicyPtr(v string) OutboundDefaultPolicyPtrInput {
	return (*outboundDefaultPolicyPtr)(&v)
}

func (*outboundDefaultPolicyPtr) ElementType() reflect.Type {
	return outboundDefaultPolicyPtrType
}

func (in *outboundDefaultPolicyPtr) ToOutboundDefaultPolicyPtrOutput() OutboundDefaultPolicyPtrOutput {
	return pulumi.ToOutput(in).(OutboundDefaultPolicyPtrOutput)
}

func (in *outboundDefaultPolicyPtr) ToOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) OutboundDefaultPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OutboundDefaultPolicyPtrOutput)
}

func (in *outboundDefaultPolicyPtr) ToOutput(ctx context.Context) pulumix.Output[*OutboundDefaultPolicy] {
	return pulumix.Output[*OutboundDefaultPolicy]{
		OutputState: in.ToOutboundDefaultPolicyPtrOutputWithContext(ctx).OutputState,
	}
}

// The default inbound policy
type ScalewayInstanceV1SecurityGroupInboundDefaultPolicy string

const (
	ScalewayInstanceV1SecurityGroupInboundDefaultPolicyAccept = ScalewayInstanceV1SecurityGroupInboundDefaultPolicy("accept")
	ScalewayInstanceV1SecurityGroupInboundDefaultPolicyDrop   = ScalewayInstanceV1SecurityGroupInboundDefaultPolicy("drop")
)

type ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SecurityGroupInboundDefaultPolicy)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput() ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput() ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput {
	return o.ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1SecurityGroupInboundDefaultPolicy) *ScalewayInstanceV1SecurityGroupInboundDefaultPolicy {
		return &v
	}).(ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1SecurityGroupInboundDefaultPolicy] {
	return pulumix.Output[ScalewayInstanceV1SecurityGroupInboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SecurityGroupInboundDefaultPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SecurityGroupInboundDefaultPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SecurityGroupInboundDefaultPolicy)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput() ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) ToScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1SecurityGroupInboundDefaultPolicy] {
	return pulumix.Output[*ScalewayInstanceV1SecurityGroupInboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) Elem() ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupInboundDefaultPolicy) ScalewayInstanceV1SecurityGroupInboundDefaultPolicy {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SecurityGroupInboundDefaultPolicy
		return ret
	}).(ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput)
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1SecurityGroupInboundDefaultPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// The default outbound policy
type ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy string

const (
	ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyAccept = ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy("accept")
	ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyDrop   = ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy("drop")
)

type ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput() ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput() ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o.ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy) *ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy {
		return &v
	}).(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy] {
	return pulumix.Output[ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput() ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) ToScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy] {
	return pulumix.Output[*ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) Elem() ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy) ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy
		return ret
	}).(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput)
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// Security group state
type ScalewayInstanceV1SecurityGroupState string

const (
	ScalewayInstanceV1SecurityGroupStateAvailable    = ScalewayInstanceV1SecurityGroupState("available")
	ScalewayInstanceV1SecurityGroupStateSyncing      = ScalewayInstanceV1SecurityGroupState("syncing")
	ScalewayInstanceV1SecurityGroupStateSyncingError = ScalewayInstanceV1SecurityGroupState("syncing_error")
)

type ScalewayInstanceV1SecurityGroupStateOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SecurityGroupState)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToScalewayInstanceV1SecurityGroupStateOutput() ScalewayInstanceV1SecurityGroupStateOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToScalewayInstanceV1SecurityGroupStateOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupStateOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToScalewayInstanceV1SecurityGroupStatePtrOutput() ScalewayInstanceV1SecurityGroupStatePtrOutput {
	return o.ToScalewayInstanceV1SecurityGroupStatePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToScalewayInstanceV1SecurityGroupStatePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1SecurityGroupState) *ScalewayInstanceV1SecurityGroupState {
		return &v
	}).(ScalewayInstanceV1SecurityGroupStatePtrOutput)
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1SecurityGroupState] {
	return pulumix.Output[ScalewayInstanceV1SecurityGroupState]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SecurityGroupState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SecurityGroupState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SecurityGroupStatePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SecurityGroupStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SecurityGroupState)(nil)).Elem()
}

func (o ScalewayInstanceV1SecurityGroupStatePtrOutput) ToScalewayInstanceV1SecurityGroupStatePtrOutput() ScalewayInstanceV1SecurityGroupStatePtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupStatePtrOutput) ToScalewayInstanceV1SecurityGroupStatePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SecurityGroupStatePtrOutput {
	return o
}

func (o ScalewayInstanceV1SecurityGroupStatePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1SecurityGroupState] {
	return pulumix.Output[*ScalewayInstanceV1SecurityGroupState]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SecurityGroupStatePtrOutput) Elem() ScalewayInstanceV1SecurityGroupStateOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SecurityGroupState) ScalewayInstanceV1SecurityGroupState {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SecurityGroupState
		return ret
	}).(ScalewayInstanceV1SecurityGroupStateOutput)
}

func (o ScalewayInstanceV1SecurityGroupStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SecurityGroupStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1SecurityGroupState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// Security group state
type State string

const (
	StateAvailable    = State("available")
	StateSyncing      = State("syncing")
	StateSyncingError = State("syncing_error")
)

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

func (o StateOutput) ToStatePtrOutput() StatePtrOutput {
	return o.ToStatePtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v State) *State {
		return &v
	}).(StatePtrOutput)
}

func (o StateOutput) ToOutput(ctx context.Context) pulumix.Output[State] {
	return pulumix.Output[State]{
		OutputState: o.OutputState,
	}
}

func (o StateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatePtrOutput struct{ *pulumi.OutputState }

func (StatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StatePtrOutput) ToStatePtrOutput() StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*State] {
	return pulumix.Output[*State]{
		OutputState: o.OutputState,
	}
}

func (o StatePtrOutput) Elem() StateOutput {
	return o.ApplyT(func(v *State) State {
		if v != nil {
			return *v
		}
		var ret State
		return ret
	}).(StateOutput)
}

func (o StatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *State) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InboundDefaultPolicyInput)(nil)).Elem(), InboundDefaultPolicy("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*InboundDefaultPolicyPtrInput)(nil)).Elem(), InboundDefaultPolicy("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundDefaultPolicyInput)(nil)).Elem(), OutboundDefaultPolicy("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundDefaultPolicyPtrInput)(nil)).Elem(), OutboundDefaultPolicy("accept"))
	pulumi.RegisterOutputType(InboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(InboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(OutboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(OutboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupStateOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupStatePtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
