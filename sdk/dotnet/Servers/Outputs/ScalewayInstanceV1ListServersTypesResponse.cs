// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Servers.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1ListServersTypesResponse
    {
        /// <summary>
        /// List of server types
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ScalewayInstanceV1ServerType>? Servers;

        [OutputConstructor]
        private ScalewayInstanceV1ListServersTypesResponse(ImmutableDictionary<string, Outputs.ScalewayInstanceV1ServerType>? servers)
        {
            Servers = servers;
        }
    }
}
