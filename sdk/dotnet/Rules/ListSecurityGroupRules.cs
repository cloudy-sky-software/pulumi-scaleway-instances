// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Rules
{
    public static class ListSecurityGroupRules
    {
        public static Task<Outputs.ScalewayInstanceV1ListSecurityGroupRulesResponse> InvokeAsync(ListSecurityGroupRulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayInstanceV1ListSecurityGroupRulesResponse>("scaleway-instances:rules:listSecurityGroupRules", args ?? new ListSecurityGroupRulesArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1ListSecurityGroupRulesResponse> Invoke(ListSecurityGroupRulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1ListSecurityGroupRulesResponse>("scaleway-instances:rules:listSecurityGroupRules", args ?? new ListSecurityGroupRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSecurityGroupRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the security group
        /// </summary>
        [Input("securityGroupId", required: true)]
        public string SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListSecurityGroupRulesArgs()
        {
        }
        public static new ListSecurityGroupRulesArgs Empty => new ListSecurityGroupRulesArgs();
    }

    public sealed class ListSecurityGroupRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the security group
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListSecurityGroupRulesInvokeArgs()
        {
        }
        public static new ListSecurityGroupRulesInvokeArgs Empty => new ListSecurityGroupRulesInvokeArgs();
    }
}
