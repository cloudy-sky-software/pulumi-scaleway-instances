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
    /// Information about the public IP
    /// </summary>
    [OutputType]
    public sealed class ScalewayInstanceV1ServerPublicIpProperties
    {
        /// <summary>
        /// The server public IPv4 IP-Address (IPv4 address)
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// True if the IP address is dynamic
        /// </summary>
        public readonly bool? Dynamic;
        /// <summary>
        /// The unique ID of the IP address
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ScalewayInstanceV1ServerPublicIpProperties(
            string? address,

            bool? dynamic,

            string? id)
        {
            Address = address;
            Dynamic = dynamic;
            Id = id;
        }
    }
}
