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
from ._inputs import *

__all__ = ['IpArgs', 'Ip']

@pulumi.input_type
class IpArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 organization: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input['GoogleProtobufStringValueArgs']] = None,
                 server: Optional[pulumi.Input['ScalewayInstanceV1ServerSummaryArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ip resource.
        :param pulumi.Input[str] zone: The zone you want to target
        """
        IpArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project=project,
            organization=organization,
            reverse=reverse,
            server=server,
            tags=tags,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project: pulumi.Input[str],
             organization: Optional[pulumi.Input[str]] = None,
             reverse: Optional[pulumi.Input['GoogleProtobufStringValueArgs']] = None,
             server: Optional[pulumi.Input['ScalewayInstanceV1ServerSummaryArgs']] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             zone: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("project", project)
        if organization is not None:
            _setter("organization", organization)
        if reverse is not None:
            _setter("reverse", reverse)
        if server is not None:
            _setter("server", server)
        if tags is not None:
            _setter("tags", tags)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input['GoogleProtobufStringValueArgs']]:
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input['GoogleProtobufStringValueArgs']]):
        pulumi.set(self, "reverse", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input['ScalewayInstanceV1ServerSummaryArgs']]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input['ScalewayInstanceV1ServerSummaryArgs']]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to target
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Ip(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[pulumi.InputType['GoogleProtobufStringValueArgs']]] = None,
                 server: Optional[pulumi.Input[pulumi.InputType['ScalewayInstanceV1ServerSummaryArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Ip resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ip resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IpArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[pulumi.InputType['GoogleProtobufStringValueArgs']]] = None,
                 server: Optional[pulumi.Input[pulumi.InputType['ScalewayInstanceV1ServerSummaryArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpArgs.__new__(IpArgs)

            __props__.__dict__["organization"] = organization
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if reverse is not None and not isinstance(reverse, GoogleProtobufStringValueArgs):
                reverse = reverse or {}
                def _setter(key, value):
                    reverse[key] = value
                GoogleProtobufStringValueArgs._configure(_setter, **reverse)
            __props__.__dict__["reverse"] = reverse
            if server is not None and not isinstance(server, ScalewayInstanceV1ServerSummaryArgs):
                server = server or {}
                def _setter(key, value):
                    server[key] = value
                ScalewayInstanceV1ServerSummaryArgs._configure(_setter, **server)
            __props__.__dict__["server"] = server
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["location"] = None
            __props__.__dict__["address"] = None
            __props__.__dict__["ip"] = None
        super(Ip, __self__).__init__(
            'scaleway-instances:ips:Ip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Ip':
        """
        Get an existing Ip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IpArgs.__new__(IpArgs)

        __props__.__dict__["location"] = None
        __props__.__dict__["address"] = None
        __props__.__dict__["ip"] = None
        __props__.__dict__["organization"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["reverse"] = None
        __props__.__dict__["server"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["zone"] = None
        return Ip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="Location")
    def location(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[Optional[str]]:
        """
        (IPv4 address)
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1Ip']]:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[Optional['outputs.GoogleProtobufStringValue']]:
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1ServerSummary']]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "zone")

