// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Volumes.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1VolumeType
    {
        public readonly Outputs.ScalewayInstanceV1VolumeTypeCapabilities? Capabilities;
        public readonly Outputs.ScalewayInstanceV1VolumeTypeConstraints? Constraints;
        public readonly string? Display_name;

        [OutputConstructor]
        private ScalewayInstanceV1VolumeType(
            Outputs.ScalewayInstanceV1VolumeTypeCapabilities? capabilities,

            Outputs.ScalewayInstanceV1VolumeTypeConstraints? constraints,

            string? display_name)
        {
            Capabilities = capabilities;
            Constraints = constraints;
            Display_name = display_name;
        }
    }
}