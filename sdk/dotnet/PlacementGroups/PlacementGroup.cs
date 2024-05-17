// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.PlacementGroups
{
    [ScalewayInstancesResourceType("scaleway-instances:placement_groups:PlacementGroup")]
    public partial class PlacementGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The placement group name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The placement group organization ID
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        [Output("placementGroup")]
        public Output<Outputs.ScalewayInstanceV1PlacementGroup?> PlacementGroupValue { get; private set; } = null!;

        [Output("policyMode")]
        public Output<Pulumi.ScalewayInstances.PlacementGroups.PlacementGroupPolicyMode?> PolicyMode { get; private set; } = null!;

        /// <summary>
        /// Returns true if the policy is respected, false otherwise
        /// </summary>
        [Output("policyRespected")]
        public Output<bool?> PolicyRespected { get; private set; } = null!;

        [Output("policyType")]
        public Output<Pulumi.ScalewayInstances.PlacementGroups.PlacementGroupPolicyType?> PolicyType { get; private set; } = null!;

        /// <summary>
        /// The placement group project ID
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The placement group tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The zone in which is the placement group
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a PlacementGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PlacementGroup(string name, PlacementGroupArgs args, CustomResourceOptions? options = null)
            : base("scaleway-instances:placement_groups:PlacementGroup", name, args ?? new PlacementGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PlacementGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("scaleway-instances:placement_groups:PlacementGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PlacementGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PlacementGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PlacementGroup(name, id, options);
        }
    }

    public sealed class PlacementGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The placement group name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The placement group organization ID
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("policyMode")]
        public Input<Pulumi.ScalewayInstances.PlacementGroups.PlacementGroupPolicyMode>? PolicyMode { get; set; }

        [Input("policyType")]
        public Input<Pulumi.ScalewayInstances.PlacementGroups.PlacementGroupPolicyType>? PolicyType { get; set; }

        /// <summary>
        /// The placement group project ID
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The placement group tags
        /// </summary>
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

        public PlacementGroupArgs()
        {
            PolicyMode = Pulumi.ScalewayInstances.PlacementGroups.PlacementGroupPolicyMode.Optional;
            PolicyType = Pulumi.ScalewayInstances.PlacementGroups.PlacementGroupPolicyType.MaxAvailability;
        }
        public static new PlacementGroupArgs Empty => new PlacementGroupArgs();
    }
}
