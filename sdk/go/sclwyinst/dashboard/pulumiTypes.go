// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dashboard

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1Dashboard struct {
	Images_count             *float64           `pulumi:"images_count"`
	Ips_count                *float64           `pulumi:"ips_count"`
	Ips_unused               *float64           `pulumi:"ips_unused"`
	Placement_groups_count   *float64           `pulumi:"placement_groups_count"`
	Private_nics_count       *float64           `pulumi:"private_nics_count"`
	Running_servers_count    *float64           `pulumi:"running_servers_count"`
	Security_groups_count    *float64           `pulumi:"security_groups_count"`
	Servers_by_types         map[string]float64 `pulumi:"servers_by_types"`
	Servers_count            *float64           `pulumi:"servers_count"`
	Snapshots_count          *float64           `pulumi:"snapshots_count"`
	Volumes_b_ssd_count      *float64           `pulumi:"volumes_b_ssd_count"`
	Volumes_b_ssd_total_size *float64           `pulumi:"volumes_b_ssd_total_size"`
	Volumes_count            *float64           `pulumi:"volumes_count"`
	Volumes_l_ssd_count      *float64           `pulumi:"volumes_l_ssd_count"`
	Volumes_l_ssd_total_size *float64           `pulumi:"volumes_l_ssd_total_size"`
}

type ScalewayInstanceV1DashboardOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1Dashboard)(nil)).Elem()
}

func (o ScalewayInstanceV1DashboardOutput) ToScalewayInstanceV1DashboardOutput() ScalewayInstanceV1DashboardOutput {
	return o
}

func (o ScalewayInstanceV1DashboardOutput) ToScalewayInstanceV1DashboardOutputWithContext(ctx context.Context) ScalewayInstanceV1DashboardOutput {
	return o
}

func (o ScalewayInstanceV1DashboardOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1Dashboard] {
	return pulumix.Output[ScalewayInstanceV1Dashboard]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1DashboardOutput) Images_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Images_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Ips_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Ips_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Ips_unused() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Ips_unused }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Placement_groups_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Placement_groups_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Private_nics_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Private_nics_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Running_servers_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Running_servers_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Security_groups_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Security_groups_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Servers_by_types() pulumi.Float64MapOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) map[string]float64 { return v.Servers_by_types }).(pulumi.Float64MapOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Servers_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Servers_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Snapshots_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Snapshots_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Volumes_b_ssd_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Volumes_b_ssd_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Volumes_b_ssd_total_size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Volumes_b_ssd_total_size }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Volumes_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Volumes_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Volumes_l_ssd_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Volumes_l_ssd_count }).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardOutput) Volumes_l_ssd_total_size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1Dashboard) *float64 { return v.Volumes_l_ssd_total_size }).(pulumi.Float64PtrOutput)
}

type ScalewayInstanceV1DashboardPtrOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1DashboardPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalewayInstanceV1Dashboard)(nil)).Elem()
}

func (o ScalewayInstanceV1DashboardPtrOutput) ToScalewayInstanceV1DashboardPtrOutput() ScalewayInstanceV1DashboardPtrOutput {
	return o
}

func (o ScalewayInstanceV1DashboardPtrOutput) ToScalewayInstanceV1DashboardPtrOutputWithContext(ctx context.Context) ScalewayInstanceV1DashboardPtrOutput {
	return o
}

func (o ScalewayInstanceV1DashboardPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalewayInstanceV1Dashboard] {
	return pulumix.Output[*ScalewayInstanceV1Dashboard]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1DashboardPtrOutput) Elem() ScalewayInstanceV1DashboardOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) ScalewayInstanceV1Dashboard {
		if v != nil {
			return *v
		}
		var ret ScalewayInstanceV1Dashboard
		return ret
	}).(ScalewayInstanceV1DashboardOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Images_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Images_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Ips_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Ips_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Ips_unused() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Ips_unused
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Placement_groups_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Placement_groups_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Private_nics_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Private_nics_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Running_servers_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Running_servers_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Security_groups_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Security_groups_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Servers_by_types() pulumi.Float64MapOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) map[string]float64 {
		if v == nil {
			return nil
		}
		return v.Servers_by_types
	}).(pulumi.Float64MapOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Servers_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Servers_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Snapshots_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Snapshots_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Volumes_b_ssd_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Volumes_b_ssd_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Volumes_b_ssd_total_size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Volumes_b_ssd_total_size
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Volumes_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Volumes_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Volumes_l_ssd_count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Volumes_l_ssd_count
	}).(pulumi.Float64PtrOutput)
}

func (o ScalewayInstanceV1DashboardPtrOutput) Volumes_l_ssd_total_size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalewayInstanceV1Dashboard) *float64 {
		if v == nil {
			return nil
		}
		return v.Volumes_l_ssd_total_size
	}).(pulumi.Float64PtrOutput)
}

type ScalewayInstanceV1GetDashboardResponse struct {
	Dashboard *ScalewayInstanceV1Dashboard `pulumi:"dashboard"`
}

type ScalewayInstanceV1GetDashboardResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1GetDashboardResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1GetDashboardResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1GetDashboardResponseOutput) ToScalewayInstanceV1GetDashboardResponseOutput() ScalewayInstanceV1GetDashboardResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetDashboardResponseOutput) ToScalewayInstanceV1GetDashboardResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1GetDashboardResponseOutput {
	return o
}

func (o ScalewayInstanceV1GetDashboardResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalewayInstanceV1GetDashboardResponse] {
	return pulumix.Output[ScalewayInstanceV1GetDashboardResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScalewayInstanceV1GetDashboardResponseOutput) Dashboard() ScalewayInstanceV1DashboardPtrOutput {
	return o.ApplyT(func(v ScalewayInstanceV1GetDashboardResponse) *ScalewayInstanceV1Dashboard { return v.Dashboard }).(ScalewayInstanceV1DashboardPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1DashboardOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1DashboardPtrOutput{})
	pulumi.RegisterOutputType(ScalewayInstanceV1GetDashboardResponseOutput{})
}
