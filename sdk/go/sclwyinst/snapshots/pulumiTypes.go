// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snapshots

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

// The volume on which the snapshot is based on
type BaseVolumeProperties struct {
	// The volume ID on which the snapshot is based on
	Id *string `pulumi:"id"`
	// The volume name on which the snapshot is based on
	Name *string `pulumi:"name"`
}

// The volume on which the snapshot is based on
type BaseVolumePropertiesOutput struct{ *pulumi.OutputState }

func (BaseVolumePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseVolumeProperties)(nil)).Elem()
}

func (o BaseVolumePropertiesOutput) ToBaseVolumePropertiesOutput() BaseVolumePropertiesOutput {
	return o
}

func (o BaseVolumePropertiesOutput) ToBaseVolumePropertiesOutputWithContext(ctx context.Context) BaseVolumePropertiesOutput {
	return o
}

func (o BaseVolumePropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[BaseVolumeProperties] {
	return pulumix.Output[BaseVolumeProperties]{
		OutputState: o.OutputState,
	}
}

// The volume ID on which the snapshot is based on
func (o BaseVolumePropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseVolumeProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The volume name on which the snapshot is based on
func (o BaseVolumePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseVolumeProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type BaseVolumePropertiesPtrOutput struct{ *pulumi.OutputState }

func (BaseVolumePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseVolumeProperties)(nil)).Elem()
}

func (o BaseVolumePropertiesPtrOutput) ToBaseVolumePropertiesPtrOutput() BaseVolumePropertiesPtrOutput {
	return o
}

func (o BaseVolumePropertiesPtrOutput) ToBaseVolumePropertiesPtrOutputWithContext(ctx context.Context) BaseVolumePropertiesPtrOutput {
	return o
}

func (o BaseVolumePropertiesPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*BaseVolumeProperties] {
	return pulumix.Output[*BaseVolumeProperties]{
		OutputState: o.OutputState,
	}
}

func (o BaseVolumePropertiesPtrOutput) Elem() BaseVolumePropertiesOutput {
	return o.ApplyT(func(v *BaseVolumeProperties) BaseVolumeProperties {
		if v != nil {
			return *v
		}
		var ret BaseVolumeProperties
		return ret
	}).(BaseVolumePropertiesOutput)
}

// The volume ID on which the snapshot is based on
func (o BaseVolumePropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseVolumeProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The volume name on which the snapshot is based on
func (o BaseVolumePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseVolumeProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1GetSnapshotResponse struct {
	Snapshot *ScalewayInstanceV1Snapshot `pulumi:"snapshot"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1GetSnapshotResponse
func (val *ScalewayInstanceV1GetSnapshotResponse) Defaults() *ScalewayInstanceV1GetSnapshotResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Snapshot = tmp.Snapshot.Defaults()

	return &tmp
}

type ScalewayInstanceV1GetSnapshotResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetSnapshotResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetSnapshotResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetSnapshotResponseOutput) ToScalewayInstanceV1GetSnapshotResponseOutput() ScalewayInstanceV1GetSnapshotResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetSnapshotResponseOutput) ToScalewayInstanceV1GetSnapshotResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetSnapshotResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetSnapshotResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1GetSnapshotResponse] {
	return pulumix.Output[ScalewayInstanceV1GetSnapshotResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetSnapshotResponseOutput) Snapshot() ScalewayInstanceV1SnapshotPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetSnapshotResponse) *ScalewayInstanceV1Snapshot { return v.Snapshot }).(ScalewayInstanceV1SnapshotPtrOutput)
}

type ScalewayInstanceV1ListSnapshotsResponse struct {
	// List of snapshots
	Snapshots []ScalewayInstanceV1Snapshot `pulumi:"snapshots"`
}

type ScalewayInstanceV1ListSnapshotsResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListSnapshotsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListSnapshotsResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1ListSnapshotsResponseOutput) ToScalewayInstanceV1ListSnapshotsResponseOutput() ScalewayInstanceV1ListSnapshotsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListSnapshotsResponseOutput) ToScalewayInstanceV1ListSnapshotsResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1ListSnapshotsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListSnapshotsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1ListSnapshotsResponse] {
	return pulumix.Output[ScalewayInstanceV1ListSnapshotsResponse]{
		OutputState: o.OutputState,
	}
}

// List of snapshots
func (o ScalewayInstanceV1ListSnapshotsResponseOutput) Snapshots() ScalewayInstanceV1SnapshotArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ListSnapshotsResponse) []ScalewayInstanceV1Snapshot { return v.Snapshots }).(ScalewayInstanceV1SnapshotArrayOutput)
}

type ScalewayInstanceV1Snapshot struct {
	// The volume on which the snapshot is based on
	Base_volume *ScalewayInstanceV1SnapshotBaseVolumeProperties `pulumi:"base_volume"`
	// The snapshot creation date (RFC 3339 format)
	Creation_date *string `pulumi:"creation_date"`
	// The reason for the failed snapshot import
	Error_reason *string `pulumi:"error_reason"`
	Id           *string `pulumi:"id"`
	// The snapshot modification date (RFC 3339 format)
	Modification_date *string `pulumi:"modification_date"`
	// The snapshot name
	Name *string `pulumi:"name"`
	// The snapshot organization ID
	Organization *string `pulumi:"organization"`
	// The snapshot project ID
	Project *string `pulumi:"project"`
	// The snapshot size (in bytes)
	Size  *float64                         `pulumi:"size"`
	State *ScalewayInstanceV1SnapshotState `pulumi:"state"`
	// The snapshot tags
	Tags        []string                              `pulumi:"tags"`
	Volume_type *ScalewayInstanceV1SnapshotVolumeType `pulumi:"volume_type"`
	// The snapshot zone
	Zone *string `pulumi:"zone"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1Snapshot
func (val *ScalewayInstanceV1Snapshot) Defaults() *ScalewayInstanceV1Snapshot {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.State == nil {
		state_ := ScalewayInstanceV1SnapshotState("available")
		tmp.State = &state_
	}
	if tmp.Volume_type == nil {
		volume_type_ := ScalewayInstanceV1SnapshotVolumeType("l_ssd")
		tmp.Volume_type = &volume_type_
	}
	return &tmp
}

type ScalewayInstanceV1SnapshotOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1Snapshot)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotOutput) ToScalewayInstanceV1SnapshotOutput() ScalewayInstanceV1SnapshotOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotOutput) ToScalewayInstanceV1SnapshotOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1Snapshot] {
	return pulumix.Output[ScalewayInstanceV1Snapshot]{
		OutputState: o.OutputState,
	}
}

// The volume on which the snapshot is based on
func (o ScalewayInstanceV1SnapshotOutput) Base_volume() ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *ScalewayInstanceV1SnapshotBaseVolumeProperties {
		return v.Base_volume
	}).(ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput)
}

// The snapshot creation date (RFC 3339 format)
func (o ScalewayInstanceV1SnapshotOutput) Creation_date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Creation_date }).(pulumi.StringPtrOutput)
}

// The reason for the failed snapshot import
func (o ScalewayInstanceV1SnapshotOutput) Error_reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Error_reason }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1SnapshotOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The snapshot modification date (RFC 3339 format)
func (o ScalewayInstanceV1SnapshotOutput) Modification_date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Modification_date }).(pulumi.StringPtrOutput)
}

// The snapshot name
func (o ScalewayInstanceV1SnapshotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The snapshot organization ID
func (o ScalewayInstanceV1SnapshotOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// The snapshot project ID
func (o ScalewayInstanceV1SnapshotOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// The snapshot size (in bytes)
func (o ScalewayInstanceV1SnapshotOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SnapshotOutput) State() ScalewayInstanceV1SnapshotStatePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *ScalewayInstanceV1SnapshotState { return v.State }).(ScalewayInstanceV1SnapshotStatePtrOutput)
}

// The snapshot tags
func (o ScalewayInstanceV1SnapshotOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o ScalewayInstanceV1SnapshotOutput) Volume_type() ScalewayInstanceV1SnapshotVolumeTypePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *ScalewayInstanceV1SnapshotVolumeType { return v.Volume_type }).(ScalewayInstanceV1SnapshotVolumeTypePtrOutput)
}

// The snapshot zone
func (o ScalewayInstanceV1SnapshotOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Snapshot) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SnapshotPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1Snapshot)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotPtrOutput) ToScalewayInstanceV1SnapshotPtrOutput() ScalewayInstanceV1SnapshotPtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotPtrOutput) ToScalewayInstanceV1SnapshotPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotPtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1Snapshot] {
	return pulumix.Output[*ScalewayInstanceV1Snapshot]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SnapshotPtrOutput) Elem() ScalewayInstanceV1SnapshotOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) ScalewayInstanceV1Snapshot {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1Snapshot
		return ret
	}).(ScalewayInstanceV1SnapshotOutput)
}

// The volume on which the snapshot is based on
func (o ScalewayInstanceV1SnapshotPtrOutput) Base_volume() ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *ScalewayInstanceV1SnapshotBaseVolumeProperties {
		if v == nil {
			return nil
		}
		return v.Base_volume
	}).(ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput)
}

// The snapshot creation date (RFC 3339 format)
func (o ScalewayInstanceV1SnapshotPtrOutput) Creation_date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Creation_date
	}).(pulumi.StringPtrOutput)
}

// The reason for the failed snapshot import
func (o ScalewayInstanceV1SnapshotPtrOutput) Error_reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Error_reason
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1SnapshotPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The snapshot modification date (RFC 3339 format)
func (o ScalewayInstanceV1SnapshotPtrOutput) Modification_date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Modification_date
	}).(pulumi.StringPtrOutput)
}

// The snapshot name
func (o ScalewayInstanceV1SnapshotPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The snapshot organization ID
func (o ScalewayInstanceV1SnapshotPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

// The snapshot project ID
func (o ScalewayInstanceV1SnapshotPtrOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Project
	}).(pulumi.StringPtrOutput)
}

// The snapshot size (in bytes)
func (o ScalewayInstanceV1SnapshotPtrOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *float64 {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1SnapshotPtrOutput) State() ScalewayInstanceV1SnapshotStatePtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *ScalewayInstanceV1SnapshotState {
		if v == nil {
			return nil
		}
		return v.State
	}).(ScalewayInstanceV1SnapshotStatePtrOutput)
}

// The snapshot tags
func (o ScalewayInstanceV1SnapshotPtrOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) []string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayOutput)
}

func (o ScalewayInstanceV1SnapshotPtrOutput) Volume_type() ScalewayInstanceV1SnapshotVolumeTypePtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *ScalewayInstanceV1SnapshotVolumeType {
		if v == nil {
			return nil
		}
		return v.Volume_type
	}).(ScalewayInstanceV1SnapshotVolumeTypePtrOutput)
}

// The snapshot zone
func (o ScalewayInstanceV1SnapshotPtrOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Snapshot) *string {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SnapshotArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1Snapshot)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotArrayOutput) ToScalewayInstanceV1SnapshotArrayOutput() ScalewayInstanceV1SnapshotArrayOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotArrayOutput) ToScalewayInstanceV1SnapshotArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotArrayOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ScalewayInstanceV1Snapshot] {
	return pulumix.Output[[]ScalewayInstanceV1Snapshot]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SnapshotArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1Snapshot {
		return vs[0].([]ScalewayInstanceV1Snapshot)[vs[1].(int)]
	}).(ScalewayInstanceV1SnapshotOutput)
}

// The volume on which the snapshot is based on
type ScalewayInstanceV1SnapshotBaseVolumeProperties struct {
	// The volume ID on which the snapshot is based on
	Id *string `pulumi:"id"`
	// The volume name on which the snapshot is based on
	Name *string `pulumi:"name"`
}

// The volume on which the snapshot is based on
type ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1SnapshotBaseVolumeProperties)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput) ToScalewayInstanceV1SnapshotBaseVolumePropertiesOutput() ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput) ToScalewayInstanceV1SnapshotBaseVolumePropertiesOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1SnapshotBaseVolumeProperties] {
	return pulumix.Output[ScalewayInstanceV1SnapshotBaseVolumeProperties]{
		OutputState: o.OutputState,
	}
}

// The volume ID on which the snapshot is based on
func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SnapshotBaseVolumeProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The volume name on which the snapshot is based on
func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1SnapshotBaseVolumeProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1SnapshotBaseVolumeProperties)(nil)).Elem()
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) ToScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput() ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) ToScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput {
	return o
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1SnapshotBaseVolumeProperties] {
	return pulumix.Output[*ScalewayInstanceV1SnapshotBaseVolumeProperties]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) Elem() ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SnapshotBaseVolumeProperties) ScalewayInstanceV1SnapshotBaseVolumeProperties {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1SnapshotBaseVolumeProperties
		return ret
	}).(ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput)
}

// The volume ID on which the snapshot is based on
func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SnapshotBaseVolumeProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The volume name on which the snapshot is based on
func (o ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1SnapshotBaseVolumeProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BaseVolumePropertiesOutput{})
	pulumi.RegisterOutputType(BaseVolumePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetSnapshotResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListSnapshotsResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotArrayOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotBaseVolumePropertiesOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1SnapshotBaseVolumePropertiesPtrOutput{})
}
