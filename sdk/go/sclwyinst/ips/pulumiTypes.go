// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ips

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type GoogleProtobufStringValue struct {
}

// GoogleProtobufStringValueInput is an input type that accepts GoogleProtobufStringValueArgs and GoogleProtobufStringValueOutput values.
// You can construct a concrete instance of `GoogleProtobufStringValueInput` via:
//
//	GoogleProtobufStringValueArgs{...}
type GoogleProtobufStringValueInput interface {
	pulumi.Input

	ToGoogleProtobufStringValueOutput() GoogleProtobufStringValueOutput
	ToGoogleProtobufStringValueOutputWithContext(context.Context) GoogleProtobufStringValueOutput
}

type GoogleProtobufStringValueArgs struct {
}

func (GoogleProtobufStringValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleProtobufStringValue)(nil)).Elem()
}

func (i GoogleProtobufStringValueArgs) ToGoogleProtobufStringValueOutput() GoogleProtobufStringValueOutput {
	return i.ToGoogleProtobufStringValueOutputWithContext(context.Background())
}

func (i GoogleProtobufStringValueArgs) ToGoogleProtobufStringValueOutputWithContext(ctx context.Context) GoogleProtobufStringValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleProtobufStringValueOutput)
}

func (i GoogleProtobufStringValueArgs) ToOutput(ctx context.Context) pulumix.Output[GoogleProtobufStringValue] {
	return pulumix.Output[GoogleProtobufStringValue]{
		OutputState: i.ToGoogleProtobufStringValueOutputWithContext(ctx).OutputState,
	}
}

func (i GoogleProtobufStringValueArgs) ToGoogleProtobufStringValuePtrOutput() GoogleProtobufStringValuePtrOutput {
	return i.ToGoogleProtobufStringValuePtrOutputWithContext(context.Background())
}

func (i GoogleProtobufStringValueArgs) ToGoogleProtobufStringValuePtrOutputWithContext(ctx context.Context) GoogleProtobufStringValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleProtobufStringValueOutput).ToGoogleProtobufStringValuePtrOutputWithContext(ctx)
}

// GoogleProtobufStringValuePtrInput is an input type that accepts GoogleProtobufStringValueArgs, GoogleProtobufStringValuePtr and GoogleProtobufStringValuePtrOutput values.
// You can construct a concrete instance of `GoogleProtobufStringValuePtrInput` via:
//
//	        GoogleProtobufStringValueArgs{...}
//
//	or:
//
//	        nil
type GoogleProtobufStringValuePtrInput interface {
	pulumi.Input

	ToGoogleProtobufStringValuePtrOutput() GoogleProtobufStringValuePtrOutput
	ToGoogleProtobufStringValuePtrOutputWithContext(context.Context) GoogleProtobufStringValuePtrOutput
}

type googleProtobufStringValuePtrType GoogleProtobufStringValueArgs

func GoogleProtobufStringValuePtr(v *GoogleProtobufStringValueArgs) GoogleProtobufStringValuePtrInput {
	return (*googleProtobufStringValuePtrType)(v)
}

func (*googleProtobufStringValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleProtobufStringValue)(nil)).Elem()
}

func (i *googleProtobufStringValuePtrType) ToGoogleProtobufStringValuePtrOutput() GoogleProtobufStringValuePtrOutput {
	return i.ToGoogleProtobufStringValuePtrOutputWithContext(context.Background())
}

func (i *googleProtobufStringValuePtrType) ToGoogleProtobufStringValuePtrOutputWithContext(ctx context.Context) GoogleProtobufStringValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleProtobufStringValuePtrOutput)
}

func (i *googleProtobufStringValuePtrType) ToOutput(ctx context.Context) pulumix.Output[*GoogleProtobufStringValue] {
	return pulumix.Output[*GoogleProtobufStringValue]{
		OutputState: i.ToGoogleProtobufStringValuePtrOutputWithContext(ctx).OutputState,
	}
}

type GoogleProtobufStringValueOutput struct{ *pulumi.OutputState }

func (GoogleProtobufStringValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleProtobufStringValue)(nil)).Elem()
}

func (o GoogleProtobufStringValueOutput) ToGoogleProtobufStringValueOutput() GoogleProtobufStringValueOutput {
	return o
}

func (o GoogleProtobufStringValueOutput) ToGoogleProtobufStringValueOutputWithContext(ctx context.Context) GoogleProtobufStringValueOutput {
	return o
}

func (o GoogleProtobufStringValueOutput) ToGoogleProtobufStringValuePtrOutput() GoogleProtobufStringValuePtrOutput {
	return o.ToGoogleProtobufStringValuePtrOutputWithContext(context.Background())
}

func (o GoogleProtobufStringValueOutput) ToGoogleProtobufStringValuePtrOutputWithContext(ctx context.Context) GoogleProtobufStringValuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleProtobufStringValue) *GoogleProtobufStringValue {
		return &v
	}).(GoogleProtobufStringValuePtrOutput)
}

func (o GoogleProtobufStringValueOutput) ToOutput(ctx context.Context) pulumix.Output[GoogleProtobufStringValue] {
	return pulumix.Output[GoogleProtobufStringValue]{
		OutputState: o.OutputState,
	}
}

type GoogleProtobufStringValuePtrOutput struct{ *pulumi.OutputState }

func (GoogleProtobufStringValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleProtobufStringValue)(nil)).Elem()
}

func (o GoogleProtobufStringValuePtrOutput) ToGoogleProtobufStringValuePtrOutput() GoogleProtobufStringValuePtrOutput {
	return o
}

func (o GoogleProtobufStringValuePtrOutput) ToGoogleProtobufStringValuePtrOutputWithContext(ctx context.Context) GoogleProtobufStringValuePtrOutput {
	return o
}

func (o GoogleProtobufStringValuePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*GoogleProtobufStringValue] {
	return pulumix.Output[*GoogleProtobufStringValue]{
		OutputState: o.OutputState,
	}
}

func (o GoogleProtobufStringValuePtrOutput) Elem() GoogleProtobufStringValueOutput {
	return o.ApplyT(func(v *GoogleProtobufStringValue) GoogleProtobufStringValue {
		if v != nil {
			return *v
		}
		var ret GoogleProtobufStringValue
		return ret
	}).(GoogleProtobufStringValueOutput)
}

type ScalewayInstanceV1GetIpResponse struct {
	Ip *ScalewayInstanceV1Ip `pulumi:"ip"`
}

type ScalewayInstanceV1GetIpResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetIpResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetIpResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetIpResponseOutput) ToScalewayInstanceV1GetIpResponseOutput() ScalewayInstanceV1GetIpResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetIpResponseOutput) ToScalewayInstanceV1GetIpResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetIpResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetIpResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1GetIpResponse] {
	return pulumix.Output[ScalewayInstanceV1GetIpResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetIpResponseOutput) Ip() ScalewayInstanceV1IpPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetIpResponse) *ScalewayInstanceV1Ip { return v.Ip }).(ScalewayInstanceV1IpPtrOutput)
}

type ScalewayInstanceV1Ip struct {
	// (IPv4 address)
	Address      *string                          `pulumi:"address"`
	Id           *string                          `pulumi:"id"`
	Organization *string                          `pulumi:"organization"`
	Project      string                           `pulumi:"project"`
	Reverse      *GoogleProtobufStringValue       `pulumi:"reverse"`
	Server       *ScalewayInstanceV1ServerSummary `pulumi:"server"`
	Tags         []string                         `pulumi:"tags"`
	Zone         *string                          `pulumi:"zone"`
}

type ScalewayInstanceV1IpOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1IpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1Ip)(nil)).Elem()
}

func (o ScalewayInstanceV1IpOutput) ToScalewayInstanceV1IpOutput() ScalewayInstanceV1IpOutput {
	return o
}

func (o ScalewayInstanceV1IpOutput) ToScalewayInstanceV1IpOutputWithContext(ctx context.Context) ScalewayInstanceV1IpOutput {
	return o
}

func (o ScalewayInstanceV1IpOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1Ip] {
	return pulumix.Output[ScalewayInstanceV1Ip]{
		OutputState: o.OutputState,
	}
}

// (IPv4 address)
func (o ScalewayInstanceV1IpOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) string { return v.Project }).(pulumi.StringOutput)
}

func (o ScalewayInstanceV1IpOutput) Reverse() GoogleProtobufStringValuePtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) *GoogleProtobufStringValue { return v.Reverse }).(GoogleProtobufStringValuePtrOutput)
}

func (o ScalewayInstanceV1IpOutput) Server() ScalewayInstanceV1ServerSummaryPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) *ScalewayInstanceV1ServerSummary { return v.Server }).(ScalewayInstanceV1ServerSummaryPtrOutput)
}

func (o ScalewayInstanceV1IpOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o ScalewayInstanceV1IpOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Ip) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1IpPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1IpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1Ip)(nil)).Elem()
}

func (o ScalewayInstanceV1IpPtrOutput) ToScalewayInstanceV1IpPtrOutput() ScalewayInstanceV1IpPtrOutput {
	return o
}

func (o ScalewayInstanceV1IpPtrOutput) ToScalewayInstanceV1IpPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1IpPtrOutput {
	return o
}

func (o ScalewayInstanceV1IpPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1Ip] {
	return pulumix.Output[*ScalewayInstanceV1Ip]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1IpPtrOutput) Elem() ScalewayInstanceV1IpOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) ScalewayInstanceV1Ip {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1Ip
		return ret
	}).(ScalewayInstanceV1IpOutput)
}

// (IPv4 address)
func (o ScalewayInstanceV1IpPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *string {
		if v == nil {
			return nil
		}
		return &v.Project
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Reverse() GoogleProtobufStringValuePtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *GoogleProtobufStringValue {
		if v == nil {
			return nil
		}
		return v.Reverse
	}).(GoogleProtobufStringValuePtrOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Server() ScalewayInstanceV1ServerSummaryPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *ScalewayInstanceV1ServerSummary {
		if v == nil {
			return nil
		}
		return v.Server
	}).(ScalewayInstanceV1ServerSummaryPtrOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) []string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayOutput)
}

func (o ScalewayInstanceV1IpPtrOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Ip) *string {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1IpArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1IpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1Ip)(nil)).Elem()
}

func (o ScalewayInstanceV1IpArrayOutput) ToScalewayInstanceV1IpArrayOutput() ScalewayInstanceV1IpArrayOutput {
	return o
}

func (o ScalewayInstanceV1IpArrayOutput) ToScalewayInstanceV1IpArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1IpArrayOutput {
	return o
}

func (o ScalewayInstanceV1IpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ScalewayInstanceV1Ip] {
	return pulumix.Output[[]ScalewayInstanceV1Ip]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1IpArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1IpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1Ip {
		return vs[0].([]ScalewayInstanceV1Ip)[vs[1].(int)]
	}).(ScalewayInstanceV1IpOutput)
}

type ScalewayInstanceV1ListIpsResponse struct {
	// List of ips
	Ips []ScalewayInstanceV1Ip `pulumi:"ips"`
}

type ScalewayInstanceV1ListIpsResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListIpsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListIpsResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1ListIpsResponseOutput) ToScalewayInstanceV1ListIpsResponseOutput() ScalewayInstanceV1ListIpsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListIpsResponseOutput) ToScalewayInstanceV1ListIpsResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1ListIpsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListIpsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1ListIpsResponse] {
	return pulumix.Output[ScalewayInstanceV1ListIpsResponse]{
		OutputState: o.OutputState,
	}
}

// List of ips
func (o ScalewayInstanceV1ListIpsResponseOutput) Ips() ScalewayInstanceV1IpArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ListIpsResponse) []ScalewayInstanceV1Ip { return v.Ips }).(ScalewayInstanceV1IpArrayOutput)
}

type ScalewayInstanceV1ServerSummary struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

// ScalewayInstanceV1ServerSummaryInput is an input type that accepts ScalewayInstanceV1ServerSummaryArgs and ScalewayInstanceV1ServerSummaryOutput values.
// You can construct a concrete instance of `ScalewayInstanceV1ServerSummaryInput` via:
//
//	ScalewayInstanceV1ServerSummaryArgs{...}
type ScalewayInstanceV1ServerSummaryInput interface {
	pulumi.Input

	ToScalewayInstanceV1ServerSummaryOutput() ScalewayInstanceV1ServerSummaryOutput
	ToScalewayInstanceV1ServerSummaryOutputWithContext(context.Context) ScalewayInstanceV1ServerSummaryOutput
}

type ScalewayInstanceV1ServerSummaryArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ScalewayInstanceV1ServerSummaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ServerSummary)(nil)).Elem()
}

func (i ScalewayInstanceV1ServerSummaryArgs) ToScalewayInstanceV1ServerSummaryOutput() ScalewayInstanceV1ServerSummaryOutput {
	return i.ToScalewayInstanceV1ServerSummaryOutputWithContext(context.Background())
}

func (i ScalewayInstanceV1ServerSummaryArgs) ToScalewayInstanceV1ServerSummaryOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerSummaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalewayInstanceV1ServerSummaryOutput)
}

func (i ScalewayInstanceV1ServerSummaryArgs) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1ServerSummary] {
	return pulumix.Output[ScalewayInstanceV1ServerSummary]{
		OutputState: i.ToScalewayInstanceV1ServerSummaryOutputWithContext(ctx).OutputState,
	}
}

func (i ScalewayInstanceV1ServerSummaryArgs) ToScalewayInstanceV1ServerSummaryPtrOutput() ScalewayInstanceV1ServerSummaryPtrOutput {
	return i.ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(context.Background())
}

func (i ScalewayInstanceV1ServerSummaryArgs) ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerSummaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalewayInstanceV1ServerSummaryOutput).ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(ctx)
}

// ScalewayInstanceV1ServerSummaryPtrInput is an input type that accepts ScalewayInstanceV1ServerSummaryArgs, ScalewayInstanceV1ServerSummaryPtr and ScalewayInstanceV1ServerSummaryPtrOutput values.
// You can construct a concrete instance of `ScalewayInstanceV1ServerSummaryPtrInput` via:
//
//	        ScalewayInstanceV1ServerSummaryArgs{...}
//
//	or:
//
//	        nil
type ScalewayInstanceV1ServerSummaryPtrInput interface {
	pulumi.Input

	ToScalewayInstanceV1ServerSummaryPtrOutput() ScalewayInstanceV1ServerSummaryPtrOutput
	ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(context.Context) ScalewayInstanceV1ServerSummaryPtrOutput
}

type scalewayInstanceV1ServerSummaryPtrType ScalewayInstanceV1ServerSummaryArgs

func ScalewayInstanceV1ServerSummaryPtr(v *ScalewayInstanceV1ServerSummaryArgs) ScalewayInstanceV1ServerSummaryPtrInput {
	return (*scalewayInstanceV1ServerSummaryPtrType)(v)
}

func (*scalewayInstanceV1ServerSummaryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1ServerSummary)(nil)).Elem()
}

func (i *scalewayInstanceV1ServerSummaryPtrType) ToScalewayInstanceV1ServerSummaryPtrOutput() ScalewayInstanceV1ServerSummaryPtrOutput {
	return i.ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(context.Background())
}

func (i *scalewayInstanceV1ServerSummaryPtrType) ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerSummaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalewayInstanceV1ServerSummaryPtrOutput)
}

func (i *scalewayInstanceV1ServerSummaryPtrType) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1ServerSummary] {
	return pulumix.Output[*ScalewayInstanceV1ServerSummary]{
		OutputState: i.ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(ctx).OutputState,
	}
}

type ScalewayInstanceV1ServerSummaryOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ServerSummaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ServerSummary)(nil)).Elem()
}

func (o ScalewayInstanceV1ServerSummaryOutput) ToScalewayInstanceV1ServerSummaryOutput() ScalewayInstanceV1ServerSummaryOutput {
	return o
}

func (o ScalewayInstanceV1ServerSummaryOutput) ToScalewayInstanceV1ServerSummaryOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerSummaryOutput {
	return o
}

func (o ScalewayInstanceV1ServerSummaryOutput) ToScalewayInstanceV1ServerSummaryPtrOutput() ScalewayInstanceV1ServerSummaryPtrOutput {
	return o.ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(context.Background())
}

func (o ScalewayInstanceV1ServerSummaryOutput) ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerSummaryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalewayInstanceV1ServerSummary) *ScalewayInstanceV1ServerSummary {
		return &v
	}).(ScalewayInstanceV1ServerSummaryPtrOutput)
}

func (o ScalewayInstanceV1ServerSummaryOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1ServerSummary] {
	return pulumix.Output[ScalewayInstanceV1ServerSummary]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1ServerSummaryOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ServerSummary) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1ServerSummaryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ServerSummary) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1ServerSummaryPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ServerSummaryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1ServerSummary)(nil)).Elem()
}

func (o ScalewayInstanceV1ServerSummaryPtrOutput) ToScalewayInstanceV1ServerSummaryPtrOutput() ScalewayInstanceV1ServerSummaryPtrOutput {
	return o
}

func (o ScalewayInstanceV1ServerSummaryPtrOutput) ToScalewayInstanceV1ServerSummaryPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1ServerSummaryPtrOutput {
	return o
}

func (o ScalewayInstanceV1ServerSummaryPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1ServerSummary] {
	return pulumix.Output[*ScalewayInstanceV1ServerSummary]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1ServerSummaryPtrOutput) Elem() ScalewayInstanceV1ServerSummaryOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1ServerSummary) ScalewayInstanceV1ServerSummary {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1ServerSummary
		return ret
	}).(ScalewayInstanceV1ServerSummaryOutput)
}

func (o ScalewayInstanceV1ServerSummaryPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1ServerSummary) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1ServerSummaryPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1ServerSummary) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleProtobufStringValueInput)(nil)).Elem(), GoogleProtobufStringValueArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleProtobufStringValuePtrInput)(nil)).Elem(), GoogleProtobufStringValueArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalewayInstanceV1ServerSummaryInput)(nil)).Elem(), ScalewayInstanceV1ServerSummaryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalewayInstanceV1ServerSummaryPtrInput)(nil)).Elem(), ScalewayInstanceV1ServerSummaryArgs{})
	pulumi.RegisterOutputType(GoogleProtobufStringValueOutput{})
	pulumi.RegisterOutputType(GoogleProtobufStringValuePtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetIpResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1IpOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1IpPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1IpArrayOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListIpsResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ServerSummaryOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ServerSummaryPtrOutput{})
}
