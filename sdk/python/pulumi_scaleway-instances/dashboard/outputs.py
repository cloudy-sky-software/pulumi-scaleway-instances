# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ScalewayInstanceV1Dashboard',
    'ScalewayInstanceV1GetDashboardResponse',
]

@pulumi.output_type
class ScalewayInstanceV1Dashboard(dict):
    def __init__(__self__, *,
                 images_count: Optional[float] = None,
                 ips_count: Optional[float] = None,
                 ips_unused: Optional[float] = None,
                 placement_groups_count: Optional[float] = None,
                 private_nics_count: Optional[float] = None,
                 running_servers_count: Optional[float] = None,
                 security_groups_count: Optional[float] = None,
                 servers_by_types: Optional[Mapping[str, float]] = None,
                 servers_count: Optional[float] = None,
                 snapshots_count: Optional[float] = None,
                 volumes_b_ssd_count: Optional[float] = None,
                 volumes_b_ssd_total_size: Optional[float] = None,
                 volumes_count: Optional[float] = None,
                 volumes_l_ssd_count: Optional[float] = None,
                 volumes_l_ssd_total_size: Optional[float] = None):
        if images_count is not None:
            pulumi.set(__self__, "images_count", images_count)
        if ips_count is not None:
            pulumi.set(__self__, "ips_count", ips_count)
        if ips_unused is not None:
            pulumi.set(__self__, "ips_unused", ips_unused)
        if placement_groups_count is not None:
            pulumi.set(__self__, "placement_groups_count", placement_groups_count)
        if private_nics_count is not None:
            pulumi.set(__self__, "private_nics_count", private_nics_count)
        if running_servers_count is not None:
            pulumi.set(__self__, "running_servers_count", running_servers_count)
        if security_groups_count is not None:
            pulumi.set(__self__, "security_groups_count", security_groups_count)
        if servers_by_types is not None:
            pulumi.set(__self__, "servers_by_types", servers_by_types)
        if servers_count is not None:
            pulumi.set(__self__, "servers_count", servers_count)
        if snapshots_count is not None:
            pulumi.set(__self__, "snapshots_count", snapshots_count)
        if volumes_b_ssd_count is not None:
            pulumi.set(__self__, "volumes_b_ssd_count", volumes_b_ssd_count)
        if volumes_b_ssd_total_size is not None:
            pulumi.set(__self__, "volumes_b_ssd_total_size", volumes_b_ssd_total_size)
        if volumes_count is not None:
            pulumi.set(__self__, "volumes_count", volumes_count)
        if volumes_l_ssd_count is not None:
            pulumi.set(__self__, "volumes_l_ssd_count", volumes_l_ssd_count)
        if volumes_l_ssd_total_size is not None:
            pulumi.set(__self__, "volumes_l_ssd_total_size", volumes_l_ssd_total_size)

    @property
    @pulumi.getter
    def images_count(self) -> Optional[float]:
        return pulumi.get(self, "images_count")

    @property
    @pulumi.getter
    def ips_count(self) -> Optional[float]:
        return pulumi.get(self, "ips_count")

    @property
    @pulumi.getter
    def ips_unused(self) -> Optional[float]:
        return pulumi.get(self, "ips_unused")

    @property
    @pulumi.getter
    def placement_groups_count(self) -> Optional[float]:
        return pulumi.get(self, "placement_groups_count")

    @property
    @pulumi.getter
    def private_nics_count(self) -> Optional[float]:
        return pulumi.get(self, "private_nics_count")

    @property
    @pulumi.getter
    def running_servers_count(self) -> Optional[float]:
        return pulumi.get(self, "running_servers_count")

    @property
    @pulumi.getter
    def security_groups_count(self) -> Optional[float]:
        return pulumi.get(self, "security_groups_count")

    @property
    @pulumi.getter
    def servers_by_types(self) -> Optional[Mapping[str, float]]:
        return pulumi.get(self, "servers_by_types")

    @property
    @pulumi.getter
    def servers_count(self) -> Optional[float]:
        return pulumi.get(self, "servers_count")

    @property
    @pulumi.getter
    def snapshots_count(self) -> Optional[float]:
        return pulumi.get(self, "snapshots_count")

    @property
    @pulumi.getter
    def volumes_b_ssd_count(self) -> Optional[float]:
        return pulumi.get(self, "volumes_b_ssd_count")

    @property
    @pulumi.getter
    def volumes_b_ssd_total_size(self) -> Optional[float]:
        return pulumi.get(self, "volumes_b_ssd_total_size")

    @property
    @pulumi.getter
    def volumes_count(self) -> Optional[float]:
        return pulumi.get(self, "volumes_count")

    @property
    @pulumi.getter
    def volumes_l_ssd_count(self) -> Optional[float]:
        return pulumi.get(self, "volumes_l_ssd_count")

    @property
    @pulumi.getter
    def volumes_l_ssd_total_size(self) -> Optional[float]:
        return pulumi.get(self, "volumes_l_ssd_total_size")


@pulumi.output_type
class ScalewayInstanceV1GetDashboardResponse(dict):
    def __init__(__self__, *,
                 dashboard: Optional['outputs.ScalewayInstanceV1Dashboard'] = None):
        if dashboard is not None:
            pulumi.set(__self__, "dashboard", dashboard)

    @property
    @pulumi.getter
    def dashboard(self) -> Optional['outputs.ScalewayInstanceV1Dashboard']:
        return pulumi.get(self, "dashboard")


