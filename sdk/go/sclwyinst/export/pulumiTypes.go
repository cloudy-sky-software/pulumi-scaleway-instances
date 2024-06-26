// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package export

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1Task struct {
	// The description of the task
	Description *string `pulumi:"description"`
	HrefFrom    *string `pulumi:"hrefFrom"`
	HrefResult  *string `pulumi:"hrefResult"`
	// The unique ID of the task
	Id *string `pulumi:"id"`
	// The progress of the task in percent
	Progress *float64 `pulumi:"progress"`
	// The task start date (RFC 3339 format)
	StartedAt *string `pulumi:"startedAt"`
	// The task status
	Status *ScalewayInstanceV1TaskStatus `pulumi:"status"`
	// The task end date (RFC 3339 format)
	TerminatedAt *string `pulumi:"terminatedAt"`
	// The zone in which is the task
	Zone *string `pulumi:"zone"`
}

// Defaults sets the appropriate defaults for ScalewayInstanceV1Task
func (val *ScalewayInstanceV1Task) Defaults() *ScalewayInstanceV1Task {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Status == nil {
		status_ := ScalewayInstanceV1TaskStatus("pending")
		tmp.Status = &status_
	}
	return &tmp
}

type ScalewayInstanceV1TaskOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1Task)(nil)).Elem()
}

func (o ScalewayInstanceV1TaskOutput) ToScalewayInstanceV1TaskOutput() ScalewayInstanceV1TaskOutput {
	return o
}

func (o ScalewayInstanceV1TaskOutput) ToScalewayInstanceV1TaskOutputWithContext(ctx context.Context) ScalewayInstanceV1TaskOutput {
	return o
}

// The description of the task
func (o ScalewayInstanceV1TaskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1TaskOutput) HrefFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.HrefFrom }).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1TaskOutput) HrefResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.HrefResult }).(pulumi.StringPtrOutput)
}

// The unique ID of the task
func (o ScalewayInstanceV1TaskOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The progress of the task in percent
func (o ScalewayInstanceV1TaskOutput) Progress() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *float64 { return v.Progress }).(pulumi.Float64PtrOutput)
}

// The task start date (RFC 3339 format)
func (o ScalewayInstanceV1TaskOutput) StartedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.StartedAt }).(pulumi.StringPtrOutput)
}

// The task status
func (o ScalewayInstanceV1TaskOutput) Status() ScalewayInstanceV1TaskStatusPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *ScalewayInstanceV1TaskStatus { return v.Status }).(ScalewayInstanceV1TaskStatusPtrOutput)
}

// The task end date (RFC 3339 format)
func (o ScalewayInstanceV1TaskOutput) TerminatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.TerminatedAt }).(pulumi.StringPtrOutput)
}

// The zone in which is the task
func (o ScalewayInstanceV1TaskOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Task) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ScalewayInstanceV1TaskPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1TaskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1Task)(nil)).Elem()
}

func (o ScalewayInstanceV1TaskPtrOutput) ToScalewayInstanceV1TaskPtrOutput() ScalewayInstanceV1TaskPtrOutput {
	return o
}

func (o ScalewayInstanceV1TaskPtrOutput) ToScalewayInstanceV1TaskPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1TaskPtrOutput {
	return o
}

func (o ScalewayInstanceV1TaskPtrOutput) Elem() ScalewayInstanceV1TaskOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) ScalewayInstanceV1Task {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1Task
		return ret
	}).(ScalewayInstanceV1TaskOutput)
}

// The description of the task
func (o ScalewayInstanceV1TaskPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1TaskPtrOutput) HrefFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.HrefFrom
	}).(pulumi.StringPtrOutput)
}

func (o ScalewayInstanceV1TaskPtrOutput) HrefResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.HrefResult
	}).(pulumi.StringPtrOutput)
}

// The unique ID of the task
func (o ScalewayInstanceV1TaskPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The progress of the task in percent
func (o ScalewayInstanceV1TaskPtrOutput) Progress() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *float64 {
		if v == nil {
			return nil
		}
		return v.Progress
	}).(pulumi.Float64PtrOutput)
}

// The task start date (RFC 3339 format)
func (o ScalewayInstanceV1TaskPtrOutput) StartedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.StartedAt
	}).(pulumi.StringPtrOutput)
}

// The task status
func (o ScalewayInstanceV1TaskPtrOutput) Status() ScalewayInstanceV1TaskStatusPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *ScalewayInstanceV1TaskStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(ScalewayInstanceV1TaskStatusPtrOutput)
}

// The task end date (RFC 3339 format)
func (o ScalewayInstanceV1TaskPtrOutput) TerminatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.TerminatedAt
	}).(pulumi.StringPtrOutput)
}

// The zone in which is the task
func (o ScalewayInstanceV1TaskPtrOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Task) *string {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1TaskOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1TaskPtrOutput{})
}
