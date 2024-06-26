# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ScalewayInstanceV1ListServerUserDataResponse',
    'ScalewayStdFile',
]

@pulumi.output_type
class ScalewayInstanceV1ListServerUserDataResponse(dict):
    def __init__(__self__, *,
                 user_data: Optional[Sequence[str]] = None):
        if user_data is not None:
            pulumi.set(__self__, "user_data", user_data)

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "user_data")


@pulumi.output_type
class ScalewayStdFile(dict):
    def __init__(__self__, *,
                 content: Optional[str] = None,
                 content_type: Optional[str] = None,
                 name: Optional[str] = None):
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> Optional[str]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


