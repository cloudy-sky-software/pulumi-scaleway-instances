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

__all__ = ['ServerActionArgs', 'ServerAction']

@pulumi.input_type
class ServerActionArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input['Action']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs']]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServerAction resource.
        :param pulumi.Input['Action'] action: The action to perform on the server
        :param pulumi.Input[str] name: The name of the backup you want to create.
               This field should only be specified when performing a backup action.
        :param pulumi.Input[str] server_id: UUID of the server
        :param pulumi.Input[str] zone: The zone you want to target
        """
        if action is None:
            action = 'poweron'
        if action is not None:
            pulumi.set(__self__, "action", action)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if volumes is not None:
            pulumi.set(__self__, "volumes", volumes)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['Action']]:
        """
        The action to perform on the server
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['Action']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the backup you want to create.
        This field should only be specified when performing a backup action.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the server
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def volumes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs']]]]:
        return pulumi.get(self, "volumes")

    @volumes.setter
    def volumes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs']]]]):
        pulumi.set(self, "volumes", value)

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


class ServerAction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['Action']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs']]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ServerAction resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Action'] action: The action to perform on the server
        :param pulumi.Input[str] name: The name of the backup you want to create.
               This field should only be specified when performing a backup action.
        :param pulumi.Input[str] server_id: UUID of the server
        :param pulumi.Input[str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServerActionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ServerAction resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ServerActionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerActionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['Action']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs']]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerActionArgs.__new__(ServerActionArgs)

            if action is None:
                action = 'poweron'
            __props__.__dict__["action"] = action
            __props__.__dict__["name"] = name
            __props__.__dict__["server_id"] = server_id
            __props__.__dict__["volumes"] = volumes
            __props__.__dict__["zone"] = zone
            __props__.__dict__["task"] = None
        super(ServerAction, __self__).__init__(
            'scaleway-instances:action:ServerAction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServerAction':
        """
        Get an existing ServerAction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServerActionArgs.__new__(ServerActionArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["task"] = None
        __props__.__dict__["volumes"] = None
        return ServerAction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['Action']]:
        """
        The action to perform on the server
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the backup you want to create.
        This field should only be specified when performing a backup action.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def task(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1Task']]:
        return pulumi.get(self, "task")

    @property
    @pulumi.getter
    def volumes(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate']]]:
        return pulumi.get(self, "volumes")

