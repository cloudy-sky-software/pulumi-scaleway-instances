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
    'ScalewayInstanceV1GetImageResponse',
    'AwaitableScalewayInstanceV1GetImageResponse',
    'get_image',
    'get_image_output',
]

@pulumi.output_type
class ScalewayInstanceV1GetImageResponse:
    def __init__(__self__, image=None):
        if image and not isinstance(image, dict):
            raise TypeError("Expected argument 'image' to be a dict")
        pulumi.set(__self__, "image", image)

    @property
    @pulumi.getter
    def image(self) -> Optional['outputs.ScalewayInstanceV1Image']:
        return pulumi.get(self, "image")


class AwaitableScalewayInstanceV1GetImageResponse(ScalewayInstanceV1GetImageResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1GetImageResponse(
            image=self.image)


def get_image(id: Optional[builtins.str] = None,
              zone: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1GetImageResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str id: UUID of the image you want to get
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:images:getImage', __args__, opts=opts, typ=ScalewayInstanceV1GetImageResponse).value

    return AwaitableScalewayInstanceV1GetImageResponse(
        image=pulumi.get(__ret__, 'image'))
def get_image_output(id: Optional[pulumi.Input[builtins.str]] = None,
                     zone: Optional[pulumi.Input[builtins.str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1GetImageResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str id: UUID of the image you want to get
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:images:getImage', __args__, opts=opts, typ=ScalewayInstanceV1GetImageResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1GetImageResponse(
        image=pulumi.get(__response__, 'image')))
