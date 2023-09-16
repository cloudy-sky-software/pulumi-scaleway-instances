// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_nics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The private NIC state
type ScalewayInstanceV1PrivateNICState string

const (
	ScalewayInstanceV1PrivateNICStateAvailable    = ScalewayInstanceV1PrivateNICState("available")
	ScalewayInstanceV1PrivateNICStateSyncing      = ScalewayInstanceV1PrivateNICState("syncing")
	ScalewayInstanceV1PrivateNICStateSyncingError = ScalewayInstanceV1PrivateNICState("syncing_error")
)

type ScalewayInstanceV1PrivateNICStateOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PrivateNICStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1PrivateNICState)(nil)).Elem()
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToScalewayInstanceV1PrivateNICStateOutput() ScalewayInstanceV1PrivateNICStateOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToScalewayInstanceV1PrivateNICStateOutputWithContext(ctx context.Context) ScalewayInstanceV1PrivateNICStateOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToScalewayInstanceV1PrivateNICStatePtrOutput() ScalewayInstanceV1PrivateNICStatePtrOutput {
	return o.ToScalewayInstanceV1PrivateNICStatePtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToScalewayInstanceV1PrivateNICStatePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PrivateNICStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1PrivateNICState) *ScalewayInstanceV1PrivateNICState {
		return &v
	}).(ScalewayInstanceV1PrivateNICStatePtrOutput)
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1PrivateNICState] {
	return pulumix.Output[ScalewayInstanceV1PrivateNICState]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1PrivateNICState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PrivateNICStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScalewayInstanceV1PrivateNICState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1PrivateNICStatePtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PrivateNICStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1PrivateNICState)(nil)).Elem()
}

func (o ScalewayInstanceV1PrivateNICStatePtrOutput) ToScalewayInstanceV1PrivateNICStatePtrOutput() ScalewayInstanceV1PrivateNICStatePtrOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICStatePtrOutput) ToScalewayInstanceV1PrivateNICStatePtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PrivateNICStatePtrOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICStatePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1PrivateNICState] {
	return pulumix.Output[*ScalewayInstanceV1PrivateNICState]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1PrivateNICStatePtrOutput) Elem() ScalewayInstanceV1PrivateNICStateOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNICState) ScalewayInstanceV1PrivateNICState {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1PrivateNICState
		return ret
	}).(ScalewayInstanceV1PrivateNICStateOutput)
}

func (o ScalewayInstanceV1PrivateNICStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1PrivateNICStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScalewayInstanceV1PrivateNICState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1PrivateNICStateOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PrivateNICStatePtrOutput{})
}
