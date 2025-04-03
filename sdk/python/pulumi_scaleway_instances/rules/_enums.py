# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'Action',
    'Direction',
    'Protocol',
    'ScalewayInstanceV1SecurityGroupRuleAction',
    'ScalewayInstanceV1SecurityGroupRuleDirection',
    'ScalewayInstanceV1SecurityGroupRuleProtocol',
]


class Action(builtins.str, Enum):
    ACCEPT = "accept"
    DROP = "drop"


class Direction(builtins.str, Enum):
    INBOUND = "inbound"
    OUTBOUND = "outbound"


class Protocol(builtins.str, Enum):
    TCP = "TCP"
    UDP = "UDP"
    ICMP = "ICMP"
    ANY = "ANY"


class ScalewayInstanceV1SecurityGroupRuleAction(builtins.str, Enum):
    ACCEPT = "accept"
    DROP = "drop"


class ScalewayInstanceV1SecurityGroupRuleDirection(builtins.str, Enum):
    INBOUND = "inbound"
    OUTBOUND = "outbound"


class ScalewayInstanceV1SecurityGroupRuleProtocol(builtins.str, Enum):
    TCP = "TCP"
    UDP = "UDP"
    ICMP = "ICMP"
    ANY = "ANY"
