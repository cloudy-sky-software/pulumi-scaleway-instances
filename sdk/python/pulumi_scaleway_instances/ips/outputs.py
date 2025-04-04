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
from . import outputs

__all__ = [
    'ScalewayInstanceV1Ip',
    'ScalewayInstanceV1ServerSummary',
]

@pulumi.output_type
class ScalewayInstanceV1Ip(dict):
    def __init__(__self__, *,
                 project: builtins.str,
                 address: Optional[builtins.str] = None,
                 id: Optional[builtins.str] = None,
                 organization: Optional[builtins.str] = None,
                 reverse: Optional[builtins.str] = None,
                 server: Optional['outputs.ScalewayInstanceV1ServerSummary'] = None,
                 tags: Optional[Sequence[builtins.str]] = None,
                 zone: Optional[builtins.str] = None):
        """
        :param builtins.str address: (IPv4 address)
        """
        pulumi.set(__self__, "project", project)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def project(self) -> builtins.str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def address(self) -> Optional[builtins.str]:
        """
        (IPv4 address)
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def organization(self) -> Optional[builtins.str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def reverse(self) -> Optional[builtins.str]:
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def server(self) -> Optional['outputs.ScalewayInstanceV1ServerSummary']:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[builtins.str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[builtins.str]:
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1ServerSummary(dict):
    def __init__(__self__, *,
                 id: Optional[builtins.str] = None,
                 name: Optional[builtins.str] = None):
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")


