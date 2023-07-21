// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Image struct {
	pulumi.CustomResourceState

	Arch ArchPtrOutput `pulumi:"arch"`
	// (RFC 3339 format)
	Creation_date      pulumi.StringPtrOutput                `pulumi:"creation_date"`
	Default_bootscript ScalewayInstanceV1BootscriptPtrOutput `pulumi:"default_bootscript"`
	Extra_volumes      ScalewayInstanceV1VolumeMapOutput     `pulumi:"extra_volumes"`
	From_server        pulumi.StringPtrOutput                `pulumi:"from_server"`
	// (RFC 3339 format)
	Modification_date pulumi.StringPtrOutput                `pulumi:"modification_date"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	Organization      pulumi.StringPtrOutput                `pulumi:"organization"`
	Project           pulumi.StringOutput                   `pulumi:"project"`
	Public            pulumi.BoolPtrOutput                  `pulumi:"public"`
	Root_volume       ScalewayInstanceV1VolumeSummaryOutput `pulumi:"root_volume"`
	State             StatePtrOutput                        `pulumi:"state"`
	Tags              pulumi.StringArrayOutput              `pulumi:"tags"`
	Zone              pulumi.StringPtrOutput                `pulumi:"zone"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Root_volume == nil {
		return nil, errors.New("invalid value for required argument 'Root_volume'")
	}
	if args.Arch == nil {
		args.Arch = Arch("x86_64")
	}
	if args.Default_bootscript != nil {
		args.Default_bootscript = args.Default_bootscript.ToScalewayInstanceV1BootscriptPtrOutput().ApplyT(func(v *ScalewayInstanceV1Bootscript) *ScalewayInstanceV1Bootscript { return v.Defaults() }).(ScalewayInstanceV1BootscriptPtrOutput)
	}
	args.Root_volume = args.Root_volume.ToScalewayInstanceV1VolumeSummaryOutput().ApplyT(func(v ScalewayInstanceV1VolumeSummary) ScalewayInstanceV1VolumeSummary { return *v.Defaults() }).(ScalewayInstanceV1VolumeSummaryOutput)
	if args.State == nil {
		args.State = State("available")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("scaleway-instances:images:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("scaleway-instances:images:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	Arch               *Arch                               `pulumi:"arch"`
	Default_bootscript *ScalewayInstanceV1Bootscript       `pulumi:"default_bootscript"`
	Extra_volumes      map[string]ScalewayInstanceV1Volume `pulumi:"extra_volumes"`
	Name               *string                             `pulumi:"name"`
	Organization       *string                             `pulumi:"organization"`
	Project            string                              `pulumi:"project"`
	Public             *bool                               `pulumi:"public"`
	Root_volume        ScalewayInstanceV1VolumeSummary     `pulumi:"root_volume"`
	State              *State                              `pulumi:"state"`
	Tags               []string                            `pulumi:"tags"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	Arch               ArchPtrInput
	Default_bootscript ScalewayInstanceV1BootscriptPtrInput
	Extra_volumes      ScalewayInstanceV1VolumeMapInput
	Name               pulumi.StringPtrInput
	Organization       pulumi.StringPtrInput
	Project            pulumi.StringInput
	Public             pulumi.BoolPtrInput
	Root_volume        ScalewayInstanceV1VolumeSummaryInput
	State              StatePtrInput
	Tags               pulumi.StringArrayInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func (o ImageOutput) Arch() ArchPtrOutput {
	return o.ApplyT(func(v *Image) ArchPtrOutput { return v.Arch }).(ArchPtrOutput)
}

// (RFC 3339 format)
func (o ImageOutput) Creation_date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Creation_date }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Default_bootscript() ScalewayInstanceV1BootscriptPtrOutput {
	return o.ApplyT(func(v *Image) ScalewayInstanceV1BootscriptPtrOutput { return v.Default_bootscript }).(ScalewayInstanceV1BootscriptPtrOutput)
}

func (o ImageOutput) Extra_volumes() ScalewayInstanceV1VolumeMapOutput {
	return o.ApplyT(func(v *Image) ScalewayInstanceV1VolumeMapOutput { return v.Extra_volumes }).(ScalewayInstanceV1VolumeMapOutput)
}

func (o ImageOutput) From_server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.From_server }).(pulumi.StringPtrOutput)
}

// (RFC 3339 format)
func (o ImageOutput) Modification_date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Modification_date }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ImageOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ImageOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o ImageOutput) Root_volume() ScalewayInstanceV1VolumeSummaryOutput {
	return o.ApplyT(func(v *Image) ScalewayInstanceV1VolumeSummaryOutput { return v.Root_volume }).(ScalewayInstanceV1VolumeSummaryOutput)
}

func (o ImageOutput) State() StatePtrOutput {
	return o.ApplyT(func(v *Image) StatePtrOutput { return v.State }).(StatePtrOutput)
}

func (o ImageOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o ImageOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterOutputType(ImageOutput{})
}
