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
    'ListPlacementGroupsResult',
    'AwaitableListPlacementGroupsResult',
    'list_placement_groups',
    'list_placement_groups_output',
]

@pulumi.output_type
class ListPlacementGroupsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ScalewayInstanceV1ListPlacementGroupsResponse':
        return pulumi.get(self, "items")


class AwaitableListPlacementGroupsResult(ListPlacementGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListPlacementGroupsResult(
            items=self.items)


def list_placement_groups(zone: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListPlacementGroupsResult:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway-instances:placement_groups:listPlacementGroups', __args__, opts=opts, typ=ListPlacementGroupsResult).value

    return AwaitableListPlacementGroupsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_placement_groups)
def list_placement_groups_output(zone: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListPlacementGroupsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The zone you want to target
    """
    ...
