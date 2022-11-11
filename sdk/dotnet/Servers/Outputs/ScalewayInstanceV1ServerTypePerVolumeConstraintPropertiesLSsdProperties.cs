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
        public readonly double? Max_size;
        /// <summary>
        /// Minimum volume size in bytes (in bytes)
        /// </summary>
        public readonly double? Min_size;

        [OutputConstructor]
        private ScalewayInstanceV1ServerTypePerVolumeConstraintPropertiesLSsdProperties(
            double? max_size,

            double? min_size)
        {
            Max_size = max_size;
            Min_size = min_size;
        }
    }
}
