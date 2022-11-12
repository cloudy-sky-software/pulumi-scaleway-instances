# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ScalewayInstanceV1VolumeState',
    'ScalewayInstanceV1VolumeVolumeType',
    'State',
    'VolumeType',
]


class ScalewayInstanceV1VolumeState(str, Enum):
    AVAILABLE = "available"
    SNAPSHOTTING = "snapshotting"
    ERROR = "error"
    FETCHING = "fetching"
    RESIZING = "resizing"
    SAVING = "saving"
    HOTSYNCING = "hotsyncing"


class ScalewayInstanceV1VolumeVolumeType(str, Enum):
    LSSD = "l_ssd"
    BSSD = "b_ssd"
    UNIFIED = "unified"


class State(str, Enum):
    AVAILABLE = "available"
    SNAPSHOTTING = "snapshotting"
    ERROR = "error"
    FETCHING = "fetching"
    RESIZING = "resizing"
    SAVING = "saving"
    HOTSYNCING = "hotsyncing"


class VolumeType(str, Enum):
    LSSD = "l_ssd"
    BSSD = "b_ssd"
    UNIFIED = "unified"
