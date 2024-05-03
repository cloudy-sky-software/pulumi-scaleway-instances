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
        public readonly Outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties? BaseVolume;
        /// <summary>
        /// The snapshot creation date (RFC 3339 format)
        /// </summary>
        public readonly string? CreationDate;
        /// <summary>
        /// The reason for the failed snapshot import
        /// </summary>
        public readonly string? ErrorReason;
        public readonly string? Id;
        /// <summary>
        /// The snapshot modification date (RFC 3339 format)
        /// </summary>
        public readonly string? ModificationDate;
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
        public readonly Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotVolumeType? VolumeType;
        /// <summary>
        /// The snapshot zone
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private ScalewayInstanceV1Snapshot(
            Outputs.ScalewayInstanceV1SnapshotBaseVolumeProperties? baseVolume,

            string? creationDate,

            string? errorReason,

            string? id,

            string? modificationDate,

            string? name,

            string? organization,

            string? project,

            double? size,

            Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotState? state,

            ImmutableArray<string> tags,

            Pulumi.ScalewayInstances.Snapshots.ScalewayInstanceV1SnapshotVolumeType? volumeType,

            string? zone)
        {
            BaseVolume = baseVolume;
            CreationDate = creationDate;
            ErrorReason = errorReason;
            Id = id;
            ModificationDate = modificationDate;
            Name = name;
            Organization = organization;
            Project = project;
            Size = size;
            State = state;
            Tags = tags;
            VolumeType = volumeType;
            Zone = zone;
        }
    }
}
