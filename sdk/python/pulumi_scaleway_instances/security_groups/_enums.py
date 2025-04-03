# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'InboundDefaultPolicy',
    'OutboundDefaultPolicy',
    'ScalewayInstanceV1SecurityGroupInboundDefaultPolicy',
    'ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy',
    'ScalewayInstanceV1SecurityGroupState',
    'State',
]


class InboundDefaultPolicy(builtins.str, Enum):
    """
    The default inbound policy
    """
    ACCEPT = "accept"
    DROP = "drop"


class OutboundDefaultPolicy(builtins.str, Enum):
    """
    The default outbound policy
    """
    ACCEPT = "accept"
    DROP = "drop"


class ScalewayInstanceV1SecurityGroupInboundDefaultPolicy(builtins.str, Enum):
    """
    The default inbound policy
    """
    ACCEPT = "accept"
    DROP = "drop"


class ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy(builtins.str, Enum):
    """
    The default outbound policy
    """
    ACCEPT = "accept"
    DROP = "drop"


class ScalewayInstanceV1SecurityGroupState(builtins.str, Enum):
    """
    Security group state
    """
    AVAILABLE = "available"
    SYNCING = "syncing"
    SYNCING_ERROR = "syncing_error"


class State(builtins.str, Enum):
    """
    Security group state
    """
    AVAILABLE = "available"
    SYNCING = "syncing"
    SYNCING_ERROR = "syncing_error"
