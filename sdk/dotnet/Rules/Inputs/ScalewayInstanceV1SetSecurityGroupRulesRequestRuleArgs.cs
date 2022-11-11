// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Rules.Inputs
{

    public sealed class ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to apply when the rule matches a packet
        /// </summary>
        [Input("action")]
        public Input<Pulumi.ScalewayInstances.Rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction>? Action { get; set; }

        /// <summary>
        /// Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
        /// </summary>
        [Input("dest_port_from")]
        public Input<double>? Dest_port_from { get; set; }

        /// <summary>
        /// End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
        /// </summary>
        [Input("dest_port_to")]
        public Input<double>? Dest_port_to { get; set; }

        /// <summary>
        /// Direction the rule applies to
        /// </summary>
        [Input("direction")]
        public Input<Pulumi.ScalewayInstances.Rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection>? Direction { get; set; }

        /// <summary>
        /// Indicates if this rule is editable. Rules with the value false will be ignored
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        /// <summary>
        /// UUID of the security rule to update. If no value is provided, a new rule will be created
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The range of IP address this rules applies to (IP network)
        /// </summary>
        [Input("ip_range")]
        public Input<string>? Ip_range { get; set; }

        /// <summary>
        /// Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
        /// </summary>
        [Input("position")]
        public Input<double>? Position { get; set; }

        /// <summary>
        /// Protocol family this rule applies to
        /// </summary>
        [Input("protocol")]
        public Input<Pulumi.ScalewayInstances.Rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol>? Protocol { get; set; }

        /// <summary>
        /// Zone of the rule. This field is ignored
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs()
        {
            Action = Pulumi.ScalewayInstances.Rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction.Accept;
            Direction = Pulumi.ScalewayInstances.Rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection.Inbound;
            Protocol = Pulumi.ScalewayInstances.Rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol.Tcp;
        }
        public static new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs Empty => new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs();
    }
}