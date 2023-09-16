# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetPlacementGroupServersResult',
    'AwaitableGetPlacementGroupServersResult',
    'get_placement_group_servers',
    'get_placement_group_servers_output',
]

@pulumi.output_type
class GetPlacementGroupServersResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1GetPlacementGroupServersResponse':
        return pulumi.get(self, "items")


class AwaitableGetPlacementGroupServersResult(GetPlacementGroupServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlacementGroupServersResult(
            items=self.items)


def get_placement_group_servers(placement_group_id: Optional[str] = None,
                                zone: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlacementGroupServersResult:
    """
    Use this data source to access information about an existing resource.

    :param str placement_group_id: UUID of the placement group
    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['placement_group_id'] = placement_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:servers:getPlacementGroupServers', __args__, opts=opts, typ=GetPlacementGroupServersResult).value

    return AwaitableGetPlacementGroupServersResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_placement_group_servers)
def get_placement_group_servers_output(placement_group_id: Optional[pulumi.Input[str]] = None,
                                       zone: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlacementGroupServersResult]:
    """
    Use this data source to access information about an existing resource.

    :param str placement_group_id: UUID of the placement group
    :param str zone: The zone you want to target
    """
    ...
