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
    /// Network available for the instance
    /// </summary>
    [OutputType]
    public sealed class ScalewayInstanceV1ServerTypeNetworkProperties
    {
        /// <summary>
        /// True if IPv6 is enabled
        /// </summary>
        public readonly bool? _ipv6Support;
        /// <summary>
        /// List of available network interfaces
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalewayInstanceV1ServerTypeNetworkInterface> Interfaces;
        /// <summary>
        /// Total maximum internal bandwidth in bits per seconds
        /// </summary>
        public readonly double? SumInternalBandwidth;
        /// <summary>
        /// Total maximum internet bandwidth in bits per seconds
        /// </summary>
        public readonly double? SumInternetBandwidth;

        [OutputConstructor]
        private ScalewayInstanceV1ServerTypeNetworkProperties(
            bool? _ipv6Support,

            ImmutableArray<Outputs.ScalewayInstanceV1ServerTypeNetworkInterface> interfaces,

            double? sumInternalBandwidth,

            double? sumInternetBandwidth)
        {
            this._ipv6Support = _ipv6Support;
            Interfaces = interfaces;
            SumInternalBandwidth = sumInternalBandwidth;
            SumInternetBandwidth = sumInternetBandwidth;
        }
    }
}
