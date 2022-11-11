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
    'ListSecurityGroupsResult',
    'AwaitableListSecurityGroupsResult',
    'list_security_groups',
    'list_security_groups_output',
]

@pulumi.output_type
class ListSecurityGroupsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1ListSecurityGroupsResponse':
        return pulumi.get(self, "items")


class AwaitableListSecurityGroupsResult(ListSecurityGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSecurityGroupsResult(
            items=self.items)


def list_security_groups(zone: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListSecurityGroupsResult:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:security_groups:listSecurityGroups', __args__, opts=opts, typ=ListSecurityGroupsResult).value

    return AwaitableListSecurityGroupsResult(
        items=__ret__.items)


@_utilities.lift_output_func(list_security_groups)
def list_security_groups_output(zone: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListSecurityGroupsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    ...