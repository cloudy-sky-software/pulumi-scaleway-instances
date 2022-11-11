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
    'ListPrivateNICsResult',
    'AwaitableListPrivateNICsResult',
    'list_private_nics',
    'list_private_nics_output',
]

@pulumi.output_type
class ListPrivateNICsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1ListPrivateNICsResponse':
        return pulumi.get(self, "items")


class AwaitableListPrivateNICsResult(ListPrivateNICsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListPrivateNICsResult(
            items=self.items)


def list_private_nics(id: Optional[str] = None,
                      zone: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListPrivateNICsResult:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:private_nics:listPrivateNICs', __args__, opts=opts, typ=ListPrivateNICsResult).value

    return AwaitableListPrivateNICsResult(
        items=__ret__.items)


@_utilities.lift_output_func(list_private_nics)
def list_private_nics_output(id: Optional[pulumi.Input[str]] = None,
                             zone: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListPrivateNICsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    ...
