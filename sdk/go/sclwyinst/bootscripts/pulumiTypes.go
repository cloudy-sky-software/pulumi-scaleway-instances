// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bootscripts

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1Bootscript struct {
	// The bootscript arch
	Arch *ScalewayInstanceV1BootscriptArch `pulumi:"arch"`
	// The bootscript arguments
	Bootcmdargs *string `pulumi:"bootcmdargs"`
	// Dispmay if the bootscript is the default bootscript if no other boot option is configured
	Default *bool `pulumi:"default"`
	// Provide information regarding a Device Tree Binary (dtb) for use with C1 servers
	Dtb *string `pulumi:"dtb"`
	// The bootscript ID
	Id *string `pulumi:"id"`
	// The initrd (initial ramdisk) configuration
	Initrd *string `pulumi:"initrd"`
	// The server kernel version
	Kernel *string `pulumi:"kernel"`
	// The bootscript organization ID
	Organization *string `pulumi:"organization"`
	// The bootscript project ID
	Project *string `pulumi:"project"`
	// Provide information if the bootscript is public
	Public *bool `pulumi:"public"`
	// The bootscript title
	Title *string `pulumi:"title"`
	// The zone in which is the bootscript
	Zone *string `pulumi:"zone"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1Bootscript
func (val *ScalewayInstanceV1Bootscript) Defaults() *ScalewayInstanceV1Bootscript {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Arch == nil {
		arch_ := ScalewayInstanceV1BootscriptArch("x86_64")
		tmp.Arch = &arch_
	}
	return &tmp
}

type ScalewayInstanceV1BootscriptOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1BootscriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1Bootscript)(nil)).Elem()
}

func (o ScalewayInstanceV1BootscriptOutput) ToScalewayInstanceV1BootscriptOutput() ScalewayInstanceV1BootscriptOutput {
	return o
}

func (o ScalewayInstanceV1BootscriptOutput) ToScalewayInstanceV1BootscriptOutputWithContext(ctx context.Context) ScalewayInstanceV1BootscriptOutput {
	return o
}

// The bootscript arch
func (o ScalewayInstanceV1BootscriptOutput) Arch() ScalewayInstanceV1BootscriptArchPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *ScalewayInstanceV1BootscriptArch { return v.Arch }).(ScalewayInstanceV1BootscriptArchPtrOutput)
}

// The bootscript arguments
func (o ScalewayInstanceV1BootscriptOutput) Bootcmdargs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Bootcmdargs }).(pulumi.StringPtrOutput)
}

// Dispmay if the bootscript is the default bootscript if no other boot option is configured
func (o ScalewayInstanceV1BootscriptOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *bool { return v.Default }).(pulumi.BoolPtrOutput)
}

// Provide information regarding a Device Tree Binary (dtb) for use with C1 servers
func (o ScalewayInstanceV1BootscriptOutput) Dtb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Dtb }).(pulumi.StringPtrOutput)
}

// The bootscript ID
func (o ScalewayInstanceV1BootscriptOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The initrd (initial ramdisk) configuration
func (o ScalewayInstanceV1BootscriptOutput) Initrd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Initrd }).(pulumi.StringPtrOutput)
}

// The server kernel version
func (o ScalewayInstanceV1BootscriptOutput) Kernel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Kernel }).(pulumi.StringPtrOutput)
}

// The bootscript organization ID
func (o ScalewayInstanceV1BootscriptOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// The bootscript project ID
func (o ScalewayInstanceV1BootscriptOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// Provide information if the bootscript is public
func (o ScalewayInstanceV1BootscriptOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

// The bootscript title
func (o ScalewayInstanceV1BootscriptOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The zone in which is the bootscript
func (o ScalewayInstanceV1BootscriptOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Bootscript) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1BootscriptPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1BootscriptPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1Bootscript)(nil)).Elem()
}

func (o ScalewayInstanceV1BootscriptPtrOutput) ToScalewayInstanceV1BootscriptPtrOutput() ScalewayInstanceV1BootscriptPtrOutput {
	return o
}

func (o ScalewayInstanceV1BootscriptPtrOutput) ToScalewayInstanceV1BootscriptPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1BootscriptPtrOutput {
	return o
}

func (o ScalewayInstanceV1BootscriptPtrOutput) Elem() ScalewayInstanceV1BootscriptOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) ScalewayInstanceV1Bootscript {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1Bootscript
		return ret
	}).(ScalewayInstanceV1BootscriptOutput)
}

// The bootscript arch
func (o ScalewayInstanceV1BootscriptPtrOutput) Arch() ScalewayInstanceV1BootscriptArchPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *ScalewayInstanceV1BootscriptArch {
		if v == nil {
			return nil
		}
		return v.Arch
	}).(ScalewayInstanceV1BootscriptArchPtrOutput)
}

// The bootscript arguments
func (o ScalewayInstanceV1BootscriptPtrOutput) Bootcmdargs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Bootcmdargs
	}).(pulumi.StringPtrOutput)
}

// Dispmay if the bootscript is the default bootscript if no other boot option is configured
func (o ScalewayInstanceV1BootscriptPtrOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *bool {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.BoolPtrOutput)
}

// Provide information regarding a Device Tree Binary (dtb) for use with C1 servers
func (o ScalewayInstanceV1BootscriptPtrOutput) Dtb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Dtb
	}).(pulumi.StringPtrOutput)
}

// The bootscript ID
func (o ScalewayInstanceV1BootscriptPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The initrd (initial ramdisk) configuration
func (o ScalewayInstanceV1BootscriptPtrOutput) Initrd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Initrd
	}).(pulumi.StringPtrOutput)
}

// The server kernel version
func (o ScalewayInstanceV1BootscriptPtrOutput) Kernel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Kernel
	}).(pulumi.StringPtrOutput)
}

// The bootscript organization ID
func (o ScalewayInstanceV1BootscriptPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

// The bootscript project ID
func (o ScalewayInstanceV1BootscriptPtrOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Project
	}).(pulumi.StringPtrOutput)
}

// Provide information if the bootscript is public
func (o ScalewayInstanceV1BootscriptPtrOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *bool {
		if v == nil {
			return nil
		}
		return v.Public
	}).(pulumi.BoolPtrOutput)
}

// The bootscript title
func (o ScalewayInstanceV1BootscriptPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

// The zone in which is the bootscript
func (o ScalewayInstanceV1BootscriptPtrOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Bootscript) *string {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1BootscriptArrayOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1BootscriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalewayInstanceV1Bootscript)(nil)).Elem()
}

func (o ScalewayInstanceV1BootscriptArrayOutput) ToScalewayInstanceV1BootscriptArrayOutput() ScalewayInstanceV1BootscriptArrayOutput {
	return o
}

func (o ScalewayInstanceV1BootscriptArrayOutput) ToScalewayInstanceV1BootscriptArrayOutputWithContext(ctx context.Context) ScalewayInstanceV1BootscriptArrayOutput {
	return o
}

func (o ScalewayInstanceV1BootscriptArrayOutput) Index(i pulumi.IntInput) ScalewayInstanceV1BootscriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalewayInstanceV1Bootscript {
		return vs[0].([]ScalewayInstanceV1Bootscript)[vs[1].(int)]
	}).(ScalewayInstanceV1BootscriptOutput)
}

type ScalewayInstanceV1GetBootscriptResponse struct {
	Bootscript *ScalewayInstanceV1Bootscript `pulumi:"bootscript"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1GetBootscriptResponse
func (val *ScalewayInstanceV1GetBootscriptResponse) Defaults() *ScalewayInstanceV1GetBootscriptResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Bootscript = tmp.Bootscript.Defaults()

	return &tmp
}

type ScalewayInstanceV1GetBootscriptResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetBootscriptResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetBootscriptResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetBootscriptResponseOutput) ToScalewayInstanceV1GetBootscriptResponseOutput() ScalewayInstanceV1GetBootscriptResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetBootscriptResponseOutput) ToScalewayInstanceV1GetBootscriptResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetBootscriptResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetBootscriptResponseOutput) Bootscript() ScalewayInstanceV1BootscriptPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetBootscriptResponse) *ScalewayInstanceV1Bootscript { return v.Bootscript }).(ScalewayInstanceV1BootscriptPtrOutput)
}

type ScalewayInstanceV1ListBootscriptsResponse struct {
	// List of bootscripts
	Bootscripts []ScalewayInstanceV1Bootscript `pulumi:"bootscripts"`
}

type ScalewayInstanceV1ListBootscriptsResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListBootscriptsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListBootscriptsResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1ListBootscriptsResponseOutput) ToScalewayInstanceV1ListBootscriptsResponseOutput() ScalewayInstanceV1ListBootscriptsResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListBootscriptsResponseOutput) ToScalewayInstanceV1ListBootscriptsResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1ListBootscriptsResponseOutput {
	return o
}

// List of bootscripts
func (o ScalewayInstanceV1ListBootscriptsResponseOutput) Bootscripts() ScalewayInstanceV1BootscriptArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ListBootscriptsResponse) []ScalewayInstanceV1Bootscript { return v.Bootscripts }).(ScalewayInstanceV1BootscriptArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1BootscriptOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1BootscriptPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1BootscriptArrayOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetBootscriptResponseOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1ListBootscriptsResponseOutput{})
}
