// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package placement_groups

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type PlacementGroup struct {
	pulumi.CustomResourceState

	// The placement group name
	Name pulumi.StringOutput `pulumi:"name"`
	// The placement group organization ID
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	Policy_mode  PolicyModePtrOutput    `pulumi:"policy_mode"`
	// Returns true if the policy is respected, false otherwise
	Policy_respected pulumi.BoolPtrOutput `pulumi:"policy_respected"`
	Policy_type      PolicyTypePtrOutput  `pulumi:"policy_type"`
	// The placement group project ID
	Project pulumi.StringOutput `pulumi:"project"`
	// The placement group tags
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The zone in which is the placement group
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Policy_mode == nil {
		args.Policy_mode = PolicyMode("optional")
	}
	if args.Policy_type == nil {
		args.Policy_type = PolicyType("max_availability")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PlacementGroup
	err := ctx.RegisterResource("scaleway-instances:placement_groups:PlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlacementGroup gets an existing PlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementGroupState, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	var resource PlacementGroup
	err := ctx.ReadResource("scaleway-instances:placement_groups:PlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlacementGroup resources.
type placementGroupState struct {
}

type PlacementGroupState struct {
}

func (PlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupState)(nil)).Elem()
}

type placementGroupArgs struct {
	// The placement group name
	Name *string `pulumi:"name"`
	// The placement group organization ID
	Organization *string     `pulumi:"organization"`
	Policy_mode  *PolicyMode `pulumi:"policy_mode"`
	Policy_type  *PolicyType `pulumi:"policy_type"`
	// The placement group project ID
	Project string `pulumi:"project"`
	// The placement group tags
	Tags []string `pulumi:"tags"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// The placement group name
	Name pulumi.StringPtrInput
	// The placement group organization ID
	Organization pulumi.StringPtrInput
	Policy_mode  PolicyModePtrInput
	Policy_type  PolicyTypePtrInput
	// The placement group project ID
	Project pulumi.StringInput
	// The placement group tags
	Tags pulumi.StringArrayInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (PlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupArgs)(nil)).Elem()
}

type PlacementGroupInput interface {
	pulumi.Input

	ToPlacementGroupOutput() PlacementGroupOutput
	ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput
}

func (*PlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementGroup)(nil)).Elem()
}

func (i *PlacementGroup) ToPlacementGroupOutput() PlacementGroupOutput {
	return i.ToPlacementGroupOutputWithContext(context.Background())
}

func (i *PlacementGroup) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupOutput)
}

func (i *PlacementGroup) ToOutput(ctx context.Context) pulumix.Output[*PlacementGroup] {
	return pulumix.Output[*PlacementGroup]{
		OutputState: i.ToPlacementGroupOutputWithContext(ctx).OutputState,
	}
}

type PlacementGroupOutput struct{ *pulumi.OutputState }

func (PlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementGroup)(nil)).Elem()
}

func (o PlacementGroupOutput) ToPlacementGroupOutput() PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*PlacementGroup] {
	return pulumix.Output[*PlacementGroup]{
		OutputState: o.OutputState,
	}
}

// The placement group name
func (o PlacementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PlacementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The placement group organization ID
func (o PlacementGroupOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementGroup) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o PlacementGroupOutput) Policy_mode() PolicyModePtrOutput {
	return o.ApplyT(func(v *PlacementGroup) PolicyModePtrOutput { return v.Policy_mode }).(PolicyModePtrOutput)
}

// Returns true if the policy is respected, false otherwise
func (o PlacementGroupOutput) Policy_respected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PlacementGroup) pulumi.BoolPtrOutput { return v.Policy_respected }).(pulumi.BoolPtrOutput)
}

func (o PlacementGroupOutput) Policy_type() PolicyTypePtrOutput {
	return o.ApplyT(func(v *PlacementGroup) PolicyTypePtrOutput { return v.Policy_type }).(PolicyTypePtrOutput)
}

// The placement group project ID
func (o PlacementGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PlacementGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The placement group tags
func (o PlacementGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PlacementGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone in which is the placement group
func (o PlacementGroupOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementGroup) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlacementGroupInput)(nil)).Elem(), &PlacementGroup{})
	pulumi.RegisterOutputType(PlacementGroupOutput{})
}
