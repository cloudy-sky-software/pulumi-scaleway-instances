// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.PlacementGroups.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1GetPlacementGroupResponse
    {
        public readonly Outputs.ScalewayInstanceV1PlacementGroup? Placement_group;

        [OutputConstructor]
        private ScalewayInstanceV1GetPlacementGroupResponse(Outputs.ScalewayInstanceV1PlacementGroup? placement_group)
        {
            Placement_group = placement_group;
        }
    }
}