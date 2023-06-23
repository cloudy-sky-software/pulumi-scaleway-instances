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
    'ListServerUserDataResult',
    'AwaitableListServerUserDataResult',
    'list_server_user_data',
    'list_server_user_data_output',
]

@pulumi.output_type
class ListServerUserDataResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1ListServerUserDataResponse':
        return pulumi.get(self, "items")


class AwaitableListServerUserDataResult(ListServerUserDataResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListServerUserDataResult(
            items=self.items)


def list_server_user_data(server_id: Optional[str] = None,
                          zone: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListServerUserDataResult:
    """
    Use this data source to access information about an existing resource.

    :param str server_id: UUID of the server
    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['server_id'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:user_data:listServerUserData', __args__, opts=opts, typ=ListServerUserDataResult).value

    return AwaitableListServerUserDataResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_server_user_data)
def list_server_user_data_output(server_id: Optional[pulumi.Input[str]] = None,
                                 zone: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListServerUserDataResult]:
    """
    Use this data source to access information about an existing resource.

    :param str server_id: UUID of the server
    :param str zone: The zone you want to target
    """
    ...
