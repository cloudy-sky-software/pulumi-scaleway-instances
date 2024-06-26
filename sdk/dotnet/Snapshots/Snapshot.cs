// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Snapshots
{
    [ScalewayInstancesResourceType("scaleway-instances:snapshots:Snapshot")]
    public partial class Snapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The volume on which the snapshot is based on
        /// </summary>
        [Output("baseVolume")]
        public Output<Outputs.BaseVolumeProperties?> BaseVolume { get; private set; } = null!;

        /// <summary>
        /// The snapshot creation date (RFC 3339 format)
        /// </summary>
        [Output("creationDate")]
        public Output<string?> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The reason for the failed snapshot import
        /// </summary>
        [Output("errorReason")]
        public Output<string?> ErrorReason { get; private set; } = null!;

        /// <summary>
        /// The snapshot modification date (RFC 3339 format)
        /// </summary>
        [Output("modificationDate")]
        public Output<string?> ModificationDate { get; private set; } = null!;

        /// <summary>
        /// The snapshot name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The snapshot organization ID
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        /// <summary>
        /// The snapshot project ID
        /// </summary>
        [Output("project")]
        public Output<string?> Project { get; private set; } = null!;

        /// <summary>
        /// The snapshot size (in bytes)
        /// </summary>
        [Output("size")]
        public Output<double?> Size { get; private set; } = null!;

        [Output("snapshot")]
        public Output<Outputs.ScalewayInstanceV1Snapshot?> SnapshotValue { get; private set; } = null!;

        [Output("state")]
        public Output<Pulumi.ScalewayInstances.Snapshots.State?> State { get; private set; } = null!;

        /// <summary>
        /// The snapshot tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("task")]
        public Output<Outputs.ScalewayInstanceV1Task?> Task { get; private set; } = null!;

        [Output("volumeType")]
        public Output<Pulumi.ScalewayInstances.Snapshots.VolumeType?> VolumeType { get; private set; } = null!;

        /// <summary>
        /// The snapshot zone
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway-instances:snapshots:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("scaleway-instances:snapshots:Snapshot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-scaleway-instances",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, options);
        }
    }

    public sealed class SnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The snapshot name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The snapshot organization ID
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The snapshot project ID
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The snapshot size (in bytes)
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        [Input("state")]
        public Input<Pulumi.ScalewayInstances.Snapshots.State>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The snapshot tags
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("volumeType")]
        public Input<Pulumi.ScalewayInstances.Snapshots.VolumeType>? VolumeType { get; set; }

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public SnapshotArgs()
        {
            State = Pulumi.ScalewayInstances.Snapshots.State.Available;
            VolumeType = Pulumi.ScalewayInstances.Snapshots.VolumeType.LSsd;
        }
        public static new SnapshotArgs Empty => new SnapshotArgs();
    }
}
