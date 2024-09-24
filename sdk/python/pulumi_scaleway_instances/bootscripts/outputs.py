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
from ._enums import *

__all__ = [
    'ScalewayInstanceV1Bootscript',
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


