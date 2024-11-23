# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *

__all__ = [
    'ScalewayInstanceV1ListServersResponse',
    'AwaitableScalewayInstanceV1ListServersResponse',
    'list_servers',
    'list_servers_output',
]

@pulumi.output_type
class ScalewayInstanceV1ListServersResponse:
    def __init__(__self__, servers=None):
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter
    def servers(self) -> Optional[Sequence['outputs.ScalewayInstanceV1Server']]:
        """
        List of servers
        """
        return pulumi.get(self, "servers")


class AwaitableScalewayInstanceV1ListServersResponse(ScalewayInstanceV1ListServersResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1ListServersResponse(
            servers=self.servers)


def list_servers(zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1ListServersResponse:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:servers:listServers', __args__, opts=opts, typ=ScalewayInstanceV1ListServersResponse).value

    return AwaitableScalewayInstanceV1ListServersResponse(
        servers=pulumi.get(__ret__, 'servers'))
def list_servers_output(zone: Optional[pulumi.Input[str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1ListServersResponse]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:servers:listServers', __args__, opts=opts, typ=ScalewayInstanceV1ListServersResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1ListServersResponse(
        servers=pulumi.get(__response__, 'servers')))
