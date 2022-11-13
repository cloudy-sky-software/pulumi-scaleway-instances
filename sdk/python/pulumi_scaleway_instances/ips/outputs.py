# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GoogleProtobufStringValue',
    'ScalewayInstanceV1GetIpResponse',
    'ScalewayInstanceV1Ip',
    'ScalewayInstanceV1ListIpsResponse',
    'ScalewayInstanceV1ServerSummary',
]

@pulumi.output_type
class GoogleProtobufStringValue(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class ScalewayInstanceV1GetIpResponse(dict):
    def __init__(__self__, *,
                 ip: Optional['outputs.ScalewayInstanceV1Ip'] = None):
        if ip is not None:
            pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter
    def ip(self) -> Optional['outputs.ScalewayInstanceV1Ip']:
        return pulumi.get(self, "ip")


@pulumi.output_type
class ScalewayInstanceV1Ip(dict):
    def __init__(__self__, *,
                 project: str,
                 address: Optional[str] = None,
                 id: Optional[str] = None,
                 organization: Optional[str] = None,
                 reverse: Optional['outputs.GoogleProtobufStringValue'] = None,
                 server: Optional['outputs.ScalewayInstanceV1ServerSummary'] = None,
                 tags: Optional[Sequence[str]] = None,
                 zone: Optional[str] = None):
        """
        :param str address: (IPv4 address)
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
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        (IPv4 address)
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def reverse(self) -> Optional['outputs.GoogleProtobufStringValue']:
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def server(self) -> Optional['outputs.ScalewayInstanceV1ServerSummary']:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


@pulumi.output_type
class ScalewayInstanceV1ListIpsResponse(dict):
    def __init__(__self__, *,
                 ips: Optional[Sequence['outputs.ScalewayInstanceV1Ip']] = None):
        """
        :param Sequence['ScalewayInstanceV1Ip'] ips: List of ips
        """
        if ips is not None:
            pulumi.set(__self__, "ips", ips)

    @property
    @pulumi.getter
    def ips(self) -> Optional[Sequence['outputs.ScalewayInstanceV1Ip']]:
        """
        List of ips
        """
        return pulumi.get(self, "ips")


@pulumi.output_type
class ScalewayInstanceV1ServerSummary(dict):
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
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


