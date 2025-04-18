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
from ._inputs import *

__all__ = ['ServerArgs', 'Server']

@pulumi.input_type
class ServerArgs:
    def __init__(__self__, *,
                 commercial_type: pulumi.Input[builtins.str],
                 boot_type: Optional[pulumi.Input['BootType']] = None,
                 bootscript: Optional[pulumi.Input[builtins.str]] = None,
                 dynamic_ip_required: Optional[pulumi.Input[builtins.bool]] = None,
                 enable_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 image: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 organization: Optional[pulumi.Input[builtins.str]] = None,
                 placement_group: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 public_ip: Optional[pulumi.Input[builtins.str]] = None,
                 security_group: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1VolumeServerTemplateArgs']]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Server resource.
        :param pulumi.Input[builtins.str] commercial_type: Define the server commercial type (i.e. GP1-S)
        :param pulumi.Input['BootType'] boot_type: The boot type to use
        :param pulumi.Input[builtins.str] bootscript: The bootscript ID to use when `boot_type` is set to `bootscript`
        :param pulumi.Input[builtins.bool] dynamic_ip_required: Define if a dynamic IP is required for the instance
        :param pulumi.Input[builtins.bool] enable_ipv6: True if IPv6 is enabled on the server
        :param pulumi.Input[builtins.str] image: The server image ID
        :param pulumi.Input[builtins.str] name: The server name
        :param pulumi.Input[builtins.str] organization: The server organization ID
        :param pulumi.Input[builtins.str] placement_group: Placement group ID if server must be part of a placement group
        :param pulumi.Input[builtins.str] project: The server project ID
        :param pulumi.Input[builtins.str] public_ip: The ID of the reserved IP to attach to the server
        :param pulumi.Input[builtins.str] security_group: The security group ID
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: The server tags
        :param pulumi.Input[builtins.str] zone: The zone you want to target
        """
        pulumi.set(__self__, "commercial_type", commercial_type)
        if boot_type is None:
            boot_type = 'local'
        if boot_type is not None:
            pulumi.set(__self__, "boot_type", boot_type)
        if bootscript is not None:
            pulumi.set(__self__, "bootscript", bootscript)
        if dynamic_ip_required is not None:
            pulumi.set(__self__, "dynamic_ip_required", dynamic_ip_required)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if image is not None:
            pulumi.set(__self__, "image", image)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if placement_group is not None:
            pulumi.set(__self__, "placement_group", placement_group)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if public_ip is not None:
            pulumi.set(__self__, "public_ip", public_ip)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if volumes is not None:
            pulumi.set(__self__, "volumes", volumes)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="commercialType")
    def commercial_type(self) -> pulumi.Input[builtins.str]:
        """
        Define the server commercial type (i.e. GP1-S)
        """
        return pulumi.get(self, "commercial_type")

    @commercial_type.setter
    def commercial_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "commercial_type", value)

    @property
    @pulumi.getter(name="bootType")
    def boot_type(self) -> Optional[pulumi.Input['BootType']]:
        """
        The boot type to use
        """
        return pulumi.get(self, "boot_type")

    @boot_type.setter
    def boot_type(self, value: Optional[pulumi.Input['BootType']]):
        pulumi.set(self, "boot_type", value)

    @property
    @pulumi.getter
    def bootscript(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The bootscript ID to use when `boot_type` is set to `bootscript`
        """
        return pulumi.get(self, "bootscript")

    @bootscript.setter
    def bootscript(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "bootscript", value)

    @property
    @pulumi.getter(name="dynamicIpRequired")
    def dynamic_ip_required(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Define if a dynamic IP is required for the instance
        """
        return pulumi.get(self, "dynamic_ip_required")

    @dynamic_ip_required.setter
    def dynamic_ip_required(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "dynamic_ip_required", value)

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        True if IPv6 is enabled on the server
        """
        return pulumi.get(self, "enable_ipv6")

    @enable_ipv6.setter
    def enable_ipv6(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_ipv6", value)

    @property
    @pulumi.getter
    def image(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The server image ID
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The server name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The server organization ID
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Placement group ID if server must be part of a placement group
        """
        return pulumi.get(self, "placement_group")

    @placement_group.setter
    def placement_group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "placement_group", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The server project ID
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the reserved IP to attach to the server
        """
        return pulumi.get(self, "public_ip")

    @public_ip.setter
    def public_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "public_ip", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The security group ID
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The server tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def volumes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1VolumeServerTemplateArgs']]]]:
        return pulumi.get(self, "volumes")

    @volumes.setter
    def volumes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1VolumeServerTemplateArgs']]]]):
        pulumi.set(self, "volumes", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The zone you want to target
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


class Server(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 boot_type: Optional[pulumi.Input['BootType']] = None,
                 bootscript: Optional[pulumi.Input[builtins.str]] = None,
                 commercial_type: Optional[pulumi.Input[builtins.str]] = None,
                 dynamic_ip_required: Optional[pulumi.Input[builtins.bool]] = None,
                 enable_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 image: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 organization: Optional[pulumi.Input[builtins.str]] = None,
                 placement_group: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 public_ip: Optional[pulumi.Input[builtins.str]] = None,
                 security_group: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['ScalewayInstanceV1VolumeServerTemplateArgs', 'ScalewayInstanceV1VolumeServerTemplateArgsDict']]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a Server resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['BootType'] boot_type: The boot type to use
        :param pulumi.Input[builtins.str] bootscript: The bootscript ID to use when `boot_type` is set to `bootscript`
        :param pulumi.Input[builtins.str] commercial_type: Define the server commercial type (i.e. GP1-S)
        :param pulumi.Input[builtins.bool] dynamic_ip_required: Define if a dynamic IP is required for the instance
        :param pulumi.Input[builtins.bool] enable_ipv6: True if IPv6 is enabled on the server
        :param pulumi.Input[builtins.str] image: The server image ID
        :param pulumi.Input[builtins.str] name: The server name
        :param pulumi.Input[builtins.str] organization: The server organization ID
        :param pulumi.Input[builtins.str] placement_group: Placement group ID if server must be part of a placement group
        :param pulumi.Input[builtins.str] project: The server project ID
        :param pulumi.Input[builtins.str] public_ip: The ID of the reserved IP to attach to the server
        :param pulumi.Input[builtins.str] security_group: The security group ID
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: The server tags
        :param pulumi.Input[builtins.str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Server resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 boot_type: Optional[pulumi.Input['BootType']] = None,
                 bootscript: Optional[pulumi.Input[builtins.str]] = None,
                 commercial_type: Optional[pulumi.Input[builtins.str]] = None,
                 dynamic_ip_required: Optional[pulumi.Input[builtins.bool]] = None,
                 enable_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 image: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 organization: Optional[pulumi.Input[builtins.str]] = None,
                 placement_group: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 public_ip: Optional[pulumi.Input[builtins.str]] = None,
                 security_group: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['ScalewayInstanceV1VolumeServerTemplateArgs', 'ScalewayInstanceV1VolumeServerTemplateArgsDict']]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerArgs.__new__(ServerArgs)

            if boot_type is None:
                boot_type = 'local'
            __props__.__dict__["boot_type"] = boot_type
            __props__.__dict__["bootscript"] = bootscript
            if commercial_type is None and not opts.urn:
                raise TypeError("Missing required property 'commercial_type'")
            __props__.__dict__["commercial_type"] = commercial_type
            __props__.__dict__["dynamic_ip_required"] = dynamic_ip_required
            __props__.__dict__["enable_ipv6"] = enable_ipv6
            __props__.__dict__["image"] = image
            __props__.__dict__["name"] = name
            __props__.__dict__["organization"] = organization
            __props__.__dict__["placement_group"] = placement_group
            __props__.__dict__["project"] = project
            __props__.__dict__["public_ip"] = public_ip
            __props__.__dict__["security_group"] = security_group
            __props__.__dict__["tags"] = tags
            __props__.__dict__["volumes"] = volumes
            __props__.__dict__["zone"] = zone
            __props__.__dict__["server"] = None
        super(Server, __self__).__init__(
            'scaleway-instances:servers:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Server':
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServerArgs.__new__(ServerArgs)

        __props__.__dict__["boot_type"] = None
        __props__.__dict__["bootscript"] = None
        __props__.__dict__["commercial_type"] = None
        __props__.__dict__["dynamic_ip_required"] = None
        __props__.__dict__["enable_ipv6"] = None
        __props__.__dict__["image"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["organization"] = None
        __props__.__dict__["placement_group"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["public_ip"] = None
        __props__.__dict__["security_group"] = None
        __props__.__dict__["server"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["volumes"] = None
        return Server(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bootType")
    def boot_type(self) -> pulumi.Output[Optional['BootType']]:
        """
        The boot type to use
        """
        return pulumi.get(self, "boot_type")

    @property
    @pulumi.getter
    def bootscript(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The bootscript ID to use when `boot_type` is set to `bootscript`
        """
        return pulumi.get(self, "bootscript")

    @property
    @pulumi.getter(name="commercialType")
    def commercial_type(self) -> pulumi.Output[builtins.str]:
        """
        Define the server commercial type (i.e. GP1-S)
        """
        return pulumi.get(self, "commercial_type")

    @property
    @pulumi.getter(name="dynamicIpRequired")
    def dynamic_ip_required(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Define if a dynamic IP is required for the instance
        """
        return pulumi.get(self, "dynamic_ip_required")

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        True if IPv6 is enabled on the server
        """
        return pulumi.get(self, "enable_ipv6")

    @property
    @pulumi.getter
    def image(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The server image ID
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The server name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The server organization ID
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Placement group ID if server must be part of a placement group
        """
        return pulumi.get(self, "placement_group")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The server project ID
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the reserved IP to attach to the server
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The security group ID
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1Server']]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The server tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def volumes(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.ScalewayInstanceV1VolumeServerTemplate']]]:
        return pulumi.get(self, "volumes")

