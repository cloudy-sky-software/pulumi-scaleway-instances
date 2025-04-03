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
    'ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate',
    'ScalewayInstanceV1Task',
]

@pulumi.output_type
class ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "volumeType":
            suggest = "volume_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 volume_type: Optional['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType'] = None):
        if volume_type is None:
            volume_type = 'l_ssd'
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType']:
        return pulumi.get(self, "volume_type")


@pulumi.output_type
class ScalewayInstanceV1Task(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hrefFrom":
            suggest = "href_from"
        elif key == "hrefResult":
            suggest = "href_result"
        elif key == "startedAt":
            suggest = "started_at"
        elif key == "terminatedAt":
            suggest = "terminated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1Task. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1Task.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1Task.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[builtins.str] = None,
                 href_from: Optional[builtins.str] = None,
                 href_result: Optional[builtins.str] = None,
                 id: Optional[builtins.str] = None,
                 progress: Optional[builtins.float] = None,
                 started_at: Optional[builtins.str] = None,
                 status: Optional['ScalewayInstanceV1TaskStatus'] = None,
                 terminated_at: Optional[builtins.str] = None,
                 zone: Optional[builtins.str] = None):
        """
        :param builtins.str description: The description of the task
        :param builtins.str id: The unique ID of the task
        :param builtins.float progress: The progress of the task in percent
        :param builtins.str started_at: The task start date (RFC 3339 format)
        :param 'ScalewayInstanceV1TaskStatus' status: The task status
        :param builtins.str terminated_at: The task end date (RFC 3339 format)
        :param builtins.str zone: The zone in which is the task
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if href_from is not None:
            pulumi.set(__self__, "href_from", href_from)
        if href_result is not None:
            pulumi.set(__self__, "href_result", href_result)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if progress is not None:
            pulumi.set(__self__, "progress", progress)
        if started_at is not None:
            pulumi.set(__self__, "started_at", started_at)
        if status is None:
            status = 'pending'
        if status is not None:
            pulumi.set(__self__, "status", status)
        if terminated_at is not None:
            pulumi.set(__self__, "terminated_at", terminated_at)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of the task
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hrefFrom")
    def href_from(self) -> Optional[builtins.str]:
        return pulumi.get(self, "href_from")

    @property
    @pulumi.getter(name="hrefResult")
    def href_result(self) -> Optional[builtins.str]:
        return pulumi.get(self, "href_result")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        The unique ID of the task
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def progress(self) -> Optional[builtins.float]:
        """
        The progress of the task in percent
        """
        return pulumi.get(self, "progress")

    @property
    @pulumi.getter(name="startedAt")
    def started_at(self) -> Optional[builtins.str]:
        """
        The task start date (RFC 3339 format)
        """
        return pulumi.get(self, "started_at")

    @property
    @pulumi.getter
    def status(self) -> Optional['ScalewayInstanceV1TaskStatus']:
        """
        The task status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="terminatedAt")
    def terminated_at(self) -> Optional[builtins.str]:
        """
        The task end date (RFC 3339 format)
        """
        return pulumi.get(self, "terminated_at")

    @property
    @pulumi.getter
    def zone(self) -> Optional[builtins.str]:
        """
        The zone in which is the task
        """
        return pulumi.get(self, "zone")


