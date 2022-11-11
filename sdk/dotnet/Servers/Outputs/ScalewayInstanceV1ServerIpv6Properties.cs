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
    /// The server IPv6 address
    /// </summary>
    [OutputType]
    public sealed class ScalewayInstanceV1ServerIpv6Properties
    {
        /// <summary>
        /// The server IPv6 IP-Address (IPv6 address)
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The IPv6 IP-addresses gateway (IPv6 address)
        /// </summary>
        public readonly string? Gateway;
        /// <summary>
        /// The IPv6 IP-addresses CIDR netmask
        /// </summary>
        public readonly string? Netmask;

        [OutputConstructor]
        private ScalewayInstanceV1ServerIpv6Properties(
            string? address,

            string? gateway,

            string? netmask)
        {
            Address = address;
            Gateway = gateway;
            Netmask = netmask;
        }
    }
}