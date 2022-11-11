// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snapshots

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScalewayInstanceV1SnapshotState string

const (
	ScalewayInstanceV1SnapshotStateAvailable    = ScalewayInstanceV1SnapshotState("available")
	ScalewayInstanceV1SnapshotStateSnapshotting = ScalewayInstanceV1SnapshotState("snapshotting")
	ScalewayInstanceV1SnapshotStateError        = ScalewayInstanceV1SnapshotState("error")
	ScalewayInstanceV1SnapshotStateInvalidData  = ScalewayInstanceV1SnapshotState("invalid_data")
	ScalewayInstanceV1SnapshotStateImporting    = ScalewayInstanceV1SnapshotState("importing")
	ScalewayInstanceV1SnapshotStateExporting    = ScalewayInstanceV1SnapshotState("exporting")
)

type ScalewayInstanceV1SnapshotStateOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SnapshotState)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToScalewayInstanceV1SnapshotStateOutput() ScalewayInstanceV1SnapshotStateOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToScalewayInstanceV1SnapshotStateOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotStateOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToScalewayInstanceV1SnapshotStatePtrOutput() ScalewayInstanceV1SnapshotStatePtrOutput {
	return o.ToScalewayInstanceV1SnapshotStatePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToScalewayInstanceV1SnapshotStatePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1SnapshotState) *ScalewayInstanceV1SnapshotState {
		return &v
	}).(ScalewayInstanceV1SnapshotStatePtrOutput)
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SnapshotState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SnapshotState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SnapshotStatePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SnapshotState)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotStatePtrOutput) ToScalewayInstanceV1SnapshotStatePtrOutput() ScalewayInstanceV1SnapshotStatePtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotStatePtrOutput) ToScalewayInstanceV1SnapshotStatePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotStatePtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotStatePtrOutput) Elem() ScalewayInstanceV1SnapshotStateOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SnapshotState) ScalewayInstanceV1SnapshotState {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SnapshotState
		return ret
	}).(ScalewayInstanceV1SnapshotStateOutput)
}

func (o ScalewayInstanceV1SnapshotStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1SnapshotState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SnapshotVolumeType string

const (
	ScalewayInstanceV1SnapshotVolumeTypeLSsd    = ScalewayInstanceV1SnapshotVolumeType("l_ssd")
	ScalewayInstanceV1SnapshotVolumeTypeBSsd    = ScalewayInstanceV1SnapshotVolumeType("b_ssd")
	ScalewayInstanceV1SnapshotVolumeTypeUnified = ScalewayInstanceV1SnapshotVolumeType("unified")
)

type ScalewayInstanceV1SnapshotVolumeTypeOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotVolumeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SnapshotVolumeType)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToScalewayInstanceV1SnapshotVolumeTypeOutput() ScalewayInstanceV1SnapshotVolumeTypeOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToScalewayInstanceV1SnapshotVolumeTypeOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotVolumeTypeOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToScalewayInstanceV1SnapshotVolumeTypePtrOutput() ScalewayInstanceV1SnapshotVolumeTypePtrOutput {
	return o.ToScalewayInstanceV1SnapshotVolumeTypePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToScalewayInstanceV1SnapshotVolumeTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotVolumeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1SnapshotVolumeType) *ScalewayInstanceV1SnapshotVolumeType {
		return &v
	}).(ScalewayInstanceV1SnapshotVolumeTypePtrOutput)
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SnapshotVolumeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotVolumeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1SnapshotVolumeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SnapshotVolumeTypePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotVolumeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SnapshotVolumeType)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotVolumeTypePtrOutput) ToScalewayInstanceV1SnapshotVolumeTypePtrOutput() ScalewayInstanceV1SnapshotVolumeTypePtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotVolumeTypePtrOutput) ToScalewayInstanceV1SnapshotVolumeTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotVolumeTypePtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotVolumeTypePtrOutput) Elem() ScalewayInstanceV1SnapshotVolumeTypeOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SnapshotVolumeType) ScalewayInstanceV1SnapshotVolumeType {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SnapshotVolumeType
		return ret
	}).(ScalewayInstanceV1SnapshotVolumeTypeOutput)
}

func (o ScalewayInstanceV1SnapshotVolumeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1SnapshotVolumeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1SnapshotVolumeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type State string

const (
	StateAvailable    = State("available")
	StateSnapshotting = State("snapshotting")
	StateError        = State("error")
	StateInvalidData  = State("invalid_data")
	StateImporting    = State("importing")
	StateExporting    = State("exporting")
)

func (State) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (e State) ToStateOutput() StateOutput {
	return pulumi.ToOutput(e).(StateOutput)
}

func (e State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateOutput)
}

func (e State) ToStatePtrOutput() StatePtrOutput {
	return e.ToStatePtrOutputWithContext(context.Background())
}

func (e State) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return State(e).ToStateOutputWithContext(ctx).ToStatePtrOutputWithContext(ctx)
}

func (e State) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e State) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

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

// StateInput is an input type that accepts StateArgs and StateOutput values.
// You can construct a concrete instance of `StateInput` via:
//
//	StateArgs{...}
type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(context.Context) StateOutput
}

var statePtrType = reflect.TypeOf((**State)(nil)).Elem()

type StatePtrInput interface {
	pulumi.Input

	ToStatePtrOutput() StatePtrOutput
	ToStatePtrOutputWithContext(context.Context) StatePtrOutput
}

type statePtr string

func StatePtr(v string) StatePtrInput {
	return (*statePtr)(&v)
}

func (*statePtr) ElementType() reflect.Type {
	return statePtrType
}

func (in *statePtr) ToStatePtrOutput() StatePtrOutput {
	return pulumi.ToOutput(in).(StatePtrOutput)
}

func (in *statePtr) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatePtrOutput)
}

type VolumeType string

const (
	VolumeTypeLSsd    = VolumeType("l_ssd")
	VolumeTypeBSsd    = VolumeType("b_ssd")
	VolumeTypeUnified = VolumeType("unified")
)

func (VolumeType) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeType)(nil)).Elem()
}

func (e VolumeType) ToVolumeTypeOutput() VolumeTypeOutput {
	return pulumi.ToOutput(e).(VolumeTypeOutput)
}

func (e VolumeType) ToVolumeTypeOutputWithContext(ctx context.Context) VolumeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VolumeTypeOutput)
}

func (e VolumeType) ToVolumeTypePtrOutput() VolumeTypePtrOutput {
	return e.ToVolumeTypePtrOutputWithContext(context.Background())
}

func (e VolumeType) ToVolumeTypePtrOutputWithContext(ctx context.Context) VolumeTypePtrOutput {
	return VolumeType(e).ToVolumeTypeOutputWithContext(ctx).ToVolumeTypePtrOutputWithContext(ctx)
}

func (e VolumeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VolumeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VolumeTypeOutput struct{ *pulumi.OutputState }

func (VolumeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeType)(nil)).Elem()
}

func (o VolumeTypeOutput) ToVolumeTypeOutput() VolumeTypeOutput {
	return o
}

func (o VolumeTypeOutput) ToVolumeTypeOutputWithContext(ctx context.Context) VolumeTypeOutput {
	return o
}

func (o VolumeTypeOutput) ToVolumeTypePtrOutput() VolumeTypePtrOutput {
	return o.ToVolumeTypePtrOutputWithContext(context.Background())
}

func (o VolumeTypeOutput) ToVolumeTypePtrOutputWithContext(ctx context.Context) VolumeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeType) *VolumeType {
		return &v
	}).(VolumeTypePtrOutput)
}

func (o VolumeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VolumeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VolumeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VolumeTypePtrOutput struct{ *pulumi.OutputState }

func (VolumeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeType)(nil)).Elem()
}

func (o VolumeTypePtrOutput) ToVolumeTypePtrOutput() VolumeTypePtrOutput {
	return o
}

func (o VolumeTypePtrOutput) ToVolumeTypePtrOutputWithContext(ctx context.Context) VolumeTypePtrOutput {
	return o
}

func (o VolumeTypePtrOutput) Elem() VolumeTypeOutput {
	return o.ApplyT(func(v *VolumeType) VolumeType {
		if v != nil {
			return *v
		}
		var ret VolumeType
		return ret
	}).(VolumeTypeOutput)
}

func (o VolumeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VolumeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// VolumeTypeInput is an input type that accepts VolumeTypeArgs and VolumeTypeOutput values.
// You can construct a concrete instance of `VolumeTypeInput` via:
//
//	VolumeTypeArgs{...}
type VolumeTypeInput interface {
	pulumi.Input

	ToVolumeTypeOutput() VolumeTypeOutput
	ToVolumeTypeOutputWithContext(context.Context) VolumeTypeOutput
}

var volumeTypePtrType = reflect.TypeOf((**VolumeType)(nil)).Elem()

type VolumeTypePtrInput interface {
	pulumi.Input

	ToVolumeTypePtrOutput() VolumeTypePtrOutput
	ToVolumeTypePtrOutputWithContext(context.Context) VolumeTypePtrOutput
}

type volumeTypePtr string

func VolumeTypePtr(v string) VolumeTypePtrInput {
	return (*volumeTypePtr)(&v)
}

func (*volumeTypePtr) ElementType() reflect.Type {
	return volumeTypePtrType
}

func (in *volumeTypePtr) ToVolumeTypePtrOutput() VolumeTypePtrOutput {
	return pulumi.ToOutput(in).(VolumeTypePtrOutput)
}

func (in *volumeTypePtr) ToVolumeTypePtrOutputWithContext(ctx context.Context) VolumeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VolumeTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StateInput)(nil)).Elem(), State("available"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatePtrInput)(nil)).Elem(), State("available"))
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeInput)(nil)).Elem(), VolumeType("l_ssd"))
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypePtrInput)(nil)).Elem(), VolumeType("l_ssd"))
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotStateOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotStatePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotVolumeTypeOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotVolumeTypePtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
	pulumi.RegisterOutputType(VolumeTypeOutput{})
	pulumi.RegisterOutputType(VolumeTypePtrOutput{})
}