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
    'ScalewayStdFile',
    'AwaitableScalewayStdFile',
    'get_server_user_data',
    'get_server_user_data_output',
]

@pulumi.output_type
class ScalewayStdFile:
    def __init__(__self__, content=None, content_type=None, name=None):
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        pulumi.set(__self__, "content", content)
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        pulumi.set(__self__, "content_type", content_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> Optional[builtins.str]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")


class AwaitableScalewayStdFile(ScalewayStdFile):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayStdFile(
            content=self.content,
            content_type=self.content_type,
            name=self.name)


def get_server_user_data(key: Optional[builtins.str] = None,
                         server_id: Optional[builtins.str] = None,
                         zone: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayStdFile:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str key: Key of the user data to get
    :param builtins.str server_id: UUID of the server
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['key'] = key
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:user_data:getServerUserData', __args__, opts=opts, typ=ScalewayStdFile).value

    return AwaitableScalewayStdFile(
        content=pulumi.get(__ret__, 'content'),
        content_type=pulumi.get(__ret__, 'content_type'),
        name=pulumi.get(__ret__, 'name'))
def get_server_user_data_output(key: Optional[pulumi.Input[builtins.str]] = None,
                                server_id: Optional[pulumi.Input[builtins.str]] = None,
                                zone: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayStdFile]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str key: Key of the user data to get
    :param builtins.str server_id: UUID of the server
    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['key'] = key
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:user_data:getServerUserData', __args__, opts=opts, typ=ScalewayStdFile)
    return __ret__.apply(lambda __response__: ScalewayStdFile(
        content=pulumi.get(__response__, 'content'),
        content_type=pulumi.get(__response__, 'content_type'),
        name=pulumi.get(__response__, 'name')))
