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
    public sealed class ScalewayInstanceV1PlacementGroupServer
    {
        public readonly string? Id;
        public readonly string? Name;
        public readonly bool? PolicyRespected;

        [OutputConstructor]
        private ScalewayInstanceV1PlacementGroupServer(
            string? id,

            string? name,

            bool? policyRespected)
        {
            Id = id;
            Name = name;
            PolicyRespected = policyRespected;
        }
    }
}
