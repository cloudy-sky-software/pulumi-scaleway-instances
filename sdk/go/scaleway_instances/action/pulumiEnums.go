// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package action

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The action to perform on the server
type Action string

const (
	ActionPoweron     = Action("poweron")
	ActionBackup      = Action("backup")
	ActionStopInPlace = Action("stop_in_place")
	ActionPoweroff    = Action("poweroff")
	ActionTerminate   = Action("terminate")
	ActionReboot      = Action("reboot")
)

func (Action) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (e Action) ToActionOutput() ActionOutput {
	return pulumi.ToOutput(e).(ActionOutput)
}

func (e Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionOutput)
}

func (e Action) ToActionPtrOutput() ActionPtrOutput {
	return e.ToActionPtrOutputWithContext(context.Background())
}

func (e Action) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return Action(e).ToActionOutputWithContext(ctx).ToActionPtrOutputWithContext(ctx)
}

func (e Action) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Action) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Action) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ActionInput is an input type that accepts ActionArgs and ActionOutput values.
// You can construct a concrete instance of `ActionInput` via:
//
//	ActionArgs{...}
type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

var actionPtrType = reflect.TypeOf((**Action)(nil)).Elem()

type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtr string

func ActionPtr(v string) ActionPtrInput {
	return (*actionPtr)(&v)
}

func (*actionPtr) ElementType() reflect.Type {
	return actionPtrType
}

func (in *actionPtr) ToActionPtrOutput() ActionPtrOutput {
	return pulumi.ToOutput(in).(ActionPtrOutput)
}

func (in *actionPtr) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionPtrOutput)
}

type ScalewayInstanceV1ListServerActionsResponseActionsItem string

const (
	ScalewayInstanceV1ListServerActionsResponseActionsItemPoweron     = ScalewayInstanceV1ListServerActionsResponseActionsItem("poweron")
	ScalewayInstanceV1ListServerActionsResponseActionsItemBackup      = ScalewayInstanceV1ListServerActionsResponseActionsItem("backup")
	ScalewayInstanceV1ListServerActionsResponseActionsItemStopInPlace = ScalewayInstanceV1ListServerActionsResponseActionsItem("stop_in_place")
	ScalewayInstanceV1ListServerActionsResponseActionsItemPoweroff    = ScalewayInstanceV1ListServerActionsResponseActionsItem("poweroff")
	ScalewayInstanceV1ListServerActionsResponseActionsItemTerminate   = ScalewayInstanceV1ListServerActionsResponseActionsItem("terminate")
	ScalewayInstanceV1ListServerActionsResponseActionsItemReboot      = ScalewayInstanceV1ListServerActionsResponseActionsItem("reboot")
)

type ScalewayInstanceV1ListServerActionsResponseActionsItemOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListServerActionsResponseActionsItem)(nil)).Elem()
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemOutput() ScalewayInstanceV1ListServerActionsResponseActionsItemOutput {
	return o
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemOutputWithContext(ctx context.Context) ScalewayInstanceV1ListServerActionsResponseActionsItemOutput {
	return o
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput() ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput {
	return o.ToScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1ListServerActionsResponseActionsItem) *ScalewayInstanceV1ListServerActionsResponseActionsItem {
		return &v
	}).(ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput)
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1ListServerActionsResponseActionsItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1ListServerActionsResponseActionsItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1ListServerActionsResponseActionsItem)(nil)).Elem()
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput() ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput {
	return o
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput {
	return o
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput) Elem() ScalewayInstanceV1ListServerActionsResponseActionsItemOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1ListServerActionsResponseActionsItem) ScalewayInstanceV1ListServerActionsResponseActionsItem {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1ListServerActionsResponseActionsItem
		return ret
	}).(ScalewayInstanceV1ListServerActionsResponseActionsItemOutput)
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1ListServerActionsResponseActionsItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1ListServerActionsResponseActionsItem)(nil)).Elem()
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput() ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput {
	return o
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput) ToScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput {
	return o
}

func (o ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1ListServerActionsResponseActionsItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1ListServerActionsResponseActionsItem {
		return vs[0].([]ScalewayInstanceV1ListServerActionsResponseActionsItem)[vs[1].(int)]
	}).(ScalewayInstanceV1ListServerActionsResponseActionsItemOutput)
}

type ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType string

const (
	ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeLSsd    = ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType("l_ssd")
	ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeBSsd    = ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType("b_ssd")
	ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeUnified = ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType("unified")
)

func (ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType)(nil)).Elem()
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput {
	return pulumi.ToOutput(e).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput)
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput)
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return e.ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(context.Background())
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType(e).ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutputWithContext(ctx).ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(ctx)
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType)(nil)).Elem()
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput {
	return o
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput {
	return o
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return o.ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) *ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType {
		return &v
	}).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput)
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType)(nil)).Elem()
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return o
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return o
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput) Elem() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType
		return ret
	}).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput)
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeInput is an input type that accepts ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeArgs and ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput values.
// You can construct a concrete instance of `ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeInput` via:
//
//	ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeArgs{...}
type ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeInput interface {
	pulumi.Input

	ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput
	ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutputWithContext(context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput
}

var scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrType = reflect.TypeOf((**ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType)(nil)).Elem()

type ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrInput interface {
	pulumi.Input

	ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput
	ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput
}

type scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtr string

func ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtr(v string) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrInput {
	return (*scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtr)(&v)
}

func (*scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtr) ElementType() reflect.Type {
	return scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrType
}

func (in *scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtr) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return pulumi.ToOutput(in).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput)
}

func (in *scalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtr) ToScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionInput)(nil)).Elem(), Action("poweron"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionPtrInput)(nil)).Elem(), Action("poweron"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeInput)(nil)).Elem(), ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType("l_ssd"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrInput)(nil)).Elem(), ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType("l_ssd"))
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListServerActionsResponseActionsItemOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListServerActionsResponseActionsItemPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListServerActionsResponseActionsItemArrayOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypeOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeTypePtrOutput{})
}