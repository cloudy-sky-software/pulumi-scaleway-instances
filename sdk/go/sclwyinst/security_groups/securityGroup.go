// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security_groups

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityGroup struct {
	pulumi.CustomResourceState

	// The security group creation date (RFC 3339 format)
	CreationDate pulumi.StringPtrOutput `pulumi:"creationDate"`
	// The security groups description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
	EnableDefaultSecurity pulumi.BoolPtrOutput `pulumi:"enableDefaultSecurity"`
	// The default inbound policy
	InboundDefaultPolicy SecurityGroupInboundDefaultPolicyPtrOutput `pulumi:"inboundDefaultPolicy"`
	// The security group modification date (RFC 3339 format)
	ModificationDate pulumi.StringPtrOutput `pulumi:"modificationDate"`
	// The security groups name
	Name pulumi.StringOutput `pulumi:"name"`
	// The security groups organization ID
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// True if it is your default security group for this organization ID
	OrganizationDefault pulumi.BoolPtrOutput `pulumi:"organizationDefault"`
	// The default outbound policy
	OutboundDefaultPolicy SecurityGroupOutboundDefaultPolicyPtrOutput `pulumi:"outboundDefaultPolicy"`
	// The security group project ID
	Project pulumi.StringOutput `pulumi:"project"`
	// True if it is your default security group for this project ID
	ProjectDefault pulumi.BoolPtrOutput                     `pulumi:"projectDefault"`
	SecurityGroup  ScalewayInstanceV1SecurityGroupPtrOutput `pulumi:"securityGroup"`
	// List of servers attached to this security group
	Servers ScalewayInstanceV1ServerSummaryArrayOutput `pulumi:"servers"`
	// Security group state
	State SecurityGroupStateEnumPtrOutput `pulumi:"state"`
	// True if the security group is stateful
	Stateful pulumi.BoolPtrOutput `pulumi:"stateful"`
	// The security group tags
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The zone in which is the security group
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.InboundDefaultPolicy == nil {
		args.InboundDefaultPolicy = SecurityGroupInboundDefaultPolicy("accept")
	}
	if args.OutboundDefaultPolicy == nil {
		args.OutboundDefaultPolicy = SecurityGroupOutboundDefaultPolicy("accept")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityGroup
	err := ctx.RegisterResource("scaleway-instances:security_groups:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("scaleway-instances:security_groups:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
}

type SecurityGroupState struct {
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The security groups description
	Description *string `pulumi:"description"`
	// True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	// The default inbound policy
	InboundDefaultPolicy *SecurityGroupInboundDefaultPolicy `pulumi:"inboundDefaultPolicy"`
	// The security groups name
	Name *string `pulumi:"name"`
	// The security groups organization ID
	Organization *string `pulumi:"organization"`
	// True if it is your default security group for this organization ID
	OrganizationDefault *bool `pulumi:"organizationDefault"`
	// The default outbound policy
	OutboundDefaultPolicy *SecurityGroupOutboundDefaultPolicy `pulumi:"outboundDefaultPolicy"`
	// The security group project ID
	Project string `pulumi:"project"`
	// True if it is your default security group for this project ID
	ProjectDefault *bool `pulumi:"projectDefault"`
	// True if the security group is stateful
	Stateful *bool `pulumi:"stateful"`
	// The security group tags
	Tags []string `pulumi:"tags"`
	// The zone you want to target
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The security groups description
	Description pulumi.StringPtrInput
	// True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
	EnableDefaultSecurity pulumi.BoolPtrInput
	// The default inbound policy
	InboundDefaultPolicy SecurityGroupInboundDefaultPolicyPtrInput
	// The security groups name
	Name pulumi.StringPtrInput
	// The security groups organization ID
	Organization pulumi.StringPtrInput
	// True if it is your default security group for this organization ID
	OrganizationDefault pulumi.BoolPtrInput
	// The default outbound policy
	OutboundDefaultPolicy SecurityGroupOutboundDefaultPolicyPtrInput
	// The security group project ID
	Project pulumi.StringInput
	// True if it is your default security group for this project ID
	ProjectDefault pulumi.BoolPtrInput
	// True if the security group is stateful
	Stateful pulumi.BoolPtrInput
	// The security group tags
	Tags pulumi.StringArrayInput
	// The zone you want to target
	Zone pulumi.StringPtrInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

type SecurityGroupOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

// The security group creation date (RFC 3339 format)
func (o SecurityGroupOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.CreationDate }).(pulumi.StringPtrOutput)
}

// The security groups description
func (o SecurityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
func (o SecurityGroupOutput) EnableDefaultSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.BoolPtrOutput { return v.EnableDefaultSecurity }).(pulumi.BoolPtrOutput)
}

// The default inbound policy
func (o SecurityGroupOutput) InboundDefaultPolicy() SecurityGroupInboundDefaultPolicyPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroupInboundDefaultPolicyPtrOutput { return v.InboundDefaultPolicy }).(SecurityGroupInboundDefaultPolicyPtrOutput)
}

// The security group modification date (RFC 3339 format)
func (o SecurityGroupOutput) ModificationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.ModificationDate }).(pulumi.StringPtrOutput)
}

// The security groups name
func (o SecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The security groups organization ID
func (o SecurityGroupOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

// True if it is your default security group for this organization ID
func (o SecurityGroupOutput) OrganizationDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.BoolPtrOutput { return v.OrganizationDefault }).(pulumi.BoolPtrOutput)
}

// The default outbound policy
func (o SecurityGroupOutput) OutboundDefaultPolicy() SecurityGroupOutboundDefaultPolicyPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroupOutboundDefaultPolicyPtrOutput { return v.OutboundDefaultPolicy }).(SecurityGroupOutboundDefaultPolicyPtrOutput)
}

// The security group project ID
func (o SecurityGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// True if it is your default security group for this project ID
func (o SecurityGroupOutput) ProjectDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.BoolPtrOutput { return v.ProjectDefault }).(pulumi.BoolPtrOutput)
}

func (o SecurityGroupOutput) SecurityGroup() ScalewayInstanceV1SecurityGroupPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) ScalewayInstanceV1SecurityGroupPtrOutput { return v.SecurityGroup }).(ScalewayInstanceV1SecurityGroupPtrOutput)
}

// List of servers attached to this security group
func (o SecurityGroupOutput) Servers() ScalewayInstanceV1ServerSummaryArrayOutput {
	return o.ApplyT(func(v *SecurityGroup) ScalewayInstanceV1ServerSummaryArrayOutput { return v.Servers }).(ScalewayInstanceV1ServerSummaryArrayOutput)
}

// Security group state
func (o SecurityGroupOutput) State() SecurityGroupStateEnumPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroupStateEnumPtrOutput { return v.State }).(SecurityGroupStateEnumPtrOutput)
}

// True if the security group is stateful
func (o SecurityGroupOutput) Stateful() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.BoolPtrOutput { return v.Stateful }).(pulumi.BoolPtrOutput)
}

// The security group tags
func (o SecurityGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone in which is the security group
func (o SecurityGroupOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterOutputType(SecurityGroupOutput{})
}
