// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ips

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ip struct {
	pulumi.CustomResourceState

	// (IPv4 address)
	Address      pulumi.StringPtrOutput                   `pulumi:"address"`
	Organization pulumi.StringPtrOutput                   `pulumi:"organization"`
	Project      pulumi.StringOutput                      `pulumi:"project"`
	Reverse      GoogleProtobufStringValuePtrOutput       `pulumi:"reverse"`
	Server       ScalewayInstanceV1ServerSummaryPtrOutput `pulumi:"server"`
	Tags         pulumi.StringArrayOutput                 `pulumi:"tags"`
	Zone         pulumi.StringPtrOutput                   `pulumi:"zone"`
}

// NewIp registers a new resource with the given unique name, arguments, and options.
func NewIp(ctx *pulumi.Context,
	name string, args *IpArgs, opts ...pulumi.ResourceOption) (*Ip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Ip
	err := ctx.RegisterResource("scaleway-instances:ips:Ip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIp gets an existing Ip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpState, opts ...pulumi.ResourceOption) (*Ip, error) {
	var resource Ip
	err := ctx.ReadResource("scaleway-instances:ips:Ip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ip resources.
type ipState struct {
}

type IpState struct {
}

func (IpState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipState)(nil)).Elem()
}

type ipArgs struct {
	Organization *string                          `pulumi:"organization"`
	Project      string                           `pulumi:"project"`
	Reverse      *GoogleProtobufStringValue       `pulumi:"reverse"`
	Server       *ScalewayInstanceV1ServerSummary `pulumi:"server"`
	Tags         []string                         `pulumi:"tags"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Ip resource.
type IpArgs struct {
	Organization pulumi.StringPtrInput
	Project      pulumi.StringInput
	Reverse      GoogleProtobufStringValuePtrInput
	Server       ScalewayInstanceV1ServerSummaryPtrInput
	Tags         pulumi.StringArrayInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (IpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipArgs)(nil)).Elem()
}

type IpInput interface {
	pulumi.Input

	ToIpOutput() IpOutput
	ToIpOutputWithContext(ctx context.Context) IpOutput
}

func (*Ip) ElementType() reflect.Type {
	return reflect.TypeOf((**Ip)(nil)).Elem()
}

func (i *Ip) ToIpOutput() IpOutput {
	return i.ToIpOutputWithContext(context.Background())
}

func (i *Ip) ToIpOutputWithContext(ctx context.Context) IpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpOutput)
}

type IpOutput struct{ *pulumi.OutputState }

func (IpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ip)(nil)).Elem()
}

func (o IpOutput) ToIpOutput() IpOutput {
	return o
}

func (o IpOutput) ToIpOutputWithContext(ctx context.Context) IpOutput {
	return o
}

// (IPv4 address)
func (o IpOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ip) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IpOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ip) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o IpOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o IpOutput) Reverse() GoogleProtobufStringValuePtrOutput {
	return o.ApplyT(func(v *Ip) GoogleProtobufStringValuePtrOutput { return v.Reverse }).(GoogleProtobufStringValuePtrOutput)
}

func (o IpOutput) Server() ScalewayInstanceV1ServerSummaryPtrOutput {
	return o.ApplyT(func(v *Ip) ScalewayInstanceV1ServerSummaryPtrOutput { return v.Server }).(ScalewayInstanceV1ServerSummaryPtrOutput)
}

func (o IpOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Ip) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o IpOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ip) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpInput)(nil)).Elem(), &Ip{})
	pulumi.RegisterOutputType(IpOutput{})
}
