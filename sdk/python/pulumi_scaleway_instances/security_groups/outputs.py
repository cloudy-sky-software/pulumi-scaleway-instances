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
from ._enums import *

__all__ = [
    'ScalewayInstanceV1SecurityGroup',
    'ScalewayInstanceV1ServerSummary',
]

@pulumi.output_type
class ScalewayInstanceV1SecurityGroup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "creationDate":
            suggest = "creation_date"
        elif key == "enableDefaultSecurity":
            suggest = "enable_default_security"
        elif key == "inboundDefaultPolicy":
            suggest = "inbound_default_policy"
        elif key == "modificationDate":
            suggest = "modification_date"
        elif key == "organizationDefault":
            suggest = "organization_default"
        elif key == "outboundDefaultPolicy":
            suggest = "outbound_default_policy"
        elif key == "projectDefault":
            suggest = "project_default"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalewayInstanceV1SecurityGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalewayInstanceV1SecurityGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalewayInstanceV1SecurityGroup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: builtins.str,
                 project: builtins.str,
                 creation_date: Optional[builtins.str] = None,
                 description: Optional[builtins.str] = None,
                 enable_default_security: Optional[builtins.bool] = None,
                 id: Optional[builtins.str] = None,
                 inbound_default_policy: Optional['ScalewayInstanceV1SecurityGroupInboundDefaultPolicy'] = None,
                 modification_date: Optional[builtins.str] = None,
                 organization: Optional[builtins.str] = None,
                 organization_default: Optional[builtins.bool] = None,
                 outbound_default_policy: Optional['ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy'] = None,
                 project_default: Optional[builtins.bool] = None,
                 servers: Optional[Sequence['outputs.ScalewayInstanceV1ServerSummary']] = None,
                 state: Optional['ScalewayInstanceV1SecurityGroupState'] = None,
                 stateful: Optional[builtins.bool] = None,
                 tags: Optional[Sequence[builtins.str]] = None,
                 zone: Optional[builtins.str] = None):
        """
        :param builtins.str name: The security groups name
        :param builtins.str project: The security group project ID
        :param builtins.str creation_date: The security group creation date (RFC 3339 format)
        :param builtins.str description: The security groups description
        :param builtins.bool enable_default_security: True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
        :param 'ScalewayInstanceV1SecurityGroupInboundDefaultPolicy' inbound_default_policy: The default inbound policy
        :param builtins.str modification_date: The security group modification date (RFC 3339 format)
        :param builtins.str organization: The security groups organization ID
        :param builtins.bool organization_default: True if it is your default security group for this organization ID
        :param 'ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy' outbound_default_policy: The default outbound policy
        :param builtins.bool project_default: True if it is your default security group for this project ID
        :param Sequence['ScalewayInstanceV1ServerSummary'] servers: List of servers attached to this security group
        :param 'ScalewayInstanceV1SecurityGroupState' state: Security group state
        :param builtins.bool stateful: True if the security group is stateful
        :param Sequence[builtins.str] tags: The security group tags
        :param builtins.str zone: The zone in which is the security group
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project", project)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_default_security is not None:
            pulumi.set(__self__, "enable_default_security", enable_default_security)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if inbound_default_policy is None:
            inbound_default_policy = 'accept'
        if inbound_default_policy is not None:
            pulumi.set(__self__, "inbound_default_policy", inbound_default_policy)
        if modification_date is not None:
            pulumi.set(__self__, "modification_date", modification_date)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if organization_default is not None:
            pulumi.set(__self__, "organization_default", organization_default)
        if outbound_default_policy is None:
            outbound_default_policy = 'accept'
        if outbound_default_policy is not None:
            pulumi.set(__self__, "outbound_default_policy", outbound_default_policy)
        if project_default is not None:
            pulumi.set(__self__, "project_default", project_default)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)
        if state is None:
            state = 'available'
        if state is not None:
            pulumi.set(__self__, "state", state)
        if stateful is not None:
            pulumi.set(__self__, "stateful", stateful)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The security groups name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> builtins.str:
        """
        The security group project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[builtins.str]:
        """
        The security group creation date (RFC 3339 format)
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The security groups description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableDefaultSecurity")
    def enable_default_security(self) -> Optional[builtins.bool]:
        """
        True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
        """
        return pulumi.get(self, "enable_default_security")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inboundDefaultPolicy")
    def inbound_default_policy(self) -> Optional['ScalewayInstanceV1SecurityGroupInboundDefaultPolicy']:
        """
        The default inbound policy
        """
        return pulumi.get(self, "inbound_default_policy")

    @property
    @pulumi.getter(name="modificationDate")
    def modification_date(self) -> Optional[builtins.str]:
        """
        The security group modification date (RFC 3339 format)
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def organization(self) -> Optional[builtins.str]:
        """
        The security groups organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="organizationDefault")
    def organization_default(self) -> Optional[builtins.bool]:
        """
        True if it is your default security group for this organization ID
        """
        return pulumi.get(self, "organization_default")

    @property
    @pulumi.getter(name="outboundDefaultPolicy")
    def outbound_default_policy(self) -> Optional['ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy']:
        """
        The default outbound policy
        """
        return pulumi.get(self, "outbound_default_policy")

    @property
    @pulumi.getter(name="projectDefault")
    def project_default(self) -> Optional[builtins.bool]:
        """
        True if it is your default security group for this project ID
        """
        return pulumi.get(self, "project_default")

    @property
    @pulumi.getter
    def servers(self) -> Optional[Sequence['outputs.ScalewayInstanceV1ServerSummary']]:
        """
        List of servers attached to this security group
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter
    def state(self) -> Optional['ScalewayInstanceV1SecurityGroupState']:
        """
        Security group state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def stateful(self) -> Optional[builtins.bool]:
        """
        True if the security group is stateful
        """
        return pulumi.get(self, "stateful")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[builtins.str]]:
        """
        The security group tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[builtins.str]:
        """
        The zone in which is the security group
        """
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


