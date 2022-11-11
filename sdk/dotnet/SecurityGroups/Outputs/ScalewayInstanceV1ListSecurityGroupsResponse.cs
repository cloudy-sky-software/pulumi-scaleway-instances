// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.SecurityGroups.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1ListSecurityGroupsResponse
    {
        public readonly ImmutableArray<Outputs.ScalewayInstanceV1SecurityGroup> Security_groups;
        public readonly double? Total_count;

        [OutputConstructor]
        private ScalewayInstanceV1ListSecurityGroupsResponse(
            ImmutableArray<Outputs.ScalewayInstanceV1SecurityGroup> security_groups,

            double? total_count)
        {
            Security_groups = security_groups;
            Total_count = total_count;
        }
    }
}