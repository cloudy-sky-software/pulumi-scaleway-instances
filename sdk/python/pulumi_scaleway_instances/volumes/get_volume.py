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
    'ScalewayInstanceV1GetVolumeResponse',
    'AwaitableScalewayInstanceV1GetVolumeResponse',
    'get_volume',
    'get_volume_output',
]

@pulumi.output_type
class ScalewayInstanceV1GetVolumeResponse:
    def __init__(__self__, volume=None):
        if volume and not isinstance(volume, dict):
            raise TypeError("Expected argument 'volume' to be a dict")
        pulumi.set(__self__, "volume", volume)

    @property
    @pulumi.getter
    def volume(self) -> Optional['outputs.ScalewayInstanceV1Volume']:
        return pulumi.get(self, "volume")


class AwaitableScalewayInstanceV1GetVolumeResponse(ScalewayInstanceV1GetVolumeResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1GetVolumeResponse(
            volume=self.volume)


def get_volume(id: Optional[builtins.str] = None,
               zone: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1GetVolumeResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str id: UUID of the volume you want to get
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:volumes:getVolume', __args__, opts=opts, typ=ScalewayInstanceV1GetVolumeResponse).value

    return AwaitableScalewayInstanceV1GetVolumeResponse(
        volume=pulumi.get(__ret__, 'volume'))
def get_volume_output(id: Optional[pulumi.Input[builtins.str]] = None,
                      zone: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1GetVolumeResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str id: UUID of the volume you want to get
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:volumes:getVolume', __args__, opts=opts, typ=ScalewayInstanceV1GetVolumeResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1GetVolumeResponse(
        volume=pulumi.get(__response__, 'volume')))
