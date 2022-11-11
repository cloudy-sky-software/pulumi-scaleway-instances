// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package action

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerAction struct {
	pulumi.CustomResourceState

	// The action to perform on the server
	Action ActionPtrOutput `pulumi:"action"`
	// The name of the backup you want to create.
	// This field should only be specified when performing a backup action.
	Name    pulumi.StringPtrOutput                                             `pulumi:"name"`
	Volumes ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateMapOutput `pulumi:"volumes"`
}

// NewServerAction registers a new resource with the given unique name, arguments, and options.
func NewServerAction(ctx *pulumi.Context,
	name string, args *ServerActionArgs, opts ...pulumi.ResourceOption) (*ServerAction, error) {
	if args == nil {
		args = &ServerActionArgs{}
	}

	if isZero(args.Action) {
		args.Action = Action("poweron")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServerAction
	err := ctx.RegisterResource("scaleway-instances:action:ServerAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerAction gets an existing ServerAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerActionState, opts ...pulumi.ResourceOption) (*ServerAction, error) {
	var resource ServerAction
	err := ctx.ReadResource("scaleway-instances:action:ServerAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerAction resources.
type serverActionState struct {
}

type ServerActionState struct {
}

func (ServerActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverActionState)(nil)).Elem()
}

type serverActionArgs struct {
	// The action to perform on the server
	Action *Action `pulumi:"action"`
	// UUID of the server
	Id *string `pulumi:"id"`
	// The name of the backup you want to create.
	// This field should only be specified when performing a backup action.
	Name    *string                                                              `pulumi:"name"`
	Volumes map[string]ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate `pulumi:"volumes"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ServerAction resource.
type ServerActionArgs struct {
	// The action to perform on the server
	Action ActionPtrInput
	// UUID of the server
	Id pulumi.StringPtrInput
	// The name of the backup you want to create.
	// This field should only be specified when performing a backup action.
	Name    pulumi.StringPtrInput
	Volumes ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateMapInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (ServerActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverActionArgs)(nil)).Elem()
}

type ServerActionInput interface {
	pulumi.Input

	ToServerActionOutput() ServerActionOutput
	ToServerActionOutputWithContext(ctx context.Context) ServerActionOutput
}

func (*ServerAction) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAction)(nil)).Elem()
}

func (i *ServerAction) ToServerActionOutput() ServerActionOutput {
	return i.ToServerActionOutputWithContext(context.Background())
}

func (i *ServerAction) ToServerActionOutputWithContext(ctx context.Context) ServerActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerActionOutput)
}

type ServerActionOutput struct{ *pulumi.OutputState }

func (ServerActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAction)(nil)).Elem()
}

func (o ServerActionOutput) ToServerActionOutput() ServerActionOutput {
	return o
}

func (o ServerActionOutput) ToServerActionOutputWithContext(ctx context.Context) ServerActionOutput {
	return o
}

// The action to perform on the server
func (o ServerActionOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *ServerAction) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The name of the backup you want to create.
// This field should only be specified when performing a backup action.
func (o ServerActionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerAction) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServerActionOutput) Volumes() ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateMapOutput {
	return o.ApplyT(func(v *ServerAction) ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateMapOutput {
		return v.Volumes
	}).(ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerActionInput)(nil)).Elem(), &ServerAction{})
	pulumi.RegisterOutputType(ServerActionOutput{})
}
