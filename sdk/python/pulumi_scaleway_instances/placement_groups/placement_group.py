# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = ['PlacementGroupArgs', 'PlacementGroup']

@pulumi.input_type
class PlacementGroupArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 policy_mode: Optional[pulumi.Input['PolicyMode']] = None,
                 policy_type: Optional[pulumi.Input['PolicyType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PlacementGroup resource.
        :param pulumi.Input[str] project: The placement group project ID
        :param pulumi.Input[str] name: The placement group name
        :param pulumi.Input[str] organization: The placement group organization ID
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The placement group tags
        :param pulumi.Input[str] zone: The zone you want to target
        """
        pulumi.set(__self__, "project", project)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if policy_mode is None:
            policy_mode = 'optional'
        if policy_mode is not None:
            pulumi.set(__self__, "policy_mode", policy_mode)
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
    def project(self) -> pulumi.Input[str]:
        """
        The placement group project ID
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The placement group name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        """
        The placement group organization ID
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter(name="policyMode")
    def policy_mode(self) -> Optional[pulumi.Input['PolicyMode']]:
        return pulumi.get(self, "policy_mode")

    @policy_mode.setter
    def policy_mode(self, value: Optional[pulumi.Input['PolicyMode']]):
        pulumi.set(self, "policy_mode", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input['PolicyType']]:
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input['PolicyType']]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The placement group tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to target
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class PlacementGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 policy_mode: Optional[pulumi.Input['PolicyMode']] = None,
                 policy_type: Optional[pulumi.Input['PolicyType']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PlacementGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The placement group name
        :param pulumi.Input[str] organization: The placement group organization ID
        :param pulumi.Input[str] project: The placement group project ID
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The placement group tags
        :param pulumi.Input[str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PlacementGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PlacementGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PlacementGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PlacementGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 policy_mode: Optional[pulumi.Input['PolicyMode']] = None,
                 policy_type: Optional[pulumi.Input['PolicyType']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PlacementGroupArgs.__new__(PlacementGroupArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["organization"] = organization
            if policy_mode is None:
                policy_mode = 'optional'
            __props__.__dict__["policy_mode"] = policy_mode
            if policy_type is None:
                policy_type = 'max_availability'
            __props__.__dict__["policy_type"] = policy_type
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["placement_group"] = None
            __props__.__dict__["policy_respected"] = None
        super(PlacementGroup, __self__).__init__(
            'scaleway-instances:placement_groups:PlacementGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PlacementGroup':
        """
        Get an existing PlacementGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PlacementGroupArgs.__new__(PlacementGroupArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["organization"] = None
        __props__.__dict__["placement_group"] = None
        __props__.__dict__["policy_mode"] = None
        __props__.__dict__["policy_respected"] = None
        __props__.__dict__["policy_type"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["zone"] = None
        return PlacementGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The placement group name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[str]]:
        """
        The placement group organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1PlacementGroup']]:
        return pulumi.get(self, "placement_group")

    @property
    @pulumi.getter(name="policyMode")
    def policy_mode(self) -> pulumi.Output[Optional['PolicyMode']]:
        return pulumi.get(self, "policy_mode")

    @property
    @pulumi.getter(name="policyRespected")
    def policy_respected(self) -> pulumi.Output[Optional[bool]]:
        """
        Returns true if the policy is respected, false otherwise
        """
        return pulumi.get(self, "policy_respected")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[Optional['PolicyType']]:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The placement group project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The placement group tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        """
        The zone in which is the placement group
        """
        return pulumi.get(self, "zone")

