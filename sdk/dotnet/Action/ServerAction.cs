// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Action
{
    [ScalewayInstancesResourceType("scaleway-instances:action:ServerAction")]
    public partial class ServerAction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The action to perform on the server
        /// </summary>
        [Output("action")]
        public Output<Pulumi.ScalewayInstances.Action.Action?> Action { get; private set; } = null!;

        /// <summary>
        /// The name of the backup you want to create.
        /// This field should only be specified when performing a backup action.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("volumes")]
        public Output<ImmutableDictionary<string, Outputs.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate>?> Volumes { get; private set; } = null!;


        /// <summary>
        /// Create a ServerAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerAction(string name, ServerActionArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway-instances:action:ServerAction", name, args ?? new ServerActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerAction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("scaleway-instances:action:ServerAction", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServerAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerAction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerAction(name, id, options);
        }
    }

    public sealed class ServerActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to perform on the server
        /// </summary>
        [Input("action")]
        public Input<Pulumi.ScalewayInstances.Action.Action>? Action { get; set; }

        /// <summary>
        /// UUID of the server
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the backup you want to create.
        /// This field should only be specified when performing a backup action.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("volumes")]
        private InputMap<Inputs.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs>? _volumes;
        public InputMap<Inputs.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputMap<Inputs.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs>());
            set => _volumes = value;
        }

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ServerActionArgs()
        {
            Action = Pulumi.ScalewayInstances.Action.Action.Poweron;
        }
        public static new ServerActionArgs Empty => new ServerActionArgs();
    }
}
