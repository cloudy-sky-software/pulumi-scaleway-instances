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

__all__ = [
    'ScalewayInstanceV1ListServerUserDataResponse',
    'AwaitableScalewayInstanceV1ListServerUserDataResponse',
    'list_server_user_data',
    'list_server_user_data_output',
]

@pulumi.output_type
class ScalewayInstanceV1ListServerUserDataResponse:
    def __init__(__self__, user_data=None):
        if user_data and not isinstance(user_data, list):
            raise TypeError("Expected argument 'user_data' to be a list")
        pulumi.set(__self__, "user_data", user_data)

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> Optional[Sequence[builtins.str]]:
        return pulumi.get(self, "user_data")


class AwaitableScalewayInstanceV1ListServerUserDataResponse(ScalewayInstanceV1ListServerUserDataResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1ListServerUserDataResponse(
            user_data=self.user_data)


def list_server_user_data(server_id: Optional[builtins.str] = None,
                          zone: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1ListServerUserDataResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str server_id: UUID of the server
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:user_data:listServerUserData', __args__, opts=opts, typ=ScalewayInstanceV1ListServerUserDataResponse).value

    return AwaitableScalewayInstanceV1ListServerUserDataResponse(
        user_data=pulumi.get(__ret__, 'user_data'))
def list_server_user_data_output(server_id: Optional[pulumi.Input[builtins.str]] = None,
                                 zone: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1ListServerUserDataResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str server_id: UUID of the server
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:user_data:listServerUserData', __args__, opts=opts, typ=ScalewayInstanceV1ListServerUserDataResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1ListServerUserDataResponse(
        user_data=pulumi.get(__response__, 'user_data')))
