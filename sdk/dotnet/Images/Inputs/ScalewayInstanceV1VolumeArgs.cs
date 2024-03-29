// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Images.Inputs
{

    public sealed class ScalewayInstanceV1VolumeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The volume creation date (RFC 3339 format)
        /// </summary>
        [Input("creation_date")]
        public Input<string>? Creation_date { get; set; }

        /// <summary>
        /// Show the volume NBD export URI
        /// </summary>
        [Input("export_uri")]
        public Input<string>? Export_uri { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The volume modification date (RFC 3339 format)
        /// </summary>
        [Input("modification_date")]
        public Input<string>? Modification_date { get; set; }

        /// <summary>
        /// The volume name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The volume organization ID
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The volume project ID
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The server attached to the volume
        /// </summary>
        [Input("server")]
        public Input<Inputs.ScalewayInstanceV1VolumeServerPropertiesArgs>? Server { get; set; }

        /// <summary>
        /// The volume disk size (in bytes)
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        [Input("state")]
        public Input<Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeState>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The volume tags
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("volume_type")]
        public Input<Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeVolumeType>? Volume_type { get; set; }

        /// <summary>
        /// The zone in which is the volume
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ScalewayInstanceV1VolumeArgs()
        {
            State = Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeState.Available;
            Volume_type = Pulumi.ScalewayInstances.Images.ScalewayInstanceV1VolumeVolumeType.LSsd;
        }
        public static new ScalewayInstanceV1VolumeArgs Empty => new ScalewayInstanceV1VolumeArgs();
    }
}
