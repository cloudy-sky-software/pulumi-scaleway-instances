# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ScalewayInstanceV1ListServerActionsResponse',
    'ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate',
]

@pulumi.output_type
class ScalewayInstanceV1ListServerActionsResponse(dict):
    def __init__(__self__, *,
                 actions: Optional[Sequence['ScalewayInstanceV1ListServerActionsResponseActionsItem']] = None):
        if actions is not None:
            pulumi.set(__self__, "actions", actions)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence['ScalewayInstanceV1ListServerActionsResponseActionsItem']]:
        return pulumi.get(self, "actions")


@pulumi.output_type
class ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate(dict):
    def __init__(__self__, *,
                 volume_type: Optional['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType'] = None):
        if volume_type is None:
            volume_type = 'l_ssd'
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def volume_type(self) -> Optional['ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType']:
        return pulumi.get(self, "volume_type")

