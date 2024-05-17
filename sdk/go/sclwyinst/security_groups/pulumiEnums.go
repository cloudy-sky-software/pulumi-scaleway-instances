// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security_groups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

// The default inbound policy
type SecurityGroupInboundDefaultPolicy string

const (
	SecurityGroupInboundDefaultPolicyAccept = SecurityGroupInboundDefaultPolicy("accept")
	SecurityGroupInboundDefaultPolicyDrop   = SecurityGroupInboundDefaultPolicy("drop")
)

func (SecurityGroupInboundDefaultPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupInboundDefaultPolicy)(nil)).Elem()
}

func (e SecurityGroupInboundDefaultPolicy) ToSecurityGroupInboundDefaultPolicyOutput() SecurityGroupInboundDefaultPolicyOutput {
	return pulumi.ToOutput(e).(SecurityGroupInboundDefaultPolicyOutput)
}

func (e SecurityGroupInboundDefaultPolicy) ToSecurityGroupInboundDefaultPolicyOutputWithContext(ctx context.Context) SecurityGroupInboundDefaultPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityGroupInboundDefaultPolicyOutput)
}

func (e SecurityGroupInboundDefaultPolicy) ToSecurityGroupInboundDefaultPolicyPtrOutput() SecurityGroupInboundDefaultPolicyPtrOutput {
	return e.ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (e SecurityGroupInboundDefaultPolicy) ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupInboundDefaultPolicyPtrOutput {
	return SecurityGroupInboundDefaultPolicy(e).ToSecurityGroupInboundDefaultPolicyOutputWithContext(ctx).ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx)
}

func (e SecurityGroupInboundDefaultPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityGroupInboundDefaultPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityGroupInboundDefaultPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityGroupInboundDefaultPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityGroupInboundDefaultPolicyOutput struct{ *pulumi.OutputState }

func (SecurityGroupInboundDefaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupInboundDefaultPolicy)(nil)).Elem()
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToSecurityGroupInboundDefaultPolicyOutput() SecurityGroupInboundDefaultPolicyOutput {
	return o
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToSecurityGroupInboundDefaultPolicyOutputWithContext(ctx context.Context) SecurityGroupInboundDefaultPolicyOutput {
	return o
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToSecurityGroupInboundDefaultPolicyPtrOutput() SecurityGroupInboundDefaultPolicyPtrOutput {
	return o.ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupInboundDefaultPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityGroupInboundDefaultPolicy) *SecurityGroupInboundDefaultPolicy {
		return &v
	}).(SecurityGroupInboundDefaultPolicyPtrOutput)
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityGroupInboundDefaultPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityGroupInboundDefaultPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityGroupInboundDefaultPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityGroupInboundDefaultPolicyPtrOutput struct{ *pulumi.OutputState }

func (SecurityGroupInboundDefaultPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupInboundDefaultPolicy)(nil)).Elem()
}

func (o SecurityGroupInboundDefaultPolicyPtrOutput) ToSecurityGroupInboundDefaultPolicyPtrOutput() SecurityGroupInboundDefaultPolicyPtrOutput {
	return o
}

func (o SecurityGroupInboundDefaultPolicyPtrOutput) ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupInboundDefaultPolicyPtrOutput {
	return o
}

func (o SecurityGroupInboundDefaultPolicyPtrOutput) Elem() SecurityGroupInboundDefaultPolicyOutput {
	return o.ApplyT(func(v *SecurityGroupInboundDefaultPolicy) SecurityGroupInboundDefaultPolicy {
		if v != nil {
			return *v
		}
		var ret SecurityGroupInboundDefaultPolicy
		return ret
	}).(SecurityGroupInboundDefaultPolicyOutput)
}

func (o SecurityGroupInboundDefaultPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityGroupInboundDefaultPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityGroupInboundDefaultPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SecurityGroupInboundDefaultPolicyInput is an input type that accepts values of the SecurityGroupInboundDefaultPolicy enum
// A concrete instance of `SecurityGroupInboundDefaultPolicyInput` can be one of the following:
//
//	SecurityGroupInboundDefaultPolicyAccept
//	SecurityGroupInboundDefaultPolicyDrop
type SecurityGroupInboundDefaultPolicyInput interface {
	pulumi.Input

	ToSecurityGroupInboundDefaultPolicyOutput() SecurityGroupInboundDefaultPolicyOutput
	ToSecurityGroupInboundDefaultPolicyOutputWithContext(context.Context) SecurityGroupInboundDefaultPolicyOutput
}

var securityGroupInboundDefaultPolicyPtrType = reflect.TypeOf((**SecurityGroupInboundDefaultPolicy)(nil)).Elem()

type SecurityGroupInboundDefaultPolicyPtrInput interface {
	pulumi.Input

	ToSecurityGroupInboundDefaultPolicyPtrOutput() SecurityGroupInboundDefaultPolicyPtrOutput
	ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(context.Context) SecurityGroupInboundDefaultPolicyPtrOutput
}

type securityGroupInboundDefaultPolicyPtr string

func SecurityGroupInboundDefaultPolicyPtr(v string) SecurityGroupInboundDefaultPolicyPtrInput {
	return (*securityGroupInboundDefaultPolicyPtr)(&v)
}

func (*securityGroupInboundDefaultPolicyPtr) ElementType() reflect.Type {
	return securityGroupInboundDefaultPolicyPtrType
}

func (in *securityGroupInboundDefaultPolicyPtr) ToSecurityGroupInboundDefaultPolicyPtrOutput() SecurityGroupInboundDefaultPolicyPtrOutput {
	return pulumi.ToOutput(in).(SecurityGroupInboundDefaultPolicyPtrOutput)
}

func (in *securityGroupInboundDefaultPolicyPtr) ToSecurityGroupInboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupInboundDefaultPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityGroupInboundDefaultPolicyPtrOutput)
}

// The default outbound policy
type SecurityGroupOutboundDefaultPolicy string

const (
	SecurityGroupOutboundDefaultPolicyAccept = SecurityGroupOutboundDefaultPolicy("accept")
	SecurityGroupOutboundDefaultPolicyDrop   = SecurityGroupOutboundDefaultPolicy("drop")
)

func (SecurityGroupOutboundDefaultPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupOutboundDefaultPolicy)(nil)).Elem()
}

func (e SecurityGroupOutboundDefaultPolicy) ToSecurityGroupOutboundDefaultPolicyOutput() SecurityGroupOutboundDefaultPolicyOutput {
	return pulumi.ToOutput(e).(SecurityGroupOutboundDefaultPolicyOutput)
}

func (e SecurityGroupOutboundDefaultPolicy) ToSecurityGroupOutboundDefaultPolicyOutputWithContext(ctx context.Context) SecurityGroupOutboundDefaultPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityGroupOutboundDefaultPolicyOutput)
}

func (e SecurityGroupOutboundDefaultPolicy) ToSecurityGroupOutboundDefaultPolicyPtrOutput() SecurityGroupOutboundDefaultPolicyPtrOutput {
	return e.ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (e SecurityGroupOutboundDefaultPolicy) ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupOutboundDefaultPolicyPtrOutput {
	return SecurityGroupOutboundDefaultPolicy(e).ToSecurityGroupOutboundDefaultPolicyOutputWithContext(ctx).ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx)
}

func (e SecurityGroupOutboundDefaultPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityGroupOutboundDefaultPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityGroupOutboundDefaultPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityGroupOutboundDefaultPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityGroupOutboundDefaultPolicyOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutboundDefaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupOutboundDefaultPolicy)(nil)).Elem()
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToSecurityGroupOutboundDefaultPolicyOutput() SecurityGroupOutboundDefaultPolicyOutput {
	return o
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToSecurityGroupOutboundDefaultPolicyOutputWithContext(ctx context.Context) SecurityGroupOutboundDefaultPolicyOutput {
	return o
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToSecurityGroupOutboundDefaultPolicyPtrOutput() SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o.ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(context.Background())
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityGroupOutboundDefaultPolicy) *SecurityGroupOutboundDefaultPolicy {
		return &v
	}).(SecurityGroupOutboundDefaultPolicyPtrOutput)
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityGroupOutboundDefaultPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityGroupOutboundDefaultPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityGroupOutboundDefaultPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityGroupOutboundDefaultPolicyPtrOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutboundDefaultPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupOutboundDefaultPolicy)(nil)).Elem()
}

func (o SecurityGroupOutboundDefaultPolicyPtrOutput) ToSecurityGroupOutboundDefaultPolicyPtrOutput() SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o
}

func (o SecurityGroupOutboundDefaultPolicyPtrOutput) ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o
}

func (o SecurityGroupOutboundDefaultPolicyPtrOutput) Elem() SecurityGroupOutboundDefaultPolicyOutput {
	return o.ApplyT(func(v *SecurityGroupOutboundDefaultPolicy) SecurityGroupOutboundDefaultPolicy {
		if v != nil {
			return *v
		}
		var ret SecurityGroupOutboundDefaultPolicy
		return ret
	}).(SecurityGroupOutboundDefaultPolicyOutput)
}

func (o SecurityGroupOutboundDefaultPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityGroupOutboundDefaultPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityGroupOutboundDefaultPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SecurityGroupOutboundDefaultPolicyInput is an input type that accepts values of the SecurityGroupOutboundDefaultPolicy enum
// A concrete instance of `SecurityGroupOutboundDefaultPolicyInput` can be one of the following:
//
//	SecurityGroupOutboundDefaultPolicyAccept
//	SecurityGroupOutboundDefaultPolicyDrop
type SecurityGroupOutboundDefaultPolicyInput interface {
	pulumi.Input

	ToSecurityGroupOutboundDefaultPolicyOutput() SecurityGroupOutboundDefaultPolicyOutput
	ToSecurityGroupOutboundDefaultPolicyOutputWithContext(context.Context) SecurityGroupOutboundDefaultPolicyOutput
}

var securityGroupOutboundDefaultPolicyPtrType = reflect.TypeOf((**SecurityGroupOutboundDefaultPolicy)(nil)).Elem()

type SecurityGroupOutboundDefaultPolicyPtrInput interface {
	pulumi.Input

	ToSecurityGroupOutboundDefaultPolicyPtrOutput() SecurityGroupOutboundDefaultPolicyPtrOutput
	ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(context.Context) SecurityGroupOutboundDefaultPolicyPtrOutput
}

type securityGroupOutboundDefaultPolicyPtr string

func SecurityGroupOutboundDefaultPolicyPtr(v string) SecurityGroupOutboundDefaultPolicyPtrInput {
	return (*securityGroupOutboundDefaultPolicyPtr)(&v)
}

func (*securityGroupOutboundDefaultPolicyPtr) ElementType() reflect.Type {
	return securityGroupOutboundDefaultPolicyPtrType
}

func (in *securityGroupOutboundDefaultPolicyPtr) ToSecurityGroupOutboundDefaultPolicyPtrOutput() SecurityGroupOutboundDefaultPolicyPtrOutput {
	return pulumi.ToOutput(in).(SecurityGroupOutboundDefaultPolicyPtrOutput)
}

func (in *securityGroupOutboundDefaultPolicyPtr) ToSecurityGroupOutboundDefaultPolicyPtrOutputWithContext(ctx context.Context) SecurityGroupOutboundDefaultPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityGroupOutboundDefaultPolicyPtrOutput)
}

// Security group state
type SecurityGroupStateEnum string

const (
	SecurityGroupStateEnumAvailable    = SecurityGroupStateEnum("available")
	SecurityGroupStateEnumSyncing      = SecurityGroupStateEnum("syncing")
	SecurityGroupStateEnumSyncingError = SecurityGroupStateEnum("syncing_error")
)

type SecurityGroupStateEnumOutput struct{ *pulumi.OutputState }

func (SecurityGroupStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupStateEnum)(nil)).Elem()
}

func (o SecurityGroupStateEnumOutput) ToSecurityGroupStateEnumOutput() SecurityGroupStateEnumOutput {
	return o
}

func (o SecurityGroupStateEnumOutput) ToSecurityGroupStateEnumOutputWithContext(ctx context.Context) SecurityGroupStateEnumOutput {
	return o
}

func (o SecurityGroupStateEnumOutput) ToSecurityGroupStateEnumPtrOutput() SecurityGroupStateEnumPtrOutput {
	return o.ToSecurityGroupStateEnumPtrOutputWithContext(context.Background())
}

func (o SecurityGroupStateEnumOutput) ToSecurityGroupStateEnumPtrOutputWithContext(ctx context.Context) SecurityGroupStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityGroupStateEnum) *SecurityGroupStateEnum {
		return &v
	}).(SecurityGroupStateEnumPtrOutput)
}

func (o SecurityGroupStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityGroupStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityGroupStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityGroupStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityGroupStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityGroupStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityGroupStateEnumPtrOutput struct{ *pulumi.OutputState }

func (SecurityGroupStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupStateEnum)(nil)).Elem()
}

func (o SecurityGroupStateEnumPtrOutput) ToSecurityGroupStateEnumPtrOutput() SecurityGroupStateEnumPtrOutput {
	return o
}

func (o SecurityGroupStateEnumPtrOutput) ToSecurityGroupStateEnumPtrOutputWithContext(ctx context.Context) SecurityGroupStateEnumPtrOutput {
	return o
}

func (o SecurityGroupStateEnumPtrOutput) Elem() SecurityGroupStateEnumOutput {
	return o.ApplyT(func(v *SecurityGroupStateEnum) SecurityGroupStateEnum {
		if v != nil {
			return *v
		}
		var ret SecurityGroupStateEnum
		return ret
	}).(SecurityGroupStateEnumOutput)
}

func (o SecurityGroupStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityGroupStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityGroupStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInboundDefaultPolicyInput)(nil)).Elem(), SecurityGroupInboundDefaultPolicy("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInboundDefaultPolicyPtrInput)(nil)).Elem(), SecurityGroupInboundDefaultPolicy("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupOutboundDefaultPolicyInput)(nil)).Elem(), SecurityGroupOutboundDefaultPolicy("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupOutboundDefaultPolicyPtrInput)(nil)).Elem(), SecurityGroupOutboundDefaultPolicy("accept"))
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupInboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupInboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupStateOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SecurityGroupStatePtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupInboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(SecurityGroupInboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupOutboundDefaultPolicyOutput{})
	pulumi.RegisterOutputType(SecurityGroupOutboundDefaultPolicyPtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupStateEnumOutput{})
	pulumi.RegisterOutputType(SecurityGroupStateEnumPtrOutput{})
}
