// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Action.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate
    {
        public readonly Pulumi.ScalewayInstances.Action.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType? Volume_type;

        [OutputConstructor]
        private ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate(Pulumi.ScalewayInstances.Action.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType? volume_type)
        {
            Volume_type = volume_type;
        }
    }
}