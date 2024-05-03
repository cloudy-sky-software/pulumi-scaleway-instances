// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Rules
{
    public static class GetSecurityGroupRule
    {
        public static Task<GetSecurityGroupRuleResult> InvokeAsync(GetSecurityGroupRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupRuleResult>("scaleway-instances:rules:getSecurityGroupRule", args ?? new GetSecurityGroupRuleArgs(), options.WithDefaults());

        public static Output<GetSecurityGroupRuleResult> Invoke(GetSecurityGroupRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupRuleResult>("scaleway-instances:rules:getSecurityGroupRule", args ?? new GetSecurityGroupRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("securityGroupId", required: true)]
        public string SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetSecurityGroupRuleArgs()
        {
        }
        public static new GetSecurityGroupRuleArgs Empty => new GetSecurityGroupRuleArgs();
    }

    public sealed class GetSecurityGroupRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetSecurityGroupRuleInvokeArgs()
        {
        }
        public static new GetSecurityGroupRuleInvokeArgs Empty => new GetSecurityGroupRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupRuleResult
    {
        public readonly Outputs.ScalewayInstanceV1GetSecurityGroupRuleResponse Items;

        [OutputConstructor]
        private GetSecurityGroupRuleResult(Outputs.ScalewayInstanceV1GetSecurityGroupRuleResponse items)
        {
            Items = items;
        }
    }
}
