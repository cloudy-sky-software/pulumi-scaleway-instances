# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['SecurityGroupRuleArgs', 'SecurityGroupRule']

@pulumi.input_type
class SecurityGroupRuleArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input['Action']] = None,
                 direction: Optional[pulumi.Input['Direction']] = None,
                 ip_range: pulumi.Input[str],
                 protocol: Optional[pulumi.Input['Protocol']] = None,
                 dest_port_from: Optional[pulumi.Input[float]] = None,
                 dest_port_to: Optional[pulumi.Input[float]] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 position: Optional[pulumi.Input[float]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityGroupRule resource.
        :param pulumi.Input[str] ip_range: (IP network)
        :param pulumi.Input[float] dest_port_from: The beginning of the range of ports to apply this rule to (inclusive)
        :param pulumi.Input[float] dest_port_to: The end of the range of ports to apply this rule to (inclusive)
        :param pulumi.Input[bool] editable: Indicates if this rule is editable (will be ignored)
        :param pulumi.Input[float] position: The position of this rule in the security group rules list
        :param pulumi.Input[str] security_group_id: UUID of the security group
        :param pulumi.Input[str] zone: The zone you want to target
        """
        SecurityGroupRuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            direction=direction,
            ip_range=ip_range,
            protocol=protocol,
            dest_port_from=dest_port_from,
            dest_port_to=dest_port_to,
            editable=editable,
            position=position,
            security_group_id=security_group_id,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional[pulumi.Input['Action']] = None,
             direction: Optional[pulumi.Input['Direction']] = None,
             ip_range: pulumi.Input[str],
             protocol: Optional[pulumi.Input['Protocol']] = None,
             dest_port_from: Optional[pulumi.Input[float]] = None,
             dest_port_to: Optional[pulumi.Input[float]] = None,
             editable: Optional[pulumi.Input[bool]] = None,
             position: Optional[pulumi.Input[float]] = None,
             security_group_id: Optional[pulumi.Input[str]] = None,
             zone: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if action is None:
            action = 'accept'
        _setter("action", action)
        if direction is None:
            direction = 'inbound'
        _setter("direction", direction)
        _setter("ip_range", ip_range)
        if protocol is None:
            protocol = 'TCP'
        _setter("protocol", protocol)
        if dest_port_from is not None:
            _setter("dest_port_from", dest_port_from)
        if dest_port_to is not None:
            _setter("dest_port_to", dest_port_to)
        if editable is not None:
            _setter("editable", editable)
        if position is not None:
            _setter("position", position)
        if security_group_id is not None:
            _setter("security_group_id", security_group_id)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input['Action']:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input['Action']):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Input['Direction']:
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: pulumi.Input['Direction']):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def ip_range(self) -> pulumi.Input[str]:
        """
        (IP network)
        """
        return pulumi.get(self, "ip_range")

    @ip_range.setter
    def ip_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_range", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['Protocol']:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['Protocol']):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def dest_port_from(self) -> Optional[pulumi.Input[float]]:
        """
        The beginning of the range of ports to apply this rule to (inclusive)
        """
        return pulumi.get(self, "dest_port_from")

    @dest_port_from.setter
    def dest_port_from(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "dest_port_from", value)

    @property
    @pulumi.getter
    def dest_port_to(self) -> Optional[pulumi.Input[float]]:
        """
        The end of the range of ports to apply this rule to (inclusive)
        """
        return pulumi.get(self, "dest_port_to")

    @dest_port_to.setter
    def dest_port_to(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "dest_port_to", value)

    @property
    @pulumi.getter
    def editable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this rule is editable (will be ignored)
        """
        return pulumi.get(self, "editable")

    @editable.setter
    def editable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "editable", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[float]]:
        """
        The position of this rule in the security group rules list
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the security group
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

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


class SecurityGroupRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['Action']] = None,
                 dest_port_from: Optional[pulumi.Input[float]] = None,
                 dest_port_to: Optional[pulumi.Input[float]] = None,
                 direction: Optional[pulumi.Input['Direction']] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[float]] = None,
                 protocol: Optional[pulumi.Input['Protocol']] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SecurityGroupRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] dest_port_from: The beginning of the range of ports to apply this rule to (inclusive)
        :param pulumi.Input[float] dest_port_to: The end of the range of ports to apply this rule to (inclusive)
        :param pulumi.Input[bool] editable: Indicates if this rule is editable (will be ignored)
        :param pulumi.Input[str] ip_range: (IP network)
        :param pulumi.Input[float] position: The position of this rule in the security group rules list
        :param pulumi.Input[str] security_group_id: UUID of the security group
        :param pulumi.Input[str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityGroupRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecurityGroupRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecurityGroupRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            SecurityGroupRuleArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['Action']] = None,
                 dest_port_from: Optional[pulumi.Input[float]] = None,
                 dest_port_to: Optional[pulumi.Input[float]] = None,
                 direction: Optional[pulumi.Input['Direction']] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[float]] = None,
                 protocol: Optional[pulumi.Input['Protocol']] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityGroupRuleArgs.__new__(SecurityGroupRuleArgs)

            if action is None:
                action = 'accept'
            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["dest_port_from"] = dest_port_from
            __props__.__dict__["dest_port_to"] = dest_port_to
            if direction is None:
                direction = 'inbound'
            if direction is None and not opts.urn:
                raise TypeError("Missing required property 'direction'")
            __props__.__dict__["direction"] = direction
            __props__.__dict__["editable"] = editable
            if ip_range is None and not opts.urn:
                raise TypeError("Missing required property 'ip_range'")
            __props__.__dict__["ip_range"] = ip_range
            __props__.__dict__["position"] = position
            if protocol is None:
                protocol = 'TCP'
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["security_group_id"] = security_group_id
            __props__.__dict__["zone"] = zone
        super(SecurityGroupRule, __self__).__init__(
            'scaleway-instances:rules:SecurityGroupRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecurityGroupRule':
        """
        Get an existing SecurityGroupRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecurityGroupRuleArgs.__new__(SecurityGroupRuleArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["dest_port_from"] = None
        __props__.__dict__["dest_port_to"] = None
        __props__.__dict__["direction"] = None
        __props__.__dict__["editable"] = None
        __props__.__dict__["ip_range"] = None
        __props__.__dict__["position"] = None
        __props__.__dict__["protocol"] = None
        return SecurityGroupRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output['Action']:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def dest_port_from(self) -> pulumi.Output[Optional[float]]:
        """
        The beginning of the range of ports to apply this rule to (inclusive)
        """
        return pulumi.get(self, "dest_port_from")

    @property
    @pulumi.getter
    def dest_port_to(self) -> pulumi.Output[Optional[float]]:
        """
        The end of the range of ports to apply this rule to (inclusive)
        """
        return pulumi.get(self, "dest_port_to")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output['Direction']:
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def editable(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if this rule is editable (will be ignored)
        """
        return pulumi.get(self, "editable")

    @property
    @pulumi.getter
    def ip_range(self) -> pulumi.Output[str]:
        """
        (IP network)
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[Optional[float]]:
        """
        The position of this rule in the security group rules list
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output['Protocol']:
        return pulumi.get(self, "protocol")

