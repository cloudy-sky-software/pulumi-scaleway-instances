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
from ._enums import *

__all__ = [
    'ScalewayInstanceV1ListPrivateNICsResponse',
    'AwaitableScalewayInstanceV1ListPrivateNICsResponse',
    'list_private_nics',
    'list_private_nics_output',
]

@pulumi.output_type
class ScalewayInstanceV1ListPrivateNICsResponse:
    def __init__(__self__, private_nics=None):
        if private_nics and not isinstance(private_nics, list):
            raise TypeError("Expected argument 'private_nics' to be a list")
        pulumi.set(__self__, "private_nics", private_nics)

    @property
    @pulumi.getter(name="privateNics")
    def private_nics(self) -> Optional[Sequence['outputs.ScalewayInstanceV1PrivateNIC']]:
        return pulumi.get(self, "private_nics")


class AwaitableScalewayInstanceV1ListPrivateNICsResponse(ScalewayInstanceV1ListPrivateNICsResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1ListPrivateNICsResponse(
            private_nics=self.private_nics)


def list_private_nics(server_id: Optional[builtins.str] = None,
                      zone: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1ListPrivateNICsResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:private_nics:listPrivateNICs', __args__, opts=opts, typ=ScalewayInstanceV1ListPrivateNICsResponse).value

    return AwaitableScalewayInstanceV1ListPrivateNICsResponse(
        private_nics=pulumi.get(__ret__, 'private_nics'))
def list_private_nics_output(server_id: Optional[pulumi.Input[builtins.str]] = None,
                             zone: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1ListPrivateNICsResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:private_nics:listPrivateNICs', __args__, opts=opts, typ=ScalewayInstanceV1ListPrivateNICsResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1ListPrivateNICsResponse(
        private_nics=pulumi.get(__response__, 'private_nics')))
