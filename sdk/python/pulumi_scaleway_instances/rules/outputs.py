# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GoogleProtobufUInt32Value',
    'ScalewayInstanceV1GetSecurityGroupRuleResponse',
    'ScalewayInstanceV1ListSecurityGroupRulesResponse',
    'ScalewayInstanceV1SecurityGroupRule',
    'ScalewayInstanceV1SetSecurityGroupRulesRequestRule',
]

@pulumi.output_type
class GoogleProtobufUInt32Value(dict):
    def __init__(__self__):
        pass
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             opts: Optional[pulumi.ResourceOptions]=None):
        pass


@pulumi.output_type
class ScalewayInstanceV1GetSecurityGroupRuleResponse(dict):
    def __init__(__self__, *,
                 rule: Optional['outputs.ScalewayInstanceV1SecurityGroupRule'] = None):
        ScalewayInstanceV1GetSecurityGroupRuleResponse._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            rule=rule,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             rule: Optional['outputs.ScalewayInstanceV1SecurityGroupRule'] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if rule is not None:
            _setter("rule", rule)

    @property
    @pulumi.getter
    def rule(self) -> Optional['outputs.ScalewayInstanceV1SecurityGroupRule']:
        return pulumi.get(self, "rule")


@pulumi.output_type
class ScalewayInstanceV1ListSecurityGroupRulesResponse(dict):
    def __init__(__self__, *,
                 rules: Optional[Sequence['outputs.ScalewayInstanceV1SecurityGroupRule']] = None):
        """
        :param Sequence['ScalewayInstanceV1SecurityGroupRule'] rules: List of security rules
        """
        ScalewayInstanceV1ListSecurityGroupRulesResponse._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            rules=rules,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             rules: Optional[Sequence['outputs.ScalewayInstanceV1SecurityGroupRule']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if rules is not None:
            _setter("rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Optional[Sequence['outputs.ScalewayInstanceV1SecurityGroupRule']]:
        """
        List of security rules
        """
        return pulumi.get(self, "rules")


@pulumi.output_type
class ScalewayInstanceV1SecurityGroupRule(dict):
    def __init__(__self__, *,
                 action: Optional['ScalewayInstanceV1SecurityGroupRuleAction'] = None,
                 dest_port_from: Optional['outputs.GoogleProtobufUInt32Value'] = None,
                 dest_port_to: Optional['outputs.GoogleProtobufUInt32Value'] = None,
                 direction: Optional['ScalewayInstanceV1SecurityGroupRuleDirection'] = None,
                 editable: Optional[bool] = None,
                 id: Optional[str] = None,
                 ip_range: Optional[str] = None,
                 position: Optional[float] = None,
                 protocol: Optional['ScalewayInstanceV1SecurityGroupRuleProtocol'] = None,
                 zone: Optional[str] = None):
        """
        :param str ip_range: (IP network)
        """
        ScalewayInstanceV1SecurityGroupRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            dest_port_from=dest_port_from,
            dest_port_to=dest_port_to,
            direction=direction,
            editable=editable,
            id=id,
            ip_range=ip_range,
            position=position,
            protocol=protocol,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional['ScalewayInstanceV1SecurityGroupRuleAction'] = None,
             dest_port_from: Optional['outputs.GoogleProtobufUInt32Value'] = None,
             dest_port_to: Optional['outputs.GoogleProtobufUInt32Value'] = None,
             direction: Optional['ScalewayInstanceV1SecurityGroupRuleDirection'] = None,
             editable: Optional[bool] = None,
             id: Optional[str] = None,
             ip_range: Optional[str] = None,
             position: Optional[float] = None,
             protocol: Optional['ScalewayInstanceV1SecurityGroupRuleProtocol'] = None,
             zone: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if action is None:
            action = 'accept'
        if action is not None:
            _setter("action", action)
        if dest_port_from is not None:
            _setter("dest_port_from", dest_port_from)
        if dest_port_to is not None:
            _setter("dest_port_to", dest_port_to)
        if direction is None:
            direction = 'inbound'
        if direction is not None:
            _setter("direction", direction)
        if editable is not None:
            _setter("editable", editable)
        if id is not None:
            _setter("id", id)
        if ip_range is not None:
            _setter("ip_range", ip_range)
        if position is not None:
            _setter("position", position)
        if protocol is None:
            protocol = 'TCP'
        if protocol is not None:
            _setter("protocol", protocol)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def action(self) -> Optional['ScalewayInstanceV1SecurityGroupRuleAction']:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def dest_port_from(self) -> Optional['outputs.GoogleProtobufUInt32Value']:
        return pulumi.get(self, "dest_port_from")

    @property
    @pulumi.getter
    def dest_port_to(self) -> Optional['outputs.GoogleProtobufUInt32Value']:
        return pulumi.get(self, "dest_port_to")

    @property
    @pulumi.getter
    def direction(self) -> Optional['ScalewayInstanceV1SecurityGroupRuleDirection']:
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def editable(self) -> Optional[bool]:
        return pulumi.get(self, "editable")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip_range(self) -> Optional[str]:
        """
        (IP network)
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter
    def position(self) -> Optional[float]:
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def protocol(self) -> Optional['ScalewayInstanceV1SecurityGroupRuleProtocol']:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1SetSecurityGroupRulesRequestRule(dict):
    def __init__(__self__, *,
                 action: Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction'] = None,
                 dest_port_from: Optional[float] = None,
                 dest_port_to: Optional[float] = None,
                 direction: Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection'] = None,
                 editable: Optional[bool] = None,
                 id: Optional[str] = None,
                 ip_range: Optional[str] = None,
                 position: Optional[float] = None,
                 protocol: Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol'] = None,
                 zone: Optional[str] = None):
        """
        :param 'ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction' action: Action to apply when the rule matches a packet
        :param float dest_port_from: Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
        :param float dest_port_to: End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
        :param 'ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection' direction: Direction the rule applies to
        :param bool editable: Indicates if this rule is editable. Rules with the value false will be ignored
        :param str id: UUID of the security rule to update. If no value is provided, a new rule will be created
        :param str ip_range: The range of IP address this rules applies to (IP network)
        :param float position: Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
        :param 'ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol' protocol: Protocol family this rule applies to
        :param str zone: Zone of the rule. This field is ignored
        """
        ScalewayInstanceV1SetSecurityGroupRulesRequestRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            dest_port_from=dest_port_from,
            dest_port_to=dest_port_to,
            direction=direction,
            editable=editable,
            id=id,
            ip_range=ip_range,
            position=position,
            protocol=protocol,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction'] = None,
             dest_port_from: Optional[float] = None,
             dest_port_to: Optional[float] = None,
             direction: Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection'] = None,
             editable: Optional[bool] = None,
             id: Optional[str] = None,
             ip_range: Optional[str] = None,
             position: Optional[float] = None,
             protocol: Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol'] = None,
             zone: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if action is None:
            action = 'accept'
        if action is not None:
            _setter("action", action)
        if dest_port_from is not None:
            _setter("dest_port_from", dest_port_from)
        if dest_port_to is not None:
            _setter("dest_port_to", dest_port_to)
        if direction is None:
            direction = 'inbound'
        if direction is not None:
            _setter("direction", direction)
        if editable is not None:
            _setter("editable", editable)
        if id is not None:
            _setter("id", id)
        if ip_range is not None:
            _setter("ip_range", ip_range)
        if position is not None:
            _setter("position", position)
        if protocol is None:
            protocol = 'TCP'
        if protocol is not None:
            _setter("protocol", protocol)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def action(self) -> Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction']:
        """
        Action to apply when the rule matches a packet
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def dest_port_from(self) -> Optional[float]:
        """
        Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
        """
        return pulumi.get(self, "dest_port_from")

    @property
    @pulumi.getter
    def dest_port_to(self) -> Optional[float]:
        """
        End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
        """
        return pulumi.get(self, "dest_port_to")

    @property
    @pulumi.getter
    def direction(self) -> Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection']:
        """
        Direction the rule applies to
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def editable(self) -> Optional[bool]:
        """
        Indicates if this rule is editable. Rules with the value false will be ignored
        """
        return pulumi.get(self, "editable")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        UUID of the security rule to update. If no value is provided, a new rule will be created
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip_range(self) -> Optional[str]:
        """
        The range of IP address this rules applies to (IP network)
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter
    def position(self) -> Optional[float]:
        """
        Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def protocol(self) -> Optional['ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol']:
        """
        Protocol family this rule applies to
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        Zone of the rule. This field is ignored
        """
        return pulumi.get(self, "zone")


