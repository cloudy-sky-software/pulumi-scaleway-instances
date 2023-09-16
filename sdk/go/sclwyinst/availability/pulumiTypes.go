// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package availability

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1GetServerTypesAvailabilityResponse struct {
	Servers map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability `pulumi:"servers"`
}

type ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetServerTypesAvailabilityResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput) ToScalewayInstanceV1GetServerTypesAvailabilityResponseOutput() ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput) ToScalewayInstanceV1GetServerTypesAvailabilityResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1GetServerTypesAvailabilityResponse] {
	return pulumix.Output[ScalewayInstanceV1GetServerTypesAvailabilityResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput) Servers() ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetServerTypesAvailabilityResponse) map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability {
		return v.Servers
	}).(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput)
}

type ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability struct {
	Availability *ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability `pulumi:"availability"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability
func (val *ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability) Defaults() *ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Availability == nil {
		availability_ := ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability("available")
		tmp.Availability = &availability_
	}
	return &tmp
}

type ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability)(nil)).Elem()
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput) ToScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput() ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput {
	return o
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput) ToScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutputWithContext(ctx context.Context) ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput {
	return o
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability] {
	return pulumix.Output[ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput) Availability() ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailabilityPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability) *ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability {
		return v.Availability
	}).(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailabilityPtrOutput)
}

type ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability)(nil)).Elem()
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput) ToScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput() ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput {
	return o
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput) ToScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutputWithContext(ctx context.Context) ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput {
	return o
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability] {
	return pulumix.Output[map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput) MapIndex(k pulumi.StringInput) ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability {
		return vs[0].(map[string]ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailability)[vs[1].(string)]
	}).(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1GetServerTypesAvailabilityResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityMapOutput{})
}
