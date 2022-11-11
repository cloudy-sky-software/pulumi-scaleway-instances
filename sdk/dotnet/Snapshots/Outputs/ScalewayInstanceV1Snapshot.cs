// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Snapshots.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1Snapshot
    {
        /// <summary>
        /// The volume on which the snapshot is based on
        /// </summary>
        public readonly Outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties? Base_volume;
        /// <summary>
        /// The snapshot creation date (RFC 3339 format)
        /// </summary>
        public readonly string? Creation_date;
        /// <summary>
        /// The reason for the failed snapshot import
        /// </summary>
        public readonly string? Error_reason;
        /// <summary>
        /// The snapshot modification date (RFC 3339 format)
        /// </summary>
        public readonly string? Modification_date;
        /// <summary>
        /// The snapshot name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The snapshot organization ID
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// The snapshot project ID
        /// </summary>
        public readonly string? Project;
        /// <summary>
        /// The snapshot size (in bytes)
        /// </summary>
        public readonly double? Size;
        public readonly Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotState? State;
        /// <summary>
        /// The snapshot tags
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotVolumeType? Volume_type;
        /// <summary>
        /// The snapshot zone
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private ScalewayInstanceV1Snapshot(
            Outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties? base_volume,

            string? creation_date,

            string? error_reason,

            string? modification_date,

            string? name,

            string? organization,

            string? project,

            double? size,

            Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotState? state,

            ImmutableArray<string> tags,

            Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotVolumeType? volume_type,

            string? zone)
        {
            Base_volume = base_volume;
            Creation_date = creation_date;
            Error_reason = error_reason;
            Modification_date = modification_date;
            Name = name;
            Organization = organization;
            Project = project;
            Size = size;
            State = state;
            Tags = tags;
            Volume_type = volume_type;
            Zone = zone;
        }
    }
}
