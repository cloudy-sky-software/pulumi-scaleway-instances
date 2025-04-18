# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from ._enums import *

__all__ = [
    'ScalewayInstanceV1VolumeServerTemplateArgs',
    'ScalewayInstanceV1VolumeServerTemplateArgsDict',
]

MYPY = False

if not MYPY:
    class ScalewayInstanceV1VolumeServerTemplateArgsDict(TypedDict):
        base_snapshot: NotRequired[pulumi.Input[builtins.str]]
        """
        The ID of the snapshot on which this volume will be based
        """
        boot: NotRequired[pulumi.Input[builtins.bool]]
        """
        Force the server to boot on this volume
        """
        id: NotRequired[pulumi.Input[builtins.str]]
        """
        UUID of the volume
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        Name of the volume
        """
        organization: NotRequired[pulumi.Input[builtins.str]]
        """
        Organization ID of the volume
        """
        project: NotRequired[pulumi.Input[builtins.str]]
        """
        Project ID of the volume
        """
        size: NotRequired[pulumi.Input[builtins.float]]
        """
        Disk size of the volume, must be a multiple of 512 (in bytes)
        """
        volume_type: NotRequired[pulumi.Input['ScalewayInstanceV1VolumeServerTemplateVolumeType']]
elif False:
    ScalewayInstanceV1VolumeServerTemplateArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ScalewayInstanceV1VolumeServerTemplateArgs:
    def __init__(__self__, *,
                 base_snapshot: Optional[pulumi.Input[builtins.str]] = None,
                 boot: Optional[pulumi.Input[builtins.bool]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 organization: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 size: Optional[pulumi.Input[builtins.float]] = None,
                 volume_type: Optional[pulumi.Input['ScalewayInstanceV1VolumeServerTemplateVolumeType']] = None):
        """
        :param pulumi.Input[builtins.str] base_snapshot: The ID of the snapshot on which this volume will be based
        :param pulumi.Input[builtins.bool] boot: Force the server to boot on this volume
        :param pulumi.Input[builtins.str] id: UUID of the volume
        :param pulumi.Input[builtins.str] name: Name of the volume
        :param pulumi.Input[builtins.str] organization: Organization ID of the volume
        :param pulumi.Input[builtins.str] project: Project ID of the volume
        :param pulumi.Input[builtins.float] size: Disk size of the volume, must be a multiple of 512 (in bytes)
        """
        if base_snapshot is not None:
            pulumi.set(__self__, "base_snapshot", base_snapshot)
        if boot is None:
            boot = False
        if boot is not None:
            pulumi.set(__self__, "boot", boot)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if volume_type is None:
            volume_type = 'l_ssd'
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="baseSnapshot")
    def base_snapshot(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the snapshot on which this volume will be based
        """
        return pulumi.get(self, "base_snapshot")

    @base_snapshot.setter
    def base_snapshot(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "base_snapshot", value)

    @property
    @pulumi.getter
    def boot(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Force the server to boot on this volume
        """
        return pulumi.get(self, "boot")

    @boot.setter
    def boot(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "boot", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        UUID of the volume
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the volume
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Organization ID of the volume
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Project ID of the volume
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        Disk size of the volume, must be a multiple of 512 (in bytes)
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input['ScalewayInstanceV1VolumeServerTemplateVolumeType']]:
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input['ScalewayInstanceV1VolumeServerTemplateVolumeType']]):
        pulumi.set(self, "volume_type", value)


