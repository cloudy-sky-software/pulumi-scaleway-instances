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
from ._enums import *

__all__ = [
    'ListServersResult',
    'AwaitableListServersResult',
    'list_servers',
    'list_servers_output',
]

@pulumi.output_type
class ListServersResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1ListServersResponse':
        return pulumi.get(self, "items")


class AwaitableListServersResult(ListServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListServersResult(
            items=self.items)


def list_servers(zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListServersResult:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:servers:listServers', __args__, opts=opts, typ=ListServersResult).value

    return AwaitableListServersResult(
        items=__ret__.items)


@_utilities.lift_output_func(list_servers)
def list_servers_output(zone: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListServersResult]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    ...
