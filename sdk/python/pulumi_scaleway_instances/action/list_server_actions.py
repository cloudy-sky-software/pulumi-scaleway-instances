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
    'ListServerActionsResult',
    'AwaitableListServerActionsResult',
    'list_server_actions',
    'list_server_actions_output',
]

@pulumi.output_type
class ListServerActionsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1ListServerActionsResponse':
        return pulumi.get(self, "items")


class AwaitableListServerActionsResult(ListServerActionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListServerActionsResult(
            items=self.items)


def list_server_actions(id: Optional[str] = None,
                        zone: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListServerActionsResult:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:action:listServerActions', __args__, opts=opts, typ=ListServerActionsResult).value

    return AwaitableListServerActionsResult(
        items=__ret__.items)


@_utilities.lift_output_func(list_server_actions)
def list_server_actions_output(id: Optional[pulumi.Input[str]] = None,
                               zone: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListServerActionsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    ...