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
    'ScalewayInstanceV1GetSecurityGroupRuleResponse',
    'AwaitableScalewayInstanceV1GetSecurityGroupRuleResponse',
    'get_security_group_rule',
    'get_security_group_rule_output',
]

@pulumi.output_type
class ScalewayInstanceV1GetSecurityGroupRuleResponse:
    def __init__(__self__, rule=None):
        if rule and not isinstance(rule, dict):
            raise TypeError("Expected argument 'rule' to be a dict")
        pulumi.set(__self__, "rule", rule)

    @property
    @pulumi.getter
    def rule(self) -> Optional['outputs.ScalewayInstanceV1SecurityGroupRule']:
        return pulumi.get(self, "rule")


class AwaitableScalewayInstanceV1GetSecurityGroupRuleResponse(ScalewayInstanceV1GetSecurityGroupRuleResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalewayInstanceV1GetSecurityGroupRuleResponse(
            rule=self.rule)


def get_security_group_rule(id: Optional[builtins.str] = None,
                            security_group_id: Optional[builtins.str] = None,
                            zone: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalewayInstanceV1GetSecurityGroupRuleResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['securityGroupId'] = security_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:rules:getSecurityGroupRule', __args__, opts=opts, typ=ScalewayInstanceV1GetSecurityGroupRuleResponse).value

    return AwaitableScalewayInstanceV1GetSecurityGroupRuleResponse(
        rule=pulumi.get(__ret__, 'rule'))
def get_security_group_rule_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                   security_group_id: Optional[pulumi.Input[builtins.str]] = None,
                                   zone: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ScalewayInstanceV1GetSecurityGroupRuleResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['securityGroupId'] = security_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway-instances:rules:getSecurityGroupRule', __args__, opts=opts, typ=ScalewayInstanceV1GetSecurityGroupRuleResponse)
    return __ret__.apply(lambda __response__: ScalewayInstanceV1GetSecurityGroupRuleResponse(
        rule=pulumi.get(__response__, 'rule')))
