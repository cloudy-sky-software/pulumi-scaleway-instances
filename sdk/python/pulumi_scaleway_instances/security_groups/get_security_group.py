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
    'GetSecurityGroupResult',
    'AwaitableGetSecurityGroupResult',
    'get_security_group',
    'get_security_group_output',
]

@pulumi.output_type
class GetSecurityGroupResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1GetSecurityGroupResponse':
        return pulumi.get(self, "items")


class AwaitableGetSecurityGroupResult(GetSecurityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupResult(
            items=self.items)


def get_security_group(security_group_id: Optional[str] = None,
                       zone: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupResult:
    """
    Use this data source to access information about an existing resource.

    :param str security_group_id: UUID of the security group you want to get
    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['securityGroupId'] = security_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:security_groups:getSecurityGroup', __args__, opts=opts, typ=GetSecurityGroupResult).value

    return AwaitableGetSecurityGroupResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_security_group)
def get_security_group_output(security_group_id: Optional[pulumi.Input[str]] = None,
                              zone: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityGroupResult]:
    """
    Use this data source to access information about an existing resource.

    :param str security_group_id: UUID of the security group you want to get
    :param str zone: The zone you want to target
    """
    ...
