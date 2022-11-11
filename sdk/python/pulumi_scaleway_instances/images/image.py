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
from ._enums import *
from ._inputs import *

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 arch: Optional[pulumi.Input['Arch']] = None,
                 default_bootscript: Optional[pulumi.Input['ScalewayInstanceV1BootscriptArgs']] = None,
                 extra_volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1VolumeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 root_volume: Optional[pulumi.Input['ScalewayInstanceV1VolumeSummaryArgs']] = None,
                 state: Optional[pulumi.Input['State']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[str] zone: The zone you want to target
        """
        if arch is None:
            arch = 'x86_64'
        if arch is not None:
            pulumi.set(__self__, "arch", arch)
        if default_bootscript is not None:
            pulumi.set(__self__, "default_bootscript", default_bootscript)
        if extra_volumes is not None:
            pulumi.set(__self__, "extra_volumes", extra_volumes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if root_volume is not None:
            pulumi.set(__self__, "root_volume", root_volume)
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
    def arch(self) -> Optional[pulumi.Input['Arch']]:
        return pulumi.get(self, "arch")

    @arch.setter
    def arch(self, value: Optional[pulumi.Input['Arch']]):
        pulumi.set(self, "arch", value)

    @property
    @pulumi.getter
    def default_bootscript(self) -> Optional[pulumi.Input['ScalewayInstanceV1BootscriptArgs']]:
        return pulumi.get(self, "default_bootscript")

    @default_bootscript.setter
    def default_bootscript(self, value: Optional[pulumi.Input['ScalewayInstanceV1BootscriptArgs']]):
        pulumi.set(self, "default_bootscript", value)

    @property
    @pulumi.getter
    def extra_volumes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1VolumeArgs']]]]:
        return pulumi.get(self, "extra_volumes")

    @extra_volumes.setter
    def extra_volumes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1VolumeArgs']]]]):
        pulumi.set(self, "extra_volumes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter
    def root_volume(self) -> Optional[pulumi.Input['ScalewayInstanceV1VolumeSummaryArgs']]:
        return pulumi.get(self, "root_volume")

    @root_volume.setter
    def root_volume(self, value: Optional[pulumi.Input['ScalewayInstanceV1VolumeSummaryArgs']]):
        pulumi.set(self, "root_volume", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['State']]:
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['State']]):
        pulumi.set(self, "state", value)

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


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arch: Optional[pulumi.Input['Arch']] = None,
                 default_bootscript: Optional[pulumi.Input[pulumi.InputType['ScalewayInstanceV1BootscriptArgs']]] = None,
                 extra_volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ScalewayInstanceV1VolumeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 root_volume: Optional[pulumi.Input[pulumi.InputType['ScalewayInstanceV1VolumeSummaryArgs']]] = None,
                 state: Optional[pulumi.Input['State']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Image resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ImageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Image resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arch: Optional[pulumi.Input['Arch']] = None,
                 default_bootscript: Optional[pulumi.Input[pulumi.InputType['ScalewayInstanceV1BootscriptArgs']]] = None,
                 extra_volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ScalewayInstanceV1VolumeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 root_volume: Optional[pulumi.Input[pulumi.InputType['ScalewayInstanceV1VolumeSummaryArgs']]] = None,
                 state: Optional[pulumi.Input['State']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            if arch is None:
                arch = 'x86_64'
            __props__.__dict__["arch"] = arch
            __props__.__dict__["default_bootscript"] = default_bootscript
            __props__.__dict__["extra_volumes"] = extra_volumes
            __props__.__dict__["name"] = name
            __props__.__dict__["organization"] = organization
            __props__.__dict__["project"] = project
            __props__.__dict__["public"] = public
            __props__.__dict__["root_volume"] = root_volume
            if state is None:
                state = 'available'
            __props__.__dict__["state"] = state
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["creation_date"] = None
            __props__.__dict__["from_server"] = None
            __props__.__dict__["modification_date"] = None
        super(Image, __self__).__init__(
            'scaleway-instances:images:Image',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Image':
        """
        Get an existing Image resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ImageArgs.__new__(ImageArgs)

        __props__.__dict__["arch"] = None
        __props__.__dict__["creation_date"] = None
        __props__.__dict__["default_bootscript"] = None
        __props__.__dict__["extra_volumes"] = None
        __props__.__dict__["from_server"] = None
        __props__.__dict__["modification_date"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["organization"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["public"] = None
        __props__.__dict__["root_volume"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["zone"] = None
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arch(self) -> pulumi.Output[Optional['Arch']]:
        return pulumi.get(self, "arch")

    @property
    @pulumi.getter
    def creation_date(self) -> pulumi.Output[Optional[str]]:
        """
        (RFC 3339 format)
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def default_bootscript(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1Bootscript']]:
        return pulumi.get(self, "default_bootscript")

    @property
    @pulumi.getter
    def extra_volumes(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.ScalewayInstanceV1Volume']]]:
        return pulumi.get(self, "extra_volumes")

    @property
    @pulumi.getter
    def from_server(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "from_server")

    @property
    @pulumi.getter
    def modification_date(self) -> pulumi.Output[Optional[str]]:
        """
        (RFC 3339 format)
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def public(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def root_volume(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1VolumeSummary']]:
        return pulumi.get(self, "root_volume")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional['State']]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "zone")
