# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'ScalewayInstanceV1GetPlacementGroupServersResponse',
    'AwaitableScalewayInstanceV1GetPlacementGroupServersResponse',
    'get_placement_group_servers',
    'get_placement_group_servers_output',
]

@pulumi.output_type
class ScalewayInstanceV1GetPlacementGroupServersResponse:
    def __init__(__self__, servers=None):
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter
    def servers(self) -> Optional[Sequence['outputs.ScalewayInstanceV1PlacementGroupServer']]:
        return pulumi.get(self, "servers")


class AwaitableScalewayInstanceV1GetPlacementGroupServersResponse(ScalewayInstanceV1GetPlacementGroupServersResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1GetPlacementGroupServersResponse(
            servers=self.servers)


def get_placement_group_servers(placement_group_id: Optional[builtins.str] = None,
                                zone: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1GetPlacementGroupServersResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str placement_group_id: UUID of the placement group
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['placementGroupId'] = placement_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:servers:getPlacementGroupServers', __args__, opts=opts, typ=ScalewayInstanceV1GetPlacementGroupServersResponse).value

    return AwaitableScalewayInstanceV1GetPlacementGroupServersResponse(
        servers=pulumi.get(__ret__, 'servers'))
def get_placement_group_servers_output(placement_group_id: Optional[pulumi.Input[builtins.str]] = None,
                                       zone: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1GetPlacementGroupServersResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str placement_group_id: UUID of the placement group
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['placementGroupId'] = placement_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:servers:getPlacementGroupServers', __args__, opts=opts, typ=ScalewayInstanceV1GetPlacementGroupServersResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1GetPlacementGroupServersResponse(
        servers=pulumi.get(__response__, 'servers')))
