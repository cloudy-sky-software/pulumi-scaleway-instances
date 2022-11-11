// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Ips
{
    [ScalewayInstancesResourceType("scaleway-instances:ips:Ip")]
    public partial class Ip : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (IPv4 address)
        /// </summary>
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        [Output("project")]
        public Output<string?> Project { get; private set; } = null!;

        [Output("reverse")]
        public Output<Outputs.GoogleProtobufStringValue?> Reverse { get; private set; } = null!;

        [Output("server")]
        public Output<Outputs.ScalewayInstanceV1ServerSummary?> Server { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Ip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ip(string name, IpArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway-instances:ips:Ip", name, args ?? new IpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ip(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("scaleway-instances:ips:Ip", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Ip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ip Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Ip(name, id, options);
        }
    }

    public sealed class IpArgs : global::Pulumi.ResourceArgs
    {
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("reverse")]
        public Input<Inputs.GoogleProtobufStringValueArgs>? Reverse { get; set; }

        [Input("server")]
        public Input<Inputs.ScalewayInstanceV1ServerSummaryArgs>? Server { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public IpArgs()
        {
        }
        public static new IpArgs Empty => new IpArgs();
    }
}
