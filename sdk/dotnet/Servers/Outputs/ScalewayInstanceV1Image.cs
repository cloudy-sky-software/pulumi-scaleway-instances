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
    public sealed class ScalewayInstanceV1Image
    {
        public readonly Pulumi.ScalewayInstances.Servers.ScalewayInstanceV1ImageArch? Arch;
        /// <summary>
        /// (RFC 3339 format)
        /// </summary>
        public readonly string? Creation_date;
        public readonly Outputs.ScalewayInstanceV1Bootscript? Default_bootscript;
        public readonly ImmutableDictionary<string, Outputs.ScalewayInstanceV1Volume>? Extra_volumes;
        public readonly string? From_server;
        public readonly string? Id;
        /// <summary>
        /// (RFC 3339 format)
        /// </summary>
        public readonly string? Modification_date;
        public readonly string Name;
        public readonly string? Organization;
        public readonly string Project;
        public readonly bool? Public;
        public readonly Outputs.ScalewayInstanceV1VolumeSummary Root_volume;
        public readonly Pulumi.ScalewayInstances.Servers.ScalewayInstanceV1ImageState? State;
        public readonly ImmutableArray<string> Tags;
        public readonly string? Zone;

        [OutputConstructor]
        private ScalewayInstanceV1Image(
            Pulumi.ScalewayInstances.Servers.ScalewayInstanceV1ImageArch? arch,

            string? creation_date,

            Outputs.ScalewayInstanceV1Bootscript? default_bootscript,

            ImmutableDictionary<string, Outputs.ScalewayInstanceV1Volume>? extra_volumes,

            string? from_server,

            string? id,

            string? modification_date,

            string name,

            string? organization,

            string project,

            bool? @public,

            Outputs.ScalewayInstanceV1VolumeSummary root_volume,

            Pulumi.ScalewayInstances.Servers.ScalewayInstanceV1ImageState? state,

            ImmutableArray<string> tags,

            string? zone)
        {
            Arch = arch;
            Creation_date = creation_date;
            Default_bootscript = default_bootscript;
            Extra_volumes = extra_volumes;
            From_server = from_server;
            Id = id;
            Modification_date = modification_date;
            Name = name;
            Organization = organization;
            Project = project;
            Public = @public;
            Root_volume = root_volume;
            State = state;
            Tags = tags;
            Zone = zone;
        }
    }
}
