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

__all__ = ['PrivateNICArgs', 'PrivateNIC']

@pulumi.input_type
class PrivateNICArgs:
    def __init__(__self__, *,
                 private_network_id: pulumi.Input[builtins.str],
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a PrivateNIC resource.
        :param pulumi.Input[builtins.str] zone: The zone you want to target
        """
        pulumi.set(__self__, "private_network_id", private_network_id)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "server_id", value)

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


class PrivateNIC(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 private_network_id: Optional[pulumi.Input[builtins.str]] = None,
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a PrivateNIC resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] zone: The zone you want to target
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateNICArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PrivateNIC resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PrivateNICArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateNICArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 private_network_id: Optional[pulumi.Input[builtins.str]] = None,
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateNICArgs.__new__(PrivateNICArgs)

            if private_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'private_network_id'")
            __props__.__dict__["private_network_id"] = private_network_id
            __props__.__dict__["server_id"] = server_id
            __props__.__dict__["zone"] = zone
            __props__.__dict__["private_nic"] = None
        super(PrivateNIC, __self__).__init__(
            'scaleway-instances:private_nics:PrivateNIC',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateNIC':
        """
        Get an existing PrivateNIC resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateNICArgs.__new__(PrivateNICArgs)

        __props__.__dict__["private_network_id"] = None
        __props__.__dict__["private_nic"] = None
        return PrivateNIC(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter(name="privateNic")
    def private_nic(self) -> pulumi.Output[Optional['outputs.ScalewayInstanceV1PrivateNIC']]:
        return pulumi.get(self, "private_nic")

