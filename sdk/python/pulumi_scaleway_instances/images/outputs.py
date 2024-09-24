# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from . import outputs
from ._enums import *

__all__ = [
    'ScalewayInstanceV1Bootscript',
    'ScalewayInstanceV1Image',
    'ScalewayInstanceV1Volume',
    'ScalewayInstanceV1VolumeServerProperties',
    'ScalewayInstanceV1VolumeSummary',
]

@pulumi.output_type
class ScalewayInstanceV1Bootscript(dict):
    def __init__(__self__, *,
                 arch: Optional['ScalewayInstanceV1BootscriptArch'] = None,
                 bootcmdargs: Optional[str] = None,
                 default: Optional[bool] = None,
                 dtb: Optional[str] = None,
                 id: Optional[str] = None,
                 initrd: Optional[str] = None,
                 kernel: Optional[str] = None,
                 organization: Optional[str] = None,
                 project: Optional[str] = None,
                 public: Optional[bool] = None,
                 title: Optional[str] = None,
                 zone: Optional[str] = None):
        """
        :param 'ScalewayInstanceV1BootscriptArch' arch: The bootscript arch
        :param str bootcmdargs: The bootscript arguments
        :param bool default: Dispmay if the bootscript is the default bootscript if no other boot option is configured
        :param str dtb: Provide information regarding a Device Tree Binary (dtb) for use with C1 servers
        :param str id: The bootscript ID
        :param str initrd: The initrd (initial ramdisk) configuration
        :param str kernel: The server kernel version
        :param str organization: The bootscript organization ID
        :param str project: The bootscript project ID
        :param bool public: Provide information if the bootscript is public
        :param str title: The bootscript title
        :param str zone: The zone in which is the bootscript
        """
        if arch is None:
            arch = 'x86_64'
        if arch is not None:
            pulumi.set(__self__, "arch", arch)
        if bootcmdargs is not None:
            pulumi.set(__self__, "bootcmdargs", bootcmdargs)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if dtb is not None:
            pulumi.set(__self__, "dtb", dtb)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if initrd is not None:
            pulumi.set(__self__, "initrd", initrd)
        if kernel is not None:
            pulumi.set(__self__, "kernel", kernel)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def arch(self) -> Optional['ScalewayInstanceV1BootscriptArch']:
        """
        The bootscript arch
        """
        return pulumi.get(self, "arch")

    @property
    @pulumi.getter
    def bootcmdargs(self) -> Optional[str]:
        """
        The bootscript arguments
        """
        return pulumi.get(self, "bootcmdargs")

    @property
    @pulumi.getter
    def default(self) -> Optional[bool]:
        """
        Dispmay if the bootscript is the default bootscript if no other boot option is configured
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def dtb(self) -> Optional[str]:
        """
        Provide information regarding a Device Tree Binary (dtb) for use with C1 servers
        """
        return pulumi.get(self, "dtb")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The bootscript ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def initrd(self) -> Optional[str]:
        """
        The initrd (initial ramdisk) configuration
        """
        return pulumi.get(self, "initrd")

    @property
    @pulumi.getter
    def kernel(self) -> Optional[str]:
        """
        The server kernel version
        """
        return pulumi.get(self, "kernel")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        """
        The bootscript organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        """
        The bootscript project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def public(self) -> Optional[bool]:
        """
        Provide information if the bootscript is public
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        """
        The bootscript title
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        The zone in which is the bootscript
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1Image(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "rootVolume":
            suggest = "root_volume"
        elif key == "creationDate":
            suggest = "creation_date"
        elif key == "defaultBootscript":
            suggest = "default_bootscript"
        elif key == "extraVolumes":
            suggest = "extra_volumes"
        elif key == "fromServer":
            suggest = "from_server"
        elif key == "modificationDate":
            suggest = "modification_date"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1Image. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1Image.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1Image.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 project: str,
                 root_volume: 'outputs.ScalewayInstanceV1VolumeSummary',
                 arch: Optional['ScalewayInstanceV1ImageArch'] = None,
                 creation_date: Optional[str] = None,
                 default_bootscript: Optional['outputs.ScalewayInstanceV1Bootscript'] = None,
                 extra_volumes: Optional[Mapping[str, 'outputs.ScalewayInstanceV1Volume']] = None,
                 from_server: Optional[str] = None,
                 id: Optional[str] = None,
                 modification_date: Optional[str] = None,
                 organization: Optional[str] = None,
                 public: Optional[bool] = None,
                 state: Optional['ScalewayInstanceV1ImageState'] = None,
                 tags: Optional[Sequence[str]] = None,
                 zone: Optional[str] = None):
        """
        :param str creation_date: (RFC 3339 format)
        :param str modification_date: (RFC 3339 format)
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "root_volume", root_volume)
        if arch is None:
            arch = 'x86_64'
        if arch is not None:
            pulumi.set(__self__, "arch", arch)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if default_bootscript is not None:
            pulumi.set(__self__, "default_bootscript", default_bootscript)
        if extra_volumes is not None:
            pulumi.set(__self__, "extra_volumes", extra_volumes)
        if from_server is not None:
            pulumi.set(__self__, "from_server", from_server)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if modification_date is not None:
            pulumi.set(__self__, "modification_date", modification_date)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if state is None:
            state = 'available'
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="rootVolume")
    def root_volume(self) -> 'outputs.ScalewayInstanceV1VolumeSummary':
        return pulumi.get(self, "root_volume")

    @property
    @pulumi.getter
    def arch(self) -> Optional['ScalewayInstanceV1ImageArch']:
        return pulumi.get(self, "arch")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[str]:
        """
        (RFC 3339 format)
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="defaultBootscript")
    def default_bootscript(self) -> Optional['outputs.ScalewayInstanceV1Bootscript']:
        return pulumi.get(self, "default_bootscript")

    @property
    @pulumi.getter(name="extraVolumes")
    def extra_volumes(self) -> Optional[Mapping[str, 'outputs.ScalewayInstanceV1Volume']]:
        return pulumi.get(self, "extra_volumes")

    @property
    @pulumi.getter(name="fromServer")
    def from_server(self) -> Optional[str]:
        return pulumi.get(self, "from_server")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="modificationDate")
    def modification_date(self) -> Optional[str]:
        """
        (RFC 3339 format)
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def public(self) -> Optional[bool]:
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def state(self) -> Optional['ScalewayInstanceV1ImageState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1Volume(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "creationDate":
            suggest = "creation_date"
        elif key == "exportUri":
            suggest = "export_uri"
        elif key == "modificationDate":
            suggest = "modification_date"
        elif key == "volumeType":
            suggest = "volume_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1Volume. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1Volume.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1Volume.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 project: str,
                 creation_date: Optional[str] = None,
                 export_uri: Optional[str] = None,
                 id: Optional[str] = None,
                 modification_date: Optional[str] = None,
                 organization: Optional[str] = None,
                 server: Optional['outputs.ScalewayInstanceV1VolumeServerProperties'] = None,
                 size: Optional[float] = None,
                 state: Optional['ScalewayInstanceV1VolumeState'] = None,
                 tags: Optional[Sequence[str]] = None,
                 volume_type: Optional['ScalewayInstanceV1VolumeVolumeType'] = None,
                 zone: Optional[str] = None):
        """
        :param str name: The volume name
        :param str project: The volume project ID
        :param str creation_date: The volume creation date (RFC 3339 format)
        :param str export_uri: Show the volume NBD export URI
        :param str modification_date: The volume modification date (RFC 3339 format)
        :param str organization: The volume organization ID
        :param 'ScalewayInstanceV1VolumeServerProperties' server: The server attached to the volume
        :param float size: The volume disk size (in bytes)
        :param Sequence[str] tags: The volume tags
        :param str zone: The zone in which is the volume
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project", project)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if export_uri is not None:
            pulumi.set(__self__, "export_uri", export_uri)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if modification_date is not None:
            pulumi.set(__self__, "modification_date", modification_date)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if state is None:
            state = 'available'
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if volume_type is None:
            volume_type = 'l_ssd'
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The volume name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The volume project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[str]:
        """
        The volume creation date (RFC 3339 format)
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="exportUri")
    def export_uri(self) -> Optional[str]:
        """
        Show the volume NBD export URI
        """
        return pulumi.get(self, "export_uri")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="modificationDate")
    def modification_date(self) -> Optional[str]:
        """
        The volume modification date (RFC 3339 format)
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        """
        The volume organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def server(self) -> Optional['outputs.ScalewayInstanceV1VolumeServerProperties']:
        """
        The server attached to the volume
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def size(self) -> Optional[float]:
        """
        The volume disk size (in bytes)
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> Optional['ScalewayInstanceV1VolumeState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The volume tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional['ScalewayInstanceV1VolumeVolumeType']:
        return pulumi.get(self, "volume_type")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        The zone in which is the volume
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1VolumeServerProperties(dict):
    """
    The server attached to the volume
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        """
        The server attached to the volume
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


@pulumi.output_type
class ScalewayInstanceV1VolumeSummary(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "volumeType":
            suggest = "volume_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1VolumeSummary. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1VolumeSummary.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1VolumeSummary.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 size: Optional[float] = None,
                 volume_type: Optional['ScalewayInstanceV1VolumeSummaryVolumeType'] = None):
        """
        :param float size: (in bytes)
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if volume_type is None:
            volume_type = 'l_ssd'
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> Optional[float]:
        """
        (in bytes)
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional['ScalewayInstanceV1VolumeSummaryVolumeType']:
        return pulumi.get(self, "volume_type")


