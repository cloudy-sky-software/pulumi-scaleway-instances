# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'BootType',
    'ScalewayInstanceV1BootscriptArch',
    'ScalewayInstanceV1ImageArch',
    'ScalewayInstanceV1ImageState',
    'ScalewayInstanceV1PlacementGroupPolicyMode',
    'ScalewayInstanceV1PlacementGroupPolicyType',
    'ScalewayInstanceV1PrivateNICState',
    'ScalewayInstanceV1ServerAllowedActionsItem',
    'ScalewayInstanceV1ServerArch',
    'ScalewayInstanceV1ServerBootType',
    'ScalewayInstanceV1ServerState',
    'ScalewayInstanceV1ServerTypeArch',
    'ScalewayInstanceV1ServerTypeCapabilitiesPropertiesBootTypesItem',
    'ScalewayInstanceV1VolumeServerState',
    'ScalewayInstanceV1VolumeServerTemplateVolumeType',
    'ScalewayInstanceV1VolumeServerVolumeType',
    'ScalewayInstanceV1VolumeState',
    'ScalewayInstanceV1VolumeSummaryVolumeType',
    'ScalewayInstanceV1VolumeVolumeType',
]


class BootType(str, Enum):
    """
    The boot type to use
    """
    LOCAL = "local"
    BOOTSCRIPT = "bootscript"
    RESCUE = "rescue"


class ScalewayInstanceV1BootscriptArch(str, Enum):
    """
    The bootscript arch
    """
    X8664 = "x86_64"
    ARM = "arm"


class ScalewayInstanceV1ImageArch(str, Enum):
    X8664 = "x86_64"
    ARM = "arm"


class ScalewayInstanceV1ImageState(str, Enum):
    AVAILABLE = "available"
    CREATING = "creating"
    ERROR = "error"


class ScalewayInstanceV1PlacementGroupPolicyMode(str, Enum):
    OPTIONAL = "optional"
    ENFORCED = "enforced"


class ScalewayInstanceV1PlacementGroupPolicyType(str, Enum):
    MAX_AVAILABILITY = "max_availability"
    LOW_LATENCY = "low_latency"


class ScalewayInstanceV1PrivateNICState(str, Enum):
    """
    The private NIC state
    """
    AVAILABLE = "available"
    SYNCING = "syncing"
    SYNCING_ERROR = "syncing_error"


class ScalewayInstanceV1ServerAllowedActionsItem(str, Enum):
    POWERON = "poweron"
    BACKUP = "backup"
    STOP_IN_PLACE = "stop_in_place"
    POWEROFF = "poweroff"
    TERMINATE = "terminate"
    REBOOT = "reboot"


class ScalewayInstanceV1ServerArch(str, Enum):
    """
    The server arch
    """
    X8664 = "x86_64"
    ARM = "arm"


class ScalewayInstanceV1ServerBootType(str, Enum):
    LOCAL = "local"
    BOOTSCRIPT = "bootscript"
    RESCUE = "rescue"


class ScalewayInstanceV1ServerState(str, Enum):
    """
    The server state
    """
    RUNNING = "running"
    STOPPED = "stopped"
    STOPPED_IN_PLACE = "stopped in place"
    STARTING = "starting"
    STOPPING = "stopping"
    LOCKED = "locked"


class ScalewayInstanceV1ServerTypeArch(str, Enum):
    """
    CPU architecture
    """
    X8664 = "x86_64"
    ARM = "arm"


class ScalewayInstanceV1ServerTypeCapabilitiesPropertiesBootTypesItem(str, Enum):
    LOCAL = "local"
    BOOTSCRIPT = "bootscript"
    RESCUE = "rescue"


class ScalewayInstanceV1VolumeServerState(str, Enum):
    AVAILABLE = "available"
    SNAPSHOTTING = "snapshotting"
    ERROR = "error"
    FETCHING = "fetching"
    RESIZING = "resizing"
    SAVING = "saving"
    HOTSYNCING = "hotsyncing"


class ScalewayInstanceV1VolumeServerTemplateVolumeType(str, Enum):
    LSSD = "l_ssd"
    BSSD = "b_ssd"


class ScalewayInstanceV1VolumeServerVolumeType(str, Enum):
    LSSD = "l_ssd"
    BSSD = "b_ssd"


class ScalewayInstanceV1VolumeState(str, Enum):
    AVAILABLE = "available"
    SNAPSHOTTING = "snapshotting"
    ERROR = "error"
    FETCHING = "fetching"
    RESIZING = "resizing"
    SAVING = "saving"
    HOTSYNCING = "hotsyncing"


class ScalewayInstanceV1VolumeSummaryVolumeType(str, Enum):
    LSSD = "l_ssd"
    BSSD = "b_ssd"
    UNIFIED = "unified"


class ScalewayInstanceV1VolumeVolumeType(str, Enum):
    LSSD = "l_ssd"
    BSSD = "b_ssd"
    UNIFIED = "unified"
