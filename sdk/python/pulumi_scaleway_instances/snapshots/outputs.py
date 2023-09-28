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
    'BaseVolumeProperties',
    'ScalewayInstanceV1GetSnapshotResponse',
    'ScalewayInstanceV1ListSnapshotsResponse',
    'ScalewayInstanceV1Snapshot',
    'ScalewayInstanceV1SnapshotBaseVolumeProperties',
    'ScalewayInstanceV1Task',
]

@pulumi.output_type
class BaseVolumeProperties(dict):
    """
    The volume on which the snapshot is based on
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        """
        The volume on which the snapshot is based on
        :param str id: The volume ID on which the snapshot is based on
        :param str name: The volume name on which the snapshot is based on
        """
        BaseVolumeProperties._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            id=id,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             id: Optional[str] = None,
             name: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if id is not None:
            _setter("id", id)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The volume ID on which the snapshot is based on
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The volume name on which the snapshot is based on
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ScalewayInstanceV1GetSnapshotResponse(dict):
    def __init__(__self__, *,
                 snapshot: Optional['outputs.ScalewayInstanceV1Snapshot'] = None):
        ScalewayInstanceV1GetSnapshotResponse._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            snapshot=snapshot,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             snapshot: Optional['outputs.ScalewayInstanceV1Snapshot'] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if snapshot is not None:
            _setter("snapshot", snapshot)

    @property
    @pulumi.getter
    def snapshot(self) -> Optional['outputs.ScalewayInstanceV1Snapshot']:
        return pulumi.get(self, "snapshot")


@pulumi.output_type
class ScalewayInstanceV1ListSnapshotsResponse(dict):
    def __init__(__self__, *,
                 snapshots: Optional[Sequence['outputs.ScalewayInstanceV1Snapshot']] = None):
        """
        :param Sequence['ScalewayInstanceV1Snapshot'] snapshots: List of snapshots
        """
        ScalewayInstanceV1ListSnapshotsResponse._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            snapshots=snapshots,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             snapshots: Optional[Sequence['outputs.ScalewayInstanceV1Snapshot']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if snapshots is not None:
            _setter("snapshots", snapshots)

    @property
    @pulumi.getter
    def snapshots(self) -> Optional[Sequence['outputs.ScalewayInstanceV1Snapshot']]:
        """
        List of snapshots
        """
        return pulumi.get(self, "snapshots")


@pulumi.output_type
class ScalewayInstanceV1Snapshot(dict):
    def __init__(__self__, *,
                 base_volume: Optional['outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties'] = None,
                 creation_date: Optional[str] = None,
                 error_reason: Optional[str] = None,
                 id: Optional[str] = None,
                 modification_date: Optional[str] = None,
                 name: Optional[str] = None,
                 organization: Optional[str] = None,
                 project: Optional[str] = None,
                 size: Optional[float] = None,
                 state: Optional['ScalewayInstanceV1SnapshotState'] = None,
                 tags: Optional[Sequence[str]] = None,
                 volume_type: Optional['ScalewayInstanceV1SnapshotVolumeType'] = None,
                 zone: Optional[str] = None):
        """
        :param 'ScalewayInstanceV1SnapshotBaseVolumeProperties' base_volume: The volume on which the snapshot is based on
        :param str creation_date: The snapshot creation date (RFC 3339 format)
        :param str error_reason: The reason for the failed snapshot import
        :param str modification_date: The snapshot modification date (RFC 3339 format)
        :param str name: The snapshot name
        :param str organization: The snapshot organization ID
        :param str project: The snapshot project ID
        :param float size: The snapshot size (in bytes)
        :param Sequence[str] tags: The snapshot tags
        :param str zone: The snapshot zone
        """
        ScalewayInstanceV1Snapshot._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            base_volume=base_volume,
            creation_date=creation_date,
            error_reason=error_reason,
            id=id,
            modification_date=modification_date,
            name=name,
            organization=organization,
            project=project,
            size=size,
            state=state,
            tags=tags,
            volume_type=volume_type,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             base_volume: Optional['outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties'] = None,
             creation_date: Optional[str] = None,
             error_reason: Optional[str] = None,
             id: Optional[str] = None,
             modification_date: Optional[str] = None,
             name: Optional[str] = None,
             organization: Optional[str] = None,
             project: Optional[str] = None,
             size: Optional[float] = None,
             state: Optional['ScalewayInstanceV1SnapshotState'] = None,
             tags: Optional[Sequence[str]] = None,
             volume_type: Optional['ScalewayInstanceV1SnapshotVolumeType'] = None,
             zone: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if base_volume is not None:
            _setter("base_volume", base_volume)
        if creation_date is not None:
            _setter("creation_date", creation_date)
        if error_reason is not None:
            _setter("error_reason", error_reason)
        if id is not None:
            _setter("id", id)
        if modification_date is not None:
            _setter("modification_date", modification_date)
        if name is not None:
            _setter("name", name)
        if organization is not None:
            _setter("organization", organization)
        if project is not None:
            _setter("project", project)
        if size is not None:
            _setter("size", size)
        if state is None:
            state = 'available'
        if state is not None:
            _setter("state", state)
        if tags is not None:
            _setter("tags", tags)
        if volume_type is None:
            volume_type = 'l_ssd'
        if volume_type is not None:
            _setter("volume_type", volume_type)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def base_volume(self) -> Optional['outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties']:
        """
        The volume on which the snapshot is based on
        """
        return pulumi.get(self, "base_volume")

    @property
    @pulumi.getter
    def creation_date(self) -> Optional[str]:
        """
        The snapshot creation date (RFC 3339 format)
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def error_reason(self) -> Optional[str]:
        """
        The reason for the failed snapshot import
        """
        return pulumi.get(self, "error_reason")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def modification_date(self) -> Optional[str]:
        """
        The snapshot modification date (RFC 3339 format)
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The snapshot name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        """
        The snapshot organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        """
        The snapshot project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def size(self) -> Optional[float]:
        """
        The snapshot size (in bytes)
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> Optional['ScalewayInstanceV1SnapshotState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The snapshot tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def volume_type(self) -> Optional['ScalewayInstanceV1SnapshotVolumeType']:
        return pulumi.get(self, "volume_type")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        The snapshot zone
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1SnapshotBaseVolumeProperties(dict):
    """
    The volume on which the snapshot is based on
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        """
        The volume on which the snapshot is based on
        :param str id: The volume ID on which the snapshot is based on
        :param str name: The volume name on which the snapshot is based on
        """
        ScalewayInstanceV1SnapshotBaseVolumeProperties._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            id=id,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             id: Optional[str] = None,
             name: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if id is not None:
            _setter("id", id)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The volume ID on which the snapshot is based on
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The volume name on which the snapshot is based on
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ScalewayInstanceV1Task(dict):
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 href_from: Optional[str] = None,
                 href_result: Optional[str] = None,
                 id: Optional[str] = None,
                 progress: Optional[float] = None,
                 started_at: Optional[str] = None,
                 status: Optional['ScalewayInstanceV1TaskStatus'] = None,
                 terminated_at: Optional[str] = None,
                 zone: Optional[str] = None):
        """
        :param str description: The description of the task
        :param str id: The unique ID of the task
        :param float progress: The progress of the task in percent
        :param str started_at: The task start date (RFC 3339 format)
        :param 'ScalewayInstanceV1TaskStatus' status: The task status
        :param str terminated_at: The task end date (RFC 3339 format)
        :param str zone: The zone in which is the task
        """
        ScalewayInstanceV1Task._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            href_from=href_from,
            href_result=href_result,
            id=id,
            progress=progress,
            started_at=started_at,
            status=status,
            terminated_at=terminated_at,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[str] = None,
             href_from: Optional[str] = None,
             href_result: Optional[str] = None,
             id: Optional[str] = None,
             progress: Optional[float] = None,
             started_at: Optional[str] = None,
             status: Optional['ScalewayInstanceV1TaskStatus'] = None,
             terminated_at: Optional[str] = None,
             zone: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if description is not None:
            _setter("description", description)
        if href_from is not None:
            _setter("href_from", href_from)
        if href_result is not None:
            _setter("href_result", href_result)
        if id is not None:
            _setter("id", id)
        if progress is not None:
            _setter("progress", progress)
        if started_at is not None:
            _setter("started_at", started_at)
        if status is None:
            status = 'pending'
        if status is not None:
            _setter("status", status)
        if terminated_at is not None:
            _setter("terminated_at", terminated_at)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the task
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def href_from(self) -> Optional[str]:
        return pulumi.get(self, "href_from")

    @property
    @pulumi.getter
    def href_result(self) -> Optional[str]:
        return pulumi.get(self, "href_result")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The unique ID of the task
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def progress(self) -> Optional[float]:
        """
        The progress of the task in percent
        """
        return pulumi.get(self, "progress")

    @property
    @pulumi.getter
    def started_at(self) -> Optional[str]:
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
    @pulumi.getter
    def terminated_at(self) -> Optional[str]:
        """
        The task end date (RFC 3339 format)
        """
        return pulumi.get(self, "terminated_at")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        The zone in which is the task
        """
        return pulumi.get(self, "zone")


