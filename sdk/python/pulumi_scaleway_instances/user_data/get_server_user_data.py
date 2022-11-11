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
    'GetServerUserDataResult',
    'AwaitableGetServerUserDataResult',
    'get_server_user_data',
    'get_server_user_data_output',
]

@pulumi.output_type
class GetServerUserDataResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayStdFile':
        return pulumi.get(self, "items")


class AwaitableGetServerUserDataResult(GetServerUserDataResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerUserDataResult(
            items=self.items)


def get_server_user_data(id: Optional[str] = None,
                         key: Optional[str] = None,
                         zone: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerUserDataResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: UUID of the server
    :param str key: Key of the user data to get
    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['key'] = key
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:user_data:getServerUserData', __args__, opts=opts, typ=GetServerUserDataResult).value

    return AwaitableGetServerUserDataResult(
        items=__ret__.items)


@_utilities.lift_output_func(get_server_user_data)
def get_server_user_data_output(id: Optional[pulumi.Input[str]] = None,
                                key: Optional[pulumi.Input[str]] = None,
                                zone: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerUserDataResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: UUID of the server
    :param str key: Key of the user data to get
    :param str zone: The zone you want to target
    """
    ...