// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Images.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1Volume
    {
        /// <summary>
        /// The volume creation date (RFC 3339 format)
        /// </summary>
        public readonly string? Creation_date;
        /// <summary>
        /// Show the volume NBD export URI
        /// </summary>
        public readonly string? Export_uri;
        /// <summary>
        /// The volume modification date (RFC 3339 format)
        /// </summary>
        public readonly string? Modification_date;
        /// <summary>
        /// The volume name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The volume organization ID
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// The volume project ID
        /// </summary>
        public readonly string? Project;
        /// <summary>
        /// The server attached to the volume
        /// </summary>
        public readonly Outputs.ScalewayInstanceV1VolumeServerProperties? Server;
        /// <summary>
        /// The volume disk size (in bytes)
        /// </summary>
        public readonly double? Size;
        public readonly Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeState? State;
        /// <summary>
        /// The volume tags
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeVolumeType? Volume_type;
        /// <summary>
        /// The zone in which is the volume
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private ScalewayInstanceV1Volume(
            string? creation_date,

            string? export_uri,

            string? modification_date,

            string? name,

            string? organization,

            string? project,

            Outputs.ScalewayInstanceV1VolumeServerProperties? server,

            double? size,

            Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeState? state,

            ImmutableArray<string> tags,

            Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeVolumeType? volume_type,

            string? zone)
        {
            Creation_date = creation_date;
            Export_uri = export_uri;
            Modification_date = modification_date;
            Name = name;
            Organization = organization;
            Project = project;
            Server = server;
            Size = size;
            State = state;
            Tags = tags;
            Volume_type = volume_type;
            Zone = zone;
        }
    }
}
