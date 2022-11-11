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
    /// The server security group
    /// </summary>
    [OutputType]
    public sealed class ScalewayInstanceV1ServerSecurityGroupProperties
    {
        public readonly string? Id;
        public readonly string? Name;

        [OutputConstructor]
        private ScalewayInstanceV1ServerSecurityGroupProperties(
            string? id,

            string? name)
        {
            Id = id;
            Name = name;
        }
    }
}
