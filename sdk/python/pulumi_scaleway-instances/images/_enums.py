# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'Arch',
    'ScalewayInstanceV1BootscriptArch',
    'ScalewayInstanceV1ImageArch',
    'ScalewayInstanceV1ImageState',
    'ScalewayInstanceV1VolumeState',
    'ScalewayInstanceV1VolumeSummaryVolumeType',
    'ScalewayInstanceV1VolumeVolumeType',
    'State',
]


class Arch(str, Enum):
    X8664 = "x86_64"
    ARM = "arm"


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


class State(str, Enum):
    AVAILABLE = "available"
    CREATING = "creating"
    ERROR = "error"
