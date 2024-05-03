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
    'ScalewayInstanceV1GetPlacementGroupResponse',
    'ScalewayInstanceV1ListPlacementGroupsResponse',
    'ScalewayInstanceV1PlacementGroup',
]

@pulumi.output_type
class ScalewayInstanceV1GetPlacementGroupResponse(dict):
    def __init__(__self__, *,
                 placement_group: Optional['outputs.ScalewayInstanceV1PlacementGroup'] = None):
        if placement_group is not None:
            pulumi.set(__self__, "placement_group", placement_group)

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> Optional['outputs.ScalewayInstanceV1PlacementGroup']:
        return pulumi.get(self, "placement_group")


@pulumi.output_type
class ScalewayInstanceV1ListPlacementGroupsResponse(dict):
    def __init__(__self__, *,
                 placement_groups: Optional[Sequence['outputs.ScalewayInstanceV1PlacementGroup']] = None):
        """
        :param Sequence['ScalewayInstanceV1PlacementGroup'] placement_groups: List of placement groups
        """
        if placement_groups is not None:
            pulumi.set(__self__, "placement_groups", placement_groups)

    @property
    @pulumi.getter(name="placementGroups")
    def placement_groups(self) -> Optional[Sequence['outputs.ScalewayInstanceV1PlacementGroup']]:
        """
        List of placement groups
        """
        return pulumi.get(self, "placement_groups")


@pulumi.output_type
class ScalewayInstanceV1PlacementGroup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyMode":
            suggest = "policy_mode"
        elif key == "policyRespected":
            suggest = "policy_respected"
        elif key == "policyType":
            suggest = "policy_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1PlacementGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1PlacementGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1PlacementGroup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 project: str,
                 id: Optional[str] = None,
                 organization: Optional[str] = None,
                 policy_mode: Optional['ScalewayInstanceV1PlacementGroupPolicyMode'] = None,
                 policy_respected: Optional[bool] = None,
                 policy_type: Optional['ScalewayInstanceV1PlacementGroupPolicyType'] = None,
                 tags: Optional[Sequence[str]] = None,
                 zone: Optional[str] = None):
        """
        :param str name: The placement group name
        :param str project: The placement group project ID
        :param str organization: The placement group organization ID
        :param bool policy_respected: Returns true if the policy is respected, false otherwise
        :param Sequence[str] tags: The placement group tags
        :param str zone: The zone in which is the placement group
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project", project)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if policy_mode is None:
            policy_mode = 'optional'
        if policy_mode is not None:
            pulumi.set(__self__, "policy_mode", policy_mode)
        if policy_respected is not None:
            pulumi.set(__self__, "policy_respected", policy_respected)
        if policy_type is None:
            policy_type = 'max_availability'
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The placement group name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The placement group project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        """
        The placement group organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="policyMode")
    def policy_mode(self) -> Optional['ScalewayInstanceV1PlacementGroupPolicyMode']:
        return pulumi.get(self, "policy_mode")

    @property
    @pulumi.getter(name="policyRespected")
    def policy_respected(self) -> Optional[bool]:
        """
        Returns true if the policy is respected, false otherwise
        """
        return pulumi.get(self, "policy_respected")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional['ScalewayInstanceV1PlacementGroupPolicyType']:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The placement group tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        The zone in which is the placement group
        """
        return pulumi.get(self, "zone")


