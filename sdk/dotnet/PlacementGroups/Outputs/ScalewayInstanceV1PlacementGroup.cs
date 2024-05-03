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
    public sealed class ScalewayInstanceV1PlacementGroup
    {
        public readonly string? Id;
        /// <summary>
        /// The placement group name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The placement group organization ID
        /// </summary>
        public readonly string? Organization;
        public readonly Pulumi.ScalewayInstances.PlacementGroups.ScalewayInstanceV1PlacementGroupPolicyMode? PolicyMode;
        /// <summary>
        /// Returns true if the policy is respected, false otherwise
        /// </summary>
        public readonly bool? PolicyRespected;
        public readonly Pulumi.ScalewayInstances.PlacementGroups.ScalewayInstanceV1PlacementGroupPolicyType? PolicyType;
        /// <summary>
        /// The placement group project ID
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The placement group tags
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The zone in which is the placement group
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private ScalewayInstanceV1PlacementGroup(
            string? id,

            string name,

            string? organization,

            Pulumi.ScalewayInstances.PlacementGroups.ScalewayInstanceV1PlacementGroupPolicyMode? policyMode,

            bool? policyRespected,

            Pulumi.ScalewayInstances.PlacementGroups.ScalewayInstanceV1PlacementGroupPolicyType? policyType,

            string project,

            ImmutableArray<string> tags,

            string? zone)
        {
            Id = id;
            Name = name;
            Organization = organization;
            PolicyMode = policyMode;
            PolicyRespected = policyRespected;
            PolicyType = policyType;
            Project = project;
            Tags = tags;
            Zone = zone;
        }
    }
}
