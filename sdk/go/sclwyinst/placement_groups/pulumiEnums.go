// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package placement_groups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyMode string

const (
	PolicyModeOptional = PolicyMode("optional")
	PolicyModeEnforced = PolicyMode("enforced")
)

func (PolicyMode) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyMode)(nil)).Elem()
}

func (e PolicyMode) ToPolicyModeOutput() PolicyModeOutput {
	return pulumi.ToOutput(e).(PolicyModeOutput)
}

func (e PolicyMode) ToPolicyModeOutputWithContext(ctx context.Context) PolicyModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyModeOutput)
}

func (e PolicyMode) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return e.ToPolicyModePtrOutputWithContext(context.Background())
}

func (e PolicyMode) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return PolicyMode(e).ToPolicyModeOutputWithContext(ctx).ToPolicyModePtrOutputWithContext(ctx)
}

func (e PolicyMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyModeOutput struct{ *pulumi.OutputState }

func (PolicyModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyMode)(nil)).Elem()
}

func (o PolicyModeOutput) ToPolicyModeOutput() PolicyModeOutput {
	return o
}

func (o PolicyModeOutput) ToPolicyModeOutputWithContext(ctx context.Context) PolicyModeOutput {
	return o
}

func (o PolicyModeOutput) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return o.ToPolicyModePtrOutputWithContext(context.Background())
}

func (o PolicyModeOutput) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyMode) *PolicyMode {
		return &v
	}).(PolicyModePtrOutput)
}

func (o PolicyModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyModePtrOutput struct{ *pulumi.OutputState }

func (PolicyModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyMode)(nil)).Elem()
}

func (o PolicyModePtrOutput) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return o
}

func (o PolicyModePtrOutput) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return o
}

func (o PolicyModePtrOutput) Elem() PolicyModeOutput {
	return o.ApplyT(func(v *PolicyMode) PolicyMode {
		if v != nil {
			return *v
		}
		var ret PolicyMode
		return ret
	}).(PolicyModeOutput)
}

func (o PolicyModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PolicyModeInput is an input type that accepts values of the PolicyMode enum
// A concrete instance of `PolicyModeInput` can be one of the following:
//
//	PolicyModeOptional
//	PolicyModeEnforced
type PolicyModeInput interface {
	pulumi.Input

	ToPolicyModeOutput() PolicyModeOutput
	ToPolicyModeOutputWithContext(context.Context) PolicyModeOutput
}

var policyModePtrType = reflect.TypeOf((**PolicyMode)(nil)).Elem()

type PolicyModePtrInput interface {
	pulumi.Input

	ToPolicyModePtrOutput() PolicyModePtrOutput
	ToPolicyModePtrOutputWithContext(context.Context) PolicyModePtrOutput
}

type policyModePtr string

func PolicyModePtr(v string) PolicyModePtrInput {
	return (*policyModePtr)(&v)
}

func (*policyModePtr) ElementType() reflect.Type {
	return policyModePtrType
}

func (in *policyModePtr) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return pulumi.ToOutput(in).(PolicyModePtrOutput)
}

func (in *policyModePtr) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyModePtrOutput)
}

type PolicyType string

const (
	PolicyTypeMaxAvailability = PolicyType("max_availability")
	PolicyTypeLowLatency      = PolicyType("low_latency")
)

func (PolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (e PolicyType) ToPolicyTypeOutput() PolicyTypeOutput {
	return pulumi.ToOutput(e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return e.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (e PolicyType) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return PolicyType(e).ToPolicyTypeOutputWithContext(ctx).ToPolicyTypePtrOutputWithContext(ctx)
}

func (e PolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyTypeOutput struct{ *pulumi.OutputState }

func (PolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (o PolicyTypeOutput) ToPolicyTypeOutput() PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyType) *PolicyType {
		return &v
	}).(PolicyTypePtrOutput)
}

func (o PolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyTypePtrOutput struct{ *pulumi.OutputState }

func (PolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyType)(nil)).Elem()
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) Elem() PolicyTypeOutput {
	return o.ApplyT(func(v *PolicyType) PolicyType {
		if v != nil {
			return *v
		}
		var ret PolicyType
		return ret
	}).(PolicyTypeOutput)
}

func (o PolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PolicyTypeInput is an input type that accepts values of the PolicyType enum
// A concrete instance of `PolicyTypeInput` can be one of the following:
//
//	PolicyTypeMaxAvailability
//	PolicyTypeLowLatency
type PolicyTypeInput interface {
	pulumi.Input

	ToPolicyTypeOutput() PolicyTypeOutput
	ToPolicyTypeOutputWithContext(context.Context) PolicyTypeOutput
}

var policyTypePtrType = reflect.TypeOf((**PolicyType)(nil)).Elem()

type PolicyTypePtrInput interface {
	pulumi.Input

	ToPolicyTypePtrOutput() PolicyTypePtrOutput
	ToPolicyTypePtrOutputWithContext(context.Context) PolicyTypePtrOutput
}

type policyTypePtr string

func PolicyTypePtr(v string) PolicyTypePtrInput {
	return (*policyTypePtr)(&v)
}

func (*policyTypePtr) ElementType() reflect.Type {
	return policyTypePtrType
}

func (in *policyTypePtr) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return pulumi.ToOutput(in).(PolicyTypePtrOutput)
}

func (in *policyTypePtr) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyTypePtrOutput)
}

type ScalewayInstanceV1PlacementGroupPolicyMode string

const (
	ScalewayInstanceV1PlacementGroupPolicyModeOptional = ScalewayInstanceV1PlacementGroupPolicyMode("optional")
	ScalewayInstanceV1PlacementGroupPolicyModeEnforced = ScalewayInstanceV1PlacementGroupPolicyMode("enforced")
)

type ScalewayInstanceV1PlacementGroupPolicyModeOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PlacementGroupPolicyModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1PlacementGroupPolicyMode)(nil)).Elem()
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToScalewayInstanceV1PlacementGroupPolicyModeOutput() ScalewayInstanceV1PlacementGroupPolicyModeOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToScalewayInstanceV1PlacementGroupPolicyModeOutputWithContext(ctx context.Context) ScalewayInstanceV1PlacementGroupPolicyModeOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToScalewayInstanceV1PlacementGroupPolicyModePtrOutput() ScalewayInstanceV1PlacementGroupPolicyModePtrOutput {
	return o.ToScalewayInstanceV1PlacementGroupPolicyModePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToScalewayInstanceV1PlacementGroupPolicyModePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PlacementGroupPolicyModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1PlacementGroupPolicyMode) *ScalewayInstanceV1PlacementGroupPolicyMode {
		return &v
	}).(ScalewayInstanceV1PlacementGroupPolicyModePtrOutput)
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1PlacementGroupPolicyMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1PlacementGroupPolicyMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1PlacementGroupPolicyModePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PlacementGroupPolicyModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1PlacementGroupPolicyMode)(nil)).Elem()
}

func (o ScalewayInstanceV1PlacementGroupPolicyModePtrOutput) ToScalewayInstanceV1PlacementGroupPolicyModePtrOutput() ScalewayInstanceV1PlacementGroupPolicyModePtrOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyModePtrOutput) ToScalewayInstanceV1PlacementGroupPolicyModePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PlacementGroupPolicyModePtrOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyModePtrOutput) Elem() ScalewayInstanceV1PlacementGroupPolicyModeOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PlacementGroupPolicyMode) ScalewayInstanceV1PlacementGroupPolicyMode {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1PlacementGroupPolicyMode
		return ret
	}).(ScalewayInstanceV1PlacementGroupPolicyModeOutput)
}

func (o ScalewayInstanceV1PlacementGroupPolicyModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1PlacementGroupPolicyMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1PlacementGroupPolicyType string

const (
	ScalewayInstanceV1PlacementGroupPolicyTypeMaxAvailability = ScalewayInstanceV1PlacementGroupPolicyType("max_availability")
	ScalewayInstanceV1PlacementGroupPolicyTypeLowLatency      = ScalewayInstanceV1PlacementGroupPolicyType("low_latency")
)

type ScalewayInstanceV1PlacementGroupPolicyTypeOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1PlacementGroupPolicyType)(nil)).Elem()
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToScalewayInstanceV1PlacementGroupPolicyTypeOutput() ScalewayInstanceV1PlacementGroupPolicyTypeOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToScalewayInstanceV1PlacementGroupPolicyTypeOutputWithContext(ctx context.Context) ScalewayInstanceV1PlacementGroupPolicyTypeOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToScalewayInstanceV1PlacementGroupPolicyTypePtrOutput() ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput {
	return o.ToScalewayInstanceV1PlacementGroupPolicyTypePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToScalewayInstanceV1PlacementGroupPolicyTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1PlacementGroupPolicyType) *ScalewayInstanceV1PlacementGroupPolicyType {
		return &v
	}).(ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput)
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1PlacementGroupPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1PlacementGroupPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1PlacementGroupPolicyType)(nil)).Elem()
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput) ToScalewayInstanceV1PlacementGroupPolicyTypePtrOutput() ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput) ToScalewayInstanceV1PlacementGroupPolicyTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput {
	return o
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput) Elem() ScalewayInstanceV1PlacementGroupPolicyTypeOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PlacementGroupPolicyType) ScalewayInstanceV1PlacementGroupPolicyType {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1PlacementGroupPolicyType
		return ret
	}).(ScalewayInstanceV1PlacementGroupPolicyTypeOutput)
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1PlacementGroupPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyModeInput)(nil)).Elem(), PolicyMode("optional"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyModePtrInput)(nil)).Elem(), PolicyMode("optional"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTypeInput)(nil)).Elem(), PolicyType("max_availability"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTypePtrInput)(nil)).Elem(), PolicyType("max_availability"))
	pulumi.RegisterOutputType(PolicyModeOutput{})
	pulumi.RegisterOutputType(PolicyModePtrOutput{})
	pulumi.RegisterOutputType(PolicyTypeOutput{})
	pulumi.RegisterOutputType(PolicyTypePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PlacementGroupPolicyModeOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PlacementGroupPolicyModePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PlacementGroupPolicyTypeOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PlacementGroupPolicyTypePtrOutput{})
}
