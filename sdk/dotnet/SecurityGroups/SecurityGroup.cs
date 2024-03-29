// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.SecurityGroups
{
    [ScalewayInstancesResourceType("scaleway-instances:security_groups:SecurityGroup")]
    public partial class SecurityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The security group creation date (RFC 3339 format)
        /// </summary>
        [Output("creation_date")]
        public Output<string?> Creation_date { get; private set; } = null!;

        /// <summary>
        /// The security groups description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
        /// </summary>
        [Output("enable_default_security")]
        public Output<bool?> Enable_default_security { get; private set; } = null!;

        /// <summary>
        /// The default inbound policy
        /// </summary>
        [Output("inbound_default_policy")]
        public Output<Pulumi.ScalewayInstances.SecurityGroups.InboundDefaultPolicy?> Inbound_default_policy { get; private set; } = null!;

        /// <summary>
        /// The security group modification date (RFC 3339 format)
        /// </summary>
        [Output("modification_date")]
        public Output<string?> Modification_date { get; private set; } = null!;

        /// <summary>
        /// The security groups name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The security groups organization ID
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        /// <summary>
        /// True if it is your default security group for this organization ID
        /// </summary>
        [Output("organization_default")]
        public Output<bool?> Organization_default { get; private set; } = null!;

        /// <summary>
        /// The default outbound policy
        /// </summary>
        [Output("outbound_default_policy")]
        public Output<Pulumi.ScalewayInstances.SecurityGroups.OutboundDefaultPolicy?> Outbound_default_policy { get; private set; } = null!;

        /// <summary>
        /// The security group project ID
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// True if it is your default security group for this project ID
        /// </summary>
        [Output("project_default")]
        public Output<bool?> Project_default { get; private set; } = null!;

        [Output("security_group")]
        public Output<Outputs.ScalewayInstanceV1SecurityGroup?> SecurityGroupValue { get; private set; } = null!;

        /// <summary>
        /// List of servers attached to this security group
        /// </summary>
        [Output("servers")]
        public Output<ImmutableArray<Outputs.ScalewayInstanceV1ServerSummary>> Servers { get; private set; } = null!;

        /// <summary>
        /// Security group state
        /// </summary>
        [Output("state")]
        public Output<Pulumi.ScalewayInstances.SecurityGroups.State?> State { get; private set; } = null!;

        /// <summary>
        /// True if the security group is stateful
        /// </summary>
        [Output("stateful")]
        public Output<bool?> Stateful { get; private set; } = null!;

        /// <summary>
        /// The security group tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The zone in which is the security group
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroup(string name, SecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("scaleway-instances:security_groups:SecurityGroup", name, args ?? new SecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("scaleway-instances:security_groups:SecurityGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityGroup(name, id, options);
        }
    }

    public sealed class SecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The security groups description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
        /// </summary>
        [Input("enable_default_security")]
        public Input<bool>? Enable_default_security { get; set; }

        /// <summary>
        /// The default inbound policy
        /// </summary>
        [Input("inbound_default_policy")]
        public Input<Pulumi.ScalewayInstances.SecurityGroups.InboundDefaultPolicy>? Inbound_default_policy { get; set; }

        /// <summary>
        /// The security groups name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The security groups organization ID
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// True if it is your default security group for this organization ID
        /// </summary>
        [Input("organization_default")]
        public Input<bool>? Organization_default { get; set; }

        /// <summary>
        /// The default outbound policy
        /// </summary>
        [Input("outbound_default_policy")]
        public Input<Pulumi.ScalewayInstances.SecurityGroups.OutboundDefaultPolicy>? Outbound_default_policy { get; set; }

        /// <summary>
        /// The security group project ID
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// True if it is your default security group for this project ID
        /// </summary>
        [Input("project_default")]
        public Input<bool>? Project_default { get; set; }

        /// <summary>
        /// True if the security group is stateful
        /// </summary>
        [Input("stateful")]
        public Input<bool>? Stateful { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The security group tags
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

        public SecurityGroupArgs()
        {
            Inbound_default_policy = Pulumi.ScalewayInstances.SecurityGroups.InboundDefaultPolicy.Accept;
            Outbound_default_policy = Pulumi.ScalewayInstances.SecurityGroups.OutboundDefaultPolicy.Accept;
        }
        public static new SecurityGroupArgs Empty => new SecurityGroupArgs();
    }
}
