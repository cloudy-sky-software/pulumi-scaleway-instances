// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Servers.Outputs
{

    /// <summary>
    /// Local SSD volumes
    /// </summary>
    [OutputType]
    public sealed class ScalewayInstanceV1ServerTypePerVolumeConstraintPropertiesLSsdProperties
    {
        /// <summary>
        /// Maximum volume size in bytes (in bytes)
        /// </summary>
        public readonly double? MaxSize;
        /// <summary>
        /// Minimum volume size in bytes (in bytes)
        /// </summary>
        public readonly double? MinSize;

        [OutputConstructor]
        private ScalewayInstanceV1ServerTypePerVolumeConstraintPropertiesLSsdProperties(
            double? maxSize,

            double? minSize)
        {
            MaxSize = maxSize;
            MinSize = minSize;
        }
    }
}
