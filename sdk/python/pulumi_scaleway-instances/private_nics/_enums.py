# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ScalewayInstanceV1PrivateNICState',
]


class ScalewayInstanceV1PrivateNICState(str, Enum):
    """
    The private NIC state
    """
    AVAILABLE = "available"
    SYNCING = "syncing"
    SYNCING_ERROR = "syncing_error"
