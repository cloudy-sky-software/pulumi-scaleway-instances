# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from .. import _utilities
from . import outputs

__all__ = [
    'ScalewayInstanceV1GetDashboardResponse',
    'AwaitableScalewayInstanceV1GetDashboardResponse',
    'get_dashboard',
    'get_dashboard_output',
]

@pulumi.output_type
class ScalewayInstanceV1GetDashboardResponse:
    def __init__(__self__, dashboard=None):
        if dashboard and not isinstance(dashboard, dict):
            raise TypeError("Expected argument 'dashboard' to be a dict")
        pulumi.set(__self__, "dashboard", dashboard)

    @property
    @pulumi.getter
    def dashboard(self) -> Optional['outputs.ScalewayInstanceV1Dashboard']:
        return pulumi.get(self, "dashboard")


class AwaitableScalewayInstanceV1GetDashboardResponse(ScalewayInstanceV1GetDashboardResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1GetDashboardResponse(
            dashboard=self.dashboard)


def get_dashboard(zone: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1GetDashboardResponse:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:dashboard:getDashboard', __args__, opts=opts, typ=ScalewayInstanceV1GetDashboardResponse).value

    return AwaitableScalewayInstanceV1GetDashboardResponse(
        dashboard=pulumi.get(__ret__, 'dashboard'))


@_utilities.lift_output_func(get_dashboard)
def get_dashboard_output(zone: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ScalewayInstanceV1GetDashboardResponse]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    ...
