// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servers

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Server struct {
	pulumi.CustomResourceState

	// The boot type to use
	Boot_type BootTypePtrOutput `pulumi:"boot_type"`
	// The bootscript ID to use when `boot_type` is set to `bootscript`
	Bootscript pulumi.StringPtrOutput `pulumi:"bootscript"`
	// Define the server commercial type (i.e. GP1-S)
	Commercial_type pulumi.StringOutput `pulumi:"commercial_type"`
	// Define if a dynamic IP is required for the instance
	Dynamic_ip_required pulumi.BoolPtrOutput `pulumi:"dynamic_ip_required"`
	// True if IPv6 is enabled on the server
	Enable_ipv6 pulumi.BoolPtrOutput `pulumi:"enable_ipv6"`
	// The server image ID
	Image pulumi.StringPtrOutput `pulumi:"image"`
	// The server name
	Name pulumi.StringOutput `pulumi:"name"`
	// The server organization ID
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// Placement group ID if server must be part of a placement group
	Placement_group pulumi.StringPtrOutput `pulumi:"placement_group"`
	// The server project ID
	Project pulumi.StringPtrOutput `pulumi:"project"`
	// The ID of the reserved IP to attach to the server
	Public_ip pulumi.StringPtrOutput `pulumi:"public_ip"`
	// The security group ID
	Security_group pulumi.StringPtrOutput            `pulumi:"security_group"`
	Server         ScalewayInstanceV1ServerPtrOutput `pulumi:"server"`
	// The server tags
	Tags    pulumi.StringArrayOutput                        `pulumi:"tags"`
	Volumes ScalewayInstanceV1VolumeServerTemplateMapOutput `pulumi:"volumes"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Commercial_type == nil {
		return nil, errors.New("invalid value for required argument 'Commercial_type'")
	}
	if args.Boot_type == nil {
		args.Boot_type = BootType("local")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("scaleway-instances:servers:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("scaleway-instances:servers:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
}

type ServerState struct {
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The boot type to use
	Boot_type *BootType `pulumi:"boot_type"`
	// The bootscript ID to use when `boot_type` is set to `bootscript`
	Bootscript *string `pulumi:"bootscript"`
	// Define the server commercial type (i.e. GP1-S)
	Commercial_type string `pulumi:"commercial_type"`
	// Define if a dynamic IP is required for the instance
	Dynamic_ip_required *bool `pulumi:"dynamic_ip_required"`
	// True if IPv6 is enabled on the server
	Enable_ipv6 *bool `pulumi:"enable_ipv6"`
	// The server image ID
	Image *string `pulumi:"image"`
	// The server name
	Name *string `pulumi:"name"`
	// The server organization ID
	Organization *string `pulumi:"organization"`
	// Placement group ID if server must be part of a placement group
	Placement_group *string `pulumi:"placement_group"`
	// The server project ID
	Project *string `pulumi:"project"`
	// The ID of the reserved IP to attach to the server
	Public_ip *string `pulumi:"public_ip"`
	// The security group ID
	Security_group *string `pulumi:"security_group"`
	// The server tags
	Tags    []string                                          `pulumi:"tags"`
	Volumes map[string]ScalewayInstanceV1VolumeServerTemplate `pulumi:"volumes"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The boot type to use
	Boot_type BootTypePtrInput
	// The bootscript ID to use when `boot_type` is set to `bootscript`
	Bootscript pulumi.StringPtrInput
	// Define the server commercial type (i.e. GP1-S)
	Commercial_type pulumi.StringInput
	// Define if a dynamic IP is required for the instance
	Dynamic_ip_required pulumi.BoolPtrInput
	// True if IPv6 is enabled on the server
	Enable_ipv6 pulumi.BoolPtrInput
	// The server image ID
	Image pulumi.StringPtrInput
	// The server name
	Name pulumi.StringPtrInput
	// The server organization ID
	Organization pulumi.StringPtrInput
	// Placement group ID if server must be part of a placement group
	Placement_group pulumi.StringPtrInput
	// The server project ID
	Project pulumi.StringPtrInput
	// The ID of the reserved IP to attach to the server
	Public_ip pulumi.StringPtrInput
	// The security group ID
	Security_group pulumi.StringPtrInput
	// The server tags
	Tags    pulumi.StringArrayInput
	Volumes ScalewayInstanceV1VolumeServerTemplateMapInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

func (i *Server) ToOutput(ctx context.Context) pulumix.Output[*Server] {
	return pulumix.Output[*Server]{
		OutputState: i.ToServerOutputWithContext(ctx).OutputState,
	}
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func (o ServerOutput) ToOutput(ctx context.Context) pulumix.Output[*Server] {
	return pulumix.Output[*Server]{
		OutputState: o.OutputState,
	}
}

// The boot type to use
func (o ServerOutput) Boot_type() BootTypePtrOutput {
	return o.ApplyT(func(v *Server) BootTypePtrOutput { return v.Boot_type }).(BootTypePtrOutput)
}

// The bootscript ID to use when `boot_type` is set to `bootscript`
func (o ServerOutput) Bootscript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Bootscript }).(pulumi.StringPtrOutput)
}

// Define the server commercial type (i.e. GP1-S)
func (o ServerOutput) Commercial_type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Commercial_type }).(pulumi.StringOutput)
}

// Define if a dynamic IP is required for the instance
func (o ServerOutput) Dynamic_ip_required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.BoolPtrOutput { return v.Dynamic_ip_required }).(pulumi.BoolPtrOutput)
}

// True if IPv6 is enabled on the server
func (o ServerOutput) Enable_ipv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.BoolPtrOutput { return v.Enable_ipv6 }).(pulumi.BoolPtrOutput)
}

// The server image ID
func (o ServerOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Image }).(pulumi.StringPtrOutput)
}

// The server name
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The server organization ID
func (o ServerOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

// Placement group ID if server must be part of a placement group
func (o ServerOutput) Placement_group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Placement_group }).(pulumi.StringPtrOutput)
}

// The server project ID
func (o ServerOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Project }).(pulumi.StringPtrOutput)
}

// The ID of the reserved IP to attach to the server
func (o ServerOutput) Public_ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Public_ip }).(pulumi.StringPtrOutput)
}

// The security group ID
func (o ServerOutput) Security_group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Security_group }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Server() ScalewayInstanceV1ServerPtrOutput {
	return o.ApplyT(func(v *Server) ScalewayInstanceV1ServerPtrOutput { return v.Server }).(ScalewayInstanceV1ServerPtrOutput)
}

// The server tags
func (o ServerOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Server) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o ServerOutput) Volumes() ScalewayInstanceV1VolumeServerTemplateMapOutput {
	return o.ApplyT(func(v *Server) ScalewayInstanceV1VolumeServerTemplateMapOutput { return v.Volumes }).(ScalewayInstanceV1VolumeServerTemplateMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterOutputType(ServerOutput{})
}
