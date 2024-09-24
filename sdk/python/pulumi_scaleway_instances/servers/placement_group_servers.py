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

__all__ = ['PlacementGroupServersArgs', 'PlacementGroupServers']

@pulumi.input_type
class PlacementGroupServersArgs:
    def __init__(__self__, *,
                 placement_group_id: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PlacementGroupServers resource.
        :param pulumi.Input[str] placement_group_id: UUID of the placement group
        :param pulumi.Input[str] zone: The zone you want to target
        """
        if placement_group_id is not None:
            pulumi.set(__self__, "placement_group_id", placement_group_id)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the placement group
        """
        return pulumi.get(self, "placement_group_id")

    @placement_group_id.setter
    def placement_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "placement_group_id", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "servers", value)

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


class PlacementGroupServers(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 placement_group_id: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PlacementGroupServers resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] placement_group_id: UUID of the placement group
        :param pulumi.Input[str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PlacementGroupServersArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PlacementGroupServers resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PlacementGroupServersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PlacementGroupServersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 placement_group_id: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PlacementGroupServersArgs.__new__(PlacementGroupServersArgs)

            __props__.__dict__["placement_group_id"] = placement_group_id
            __props__.__dict__["servers"] = servers
            __props__.__dict__["zone"] = zone
        super(PlacementGroupServers, __self__).__init__(
            'scaleway-instances:servers:PlacementGroupServers',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PlacementGroupServers':
        """
        Get an existing PlacementGroupServers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PlacementGroupServersArgs.__new__(PlacementGroupServersArgs)

        __props__.__dict__["servers"] = None
        return PlacementGroupServers(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "servers")

