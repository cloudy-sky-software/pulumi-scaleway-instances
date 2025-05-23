// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Servers
{
    [ScalewayInstancesResourceType("scaleway-instances:servers:PlacementGroupServers")]
    public partial class PlacementGroupServers : global::Pulumi.CustomResource
    {
        [Output("servers")]
        public Output<ImmutableArray<Outputs.ScalewayInstanceV1PlacementGroupServer>> Servers { get; private set; } = null!;


        /// <summary>
        /// Create a PlacementGroupServers resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PlacementGroupServers(string name, PlacementGroupServersArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway-instances:servers:PlacementGroupServers", name, args ?? new PlacementGroupServersArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PlacementGroupServers(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("scaleway-instances:servers:PlacementGroupServers", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PlacementGroupServers resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PlacementGroupServers Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PlacementGroupServers(name, id, options);
        }
    }

    public sealed class PlacementGroupServersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the placement group
        /// </summary>
        [Input("placementGroupId")]
        public Input<string>? PlacementGroupId { get; set; }

        [Input("servers")]
        private InputList<string>? _servers;
        public InputList<string> Servers
        {
            get => _servers ?? (_servers = new InputList<string>());
            set => _servers = value;
        }

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public PlacementGroupServersArgs()
        {
        }
        public static new PlacementGroupServersArgs Empty => new PlacementGroupServersArgs();
    }
}
