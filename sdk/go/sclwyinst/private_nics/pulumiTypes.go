// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_nics

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1GetPrivateNICResponse struct {
	Private_nic *ScalewayInstanceV1PrivateNIC `pulumi:"private_nic"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1GetPrivateNICResponse
func (val *ScalewayInstanceV1GetPrivateNICResponse) Defaults() *ScalewayInstanceV1GetPrivateNICResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Private_nic = tmp.Private_nic.Defaults()

	return &tmp
}

type ScalewayInstanceV1GetPrivateNICResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetPrivateNICResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetPrivateNICResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetPrivateNICResponseOutput) ToScalewayInstanceV1GetPrivateNICResponseOutput() ScalewayInstanceV1GetPrivateNICResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetPrivateNICResponseOutput) ToScalewayInstanceV1GetPrivateNICResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetPrivateNICResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetPrivateNICResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1GetPrivateNICResponse] {
	return pulumix.Output[ScalewayInstanceV1GetPrivateNICResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetPrivateNICResponseOutput) Private_nic() ScalewayInstanceV1PrivateNICPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetPrivateNICResponse) *ScalewayInstanceV1PrivateNIC { return v.Private_nic }).(ScalewayInstanceV1PrivateNICPtrOutput)
}

type ScalewayInstanceV1ListPrivateNICsResponse struct {
	Private_nics []ScalewayInstanceV1PrivateNIC `pulumi:"private_nics"`
}

type ScalewayInstanceV1ListPrivateNICsResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListPrivateNICsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListPrivateNICsResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1ListPrivateNICsResponseOutput) ToScalewayInstanceV1ListPrivateNICsResponseOutput() ScalewayInstanceV1ListPrivateNICsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListPrivateNICsResponseOutput) ToScalewayInstanceV1ListPrivateNICsResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1ListPrivateNICsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListPrivateNICsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1ListPrivateNICsResponse] {
	return pulumix.Output[ScalewayInstanceV1ListPrivateNICsResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1ListPrivateNICsResponseOutput) Private_nics() ScalewayInstanceV1PrivateNICArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ListPrivateNICsResponse) []ScalewayInstanceV1PrivateNIC {
		return v.Private_nics
	}).(ScalewayInstanceV1PrivateNICArrayOutput)
}

type ScalewayInstanceV1PrivateNIC struct {
	// The private NIC unique ID
	Id *string `pulumi:"id"`
	// The private NIC MAC address
	Mac_address *string `pulumi:"mac_address"`
	// The private network where the private NIC is attached
	Private_network_id *string `pulumi:"private_network_id"`
	// The server the private NIC is attached to
	Server_id *string `pulumi:"server_id"`
	// The private NIC state
	State *ScalewayInstanceV1PrivateNICState `pulumi:"state"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1PrivateNIC
func (val *ScalewayInstanceV1PrivateNIC) Defaults() *ScalewayInstanceV1PrivateNIC {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.State == nil {
		state_ := ScalewayInstanceV1PrivateNICState("available")
		tmp.State = &state_
	}
	return &tmp
}

type ScalewayInstanceV1PrivateNICOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PrivateNICOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1PrivateNIC)(nil)).Elem()
}

func (o ScalewayInstanceV1PrivateNICOutput) ToScalewayInstanceV1PrivateNICOutput() ScalewayInstanceV1PrivateNICOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICOutput) ToScalewayInstanceV1PrivateNICOutputWithContext(ctx context.Context) ScalewayInstanceV1PrivateNICOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1PrivateNIC] {
	return pulumix.Output[ScalewayInstanceV1PrivateNIC]{
		OutputState: o.OutputState,
	}
}

// The private NIC unique ID
func (o ScalewayInstanceV1PrivateNICOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1PrivateNIC) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The private NIC MAC address
func (o ScalewayInstanceV1PrivateNICOutput) Mac_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1PrivateNIC) *string { return v.Mac_address }).(pulumi.StringPtrOutput)
}

// The private network where the private NIC is attached
func (o ScalewayInstanceV1PrivateNICOutput) Private_network_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1PrivateNIC) *string { return v.Private_network_id }).(pulumi.StringPtrOutput)
}

// The server the private NIC is attached to
func (o ScalewayInstanceV1PrivateNICOutput) Server_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1PrivateNIC) *string { return v.Server_id }).(pulumi.StringPtrOutput)
}

// The private NIC state
func (o ScalewayInstanceV1PrivateNICOutput) State() ScalewayInstanceV1PrivateNICStatePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1PrivateNIC) *ScalewayInstanceV1PrivateNICState { return v.State }).(ScalewayInstanceV1PrivateNICStatePtrOutput)
}

type ScalewayInstanceV1PrivateNICPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PrivateNICPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1PrivateNIC)(nil)).Elem()
}

func (o ScalewayInstanceV1PrivateNICPtrOutput) ToScalewayInstanceV1PrivateNICPtrOutput() ScalewayInstanceV1PrivateNICPtrOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICPtrOutput) ToScalewayInstanceV1PrivateNICPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1PrivateNICPtrOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1PrivateNIC] {
	return pulumix.Output[*ScalewayInstanceV1PrivateNIC]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1PrivateNICPtrOutput) Elem() ScalewayInstanceV1PrivateNICOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNIC) ScalewayInstanceV1PrivateNIC {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1PrivateNIC
		return ret
	}).(ScalewayInstanceV1PrivateNICOutput)
}

// The private NIC unique ID
func (o ScalewayInstanceV1PrivateNICPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNIC) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The private NIC MAC address
func (o ScalewayInstanceV1PrivateNICPtrOutput) Mac_address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNIC) *string {
		if v == nil {
			return nil
		}
		return v.Mac_address
	}).(pulumi.StringPtrOutput)
}

// The private network where the private NIC is attached
func (o ScalewayInstanceV1PrivateNICPtrOutput) Private_network_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNIC) *string {
		if v == nil {
			return nil
		}
		return v.Private_network_id
	}).(pulumi.StringPtrOutput)
}

// The server the private NIC is attached to
func (o ScalewayInstanceV1PrivateNICPtrOutput) Server_id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNIC) *string {
		if v == nil {
			return nil
		}
		return v.Server_id
	}).(pulumi.StringPtrOutput)
}

// The private NIC state
func (o ScalewayInstanceV1PrivateNICPtrOutput) State() ScalewayInstanceV1PrivateNICStatePtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1PrivateNIC) *ScalewayInstanceV1PrivateNICState {
		if v == nil {
			return nil
		}
		return v.State
	}).(ScalewayInstanceV1PrivateNICStatePtrOutput)
}

type ScalewayInstanceV1PrivateNICArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1PrivateNICArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1PrivateNIC)(nil)).Elem()
}

func (o ScalewayInstanceV1PrivateNICArrayOutput) ToScalewayInstanceV1PrivateNICArrayOutput() ScalewayInstanceV1PrivateNICArrayOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICArrayOutput) ToScalewayInstanceV1PrivateNICArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1PrivateNICArrayOutput {
	return o
}

func (o ScalewayInstanceV1PrivateNICArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ScalewayInstanceV1PrivateNIC] {
	return pulumix.Output[[]ScalewayInstanceV1PrivateNIC]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1PrivateNICArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1PrivateNICOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1PrivateNIC {
		return vs[0].([]ScalewayInstanceV1PrivateNIC)[vs[1].(int)]
	}).(ScalewayInstanceV1PrivateNICOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1GetPrivateNICResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListPrivateNICsResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PrivateNICOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PrivateNICPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1PrivateNICArrayOutput{})
}
